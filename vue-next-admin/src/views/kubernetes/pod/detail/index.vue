<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<el-row :gutter="20">
				<el-col :span="18">
					<el-button type="info" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ podStore.state.podDetail?.metadata?.name }}</span></el-col
				>
				<el-col :span="6"
					><el-button type="primary" size="small" :icon="Edit">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button>
					<el-button type="primary" size="small" :icon="RefreshRight" @click="reDeploy">重新部署</el-button></el-col
				>
			</el-row>

			<el-descriptions :column="3" border>
				<el-descriptions-item label="名称" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">
					{{ podStore.state.podDetail?.metadata?.name }}
				</el-descriptions-item>
				<el-descriptions-item label="命名空间" label-align="right" align="center">{{
					podStore.state.podDetail?.metadata?.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center">{{
					dateStrFormat(podStore.state.podDetail?.metadata?.creationTimestamp)
				}}</el-descriptions-item>
				<el-descriptions-item label="所在节点及IP" label-align="right" align="center">
					<div>{{ podStore.state.podDetail?.spec?.nodeName }}</div>
					<div>{{ podStore.state.podDetail?.status?.hostIP }}</div>
				</el-descriptions-item>
				<!-- <el-descriptions-item label="选择器" label-align="right" align="center">
					<div class="tag-center">
						<el-tag effect="plain" round v-for="(item, key, index) in k8sStore.state.activeDeployment?.spec?.selector.matchLabels" :key="index">
							{{ key }}:{{ item }}
						</el-tag>
					</div>
				</el-descriptions-item> -->
				<el-descriptions-item label="镜像" label-align="right" align="center">
					<div class="tag-center">
						<el-tag round effect="plain" v-for="(item, index) in podStore.state.podDetail?.spec?.containers" :key="index">{{
							item.image?.split('@')[0]
						}}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="Pod IP" label-align="right" align="center">
					{{ podStore.state.podDetail?.status?.podIP }}
				</el-descriptions-item>
				<el-descriptions-item label="重启次数" label-align="right" align="center">
					<template #default="scope">
						<div v-if="podStore.state.podDetail?.status?.containerStatuses">{{ podRestart(podStore.state.podDetail?.status) }}</div>
					</template>
				</el-descriptions-item>
				<el-descriptions-item label="状态" label-align="right" align="center">
					<p v-html="podStatus(podStore.state.podDetail?.status!)" />
					<el-link type="primary" :underline="false" @click="data.iShow = !data.iShow" style="font-size: 10px; margin-left: 5px"
						>展开现状详情<el-icon> <CaretBottom /> </el-icon
					></el-link>
				</el-descriptions-item>
			</el-descriptions>

			<div v-show="data.iShow">
				<el-divider />
				<el-table :data="podStore.state.podDetail.status?.conditions" stripe style="width: 100%">
					<el-table-column prop="type" label="类型" />
					<el-table-column prop="status" label="状态">
						<template #default="scope">
							<el-tag v-if="scope.row.status === 'True'" type="success"> 正常</el-tag>
							<el-tag type="danger" v-else> 异常</el-tag>
						</template>
					</el-table-column>
					<el-table-column prop="lastTransitionTime" label="更新时间">
						<template #default="scope">
							{{ dateStrFormat(scope.row.lastTransitionTime) }}
						</template>
					</el-table-column>
					<el-table-column prop="reason" label="内容" />
					<el-table-column prop="message" label="消息" />
				</el-table>
			</div>

			<!-- <el-divider /> -->
			<el-tabs v-model="data.activeName" class="demo-tabs" @tab-click="handleClick">
				<el-tab-pane label="容器" name="first">
					<el-table :data="podStore.state.podDetail.status?.containerStatuses" stripe style="width: 100%" max-height="350px">
						<el-table-column label="名称">
							<template #default="scope">
								<el-text :type="scope.row.ready === true ? 'success' : 'danger'">{{ scope.row.name }}</el-text>
							</template>
						</el-table-column>
						<el-table-column label="状态">
							<template #default="scope">
								<el-text v-for="(item, key) in scope.row.state"> {{ key }}</el-text>
							</template>
						</el-table-column>
						<el-table-column label="镜像" prop="image" />
						<!-- <el-table-column label="命令" prop="command" />
						<el-table-column label="参数" prop="args" /> -->
						<el-table-column label="端口">
							<template #default="scope">
								<a v-for="item in scope.row.ports"> {{ item.name }}:{{ item.containerPort }}/{{ item.protocol }}</a>
								<!-- {{ dateStrFormat(scope.row.metadata.creationTimestamp) }} -->
							</template>
						</el-table-column>
						<el-table-column label="参数" prop="args" />

						<el-table-column label="创建时间" width="180px">
							<template #default="scope">
								<!-- {{ dateStrFormat(scope.row.metadata.creationTimestamp) }} -->
							</template>
						</el-table-column>
						<el-table-column fixed="right" label="操作" width="160">
							<template #default="scope">
								<el-button link type="primary" size="small">详情</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small">编辑</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="deletePod(scope.row)">删除</el-button>
								<el-button link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="元数据" name="second">
					<MetaDetail :metaData="podStore.state.podDetail!.metadata" />
				</el-tab-pane>
				<el-tab-pane label="环境变量" name="third">
					<el-descriptions :column="1" direction="vertical">
						<el-descriptions-item :label="'容器: ' + item.name" v-for="item in podStore.state.podDetail?.spec?.containers">
							<el-card class="card" :body-style="{ height: '200px' }">
								<div v-if="item.env" v-for="(value, key, index) in item.env" :key="index" style="margin-bottom: 8px">
									<el-tag type="info" size="default"> {{ value }} </el-tag>
								</div>
								<div v-else>无数据</div>
							</el-card>
						</el-descriptions-item>
					</el-descriptions>
				</el-tab-pane>
				<el-tab-pane label="历史版本" name="fourth">
					<el-table :data="data.replicasets" style="width: 100%">
						<el-table-column p label="版本">
							<template #default="scope">
								#{{ scope.row.metadata.annotations['deployment.kubernetes.io/revision'] }}
								<el-tag v-if="scope.row.status.replicas != 0" plain size="small" type="success" style="margin-left: 15px">当前版本</el-tag>
							</template>
						</el-table-column>
						<el-table-column label="镜像">
							<template #default="scope"> {{ scope.row.spec.template.spec.containers[0].image }} </template>
						</el-table-column>
						<el-table-column label="创建时间">
							<template #default="scope"> {{ dateStrFormat(scope.row.metadata.creationTimestamp) }} </template>
						</el-table-column>

						<el-table-column fixed="right" label="操作">
							<template #default="scope">
								<el-button link type="primary" size="small" @click="showRsYaml(scope.row)">详情</el-button>
								<el-button link type="primary" size="small" @click="rollBack(scope.row)" :disabled="data.replicasets.length == 1"
									>回滚到该版本</el-button
								>
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="监控" name="five">监控</el-tab-pane>
				<el-tab-pane label="事件" name="six">
					<el-alert title="资源事件只保存最近1小时内发生的事件" :closable="false" type="info" class="mb15" show-icon />
					<el-table :data="data.events">
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
		<YamlMegeDialog :code="code" :dialogVisible="dialogVisible" v-if="dialogVisible" />

		<YamlDialog ref="yamlRef" :resourceType="'pod'" :update-resource="updateDeployment" />
	</div>
</template>
<script lang="ts" setup name="k8sDeploymentDetail">
import { reactive, onMounted, ref, onBeforeUnmount, defineAsyncComponent } from 'vue';
import { ArrowLeft, CaretBottom, Edit, View, Minus, Plus, RefreshRight } from '@element-plus/icons-vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1ContainerStatus, V1Deployment, V1Pod, V1PodCondition, V1PodStatus, V1ReplicaSet, V1ReplicaSetCondition } from '@kubernetes/client-node';
import router from '/@/router';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox, TabsPaneContext } from 'element-plus';
import { usePodApi } from '/@/api/kubernetes/pod';
import { podInfo } from '/@/stores/pod';
import YAML from 'js-yaml';
import { deepClone } from '/@/utils/other';

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));
const YamlMegeDialog = defineAsyncComponent(() => import('/@/components/yaml/matchCode.vue'));
const MetaDetail = defineAsyncComponent(() => import('/@/components/kubernetes/metaDeail.vue'));

const yamlRef = ref();
const route = useRoute();
const podStore = podInfo();
const code = ref({});
const dialogVisible = ref(false);
const timer = ref();
const k8sStore = kubernetesInfo();
const podApi = usePodApi();
const deploymentApi = useDeploymentApi();
const data = reactive({
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

const rollBack = (rs: V1ReplicaSet) => {
	deploymentApi
		.rollBackDeployment(
			k8sStore.state.activeDeployment.metadata!.namespace!,
			k8sStore.state.activeDeployment.metadata!.name!,
			rs.metadata!.annotations!['deployment.kubernetes.io/revision'],
			{
				cloud: k8sStore.state.activeCluster,
			}
		)
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('回滚成功');
			} else {
				ElMessage.error('回滚失败,' + res.message);
			}
		})
		.catch((res) => {
			ElMessage.error('回滚失败,' + res.message);
		});
};
const handleClick = (tab: TabsPaneContext, event: Event) => {
	if (tab.paneName === 'six') {
		getEvents();
	}
};

const reDeploy = () => {
	const deployment = k8sStore.state.activeDeployment;
	deploymentApi
		.reDeployDeployment(deployment.metadata!.namespace!, deployment.metadata!.name!, { cloud: k8sStore.state.activeCluster })
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('操作成功');
			} else {
				ElMessage.error(res.message);
			}
		})
		.catch((res: any) => {
			ElMessage.error(res.message);
		});
};
const scaleDeploy = (action: string) => {
	if (action === 'plus') {
		k8sStore.state.activeDeployment.spec!.replicas!++;
	} else {
		k8sStore.state.activeDeployment.spec!.replicas!--;
	}
	deploymentApi
		.scaleDeployment(
			k8sStore.state.activeDeployment.metadata!.namespace!,
			k8sStore.state.activeDeployment.metadata!.name!,
			k8sStore.state.activeDeployment.spec?.replicas!,
			{ cloud: k8sStore.state.activeCluster }
		)
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('伸缩成功');
			} else {
				ElMessage.error('伸缩失败');
			}
		});
};

const handleEnvent = () => {
	data.replicasets.forEach((item: V1ReplicaSet) => {
		if (item.status) {
			if (item.status.conditions) {
				item.status.conditions.forEach((it: V1ReplicaSetCondition) => {
					data.events.push(it);
				});
			}
		}
	});
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
							res = `<div>${c.state.waiting.reason}</div>`;
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
		s = '<span style="color: red">ERROR</span>';
	}

	return s;
};
onMounted(() => {
	// getPods();
	// // buildWebsocket();
	// timer.value = window.setInterval(() => {
	// 	getPods();
	// }, 5000);
	// onBeforeUnmount(() => {
	// 	window.clearInterval(timer.value);
	// });
});
const updateDeployment = () => {
	const updateData = YAML.load(yamlRef.value.code) as V1Deployment;
	delete updateData.status;
	delete updateData.metadata?.managedFields;
	deploymentApi
		.updateDeployment(updateData, { cloud: k8sStore.state.activeCluster })
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
	yamlRef.value.handleClose();
};

const getPods = async () => {
	const res = await deploymentApi.detailDeployment(
		k8sStore.state.activeDeployment.metadata!.namespace!.toString(),
		k8sStore.state.activeDeployment?.metadata!.name!.toString(),
		data.param
	);
	data.pods = res.data.pods;
	data.replicasets = res.data.replicaSets;
};

const getEvents = async () => {
	const pod = podStore.state.podDetail;
	const res = await podApi.podEvents(pod.metadata!.namespace, pod.metadata!.name, data.param);
	data.events = res.data;
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
		name: 'k8sPod',
	});
};
const deletePod = async (pod: V1Pod) => {
	ElMessageBox.confirm(`此操作将删除[ ${pod.metadata?.name} ] 容器 . 是否继续?`, '警告', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			podApi.deletePod(pod.metadata?.namespace, pod.metadata?.name, data.param);
			getPods();
			ElMessage({
				type: 'success',
				message: `${pod.metadata?.name}` + ' 已删除',
			});
		})
		.catch(); // 取消
};
const showRsYaml = async (replicaSets: V1ReplicaSet) => {
	dialogVisible.value = true;
	const data = deepClone(replicaSets);
	delete data.metadata.managedFields;
	code.value = data;
};

const showYaml = async () => {
	delete podStore.state.podDetail.metadata?.managedFields;
	yamlRef.value.openDialog(podStore.state.podDetail);
};
// const buildWebsocket = () => {
// 	const ws = websocketApi.createWebsocket('deployment');

// 	ws.onmessage = (e) => {
// 		if (e.data === 'ping') {
// 			return;
// 		} else {
// 			const object = JSON.parse(e.data);
// 			if (
// 				object.type === 'deployment' &&
// 				object.result.namespace === k8sStore.state.activeDeployment?.metadata?.namespace &&
// 				object.cluster == k8sStore.state.activeCluster
// 			) {
// 				data.deployment = object.result.data;
// 				data.deployment.forEach((item: V1Deployment) => {
// 					if (item.metadata!.name == k8sStore.state.activeDeployment?.metadata?.name) {
// 						k8sStore.state.activeDeployment = item;
// 						return;
// 					}
// 				});
// 			}
// 		}
// 	};
// };
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
