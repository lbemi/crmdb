<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto layout-padding-view">
			<div class="mb15">
				<el-text class="mx-1" :size="theme.themeConfig.globalComponentSize">命名空间：</el-text>
				<el-select
					v-model="k8sStore.state.activeNamespace"
					style="max-width: 280px"
					class="m-2"
					placeholder="Select"
					size="small"
					@change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata!.name!" />
				</el-select>
				<el-input
					v-model="data.inputValue"
					placeholder="输入标签或者名称"
					size="small"
					clearable
					@change="search"
					style="max-width: 300px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="data.type" placeholder="输入标签或者名称" style="max-width: 120px" size="small">
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
				<el-button type="primary" size="small" class="ml10" @click="createVirtualService" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="data.virtualServices"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="data.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button size="small" type="primary" text @click="virtualServiceDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
					</template>
				</el-table-column>
				<el-table-column label="标签">
					<template #default="scope">
						<el-tooltip placement="right" effect="light" v-if="scope.row.metadata.labels">
							<template #content>
								<div style="display: flex; flex-direction: column">
									<el-tag
										class="label"
										effect="plain"
										type="info"
										v-for="(item, key, index) in scope.row.metadata.labels"
										:key="index"
										:size="theme.themeConfig.globalComponentSize"
									>
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<el-icon><List /></el-icon>
						</el-tooltip>
						<a v-else>无</a>
					</template>
				</el-table-column>

				<el-table-column label="创建时间">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>

				<el-table-column fixed="right" label="操作" width="260px" flex>
					<template #default="scope">
						<el-button link type="primary" size="small" @click="virtualServiceDetail(scope.row)">详情</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="updateVirtualService(scope.row)">编辑</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button><el-divider direction="vertical" />
						<el-button :disabled="scope.row.metadata.name === 'kubernetes'" link type="danger" size="small" @click="deleteVirtualService(scope.row)"
							>删除</el-button
						>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<pagination :total="data.total" @handlePageChange="handlePageChange"></pagination>
		</el-card>
		<YamlDialog v-model:dialogVisible="data.dialogVisible" :code-data="data.codeData" @update="updateVirtualServiceYaml" v-if="data.dialogVisible" />
		<DrawDialog
			v-model:visible="data.draw.visible"
			:virtualService="data.draw.virtualService"
			:title="data.draw.title"
			@refresh="listVirtualService"
			v-if="data.draw.visible"
		/>
		<VirtualServiceDetail
			v-model:visible="data.detail.visible"
			:virtualService="data.detail.virtualService"
			:title="data.detail.title"
			v-if="data.detail.visible"
		/>
	</div>
</template>

<script setup lang="ts" name="virtualservice">
import { VirtualService } from '@kubernetes-models/istio/networking.istio.io/v1beta1/VirtualService';
import { defineAsyncComponent, h, onBeforeUnmount, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete, List } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { useVirtualServiceApi } from '@/api/istio/virtualService';
import { deepClone } from '@/utils/other';
import { useWebsocketApi } from '@/api/kubernetes/websocket';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const DrawDialog = defineAsyncComponent(() => import('./component/draw.vue'));
const VirtualServiceDetail = defineAsyncComponent(() => import('./component/detail.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const VirtualServiceApi = useVirtualServiceApi();
const route = useRoute();
const theme = useThemeConfig();
const socketApi = useWebsocketApi();

//定义数据
const data = reactive({
	detail: {
		title: '',
		visible: false,
		virtualService: {} as VirtualService,
	},
	draw: {
		title: '',
		visible: false,
		virtualService: {} as VirtualService,
	},
	dialogVisible: false,
	codeData: {} as VirtualService,
	loading: false,
	selectData: [] as VirtualService[],
	virtualServices: [] as VirtualService[],
	tmpVirtualService: [] as VirtualService[],
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
	listVirtualService();
});

const ws = socketApi.createWebsocket('virtualService');
ws.onmessage = (e: any) => {
	if (e.data === 'ping') {
		return;
	} else {
		const object: WebsocketResult = JSON.parse(e.data);
		if (
			object.type === 'virtualService' &&
			object.result.namespace === k8sStore.state.activeNamespace &&
			object.cluster == k8sStore.state.activeCluster
		) {
			data.virtualServices = object.result.data as VirtualService[];
		}
	}
};
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

	listVirtualService();
};
const handleChange = () => {
	listVirtualService();
};
// const filterVirtualService = (virtualServices: Array<VirtualService>) => {
// 	const virtualServiceList = [] as VirtualService[];
// 	if (data.type === '1') {
// 		virtualServices.forEach((virtualService: VirtualService) => {
// 			if (virtualService.metadata?.name?.includes(data.query.key)) {
// 				virtualServiceList.push(virtualService);
// 			}
// 		});
// 	} else {
// 		virtualServices.forEach((virtualService: VirtualService) => {
// 			if (virtualService.metadata?.labels) {
// 				for (let k in virtualService.metadata.labels) {
// 					if (k.includes(data.query.key) || virtualService.metadata.labels[k].includes(data.query.key)) {
// 						virtualServiceList.push(virtualService);
// 						break;
// 					}
// 				}
// 			}
// 		});
// 	}
// 	data.virtualServices = virtualServiceList;
// };
const createVirtualService = () => {
	data.draw.title = '创建虚拟服务';
	data.draw.visible = true;
};
const deleteVirtualService = (virtualService: VirtualService) => {
	data.loading = true;
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${virtualService.metadata!.name}`),
			h('span', null, ' 虚拟服务. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			VirtualServiceApi.deleteVirtualService(virtualService.metadata!.namespace!, virtualService.metadata!.name!, {
				cloud: k8sStore.state.activeCluster,
			})
				.then((res: any) => {
					listVirtualService();
					ElMessage.success(res.message);
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
	data.loading = false;
};
const virtualServiceDetail = (virtualService: VirtualService) => {
	data.detail.title = '详情';
	data.detail.virtualService = virtualService;
	data.detail.visible = true;
};

const showYaml = (VirtualService: VirtualService) => {
	data.dialogVisible = true;
	delete VirtualService.metadata?.managedFields;
	data.codeData = VirtualService;
};
const updateVirtualServiceYaml = (code: any) => {
	console.log('更新VirtualService', code);
};

const handleSelectionChange = () => {};
const updateVirtualService = (virtualService: VirtualService) => {
	data.draw.title = '编辑';
	data.draw.virtualService = deepClone(virtualService) as VirtualService;
	data.draw.visible = true;
};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listVirtualService();
};
const listVirtualService = () => {
	data.loading = true;
	VirtualServiceApi.listVirtualService(k8sStore.state.activeNamespace, data.query)
		.then((res: any) => {
			data.virtualServices = res.data.data;
			data.tmpVirtualService = res.data.data;
			data.total = res.data.total;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	data.loading = false;
};

const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};

onBeforeUnmount(() => {
	ws.close();
});
</script>

<style scoped lang="scss">
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
