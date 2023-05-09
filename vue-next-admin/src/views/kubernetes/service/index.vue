<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-text class="mx-1" size="small">命名空间：</el-text>
				<el-select
					v-model="k8sStore.state.activeNamespace"
					style="max-width: 180px"
					class="m-2"
					placeholder="Select"
					size="small"
					@change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata!.name!" />
				</el-select>
				<el-input
					v-model="data.query.key"
					placeholder="输入标签或者名称"
					size="small"
					clearable
					@change="search"
					style="width: 250px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="data.query.type" placeholder="输入标签或者名称" style="width: 60px" size="small">
							<el-option label="标签" value="0" size="small" />
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
				<el-button type="primary" size="small" class="ml10" @click="createService" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
				<el-table :data="data.services" @selection-change="handleSelectionChange" style="width: 100%" max-height="100vh - 235px">
					<el-table-column type="selection" width="35" />
					<el-table-column prop="metadata.name" label="名称" />
					<el-table-column prop="metadata.namespace" label="命名空间" />
					<el-table-column prop="spec.type" label="类型" />
					<el-table-column prop="spec.clusterIP" label="集群IP" />
					<el-table-column label="外部访问IP">
						<template #default="scope">
							<a href="" v-for="item in scope.row.status.loadBalancer.ingress"> {{ item.ip }}</a>
						</template>
					</el-table-column>

					<el-table-column label="标签">
						<template #default="scope">
							<el-tooltip placement="right" effect="light">
								<template #content>
									<div style="display: flex; flex-direction: column">
										<el-tag
											class="label"
											effect="plain"
											type="info"
											v-for="(item, key, index) in scope.row.metadata.labels"
											:key="index"
											size="small"
										>
											{{ key }}:{{ item }}
										</el-tag>
									</div>
								</template>
								<el-tag type="info" effect="plain" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" size="small">
									<div>{{ key }}:{{ item }}</div>
								</el-tag>
							</el-tooltip>
						</template>
					</el-table-column>

					<el-table-column label="创建时间">
						<template #default="scope">
							{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
						</template>
					</el-table-column>

					<el-table-column fixed="right" label="操作" width="220">
						<template #default="scope">
							<el-button link type="primary" size="small" @click="updateService(scope.row)">编辑</el-button><el-divider direction="vertical" />
							<el-button :disabled="scope.row.metadata.name === 'default'" link type="danger" size="small" @click="deleteService(scope.row)"
								>删除</el-button
							>
						</template>
					</el-table-column>
				</el-table>
				<!-- 分页区域 -->
				<pagination :total="data.total" @handlePageChange="handlePageChange"></pagination>
			</div>
		</el-card>
		<!-- <NamespaceDialog
			:title="data.title"
			v-model:visible="data.visible"
			:namespace="data.activeNamespace"
			@value-change="listNamespace()"
			v-if="data.visible"
		/> -->
	</div>
</template>

<script setup lang="ts" name="k8sService">
import { V1Service } from '@kubernetes/client-node';
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useServiceApi } from '/@/api/kubernetes/service';
import { ResponseType } from '/@/types/response';
import { ElMessage } from 'element-plus';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '/@/utils/formatTime';
import { PageInfo } from '/@/types/kubernetes/common';
import { Edit, Delete } from '@element-plus/icons-vue';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));

const k8sStore = kubernetesInfo();
const servieApi = useServiceApi();
const route = useRoute();

const data = reactive({
	loading: false,
	selectData: [] as V1Service[],
	services: [] as V1Service[],
	total: 0,
	query: {
		page: 1,
		limit: 10,
		key: '',
		type: '1',
		cloud: k8sStore.state.activeCluster,
	},
});
onMounted(() => {
	listService();
});
const search = () => {};
const handleChange = () => {
	listService();
};
const createService = () => {};
const deleteService = (service: V1Service) => {};
const handleSelectionChange = () => {};
const updateService = (ervice: V1Service) => {};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listService();
};
const listService = () => {
	servieApi
		.listService(k8sStore.state.activeNamespace, data.query)
		.then((res: ResponseType) => {
			if (res.code === 200) {
				data.services = res.data.data;
				data.total = res.data.total;
			}
		})
		.catch((e) => {
			ElMessage.error(e);
		});
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
