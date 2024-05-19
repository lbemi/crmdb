<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto layout-padding-view">
			<div class="mb15">
				<el-input
					v-model="data.inputValue"
					placeholder="输入名称"
					size="small"
					clearable
					@change="search"
					style="max-width: 300px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="data.type" placeholder="输入者名称" style="max-width: 120px" size="small">
							<el-option label="名称" value="1" size="small" />
						</el-select>
					</template>
					<template #append>
						<el-button size="small" @click="search">
							<el-icon>
								<ele-Search />
							</el-icon>
							查询
						</el-button>
					</template>
				</el-input>
				<el-button type="primary" size="small" class="ml10" @click="createPersistentVolume" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="data.PersistentVolumes"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="data.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column label="名称" show-overflow-tooltip>
					<template #default="scope">
						<el-button :size="theme.themeConfig.globalComponentSize" type="primary" text @click="persistentVolumeDetail(scope.row)">
							{{ scope.row.metadata.name }}</el-button
						>
						<div>
							<el-text type="info">{{ scope.row.spec.storageClassName }}</el-text>
						</div>
					</template>
				</el-table-column>
				<el-table-column label="状态">
					<template #default="scope">
						<el-tag :type="scope.row.status.phase === 'Bound' ? 'success' : 'danger'" effect="plain">{{ scope.row.status.phase }}</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="spec.capacity.storage" label="容量" />
				<el-table-column prop="spec.persistentVolumeReclaimPolicy" label="回收策略" />
				<el-table-column prop="spec.claimRef" label="绑定信息">
					<template #default="scope">
						<div>
							<el-text>命名空间: {{ scope.row.spec.claimRef.namespace }}</el-text>
						</div>
						<div>
							<el-text>名称: {{ scope.row.spec.claimRef.name }}</el-text>
						</div>
					</template>
				</el-table-column>
				<el-table-column prop="spec.accessModes" label="访问模式" />
				<el-table-column label="创建时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>

				<el-table-column fixed="right" label="操作" width="190px" flex>
					<template #default="scope">
						<el-button link type="primary" size="small" @click="persistentVolumeDetail(scope.row)">详情</el-button
						><el-divider direction="vertical" /> <el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button
						><el-divider direction="vertical" />
						<el-button :disabled="scope.row.metadata.name === 'kubernetes'" link type="danger" size="small" @click="deletePersistentVolume(scope.row)"
							>删除</el-button
						>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<pagination :total="data.total" @handlePageChange="handlePageChange"></pagination>
		</el-card>
		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:disabledUpdate="true"
			:code-data="data.codeData"
			:resourceType="'PersistentVolume'"
			v-if="data.dialogVisible"
		/>
		<DrawDialog
			v-model:visible="data.draw.visible"
			:persistentVolume="data.draw.persistentVolume"
			:title="data.draw.title"
			@refresh="listPersistentVolume"
			v-if="data.draw.visible"
		/>
		<PersistentVolumeDetail
			v-model:visible="data.detail.visible"
			:persistentVolume="data.detail.persistentVolume"
			:title="data.detail.title"
			v-if="data.detail.visible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sPersistentVolume">
import { PersistentVolume } from 'kubernetes-models/v1';
import { defineAsyncComponent, h, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { usePersistentVolumeApi } from '@/api/kubernetes/persistentVolume';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const DrawDialog = defineAsyncComponent(() => import('./component/draw.vue'));
const PersistentVolumeDetail = defineAsyncComponent(() => import('./component/detail.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const PersistentVolumeApi = usePersistentVolumeApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	detail: {
		title: '',
		visible: false,
		persistentVolume: {} as PersistentVolume,
	},
	draw: {
		title: '',
		visible: false,
		persistentVolume: {} as PersistentVolume,
	},
	dialogVisible: false,
	codeData: {} as PersistentVolume,
	loading: false,
	selectData: [] as PersistentVolume[],
	PersistentVolumes: [] as PersistentVolume[],
	tmpPersistentVolume: [] as PersistentVolume[],
	total: 0,
	type: '1',
	inputValue: '',
	query: <queryType>{
		page: 1,
		limit: 10,
		cloud: k8sStore.state.activeCluster,
	},
});
onMounted(() => {
	listPersistentVolume();
});

const search = () => {
	if (data.type == '1') {
		data.query.name = data.inputValue;
		delete data.query.label;
	} else if (data.type == '0') {
		data.query.label = data.inputValue;
		delete data.query.name;
	}
	if (data.inputValue === '') {
		delete data.query.label;
		delete data.query.name;
	}
	listPersistentVolume();
};
const createPersistentVolume = () => {
	data.draw.visible = true;
	data.draw.title = '创建PersistentVolume';
	data.draw.persistentVolume = {};
};
const deletePersistentVolume = (PersistentVolume: PersistentVolume) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${PersistentVolume.metadata!.name}`),
			h('span', null, ' PersistentVolume. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			data.loading = true;
			PersistentVolumeApi.deletePersistentVolume(PersistentVolume.metadata!.name!, {
				cloud: k8sStore.state.activeCluster,
			})
				.then((res: any) => {
					listPersistentVolume();
					ElMessage.success(res.message);
				})
				.catch((e) => {
					ElMessage.error(e);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
	data.loading = false;
};
const showYaml = (PersistentVolume: PersistentVolume) => {
	data.dialogVisible = true;
	delete PersistentVolume.metadata?.managedFields;
	data.codeData = PersistentVolume;
	// yamlRef.value.openDialog(PersistentVolume);
};

const handleSelectionChange = () => {};

const persistentVolumeDetail = (persistentVolume: PersistentVolume) => {
	data.detail.visible = true;
	data.detail.title = '详情';
	data.detail.persistentVolume = persistentVolume;
};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listPersistentVolume();
};
const listPersistentVolume = () => {
	data.loading = true;
	PersistentVolumeApi.listPersistentVolume(data.query)
		.then((res: any) => {
			data.PersistentVolumes = res.data.data;
			data.tmpPersistentVolume = res.data.data;
			data.total = res.data.total;
		})
		.catch((e: any) => {
			if (e.code != 5003) ElMessage.error(e.message);
		});
	data.loading = false;
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
</script>

<style scoped lang="scss">
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
