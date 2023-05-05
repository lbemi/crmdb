<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<el-row :gutter="20">
				<el-col :span="16">
					<el-button type="primary" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ k8sStore.state.activeNode?.metadata?.name }}</span></el-col
				>
				<el-col :span="8"
					><el-button type="primary" size="small" :icon="Edit" @click="showYaml">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button>
					<el-button type="danger" size="small" :icon="Delete" @click="deletePod(podStore.state.podDetail)">删除</el-button>
					<el-button type="primary" size="small" @click="refreshCurrentTagsView">
						<el-icon>
							<ele-RefreshRight />
						</el-icon>
						刷新
					</el-button>
				</el-col>
			</el-row>

			<el-descriptions :column="3" border>
				<el-descriptions-item label="名称" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">
					{{ k8sStore.state.activeNode?.metadata?.name }}
				</el-descriptions-item>
				<el-descriptions-item label="角色" label-align="right" align="center">{{
					k8sStore.state.activeNode?.metadata?.labels!['kubernetes.io/role']
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center">{{
					dateStrFormat(k8sStore.state.activeNode?.metadata?.creationTimestamp)
				}}</el-descriptions-item>
				<el-descriptions-item label="IP地址" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.addresses![0].address }}
				</el-descriptions-item>
				<el-descriptions-item label="内核版本" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.nodeInfo?.kernelVersion }}
				</el-descriptions-item>
				<el-descriptions-item label="系统镜像" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.nodeInfo?.osImage }}
				</el-descriptions-item>
				<el-descriptions-item label="kubelet版本" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.nodeInfo?.kubeletVersion }}
				</el-descriptions-item>
				<el-descriptions-item label="kube-proxy版本" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.nodeInfo?.kubeProxyVersion }}
				</el-descriptions-item>
				<el-descriptions-item label="容器运行时" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.nodeInfo?.containerRuntimeVersion }}
				</el-descriptions-item>
				<el-descriptions-item label="系统架构" label-align="right" align="center">
					{{ k8sStore.state.activeNode?.status?.nodeInfo?.architecture }}
				</el-descriptions-item>
				<el-descriptions-item label="调度状态" label-align="right" align="center">
					<div v-if="k8sStore.state.activeNode?.spec?.unschedulable">
						<span style="color: red">不可调度</span>
					</div>
					<div v-else>可调度</div>
				</el-descriptions-item>
				<el-descriptions-item label="污点" label-align="right" align="center">
					<el-tag v-for="item in k8sStore.state.activeNode?.spec?.taints"> {{ item.key }}:{{ item.value }}:{{ item.effect }} </el-tag>
				</el-descriptions-item>
				<el-descriptions-item label="状态" label-align="right" align="center">
					<a v-html="podStatus(podStore.state.podDetail?.status!)" />
					<el-link type="primary" :underline="false" @click="data.iShow = !data.iShow" style="font-size: 10px; margin-left: 5px"
						>展开现状详情<el-icon> <CaretBottom /> </el-icon
					></el-link>
				</el-descriptions-item>
			</el-descriptions>

			<div v-show="data.iShow">
				<el-divider />
				<el-table :data="k8sStore.state.activeNode?.status?.conditions" stripe style="width: 100%">
					<el-table-column prop="type" label="类型" />
					<el-table-column prop="status" label="状态">
						<template #default="scope">
							<el-tag
								v-if="(scope.row.type != 'Ready' && scope.row.status === 'False') || (scope.row.type === 'Ready' && scope.row.status === 'True')"
								type="success"
							>
								正常</el-tag
							>
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
				<el-tab-pane label="dashboard" name="dashboard">
					<el-space wrap :size="20">
						<el-card id="echar-cpu" style="width: 300px; height: 300px" />
						<el-card id="echar-memory" style="width: 300px; height: 300px" />
						<el-card id="echar-pod" style="width: 300px; height: 300px" />
					</el-space>
				</el-tab-pane>
				<el-tab-pane label="容器组" name="first">
					<el-table :data="podStore.state.podDetail.status?.containerStatuses" stripe style="width: 100%" max-height="350px">
						<el-table-column label="名称">
							<template #default="scope">
								<el-text :type="scope.row.ready === true ? 'success' : 'danger'">{{ scope.row.name }}</el-text>
							</template>
						</el-table-column>
						<el-table-column label="状态">
							<template #default="scope">
								<el-text :type="scope.row.ready === true ? 'success' : 'danger'" v-for="(item, key) in scope.row.state">
									<div>
										{{ key }}
									</div>
									<div style="font-size: 10px">
										{{ item.message }}
									</div>
								</el-text>
							</template>
						</el-table-column>
						<el-table-column label="镜像" prop="image" />

						<el-table-column label="重启次数" prop="restartCount" />

						<el-table-column fixed="right" label="操作" width="160">
							<template #default="scope">
								<el-button link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button><el-divider direction="vertical" />
								<el-button link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="元数据" name="second">
					<MetaDetail :metaData="k8sStore.state.activeNode?.metadata" />
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

		<YamlDialog ref="yamlRef" :resourceType="'pod'" :update-resource="updatePod" />
	</div>
</template>
<script lang="ts" setup name="nodeDetail">
import { reactive, onMounted, ref, onBeforeUnmount, defineAsyncComponent, onUnmounted } from 'vue';
import { ArrowLeft, CaretBottom, Edit, View, Delete, Plus, RefreshRight } from '@element-plus/icons-vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1ContainerStatus, V1Deployment, V1Pod, V1PodCondition, V1PodStatus, V1ReplicaSet, V1ReplicaSetCondition } from '@kubernetes/client-node';
import router from '/@/router';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox, TabsPaneContext } from 'element-plus';
import { usePodApi } from '/@/api/kubernetes/pod';
import { podInfo } from '/@/stores/pod';
import { ECharts, EChartsOption, init } from 'echarts';

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));
const YamlMegeDialog = defineAsyncComponent(() => import('/@/components/yaml/matchCode.vue'));
const MetaDetail = defineAsyncComponent(() => import('/@/components/kubernetes/metaDeail.vue'));

onMounted(() => {
	cpuUsage();
	memoryUsage();
	podUsage();
});

const yamlRef = ref();
const route = useRoute();
const podStore = podInfo();
const code = ref({});
const dialogVisible = ref(false);
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
	activeName: 'dashboard',
	deployment: [],
	events: [] as V1ReplicaSetCondition[],
});

const cpuUsage = () => {
	const cpuDom = document.getElementById('echar-cpu') as HTMLDivElement;
	const cpuChar: ECharts = init(cpuDom);
	const option: EChartsOption = {
		textStyle: {
			fontSize: 3,
		},
		series: [
			{
				name: 'CPU使用率',
				type: 'gauge',
				axisLine: {
					lineStyle: {
						width: 10,
						color: [
							[0.8, '#529b2e'],
							[0.9, '#eebe77'],
							[1, '#fd666d'],
						],
					},
				},
				pointer: {
					itemStyle: {
						color: 'inherit',
					},
				},
				title: {
					show: true,
					fontSize: 15,
					offsetCenter: [0, '100%'],
				},
				axisLabel: {
					color: 'inherit',
					distance: 10,
					fontSize: 12,
				},
				detail: {
					valueAnimation: true,
					fontSize: 19,
					formatter: '{value}' + '%',
					color: 'inherit',
				},
				data: [
					{
						value: Math.round(k8sStore.state.activeNode.usage!.cpu * 100),
						name: 'CPU使用率 (总:' + k8sStore.state.activeNode.status?.capacity.cpu + 'vCPU)',
					},
				],
			},
		],
	};
	cpuChar.setOption(option);
};
const memoryUsage = () => {
	const cpuDom = document.getElementById('echar-memory') as HTMLDivElement;
	const cpuChar: ECharts = init(cpuDom);
	const option: EChartsOption = {
		// tooltip: {
		// 	formatter: '{b} : {c}%',
		// },
		series: [
			{
				name: '内存使用率',
				type: 'gauge',
				axisLine: {
					lineStyle: {
						width: 10,
						color: [
							[0.8, '#529b2e'],
							[0.9, '#eebe77'],
							[1, '#fd666d'],
						],
					},
				},
				title: {
					show: true,
					fontSize: 15,
					offsetCenter: [0, '100%'],
				},
				pointer: {
					itemStyle: {
						color: 'inherit',
					},
				},
				axisLabel: {
					color: 'inherit',
					distance: 10,
					fontSize: 12,
				},
				detail: {
					valueAnimation: true,
					fontSize: 19,
					formatter: '{value}' + '%',
					color: 'inherit',
				},
				data: [
					{
						value: Math.round(k8sStore.state.activeNode.usage?.memory! * 100),
						name: '内存使用率 (总:' + parseInt(k8sStore.state.activeNode.status!.capacity!.memory!.split('Ki')[0] / 1000000 + '') + 'GiB)',
					},
				],
			},
		],
	};
	cpuChar.setOption(option);
};
const podUsage = () => {
	const cpuDom = document.getElementById('echar-pod') as HTMLDivElement;
	const cpuChar: ECharts = init(cpuDom);
	const option: EChartsOption = {
		tooltip: {
			trigger: 'item',
			formatter: '{b} : {c}',
		},
		legend: {
			top: '5%',
			left: 'center',
		},
		series: [
			{
				name: 'Access From',
				type: 'pie',
				radius: ['50%', '70%'],
				avoidLabelOverlap: false,
				itemStyle: {
					borderRadius: 10,
					borderColor: '#fff',
					borderWidth: 2,
				},
				label: {
					show: false,
					position: 'center',
				},
				emphasis: {
					label: {
						show: false,
						fontSize: 10,
						fontWeight: 'bold',
					},
				},
				labelLine: {
					show: false,
				},
				data: [
					{ value: parseInt(k8sStore.state.activeNode.status!.allocatable!.pods, 10), name: '可分配pod数量' },
					{ value: k8sStore.state.activeNode.usage?.pod!, name: '已分配pod数量' },
				],
			},
		],
	};
	cpuChar.setOption(option);
};
const handleClick = (tab: TabsPaneContext, event: Event) => {
	if (tab.paneName === 'six') {
		getEvents();
	}
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
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

const updatePod = () => {
	// TODO 完善功能
	ElMessage.success('更新成功');
	// const updateData = YAML.load(yamlRef.value.code) as V1Deployment;
	// delete updateData.status;
	// delete updateData.metadata?.managedFields;
	// deploymentApi
	// 	.updateDeployment(updateData, { cloud: k8sStore.state.activeCluster })
	// 	.then((res) => {
	// 		if (res.code == 200) {
	// 			ElMessage.success('更新成功');
	// 		} else {
	// 			ElMessage.error(res.message);
	// 		}
	// 	})
	// 	.catch((e) => {
	// 		ElMessage.error(e.message);
	// 	});
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
	podStore.state.podShell = podStore.state.podDetail;
	router.push({
		name: 'podShell',
	});
};
const jumpPodLog = (p: V1Pod) => {
	podStore.state.podShell = podStore.state.podDetail;
	router.push({
		name: 'podLog',
	});
};
const backRoute = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
	router.push({
		name: 'k8sNode',
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
			backRoute();
		})
		.catch(); // 取消
};

const showYaml = async () => {
	delete podStore.state.podDetail.metadata?.managedFields;
	yamlRef.value.openDialog(podStore.state.podDetail);
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
