<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto layout-padding-view">
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
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" :icon="Delete" @click="deleteServices"
					>批量删除</el-button
				>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
				<el-table
					:data="data.services"
					@selection-change="handleSelectionChange"
					size="small"
					style="width: 100%"
					max-height="100vh - 235px"
					v-loading="data.loading"
				>
					<el-table-column type="selection" width="35" />
					<el-table-column prop="metadata.name" label="名称">
						<template #default="scope">
							<el-button link type="primary" size="small" @click="serviceDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
						</template>
					</el-table-column>
					<el-table-column prop="metadata.namespace" label="命名空间" />
					<el-table-column prop="spec.type" label="类型" />
					<el-table-column prop="spec.clusterIP" label="集群IP" />

					<el-table-column label="端口" style="display: flex" align="center" width="220px">
						<template #header>
							<span>端口</span><br /><span style="font-size: 10px; font-weight: 50">(nodePort:port/protocol->targetPort)</span>
						</template>

						<template #default="scope">
							<div v-if="scope.row.spec.ports" v-for="item in scope.row.spec.ports">
								<el-tag class="label" size="small" effect="plain">
									<a v-if="scope.row.spec.type === 'NodePort'"> {{ item.nodePort }}:</a>{{ item.port }}/{{ item.protocol }}->{{
										item.targetPort
									}}</el-tag
								>
							</div>
						</template>
					</el-table-column>
					<el-table-column label="外部访问IP">
						<template #default="scope">
							<el-link
								v-if="scope.row.status.loadBalancer.ingress"
								v-for="item in scope.row.status.loadBalancer.ingress"
								target="_blank"
								type="primary"
								:href="'http://' + item.ip"
								>{{ item.ip }}</el-link
							>
						</template>
					</el-table-column>
					<el-table-column label="标签" width="70px">
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
											size="small"
										>
											{{ key }}:{{ item }}
										</el-tag>
									</div>
								</template>
								<el-icon><List /></el-icon>
							</el-tooltip>
						</template>
					</el-table-column>

					<el-table-column label="创建时间" width="170px">
						<template #default="scope">
							{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
						</template>
					</el-table-column>

					<el-table-column fixed="right" label="操作">
						<template #default="scope">
							<el-button link type="primary" size="small" @click="serviceDetail(scope.row)">详情</el-button><el-divider direction="vertical" />
							<el-button link type="primary" size="small" @click="updateService(scope.row)">编辑</el-button><el-divider direction="vertical" />
							<el-dropdown size="small">
								<span class="el-dropdown-link" style="font-size: 12px">
									更多<el-icon class="el-icon--right"><CaretBottom /></el-icon>
								</span>
								<template #dropdown>
									<el-dropdown-menu>
										<el-dropdown-item @click="showYaml(scope.row)">查看Yaml</el-dropdown-item>
										<el-dropdown-item>日志</el-dropdown-item>
										<el-dropdown-item
											:disabled="scope.row.metadata.name === 'kubernetes'"
											link
											type="danger"
											size="small"
											@click="deleteService(scope.row)"
											>删除</el-dropdown-item
										>
									</el-dropdown-menu>
								</template>
							</el-dropdown>
						</template>
					</el-table-column>
				</el-table>
				<!-- 分页区域 -->
				<pagination :total="data.total" @handlePageChange="handlePageChange"></pagination>
			</div>
		</el-card>
		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:code-data="data.codeData"
			:resourceType="'service'"
			@update="updateServiceYaml"
			v-if="data.dialogVisible"
		/>
		<ServiceDialog
			v-model:dialogVisible="data.serviceDialog"
			:service="data.activeService"
			:title="data.title"
			@refresh="listService()"
			v-if="data.serviceDialog"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sService">
import { Service } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, h, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useServiceApi } from '@/api/kubernetes/service';
import { ResponseType } from '@/types/response';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete, List, CaretBottom } from '@element-plus/icons-vue';
import { deepClone } from '@/utils/other';
import YAML from 'js-yaml';
import router from '@/router';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const ServiceDialog = defineAsyncComponent(() => import('./component/dialog.vue'));

const k8sStore = kubernetesInfo();
const servieApi = useServiceApi();
const route = useRoute();

const data = reactive({
	title: '',
	editDialog: false,
	activeService: {},
	serviceDialog: false,
	dialogVisible: false,
	codeData: {} as Service,
	loading: false,
	selectData: [] as Service[],
	services: [] as Service[],
	tmpService: [] as Service[],
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
const search = () => {
	filterService(data.tmpService);
};
const handleChange = () => {
	listService();
};

const serviceDetail = (service: Service) => {
	k8sStore.state.activeService = service;
	router.push({
		name: 'k8sServiceDetail',
	});
};

const showYaml = async (service: Service) => {
	const svc = deepClone(service);
	delete svc.metadata?.managedFields;
	data.codeData = svc;
	data.dialogVisible = true;
};

const filterService = (services: Array<Service>) => {
	const serviceList = [] as Service[];
	if (data.query.type === '1') {
		services.forEach((service: Service) => {
			if (service.metadata?.name?.includes(data.query.key)) {
				serviceList.push(service);
			}
		});
	} else {
		services.forEach((service: Service) => {
			if (service.metadata?.labels) {
				for (let k in service.metadata.labels) {
					if (k.includes(data.query.key) || service.metadata.labels[k].includes(data.query.key)) {
						serviceList.push(service);
						break;
					}
				}
			}
		});
	}
	data.services = serviceList;
};
const createService = () => {
	data.title = '创建service';
	data.activeService = {};
	data.serviceDialog = true;
};

const deleteServices = () => {
	data.selectData.forEach(async (service: Service) => {
		await servieApi.deleteService({ cloud: k8sStore.state.activeCluster }, service.metadata!.name!, service.metadata!.namespace!);
	});
	setTimeout(() => {
		listService();
	}, 100);
};

const deleteService = (service: Service) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${service.metadata!.name}`),
			h('span', null, ' 服务. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			servieApi
				.deleteService({ cloud: k8sStore.state.activeCluster }, service.metadata!.name!, service.metadata!.namespace!)
				.then((res) => {
					listService();
					ElMessage.success(res.message);
				})
				.catch((e) => {
					ElMessage.error(e);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};
const handleSelectionChange = (value: any) => {
	data.selectData = value;
};
const updateService = (service: Service) => {
	data.activeService = service;
	data.title = '更新< ' + service.metadata?.name + ' >服务';
	data.serviceDialog = true;
};

const updateServiceYaml = async (svc: any) => {
	const updateData = deepClone(YAML.load(svc) as Service);
	delete updateData.status;
	delete updateData.metadata?.managedFields;

	await servieApi
		.updateService({ cloud: k8sStore.state.activeCluster }, updateData)
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('更新成功');
				listService();
			} else {
				ElMessage.error(res.message);
			}
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	data.dialogVisible = false;
};

const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listService();
};

const listService = () => {
	data.loading = true;
	servieApi
		.listService(k8sStore.state.activeNamespace, data.query)
		.then((res: ResponseType) => {
			if (res.code === 200) {
				data.services = res.data.data;
				data.total = res.data.total;
				data.tmpService = res.data.data;
			}
		})
		.catch((e) => {
			ElMessage.error(e);
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
