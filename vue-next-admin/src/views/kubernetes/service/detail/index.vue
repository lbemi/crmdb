<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<el-row :gutter="20">
				<el-col :span="18">
					<el-button type="info" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ k8sStore.state.activeService?.metadata?.name }}</span></el-col
				>
				<el-col :span="6"
					><el-button type="primary" size="small" :icon="Edit">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button>
					<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
						<el-icon>
							<ele-RefreshRight />
						</el-icon>
						刷新
					</el-button>
				</el-col>
			</el-row>
			<el-descriptions :column="3" border>
				<el-descriptions-item label="名称" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{
					k8sStore.state.activeService?.metadata?.name
				}}</el-descriptions-item>
				<el-descriptions-item label="命名空间" label-align="right" align="center">{{
					k8sStore.state.activeService?.metadata?.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center">{{
					dateStrFormat(k8sStore.state.activeService?.metadata?.creationTimestamp?.toString()!)
				}}</el-descriptions-item>
				<el-descriptions-item label="类型" label-align="right" align="center">{{ k8sStore.state.activeService.spec?.type }}</el-descriptions-item>
				<el-descriptions-item label="虚拟IP" label-align="right" align="center">{{
					k8sStore.state.activeService.spec?.clusterIP
				}}</el-descriptions-item>
				<el-descriptions-item label="外部IP" label-align="right" align="center">
					<span v-for="item in k8sStore.state.activeService.status?.loadBalancer?.ingress"> {{ item.ip }}</span>
				</el-descriptions-item>
				<el-descriptions-item label="端口" label-align="right" align="center">
					<el-tag
						class="label"
						size="small"
						effect="plain"
						v-if="k8sStore.state.activeService.spec?.ports"
						v-for="item in k8sStore.state.activeService.spec?.ports"
					>
						<a v-if="k8sStore.state.activeService.spec?.type === 'NodePort'">节点端口:{{ item.nodePort }}</a> 服务端口:{{ item.port }}/{{
							item.protocol
						}}
						--> 容器端口:{{ item.targetPort }}
					</el-tag>
				</el-descriptions-item>
				<el-descriptions-item label="选择器" label-align="right" align="center">
					<div v-for="(key, value) in k8sStore.state.activeService.spec?.selector" style="margin-bottom: 5px">
						<el-tag plain>{{ key }}: {{ value }}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="端点" label-align="right" align="center">
					<div v-for="(key, value) in k8sStore.state.activeService.spec?.selector">{{ key }}: {{ value }}</div>
				</el-descriptions-item>
				<el-descriptions-item label="外部流量策略" label-align="right" align="center">
					{{ k8sStore.state.activeService.spec?.internalTrafficPolicy }}
					<el-tooltip
						content='<div style="margin-left: 12px"><ul><li>Local：流量只转发给本机的Pod</li><li >Cluster：流量可以转发到集群中其他节点上的Pod</li></ul></div>'
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

			<div v-show="data.iShow">
				<el-divider />
				<el-table :data="k8sStore.state.activeService?.status?.conditions" stripe style="width: 100%">
					<el-table-column prop="type" label="类型" />
					<el-table-column prop="status" label="状态" />
					<el-table-column prop="lastUpdateTime" label="更新时间">
						<template #default="scope">
							{{ dateStrFormat(scope.row.lastUpdateTime) }}
						</template>
					</el-table-column>
					<el-table-column prop="reason" label="内容" />
					<el-table-column prop="message" label="消息" />
				</el-table>
			</div>

			<el-tabs v-model="data.activeName" class="demo-tabs" @tab-click="handleClick">
				<el-tab-pane label="容器组" name="first">
					<el-table :data="data.pods" stripe style="width: 100%" max-height="350px">
						<el-table-column prop="metadata.name" label="名称">
							<template #default="scope">
								<el-button link type="primary" @click="jumpPodDetail(scope.row)">{{ scope.row.metadata.name }}</el-button>
								<div v-if="scope.row.status.phase != 'Running'" style="color: red">
									<div v-if="scope.row.status.containerStatuses">
										{{ scope.row.status.containerStatuses[0].state }}
									</div>
									<div v-else>{{ scope.row.status.conditions[0].reason }}:{{ scope.row.status.conditions[0].message }}</div>
								</div>
							</template>
						</el-table-column>
						<el-table-column label="状态" width="180px">
							<template #default="scope">
								<p v-html="podStatus(scope.row.status)" />
							</template>
						</el-table-column>
						<el-table-column label="重启次数">
							<template #default="scope">
								<div v-if="scope.row.status.containerStatuses">{{ podRestart(scope.row.status) }}</div>
							</template>
						</el-table-column>

						<el-table-column prop="status.podIP" label="IP">
							<template #default="scope">
								{{ scope.row.status.podIP }}
							</template>
						</el-table-column>
						<el-table-column prop="spec.nodeName" label="所在节点">
							<template #default="scope">
								<div>{{ scope.row.spec.nodeName }}</div>
								<div>{{ scope.row.status.hostIP }}</div>
							</template>
						</el-table-column>
						<el-table-column label="标签" show-overflow-tooltip>
							<template #default="scope">
								<el-tooltip placement="top" effect="light">
									<template #content>
										<div style="display: flex; flex-direction: column">
											<el-tag class="label" type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index">
												{{ key }}:{{ item }}
											</el-tag>
										</div>
									</template>
									<el-tag type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index">
										<div>{{ key }}:{{ item }}</div>
									</el-tag>
								</el-tooltip>
							</template>
						</el-table-column>
						<el-table-column label="创建时间" width="180px">
							<template #default="scope">
								{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
							</template>
						</el-table-column>
						<el-table-column fixed="right" label="操作" width="160">
							<template #default="scope">
								<el-button link type="primary" size="small" @click="jumpPodDetail(scope.row)">详情</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small">编辑</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="deletePod(scope.row)">删除</el-button>
								<el-button link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="元数据" name="second">
					<MetaDetail :metaData="k8sStore.state.activeService.metadata" />
				</el-tab-pane>
				<el-tab-pane label="事件" name="fourth"> </el-tab-pane>
			</el-tabs>
		</el-card>

		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:code-data="data.codeData"
			:resourceType="'service'"
			@update="updateServiceYaml"
			v-if="data.dialogVisible"
		/>
	</div>
</template>
<script lang="ts" setup name="k8sServiceDetail">
import { reactive, onMounted, ref, onBeforeUnmount, defineAsyncComponent, h } from 'vue';
import { ArrowLeft, CaretBottom, Edit, View, Refresh, InfoFilled, Right } from '@element-plus/icons-vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1ContainerStatus, V1Pod, V1PodCondition, V1PodStatus, V1ReplicaSet, V1ReplicaSetCondition, V1Service } from '@kubernetes/client-node';
import router from '/@/router';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox, TabsPaneContext } from 'element-plus';
import { usePodApi } from '/@/api/kubernetes/pod';
import { podInfo } from '/@/stores/pod';
import YAML from 'js-yaml';
import { deepClone } from '/@/utils/other';
import { dateStrFormat } from '/@/utils/formatTime';
import { useServiceApi } from '/@/api/kubernetes/service';

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));
const MetaDetail = defineAsyncComponent(() => import('/@/components/kubernetes/metaDeail.vue'));

const route = useRoute();
const servieApi = useServiceApi();
const podStore = podInfo();
const k8sStore = kubernetesInfo();
const podApi = usePodApi();
const deploymentApi = useDeploymentApi();
const timer = ref();

const data = reactive({
	dialogVisible: false,
	codeData: {} as V1Service,
	param: {
		cloud: k8sStore.state.activeCluster,
	},
	replicasets: [] as V1ReplicaSet[],
	pods: [] as V1Pod[],
	iShow: false,
	activeName: 'first',
	deployment: [],
	events: [] as V1ReplicaSetCondition[],
});

const handleClick = (tab: TabsPaneContext, event: Event) => {
	if (tab.paneName === 'six') {
		getEvents();
	}
};

const podRestart = (status: V1PodStatus) => {
	let count = 0;
	status.containerStatuses!.forEach((item) => {
		count += item.restartCount;
	});
	return count;
};

// FIXME
const podStatus = (status: V1PodStatus) => {
	let s = '<span style="color: green">Running</span>';
	if (status.phase === 'Running') {
		status.conditions!.forEach((item: V1PodCondition) => {
			if (item.status != 'True') {
				let res = '';
				status.containerStatuses?.forEach((c: V1ContainerStatus) => {
					if (!c.ready) {
						if (c.state?.waiting) {
							res = `<div>${c.state.waiting.reason}</div> <div style="font-size: 10px">${c.state.waiting.message}</div>`;
							// res = `${c.state.waiting.reason}`;
						}
						if (c.state?.terminated) {
							res = `${c.state.terminated.reason}`;
						}
					}
				});
				return (s = `<span style="color: red">${res}</span>`);
			}
			// s = '<span style="color: green">true</span>';
		});
	} else {
		s = '<span style="color: green">ERROR</span>';
	}

	return s;
};
onMounted(() => {
	// getPods();
	// timer.value = window.setInterval(() => {
	// 	getPods();
	// }, 5000);
	// onBeforeUnmount(() => {
	// 	window.clearInterval(timer.value);
	// });
});

const getPods = async () => {
	const res = await deploymentApi.detailDeployment(
		k8sStore.state.activeService.metadata!.namespace!.toString(),
		k8sStore.state.activeService?.metadata!.name!.toString(),
		data.param
	);
	data.pods = res.data.pods;
	data.replicasets = res.data.replicaSets;
};
const updateServiceYaml = async (svc: any) => {
	const updateData = YAML.load(svc) as V1Service;
	delete updateData.status;
	delete updateData.metadata?.managedFields;

	await servieApi
		.updateService({ cloud: k8sStore.state.activeCluster }, updateData)
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('更新成功');
			} else {
				ElMessage.error(res.message);
			}
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	data.dialogVisible = false;
};
const getEvents = async () => {
	const res = await deploymentApi.getDeploymentEvents(
		k8sStore.state.activeService.metadata!.namespace!.toString(),
		k8sStore.state.activeService?.metadata!.name!.toString(),
		data.param
	);
	data.events = res.data;
};
const jumpPodDetail = (p: V1Pod) => {
	podStore.state.podDetail = p;
	router.push({
		name: 'podDetail',
	});
};
const jumpPodExec = (p: V1Pod) => {
	podStore.state.podShell = p;
	router.push({
		name: 'podShell',
	});
};
const jumpPodLog = (p: V1Pod) => {
	podStore.state.podShell = p;
	router.push({
		name: 'podLog',
	});
};
const backRoute = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
	router.push({
		name: 'k8sService',
	});
};
const deletePod = async (pod: V1Pod) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${pod.metadata!.name}`),
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
			podApi.deletePod(pod.metadata?.namespace, pod.metadata?.name, data.param);
			getPods();
			ElMessage({
				type: 'success',
				message: `${pod.metadata?.name}` + ' 已删除',
			});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
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
