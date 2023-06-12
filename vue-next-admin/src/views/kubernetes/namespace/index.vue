<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-input v-model="data.inputValue" placeholder="输入标签或者名称" size="small" clearable @change="search"
					style="width: 250px; margin-left: 10px">
					<template #prepend>
						<el-select v-model="data.type" placeholder="输入标签或者名称" style="width: 60px" size="small">
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
				<el-button v-auth="'k8s:namespace:add'" type="primary" size="small" class="ml10" @click="createNamespace" :icon="Edit">创建</el-button>
				<el-button v-auth="'k8s:namespace:del'" type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0"
					@click="deleteNamespaces()" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
				<el-table :data="data.namespace" @selection-change="handleSelectionChange" style="width: 100%"
					max-height="100vh - 235px" class="desc-body">
					<el-table-column type="selection" width="35" />
					<el-table-column prop="metadata.name" label="名称" />
					<el-table-column label="状态">
						<template #default="scope">
							<div v-if="scope.row.status.phase === 'Active'">
								<div
									style="display: inline-block; width: 12px; height: 12px; background: #67c23a; border-radius: 50%">
								</div>
								<span style="margin-left: 5px; font-size: 12px; color: #67c23a">{{ scope.row.status.phase }}
								</span>
							</div>
							<div v-else>
								<div
									style="display: inline-block; width: 12px; height: 12px; background: #f56c6c; border-radius: 50%">
								</div>
								<span style="margin-left: 5px; font-size: 12px; color: #f56c6c">{{ scope.row.status.phase }}
								</span>
							</div>
						</template>
					</el-table-column>
					<el-table-column label="标签">
						<template #default="scope">
							<el-tooltip placement="right" effect="light">
								<template #content>
									<div style="display: flex; flex-direction: column">
										<el-tag class="label" effect="plain" type="info"
											v-for="(item, key, index) in scope.row.metadata.labels" :key="index"
											size="small">
											{{ key }}:{{ item }}
										</el-tag>
									</div>
								</template>
								<el-tag type="info" effect="plain" v-for="(item, key, index) in scope.row.metadata.labels"
									:key="index" size="small">
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
							<el-button link type="primary" size="small">资源配额与限制</el-button><el-divider
								direction="vertical" />
							<el-button v-auth="'k8s:namespace:edit'" link type="primary" size="small"
								@click="updateNamespace(scope.row)">编辑</el-button><el-divider direction="vertical" />
							<el-button v-auth="'k8s:namespace:del'" :disabled="scope.row.metadata.name === 'default'" link type="danger" size="small"
								@click="deleteNamespace(scope.row)">删除</el-button>
						</template>
					</el-table-column>
				</el-table>
				<!-- 分页区域 -->
				<pagination :total="data.total" @handlePageChange="handlePageChange"></pagination>
			</div>
		</el-card>
		<NamespaceDialog :title="data.title" v-model:visible="data.visible" :namespace="data.activeNamespace"
			@value-change="listNamespace()" v-if="data.visible" />
	</div>
</template>

<script setup lang="ts" name="k8sNamespace">
import { kubernetesInfo } from '/@/stores/kubernetes';
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import { Namespace } from 'kubernetes-types/core/v1';
import { dateStrFormat } from '/@/utils/formatTime';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { Delete, Edit } from '@element-plus/icons-vue';
import { useNamespaceApi } from '/@/api/kubernetes/namespace';
import { PageInfo } from '/@/types/kubernetes/common';
import { ElMessage, ElMessageBox } from 'element-plus';
import { number } from 'echarts';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));
const NamespaceDialog = defineAsyncComponent(() => import('./component/dialog.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
}
const namespaceApi = useNamespaceApi();
const k8sStore = kubernetesInfo();
const route = useRoute();
const data = reactive({
	activeNamespace: {} as Namespace,
	selectData: [] as Namespace[],
	total: 0,
	title: '',
	visible: false,
	loading: false,
	namespace: [] as Namespace[],
	query:<queryType> {
		page: 1,
		limit: 10,
		cloud: k8sStore.state.activeCluster,
	},
	type: "1",
	inputValue: "",
});
const listNamespace = () => {
	data.loading = true;
	namespaceApi.listNamespace(data.query).then((res) => {
		data.namespace = res.data.data;
		data.total = res.data.total;
	});
	data.loading = false;
};
const search = () => {
	if (data.type =='1') {
		data.query.name = data.inputValue
		delete data.query.label;
	} else if  (data.type == "0") {
		data.query.label = data.inputValue
		delete data.query.name;
	}
	if (data.inputValue === "") {
		delete data.query.label;
		delete data.query.name;
	}
	listNamespace();
};
const handleSelectionChange = (value: any) => {
	data.selectData = value;
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
const createNamespace = () => {
	data.title = '创建命名空间';
	data.visible = true;
};
const updateNamespace = (namespace: Namespace) => {
	data.title = '更新命名空间';
	data.activeNamespace = namespace;
	data.visible = true;
};
const handlePageChange = (pageInfo: PageInfo) => {
	data.query.page = pageInfo.page;
	data.query.limit = pageInfo.limit;
	listNamespace();
};
const deleteNamespace = (namespace: Namespace) => {
	ElMessageBox.confirm(`此操作将删除[ ${namespace.metadata!.name} ]命名空间 . 是否继续?`, '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			data.loading = true;
			namespaceApi
				.deleteNamespace({ cloud: k8sStore.state.activeCluster }, namespace.metadata!.name!)
				.then((res) => {
					listNamespace();
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
const deleteNamespaces = () => {
	ElMessage.error('危险操作,暂时不考虑~.~');
};

// const filterPod = () => {
// 	const nsList = [] as Namespace[];
// 	if (data.query.type === '1') {
// 		data.namespace.forEach((namespace: Namespace) => {
// 			if (namespace.metadata?.name?.includes(data.query.key)) {
// 				nsList.push(namespace);
// 			}
// 		});
// 	} else {
// 		data.namespace.forEach((namespace: Namespace) => {
// 			if (namespace.metadata?.labels) {
// 				for (let k in namespace.metadata.labels) {
// 					if (k.includes(data.query.key) || namespace.metadata.labels[k].includes(data.query.key)) {
// 						nsList.push(namespace);
// 						break;
// 					}
// 				}
// 			}
// 		});
// 	}
// 	data.namespace = nsList;
// 	data.total = data.namespace.length;
// };
onMounted(() => {
	listNamespace();
});
</script>
<style lang="scss" scoped>
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
