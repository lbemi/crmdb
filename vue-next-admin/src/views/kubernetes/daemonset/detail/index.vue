<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<el-row :gutter="20">
				<el-col :span="18">
					<el-button type="info" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ k8sStore.state.activeDaemonSet?.metadata?.name }}</span></el-col
				>
				<el-col :span="6"
					><el-button v-auth="'k8s:daemonset:edit'" type="primary" size="small" :icon="Edit" @click="handleEdit()">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button>
					<el-button v-auth="'k8s:daemonset:redeploy'" type="primary" size="small" :icon="Refresh" @click="reDeploy">重新部署</el-button>
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
					k8sStore.state.activeDaemonSet?.metadata?.name
				}}</el-descriptions-item>
				<el-descriptions-item label="命名空间" label-align="right" align="center">{{
					k8sStore.state.activeDaemonSet?.metadata?.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center">{{
					dateStrFormat(k8sStore.state.activeDaemonSet?.metadata?.creationTimestamp?.toString() || '')
				}}</el-descriptions-item>

				<el-descriptions-item label="镜像" label-align="right" align="center">
					<div class="tag-center">
						<el-tag round effect="plain" v-for="(item, index) in k8sStore.state.activeDaemonSet?.spec?.template?.spec?.containers" :key="index">{{
							item.image?.split('@')[0]
						}}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="滚动升级策略" label-align="right" align="center">
					<div>
						不可用Pod最大数量：
						{{ k8sStore.state.activeDaemonSet?.spec?.updateStrategy?.rollingUpdate?.maxUnavailable }}
					</div>
				</el-descriptions-item>

				<el-descriptions-item label="状态" label-align="right" align="center">
					就绪：<a v-if="k8sStore.state.activeDaemonSet?.status?.numberReady">{{ k8sStore.state.activeDaemonSet?.status?.numberReady }}</a>
					<a style="color: red" v-else>0</a> /{{ k8sStore.state.activeDaemonSet?.status?.numberAvailable }} 个，已更新：{{
						k8sStore.state.activeDaemonSet?.status?.updatedNumberScheduled
					}}
					个，可用：
					<a v-if="k8sStore.state.activeDaemonSet?.status?.numberReady">{{ k8sStore.state.activeDaemonSet?.status?.numberReady }}</a>
					<a style="color: red" v-else>0</a>

					个
				</el-descriptions-item>
			</el-descriptions>

			<!-- <el-divider /> -->
			<el-tabs v-model="data.activeName" class="demo-tabs" @tab-click="handleClick">
				<el-tab-pane label="容器组" name="first">
					<el-table :data="data.pods" stripe style="width: 100%" max-height="350px">
						<el-table-column prop="metadata.name" label="名称">
							<template #default="scope">
								<el-button link type="primary" @click="jumpPodDetail(scope.row)">{{ scope.row.metadata.name }}</el-button>
								<div v-if="scope.row.status.phase != 'Running'" style="color: red">
									<div v-for="(containerStatus, index) in scope.row.status.containerStatuses" :key="index">
										<div v-if="!containerStatus.ready">
											{{ containerStatus.state.waiting?.message }}
											{{ containerStatus.state.Terminating?.message }}
										</div>
									</div>
									<div v-for="(containerStatus, index) in scope.row.status.initContainerStatuses" :key="index">
										<div v-if="!containerStatus.ready">
											{{ containerStatus.state.waiting?.message }}
											{{ containerStatus.state.Terminating?.message }}
										</div>
									</div>
								</div>
							</template>
						</el-table-column>
						<el-table-column label="状态" width="180px">
							<template #default="scope">
								<p v-if="scope.row.status.phase" v-html="podStatus(scope.row.status)" />
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
											<el-tag size="small" class="label" type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index">
												{{ key }}:{{ item }}
											</el-tag>
										</div>
									</template>
									<el-tag size="small" type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index">
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
								<el-button v-auth="'k8s:pod:del'" link type="primary" size="small" @click="deletePod(scope.row)">删除</el-button>
								<el-button v-auth="'k8s:pod:shell'" link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button
								><el-divider direction="vertical" />
								<el-button v-auth="'k8s:pod:log'" link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="元数据" name="second">
					<MetaDetail :metaData="k8sStore.state.activeDaemonSet.metadata" />
				</el-tab-pane>
				<el-tab-pane label="环境变量" name="third">
					<el-descriptions :column="1" direction="vertical">
						<el-descriptions-item
							:label="'容器: ' + item.name"
							v-for="(item, index) in k8sStore.state.activeDaemonSet.spec?.template.spec?.containers"
							:key="index"
						>
							<el-card class="card" :body-style="{ height: '200px' }">
								<div v-if="item.env" style="margin-bottom: 8px">
									<el-tag type="info" size="default" v-for="(value, key, index) in item.env" :key="index"> {{ value }} </el-tag>
								</div>
								<div v-else>无数据</div>
							</el-card>
						</el-descriptions-item>
					</el-descriptions>
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
		<YamlDialog v-model:dialogVisible="data.RsdialogVisible" :code-data="data.rscode" :disabled-update="true" v-if="data.RsdialogVisible" />
		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:code-data="data.codeData"
			:resourceType="'daemonset'"
			@update="updateDaemonset"
			v-if="data.dialogVisible"
		/>

		<CreateDialog
			v-model:dialogVisible="data.update.dialogVisible"
			:title="data.update.title"
			:daemonset="data.daemonset"
			@refresh="refreshActiveDaemonset"
			v-if="data.update.dialogVisible"
		/>
	</div>
</template>
<script lang="ts" setup name="k8sDaemonsetDetail">
import { reactive, onMounted, ref, onBeforeUnmount, defineAsyncComponent, h } from 'vue';
import { ArrowLeft, Edit, View, Refresh } from '@element-plus/icons-vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useDaemonsetApi } from '@/api/kubernetes/daemonset';
import { ContainerStatus, Pod, PodCondition, PodStatus } from 'kubernetes-types/core/v1';
import { DaemonSet, ReplicaSet, ReplicaSetCondition } from 'kubernetes-types/apps/v1';
import router from '@/router';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox, TabsPaneContext } from 'element-plus';
import { usePodApi } from '@/api/kubernetes/pod';
import { useWebsocketApi } from '@/api/kubernetes/websocket';
import { podInfo } from '@/stores/pod';
import YAML from 'js-yaml';
import { deepClone } from '@/utils/other';
import { dateStrFormat } from '@/utils/formatTime';

const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const MetaDetail = defineAsyncComponent(() => import('@/components/kubernetes/metaDetail.vue'));
const CreateDialog = defineAsyncComponent(() => import('../component/dialog.vue'));

const route = useRoute();
const websocketApi = useWebsocketApi();
const podStore = podInfo();
const k8sStore = kubernetesInfo();
const { refreshActiveDaemonset } = kubernetesInfo();
const podApi = usePodApi();
const daemonsetApi = useDaemonsetApi();
const timer = ref();

const data = reactive({
	update: {
		dialogVisible: false,
		title: '',
	},

	dialogVisible: false,
	codeData: {} as DaemonSet,
	param: {
		cloud: k8sStore.state.activeCluster,
	},
	pods: [] as Pod[],
	iShow: false,
	activeName: 'first',
	daemonsets: [],
	daemonset: {} as DaemonSet,
	events: [] as ReplicaSetCondition[],
});

//编辑daemonset
const handleEdit = () => {
	k8sStore.state.creatDaemonSet.namespace = k8sStore.state.activeDaemonSet.metadata!.namespace!;
	const dep = deepClone(k8sStore.state.activeDaemonSet) as DaemonSet;
	delete dep.status;
	delete dep.metadata?.managedFields;
	data.daemonset = dep;
	data.update.title = '更新daemonset';
	data.update.dialogVisible = true;
};

const handleClick = (tab: TabsPaneContext) => {
	if (tab.paneName === 'six') {
		getEvents();
	}
};

const reDeploy = () => {
	const daemonset = k8sStore.state.activeDaemonSet;
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '重新部署 '),
			h('i', { style: 'color: teal' }, `${daemonset.metadata?.name}`),
			h('span', null, ' . 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			daemonsetApi
				.reDeployDaemonset(daemonset.metadata!.namespace!, daemonset.metadata!.name!, { cloud: k8sStore.state.activeCluster })
				.then(() => {
					ElMessage.success('操作成功');
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};

const podRestart = (status: PodStatus) => {
	let count = 0;
	status.containerStatuses!.forEach((item) => {
		count += item.restartCount;
	});
	return count;
};

const podStatus = (status: PodStatus) => {
	let s = '<span style="color: green">Running</span>';
	if (status.phase === 'Running') {
		status.conditions!.forEach((item: PodCondition) => {
			if (item.status != 'True') {
				let res = '';
				status.containerStatuses?.forEach((c: ContainerStatus) => {
					if (!c.ready) {
						if (c.state?.waiting) {
							res = ` </div> <div>${c.state.waiting.reason}</div> <div style="font-size: 10px">${c.state.waiting.message}</div>`;
							// res = `${c.state.waiting.reason}`;
						}
						if (c.state?.terminated) {
							res = `${c.state.terminated.reason}`;
							// res = 'Terminating';
						}
					}
				});
				return (s = `<span style="color: red">${res}</span>`);
			}
		});
	} else if (status.phase === 'Succeeded') {
		let res = '';
		status.containerStatuses?.forEach((c: ContainerStatus) => {
			if (!c.ready) {
				if (c.state?.terminated) {
					res = `${c.state.terminated.reason}`;
					// res = 'Terminating';
				}
			}
		});
		return (s = `<span style="color: #E6A23C">${res}</span>`);
	} else {
		let res = status.phase;
		status.containerStatuses?.forEach((c: ContainerStatus) => {
			if (!c.ready) {
				if (c.state?.waiting) {
					res = ` </div> <div>${c.state.waiting.reason}</div>`;
					// res = `${c.state.waiting.reason}`;
				}
				if (c.state?.terminated) {
					res = `${c.state.terminated.reason}`;
					// res = 'Terminating';
				}
			}
		});
		return (s = `<span style="color: red">${res}</span>`);
	}

	return s;
};
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
const updateDaemonset = async (codeData: any) => {
	const updateData = YAML.load(codeData) as DaemonSet;
	delete updateData.status;
	delete updateData.metadata?.managedFields;

	await daemonsetApi
		.updateDaemonset(updateData, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('更新成功');
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	data.dialogVisible = false;
};

const getPods = async () => {
	const res = await daemonsetApi.detailDaemonset(
		k8sStore.state.activeDaemonSet.metadata!.namespace!.toString(),
		k8sStore.state.activeDaemonSet?.metadata!.name!.toString(),
		data.param
	);
	data.pods = res.data.pods;
};

const getEvents = async () => {
	const res = await daemonsetApi.getDaemonsetEvents(
		k8sStore.state.activeDaemonSet.metadata!.namespace!.toString(),
		k8sStore.state.activeDaemonSet?.metadata!.name!.toString(),
		data.param
	);
	data.events = res.data;
};
const jumpPodDetail = (p: Pod) => {
	podStore.state.podDetail = p;
	router.push({
		name: 'podDetail',
	});
};
const jumpPodExec = (p: Pod) => {
	podStore.state.podShell = p;
	router.push({
		name: 'podShell',
	});
};
const jumpPodLog = (p: Pod) => {
	podStore.state.podShell = p;
	router.push({
		name: 'podLog',
	});
};
const backRoute = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
	router.push({
		name: 'k8sDaemonset',
	});
};
const deletePod = async (pod: Pod) => {
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
	delete k8sStore.state.activeDaemonSet.metadata?.managedFields;
	data.codeData = k8sStore.state.activeDaemonSet;
	data.dialogVisible = true;
};
const buildWebsocket = () => {
	const ws = websocketApi.createWebsocket('daemonset');

	ws.onmessage = (e) => {
		if (e.data === 'ping') {
			return;
		} else {
			const object = JSON.parse(e.data);
			if (
				object.type === 'daemonset' &&
				object.result.namespace === k8sStore.state.activeDaemonSet?.metadata?.namespace &&
				object.cluster == k8sStore.state.activeCluster
			) {
				data.daemonsets = object.result.data;
				data.daemonsets.forEach((item: DaemonSet) => {
					if (item.metadata!.name == k8sStore.state.activeDaemonSet?.metadata?.name) {
						k8sStore.state.activeDaemonSet = item;
						return;
					}
				});
			}
		}
	};
};

const refreshCurrentTagsView = () => {
	refreshActiveDaemonset();
};
</script>
<style lang="scss">
.card {
	overflow-y: auto;
	/* 开启滚动显示溢出内容 */
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

.label {
	margin-bottom: 3px;
}
</style>
