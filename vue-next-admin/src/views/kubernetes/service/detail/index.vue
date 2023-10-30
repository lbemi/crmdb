<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<el-row :gutter="20">
				<el-col :span="18">
					<el-button type="info" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ k8sStore.state.activeService?.metadata?.name }}</span></el-col
				>
				<el-col :span="6"
					><el-button type="primary" size="small" :icon="Edit" @click="handleEdit()">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button>
					<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
						<el-icon>
							<ele-RefreshRight />
						</el-icon>
						刷新
					</el-button>
				</el-col>
			</el-row>
			<el-descriptions :column="3" border class="desc-body">
				<el-descriptions-item label="名称" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{
					k8sStore.state.activeService?.metadata?.name
				}}</el-descriptions-item>
				<el-descriptions-item label="命名空间" label-align="right" align="center">{{
					k8sStore.state.activeService?.metadata?.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center">{{
					dateStrFormat(k8sStore.state.activeService?.metadata?.creationTimestamp?.toString())
				}}</el-descriptions-item>
				<el-descriptions-item label="类型" label-align="right" align="center">{{ k8sStore.state.activeService.spec?.type }}</el-descriptions-item>
				<el-descriptions-item label="虚拟IP" label-align="right" align="center">{{
					k8sStore.state.activeService.spec?.clusterIP
				}}</el-descriptions-item>
				<el-descriptions-item label="外部IP" label-align="right" align="center">
					<span v-for="item in k8sStore.state.activeService.status?.loadBalancer?.ingress" :key="item.ip"> {{ item.ip }}</span>
				</el-descriptions-item>
				<el-descriptions-item label="端口" label-align="right" align="center">
					<el-tag class="label" size="small" effect="plain" v-for="item in k8sStore.state.activeService.spec?.ports" :key="item.port">
						<a v-if="k8sStore.state.activeService.spec?.type === 'NodePort'">节点端口:{{ item.nodePort }}</a> 服务端口:{{ item.port }}/{{
							item.protocol
						}}
						--> 容器端口:{{ item.targetPort }}
					</el-tag>
				</el-descriptions-item>
				<el-descriptions-item label="选择器" label-align="right" align="center">
					<div v-for="(key, value) in k8sStore.state.activeService.spec?.selector" :key="key" style="margin-bottom: 5px">
						<el-tag size="small" plain>{{ value }}: {{ key }}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="端点" label-align="right" align="center">
					<div v-for="item in data.serviceInfo.endPoints.subsets" :key="item.addresses">
						<div v-for="p in item.ports" :key="p.port">
							<div v-for="e in item.addresses" :key="e.ip">{{ e.ip }}:{{ p.port }}</div>
							<div v-for="e in item.notReadyAddresses" :key="e.ip">{{ e.ip }}:{{ p.port }}</div>
						</div>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="外部流量策略" label-align="right" align="center">
					{{ k8sStore.state.activeService.spec?.internalTrafficPolicy || 'Cluster' }}
					<el-tooltip
						content='<div style="margin-left: 12px"><ul><li>Local：流量只转发给本机的Pod</li><li >Cluster：流量可以转发到集群中其他节点上的Pod</li></ul>
            <a target="_blank" href="https://kubernetes.io/zh-cn/docs/tasks/access-application-cluster/create-external-load-balancer/">详细信息</a></div>'
						raw-content
						placement="right"
						effect="light"
					>
						<el-icon color="#909399"><InfoFilled /></el-icon>
					</el-tooltip>
				</el-descriptions-item>
				<el-descriptions-item label="会话亲和性" label-align="right" align="center">{{
					k8sStore.state.activeService.spec?.sessionAffinity
				}}</el-descriptions-item>
				<el-descriptions-item label="DNS" label-align="right" align="center"
					>{{ k8sStore.state.activeService.metadata?.name }}.{{ k8sStore.state.activeService.metadata?.namespace }}</el-descriptions-item
				>
			</el-descriptions>

			<el-tabs v-model="data.activeName" class="demo-tabs">
				<el-tab-pane label="工作负载" name="workload">
					<el-descriptions :column="1" direction="vertical">
						<el-descriptions-item label="无状态" v-if="data.serviceInfo.deployments">
							<el-card>
								<div v-if="data.serviceInfo.deployments">
									<el-table stripe style="width: 100%" :data="data.serviceInfo.deployments">
										<el-table-column prop="metadata.name" label="名称">
											<template #default="scope">
												<el-button link type="primary" @click="deployDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
											</template>
										</el-table-column>
										<el-table-column prop="spec.replicas" label="Pods" align="center">
											<template #header> <span>Pods</span><br /><span style="font-size: 10px; font-weight: 50">就绪/副本/失败</span> </template>

											<template #default="scope">
												<a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a
												>/ <a style="color: green">{{ scope.row.status.replicas || '0' }}</a
												>/
												<a style="color: red">{{ scope.row.status.unavailableReplicas || '0' }}</a>
											</template>
										</el-table-column>
										<el-table-column label="镜像" show-overflow-tooltip>
											<template #default="scope">
												<el-text truncated type="" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{
													item.image.split('@')[0]
												}}</el-text>
											</template>
										</el-table-column>

										<el-table-column label="标签">
											<template #default="scope">
												<el-tooltip placement="right" effect="light" v-if="scope.row.metadata.labels">
													<template #content>
														<div style="display: flex; flex-direction: column">
															<el-tag class="label" type="" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" effect="plain">
																{{ key }}:{{ item }}
															</el-tag>
														</div>
													</template>
													<el-tag class="label" type="" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" effect="plain">
														{{ key }}:{{ item }}
													</el-tag>
												</el-tooltip>
											</template>
										</el-table-column>
										<el-table-column label="创建时间">
											<template #default="scope">
												{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
											</template>
										</el-table-column>
									</el-table>
								</div>
							</el-card>
						</el-descriptions-item>
						<el-descriptions-item label="有状态" v-if="data.serviceInfo.statefulSets">
							<el-card>
								<div v-if="data.serviceInfo.statefulSets">
									<el-table stripe style="width: 100%" :data="data.serviceInfo.statefulSets">
										<el-table-column prop="metadata.name" label="名称">
											<template #default="scope">
												<el-button link type="primary" @click="deployDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
											</template>
										</el-table-column>
										<el-table-column prop="spec.replicas" label="Pods" align="center">
											<template #header> <span>Pods</span><br /><span style="font-size: 10px; font-weight: 50">就绪/副本/失败</span> </template>

											<template #default="scope">
												<a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a
												>/ <a style="color: green">{{ scope.row.status.replicas || '0' }}</a
												>/
												<a style="color: red">{{ scope.row.status.unavailableReplicas || '0' }}</a>
											</template>
										</el-table-column>
										<el-table-column label="镜像" show-overflow-tooltip>
											<template #default="scope">
												<el-text truncated type="" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{
													item.image.split('@')[0]
												}}</el-text>
											</template>
										</el-table-column>

										<el-table-column label="标签">
											<template #default="scope">
												<el-tooltip placement="right" effect="light" v-if="scope.row.metadata.labels">
													<template #content>
														<div style="display: flex; flex-direction: column">
															<el-tag class="label" type="" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" effect="plain">
																{{ key }}:{{ item }}
															</el-tag>
														</div>
													</template>
													<el-tag class="label" type="" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" effect="plain">
														{{ key }}:{{ item }}
													</el-tag>
												</el-tooltip>
											</template>
										</el-table-column>
										<el-table-column label="创建时间">
											<template #default="scope">
												{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
											</template>
										</el-table-column>
									</el-table>
								</div>
							</el-card>
						</el-descriptions-item>
					</el-descriptions>
				</el-tab-pane>
				<el-tab-pane label="容器组" name="first">
					<el-table
						v-if="data.serviceInfo.endPoints.subsets && data.serviceInfo.endPoints.subsets.length > 0"
						:data="data.serviceInfo.endPoints.subsets[0].addresses"
						stripe
						style="width: 100%"
						max-height="350px"
					>
						<el-table-column label="名称">
							<template #default="scope">
								<el-button v-if="scope.row.targetRef.name != undefined" link type="primary" @click="podDetail(scope.row.targetRef.name)">
									{{ scope.row.targetRef.name }}</el-button
								>
							</template>
						</el-table-column>
						<el-table-column prop="nodeName" label="所在节点" />
						<el-table-column prop="ip" label="IP" />
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="元数据" name="second">
					<MetaDetail :metaData="k8sStore.state.activeService.metadata" />
				</el-tab-pane>
				<el-tab-pane label="事件" name="fourth">
					<el-alert title="资源事件只保存最近1小时内发生的事件" :closable="false" type="info" class="mb15" show-icon />
					<el-table :data="data.serviceInfo.events">
						<el-table-column prop="type" label="类型" width="100px">
							<template #default="scope">
								<el-button link type="primary">{{ scope.row.type }}</el-button>
							</template>
						</el-table-column>
						<el-table-column label="原因">
							<template #default="scope">
								{{ scope.row.reason }}
							</template>
						</el-table-column>
						<el-table-column label="来源">
							<template #default="scope">
								{{ scope.row.source.component }}
							</template>
						</el-table-column>
						<el-table-column prop="spec.nodeName" label="消息">
							<template #default="scope">
								{{ scope.row.message }}
							</template>
						</el-table-column>
						<el-table-column label="时间" width="180px">
							<template #default="scope">
								{{ dateStrFormat(scope.row.firstTimestamp) }}
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
			</el-tabs>
		</el-card>

		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:code-data="data.codeData"
			:resourceType="'service'"
			@update="updateServiceYaml"
			v-if="data.dialogVisible"
		/>

		<ServiceDialog
			v-model:dialogVisible="data.update.dialogVisible"
			:service="k8sStore.state.activeService"
			:title="data.update.title"
			@refresh="k8sStore.refreshActiveService()"
			v-if="data.update.dialogVisible"
		/>
	</div>
</template>
<script lang="ts" setup name="k8sServiceDetail">
import { reactive, onMounted, defineAsyncComponent } from 'vue';
import { ArrowLeft, Edit, View, InfoFilled } from '@element-plus/icons-vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { Event, Pod, Service, Endpoints } from 'kubernetes-types/core/v1';
import { DaemonSet, Deployment, ReplicaSet, ReplicaSetCondition, StatefulSet } from 'kubernetes-types/apps/v1';
import router from '@/router';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage } from 'element-plus';
import { dateStrFormat } from '@/utils/formatTime';
import { useServiceApi } from '@/api/kubernetes/service';
import { usePodApi } from '@/api/kubernetes/pod';
import { podInfo } from '@/stores/pod';
import { deepClone } from '@/utils/other';

const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const MetaDetail = defineAsyncComponent(() => import('@/components/istio/kubernetes/metaDetail.vue'));
const ServiceDialog = defineAsyncComponent(() => import('../component/dialog.vue'));

const route = useRoute();
const serviceApi = useServiceApi();
const k8sStore = kubernetesInfo();
const podStore = podInfo();
const podApi = usePodApi();
const data = reactive({
	update: {
		title: '',
		dialogVisible: false,
	},
	serviceInfo: {
		deployments: [] as Deployment[],
		daemonSets: [] as DaemonSet[],
		statefulSets: [] as StatefulSet[],
		events: [] as Event[],
		endPoints: {} as Endpoints,
	},
	dialogVisible: false,
	codeData: {} as Service,
	param: {
		cloud: k8sStore.state.activeCluster,
	},
	replicasets: [] as ReplicaSet[],
	pods: [] as Pod[],
	iShow: false,
	activeName: 'workload',
	deployment: [],
	events: [] as ReplicaSetCondition[],
});

const deployDetail = async (dep: Deployment) => {
	k8sStore.state.activeDeployment = dep;
	await router.push({
		name: 'k8sDeploymentDetail',
		params: {
			name: dep.metadata?.name,
		},
	});
};

const podDetail = (name: string | undefined) => {
	if (name === undefined) return;
	podApi.getPod(k8sStore.state.activeService.metadata!.namespace!, name, { cloud: k8sStore.state.activeCluster }).then((res: any) => {
		if (res.code === 200) {
			podStore.state.podDetail = res.data;
			router.push({
				name: 'podDetail',
			});
		}
	});
};
const getServiceInfo = () => {
	serviceApi
		.listServiceWorkLoad(k8sStore.state.activeService.metadata!.namespace!, k8sStore.state.activeService.metadata!.name!, {
			cloud: k8sStore.state.activeCluster,
		})
		.then((res) => {
			data.serviceInfo = res.data;
		});
};

onMounted(() => {
	getServiceInfo();
});

const updateServiceYaml = async (svc: Service) => {
	const updateData = deepClone(svc) as Service;
	delete updateData.metadata?.managedFields;

	await serviceApi
		.updateService({ cloud: k8sStore.state.activeCluster }, updateData)
		.then(() => {
			// 同步更新store数据,刷新当前页面数据
			getService();
			ElMessage.success('更新成功');
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	data.dialogVisible = false;
};

const getService = () => {
	serviceApi
		.getService(k8sStore.state.activeService!.metadata!.namespace!, k8sStore.state.activeService!.metadata!.name!, {
			cloud: k8sStore.state.activeCluster,
		})
		.then((res) => {
			k8sStore.state.activeService = res.data;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
};
const backRoute = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
	router.push({
		name: 'k8sService',
	});
};

const handleEdit = () => {
	data.update.title = '更新service';
	data.update.dialogVisible = true;
};

const showYaml = async () => {
	delete k8sStore.state.activeService.metadata?.managedFields;
	data.codeData = k8sStore.state.activeService;
	data.dialogVisible = true;
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
</script>
<style lang="scss">
.card {
	overflow-y: auto; /* 开启滚动显示溢出内容 */
}
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
.tag-center {
	display: flex;
	flex-direction: column;
	align-items: center;

	.el-tag {
		margin-bottom: 3px;
		margin-bottom: 3px;
		white-space: normal;
		height: auto;
	}
}
.container {
	:deep(.el-card__body) {
		display: flex;
		flex-direction: column;
		flex: 1;
		overflow: auto;
		.el-table {
			flex: 1;
		}
	}
}
</style>
