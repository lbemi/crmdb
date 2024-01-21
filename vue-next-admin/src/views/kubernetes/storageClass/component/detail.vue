<template>
	<el-drawer v-model="state.visible" @close="handleClose" size="55%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div class="box_body">
			<el-descriptions :title="state.storageClass.metadata!.name" :column="2" border>
				<el-descriptions-item label="创建时间" label-align="right" align="center" :span="1">{{
					dateStrFormat(state.storageClass.metadata!.creationTimestamp!)
				}}</el-descriptions-item>
				<el-descriptions-item label="提供者" label-align="right" align="center" :span="1">
					{{ state.storageClass.provisioner }}
				</el-descriptions-item>
				<el-descriptions-item label="标签" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.storageClass.metadata!.labels" :key="index">
						<el-tag class="label" type="info" size="default"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.storageClass.metadata!.annotations" :key="index">
						<span> {{ key }}:{{ item }} </span>
					</div>
				</el-descriptions-item>

				<el-descriptions-item label="参数" label-align="right" align="center" :span="1">
					<div v-for="(value, key, index) in state.storageClass.parameters" :key="index">
						<el-tag effect="plain"> {{ key }} : {{ value }}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="回收策略" label-align="right" align="center" :span="1">
					{{ state.storageClass.reclaimPolicy }}
				</el-descriptions-item>
			</el-descriptions>
		</div>

		<el-card shadow="hover" class="layout-padding-auto layout-padding-view" style="margin-top: 15px">
			<el-table :data="state.persistentVolumeClaims" max-height="100vh - 80px">
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button text type="primary"> {{ scope.row.metadata.name }}</el-button>
					</template>
				</el-table-column>
				<el-table-column label="状态">
					<template #default="scope">
						<el-tag :type="scope.row.status.phase === 'Bound' ? 'success' : 'warning'" effect="plain">{{ scope.row.status.phase }}</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="status.capacity.storage" label="容量" />
				<!--				<el-table-column prop="spec.accessModes" label="访问模式" />-->
				<!--			<el-table-column label="关联的存储卷">-->
				<!--				<template #default="scope">-->
				<!--					<el-link :underline="false">{{ scope.row.spec.volumeName }}</el-link>-->
				<!--				</template>-->
				<!--			</el-table-column>-->
				<el-table-column label="创建时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>
				<!--		<el-table-column fixed="right" label="操作" width="260px" flex>-->
				<!--			<template #default="scope">-->
				<!--				<el-button link type="primary" size="small" @click="persistentVolumeClaimDetail(scope.row)">详情</el-button-->
				<!--				><el-divider direction="vertical" />-->
				<!--				<el-button link type="primary" size="small" @click="updatePersistentVolumeClaim(scope.row)">编辑</el-button-->
				<!--				><el-divider direction="vertical" /> <el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button-->
				<!--				><el-divider direction="vertical" />-->
				<!--				<el-button-->
				<!--					:disabled="scope.row.metadata.name === 'kubernetes'"-->
				<!--					link-->
				<!--					type="danger"-->
				<!--					size="small"-->
				<!--					@click="deletePersistentVolumeClaim(scope.row)"-->
				<!--					>删除</el-button-->
				<!--				>-->
				<!--			</template>-->
				<!--		</el-table-column>-->
			</el-table>
			<!-- 分页区域 -->
			<pagination :total="state.total" @handlePageChange="handlePageChange"></pagination>
		</el-card>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer, ElMessage } from 'element-plus';
import { StorageClass } from 'kubernetes-types/storage/v1';
import { onMounted, reactive } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { dateStrFormat } from '@/utils/formatTime';
import { usePersistentVolumeClaimApi } from '@/api/kubernetes/persitentVolumeClaim';
import { kubernetesInfo } from '@/stores/kubernetes';
import { PersistentVolumeClaim } from 'kubernetes-types/core/v1';
import { PageInfo } from '@/types/kubernetes/common';
import Pagination from '@/components/pagination/pagination.vue';

const persistentVolumeClaimApi = usePersistentVolumeClaimApi();
const k8sStore = kubernetesInfo();
const state = reactive({
	visible: false,
	storageClass: {
		metadata: {
			name: '',
			namespace: '',
		},
	} as StorageClass,
	keyValues: [] as Array<{ key: string; value: string }>,
	query: {
		page: 1,
		limit: 10,
		cloud: k8sStore.state.activeCluster,
		storageClassName: '',
	},
	persistentVolumeClaims: [] as Array<PersistentVolumeClaim>,
	total: 0,
});

const handlePageChange = (page: PageInfo) => {
	state.query.page = page.page;
	state.query.limit = page.limit;
	getPersistentVolumeClaimByStorageClassName();
};
const getPersistentVolumeClaimByStorageClassName = () => {
	if (!state.storageClass.metadata || !state.storageClass.metadata.name) {
		return;
	}
	state.query.storageClassName = state.storageClass.metadata.name;
	persistentVolumeClaimApi
		.getPersistentVolumeClaimsByStorageClassName(state.query)
		.then((res: any) => {
			state.persistentVolumeClaims = res.data.data;
			state.total = res.data.total;
		})
		.catch((e: any) => {
			if (e.code != 5003) ElMessage.error(e.message);
		});
};

const props = defineProps({
	visible: Boolean,
	storageClass: {
		type: Object as () => StorageClass,
	},
	title: String,
});

onMounted(() => {
	state.visible = props.visible;
	if (props.storageClass && !isObjectValueEqual(props.storageClass, {})) {
		state.storageClass = props.storageClass;

		getPersistentVolumeClaimByStorageClassName();
	}
});

const emit = defineEmits(['update:visible']);
const handleClose = () => {
	emit('update:visible', false);
};
</script>
<style scoped>
.box_body {
	margin-left: 20px;
	margin-top: 10px;
}

.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
