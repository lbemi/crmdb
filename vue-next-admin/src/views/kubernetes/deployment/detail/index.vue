<template>
  <div class="layout-padding container" >
    <el-card shadow="hover" class="layout-padding-auto">
			<el-row :gutter="20">
				<el-col :span="20">
					<el-button type="info" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ k8sStore.state.activeDeployment?.metadata?.name }}</span></el-col
				>
				<el-col :span="4"
					><el-button type="primary" size="small" :icon="Edit">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button></el-col
				>
			</el-row>

			<el-descriptions :column="3" border>
				<el-descriptions-item label="名称" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{
					k8sStore.state.activeDeployment?.metadata?.name
				}}</el-descriptions-item>
				<el-descriptions-item label="命名空间" label-align="right" align="center">{{
					k8sStore.state.activeDeployment?.metadata?.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="副本数" label-align="right" align="center">{{
					k8sStore.state.activeDeployment?.spec?.replicas
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center">{{
					dateStrFormat(k8sStore.state.activeDeployment?.metadata.creationTimestamp)
				}}</el-descriptions-item>
				<el-descriptions-item label="选择器" label-align="right" align="center">
					<div class="tag-center">
						<el-tag effect="plain" round v-for="(item, key, index) in k8sStore.state.activeDeployment?.spec?.selector.matchLabels" :key="index">
							{{ key }}:{{ item }}
						</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="镜像" label-align="right" align="center">
					<div class="tag-center">
						<el-tag round effect="plain" v-for="(item, index) in k8sStore.state.activeDeployment?.spec?.template.spec.containers" :key="index">{{
							item.image.split('@')[0]
						}}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center">
					<div class="tag-center">
						<el-tag
							effect="plain"
							type="info"
							v-for="(item, key, index) in k8sStore.state.activeDeployment?.metadata?.annotations"
							:key="index"
							show-overflow-tooltip
						>
							{{ key }}:{{ item }}</el-tag
						>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="滚动升级策略" label-align="right" align="center">
					<div>
						超过期望的Pod数量：
						{{ k8sStore.state.activeDeployment?.spec.strategy.rollingUpdate.maxSurge }}
					</div>
					<div>
						不可用Pod最大数量：
						{{ k8sStore.state.activeDeployment?.spec.strategy.rollingUpdate.maxUnavailable }}
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="策略" label-align="right" align="center">{{
					k8sStore.state.activeDeployment?.spec?.strategy?.type
				}}</el-descriptions-item>
				<el-descriptions-item label="状态" label-align="right" align="center">
					就绪：<a v-if="k8sStore.state.activeDeployment?.status?.readyReplicas">{{ k8sStore.state.activeDeployment?.status?.readyReplicas }}</a>
					<a style="color: red" v-else>0</a> /{{ k8sStore.state.activeDeployment?.status?.replicas }} 个，已更新：{{
						k8sStore.state.activeDeployment?.status?.updatedReplicas
					}}
					个，可用：
					<a v-if="k8sStore.state.activeDeployment?.status?.readyReplicas">{{ k8sStore.state.activeDeployment?.status?.readyReplicas }}</a>
					<a style="color: red" v-else>0</a>

					个
					<el-link type="primary" :underline="false" @click="data.iShow = !data.iShow" style="font-size: 10px; margin-left: 5px"
						>展开现状现状详情<el-icon> <CaretBottom /> </el-icon
					></el-link>
				</el-descriptions-item>
			</el-descriptions>

			<div v-show="data.iShow">
				<el-divider />
				<el-table :data="k8sStore.state.activeDeployment?.status?.conditions" stripe style="width: 100%">
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

			<!-- <el-divider /> -->
			<el-tabs v-model="data.activeName" class="demo-tabs" @tab-click="handleClick">
				<el-tab-pane label="容器组" name="first">
					<el-table :data="data.pods" stripe style="width: 100%">
						<el-table-column prop="metadata.name" label="名称" width="300px">
							<template #default="scope">
								<el-button link type="primary">{{ scope.row.metadata.name }}</el-button>
								<div v-if="scope.row.status.phase != 'Running'" style="color: red">
									{{ scope.row.status.containerStatuses[0].state }}
								</div>
							</template>
						</el-table-column>
						<el-table-column label="状态" width="90px">
							<template #default="scope">
								<span v-if="scope.row.status.phase == 'Running'" style="color: green"> {{ scope.row.status.phase }}</span>
								<span v-else style="color: red"> {{ scope.row.status.phase }}</span>
							</template>
						</el-table-column>
						<el-table-column label="重启次数" width="90px">
							<template #default="scope">
								{{ scope.row.status.containerStatuses[0].restartCount }}
							</template>
						</el-table-column>

						<el-table-column prop="status.podIP" label="IP" width="150px">
							<template #default="scope">
								{{ scope.row.status.podIP }}
							</template>
						</el-table-column>
						<el-table-column prop="spec.nodeName" label="所在节点" width="120px">
							<template #default="scope">
								<div>{{ scope.row.spec.nodeName }}</div>
								<div>{{ scope.row.status.hostIP }}</div>
							</template>
						</el-table-column>
						<el-table-column label="标签" width="200px" show-overflow-tooltip>
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
								<el-button link type="primary" size="small" @click="handleClick">详情</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small">编辑</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="deletePod(scope.row)">删除</el-button>
								<el-button link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="Config" name="second">Config</el-tab-pane>
				<el-tab-pane label="Role" name="third">Role</el-tab-pane>
				<el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
			</el-tabs>
		</el-card>
    <YamlDialog ref="yamlRef" />
	</div>
</template>
<script lang="ts" setup name="k8sDeploymentDetail">
import {reactive, onMounted, ref, onBeforeUnmount, defineAsyncComponent} from 'vue';
import type { TabsPaneContext } from 'element-plus';
import { ArrowLeft, CaretBottom, Edit, View } from '@element-plus/icons-vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1Pod } from '@kubernetes/client-node';
import router from '/@/router';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus';
import { usePodApi } from '/@/api/kubernetes/pod';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';
import {podInfo} from "/@/stores/pod";

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'))

const yamlRef = ref()
const deploymentApi = useDeploymentApi();
const route = useRoute();
const websocketApi = useWebsocketApi();
const podStore = podInfo()

const handleClick = (tab: TabsPaneContext, event: Event) => {
	console.log(tab, event);
};
const k8sStore = kubernetesInfo();
const podApi = usePodApi();
const data = reactive({
	param: {
		cloud: k8sStore.state.activeCluster,
	},
	pods: [] as V1Pod[],
	iShow: false,
	activeName: 'first',
	deployment: [],
});
const timer = ref();
onMounted(() => {
	getPods();
	buildWebsocket();
	timer.value = window.setInterval(() => {
		getPods();
	}, 5000);
	onBeforeUnmount(() => {
		window.clearInterval(timer.value);
	});
});

const getPods = async () => {
	const res = await deploymentApi.detailDeployment(
		k8sStore.state.activeDeployment.metadata?.namespace?.toString(),
		k8sStore.state.activeDeployment?.metadata?.name?.toString(),
		data.param
	);
	data.pods = res.data;
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
		name: 'k8sDeployment',
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

const showYaml =async () =>{
  yamlRef.value.openDialog(k8sStore.state.activeDeployment)
}
const buildWebsocket = async () => {
	const ws = await websocketApi.createWebsocket('deployment');

	ws.onmessage = (e) => {
		if (e.data === 'ping') {
			return;
		} else {
			const object = JSON.parse(e.data);
			if (
				object.type === 'deployment' &&
				object.result.namespace === k8sStore.state.activeDeployment?.metadata?.namespace &&
				object.cluster == k8sStore.state.activeCluster
			) {
				data.deployment = object.result.data;
				data.deployment.forEach((item) => {
					if (item.metadata.name == k8sStore.state.activeDeployment?.metadata?.name) {
						k8sStore.state.activeDeployment = item;
						return;
					}
				});
			}
		}
	};
};
</script>
<style lang="scss">
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
