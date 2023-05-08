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
					<el-button type="danger" size="small" :icon="Delete" @click="drain(k8sStore.state.activeNode)">排水</el-button>
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
				<el-descriptions-item label="角色" label-align="right" align="center">
					<div v-if="k8sStore.state.activeNode.metadata?.labels!['kubernetes.io/role'] === 'master'">Master</div>
					<div v-else>Work</div>
				</el-descriptions-item>
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
					<el-button type="primary" size="small" plain :icon="Edit" @click="schedulable">设置</el-button>
				</el-descriptions-item>
				<el-descriptions-item label="污点" label-align="right" align="center">
					<el-tag v-for="item in k8sStore.state.activeNode?.spec?.taints"> {{ item.key }}:{{ item.value }}:{{ item.effect }} </el-tag>
				</el-descriptions-item>
				<el-descriptions-item label="状态" label-align="right" align="center">
					<el-text type="success" v-if="k8sStore.state.activeNode?.status!.conditions!.slice(-1)[0]['status'] == 'True'">Running</el-text>
					<el-text type="danger" v-else>故障</el-text>
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
				<el-tab-pane label="容器组" name="containers">
					<div>
						<el-table :data="data.pods" style="width: 100%" @selection-change="handleSelectionChange" v-loading="data.loading" max-height="300px">
							<el-table-column type="selection" width="55" />

							<el-table-column prop="metadata.name" label="名称" width="300px" show-overflow-tooltip>
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
							<el-table-column label="状态" width="200px">
								<template #default="scope">
									<p v-html="podStatus(scope.row.status)" />
								</template>
							</el-table-column>
							<el-table-column label="重启次数" width="100px">
								<template #default="scope">
									<div v-if="scope.row.status.containerStatuses">{{ podRestart(scope.row.status) }}</div>
								</template>
							</el-table-column>
							<el-table-column label="标签" width="180px">
								<template #default="scope">
									<el-tooltip placement="right" effect="light">
										<template #content>
											<div style="display: flex; flex-direction: column">
												<el-tag class="label" type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" size="small">
													{{ key }}:{{ item }}
												</el-tag>
											</div>
										</template>
										<el-tag type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" size="small">
											<div>{{ key }}:{{ item }}</div>
										</el-tag>
									</el-tooltip>
								</template>
							</el-table-column>

							<el-table-column prop="status.podIP" label="IP" width="220px">
								<template #default="scope">
									{{ scope.row.status.podIP }}
								</template>
							</el-table-column>
							<el-table-column prop="spec.nodeName" label="所在节点" width="220px">
								<template #default="scope">
									<div>{{ scope.row.spec.nodeName }}</div>
									<div>{{ scope.row.status.hostIP }}</div>
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
									<el-button link type="primary" size="small" @click="editPod(scope.row)">编辑</el-button><el-divider direction="vertical" />
									<el-button link type="primary" size="small" @click="deletePod(scope.row)">删除</el-button>
									<el-button link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button><el-divider direction="vertical" />
									<el-button link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
								</template>
							</el-table-column>
						</el-table>
						<!-- 分页区域 -->
						<Pagination :total="data.total" @handlePageChange="handlePageChange" />
					</div>
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

		<YamlDialog ref="yamlRef" :resourceType="'node'" />
	</div>
</template>
<script lang="ts" setup name="nodeDetail">
import { reactive, onMounted, ref, onBeforeUnmount, defineAsyncComponent, onUnmounted } from 'vue';
import { ArrowLeft, CaretBottom, View, Delete, Edit, RefreshRight } from '@element-plus/icons-vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1ContainerStatus, V1Pod, V1PodCondition, V1PodStatus, V1ReplicaSet, V1ReplicaSetCondition } from '@kubernetes/client-node';
import router from '/@/router';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox, TabsPaneContext } from 'element-plus';
import { usePodApi } from '/@/api/kubernetes/pod';
import { podInfo } from '/@/stores/pod';
import { ECharts, EChartsOption, init } from 'echarts';
import { deepClone } from '/@/utils/other';
import { useNodeApi } from '/@/api/kubernetes/node';
import { Node } from '/@/types/kubernetes/cluster';
import { PageInfo } from '/@/types/kubernetes/common';

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));
const MetaDetail = defineAsyncComponent(() => import('/@/components/kubernetes/metaDeail.vue'));
const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));

onMounted(() => {
	cpuUsage();
	memoryUsage();
	podUsage();
});

const nodeApi = useNodeApi();
const yamlRef = ref();
const route = useRoute();
const podStore = podInfo();
const k8sStore = kubernetesInfo();
const podApi = usePodApi();
const deploymentApi = useDeploymentApi();
const data = reactive({
	loading: false,
	param: {
		cloud: k8sStore.state.activeCluster,
	},
	replicasets: [] as V1ReplicaSet[],
	pods: [] as V1Pod[],
	total: 0,
	iShow: false,
	activeName: 'dashboard',
	deployment: [],
	events: [] as V1ReplicaSetCondition[],
	query: {
		cloud: k8sStore.state.activeCluster,
		page: 1,
		limit: 10,
	},
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
						name: 'CPU使用率 (总:' + k8sStore.state.activeNode.status!.capacity!.cpu + 'vCPU)',
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
							res = ` </div> <div>${c.state.waiting.reason}</div> <div style="font-size: 10px">${c.state.waiting.message}</div>`;
							// res = `${c.state.waiting.reason}`;
						}
						if (c.state?.terminated) {
							res = `${c.state.terminated.reason}`;
						}
					}
				});
				return (s = `<span style="color: red">${res}</span>`);
			}
		});
	} else {
		s = '<span style="color: green">ERROR</span>';
	}

	return s;
};
const handleClick = (tab: TabsPaneContext, event: Event) => {
	if (tab.paneName === 'six') {
		getEvents();
	} else if (tab.paneName === 'containers') {
		getPodsByNode();
	}
};

const getPodsByNode = () => {
	data.loading = true;
	nodeApi.listPodByNode(k8sStore.state.activeNode.metadata!.name!, data.query).then((res) => {
		if (res.code == 200) {
			data.pods = res.data.data;
			data.total = res.data.total;
		}
	});
	data.loading = false;
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
const handleSelectionChange = () => {};
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
const jumpPodDetail = (pod: V1Pod) => {
	podStore.state.podDetail = pod;
	router.push({
		name: 'podDetail',
	});
};
const deletePod = async (p: V1Pod) => {
	ElMessageBox.confirm(`此操作将删除[ ${p.metadata?.name} ] 容器 . 是否继续?`, '警告', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			podStore.deletePod(p);
			podStore.listPod();
			ElMessage({
				type: 'success',
				message: '${pod.metadata.name} 已删除',
			});
		})
		.catch(); // 取消
};
const handlePageChange = (pageInfo: PageInfo) => {
	data.loading = true;
	data.query.page = pageInfo.page;
	data.query.limit = pageInfo.limit;
	getPodsByNode();
	data.loading = false;
};
const editPod = (pod: V1Pod) => {
	delete pod.metadata?.managedFields;
	yamlRef.value.openDialog(pod);
};
const backRoute = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
	router.push({
		name: 'k8sNode',
	});
};
const getNode = () => {
	nodeApi.getNode(k8sStore.state.activeNode.metadata!.name!, { cloud: k8sStore.state.activeCluster }).then((res) => {
		if (res.code == 200) {
			k8sStore.state.activeNode.spec = res.data.spec;
			k8sStore.state.activeNode.metadata = res.data.metadata;
		}
	});
};
//设置是否可以调度
const schedulable = () => {
	const node = k8sStore.state.activeNode;
	let status = false;
	if (node.spec?.unschedulable === true) {
		status = false;
	} else {
		status = true;
	}

	ElMessageBox.confirm(`是否修改节点的调度状态为: ${status ? '不可调度' : '可调度'}`, 'Warning', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			nodeApi.schedulable({ cloud: k8sStore.state.activeCluster }, node.metadata!.name!, status).then((res) => {
				if (res.code === 200) {
					ElMessage.success(res.message);
					getNode();
				} else {
					ElMessage.error(res.message);
				}
			});
		})
		.catch(() => {
			ElMessage({
				type: 'info',
				message: '取消',
			});
		});
};
const drain = async (node: Node) => {
	ElMessageBox.confirm(`是否对[ ${node.metadata?.name} ]执行排水操作  . 是否继续?`, '警告', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			nodeApi.drainNode(node.metadata!.name!, { cloud: k8sStore.state.activeCluster }).then((res) => {
				if (res.code == 200) {
					ElMessage({
						type: 'success',
						message: '操作成功',
					});
					getNode();
				} else {
					ElMessage.error(res.message);
				}
			});
		})
		.catch(() => {
			ElMessage({
				type: 'info',
				message: '取消',
			});
		});
};

const showYaml = async () => {
	const node = deepClone(k8sStore.state.activeNode);
	delete node.usage;
	yamlRef.value.openDialog(node);
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
