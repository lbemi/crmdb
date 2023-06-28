/** * Created by lei on 2023/03/21 */
<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-text class="mx-1" size="small">命名空间：</el-text>
				<el-select
					v-model="k8sStore.state.activeNamespace"
					style="max-width: 180px"
					size="small"
					class="m-2"
					placeholder="Select"
					@change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata!.name!" />
				</el-select>
				<el-input
					v-model="podStore.state.inputValue"
					placeholder="输入标签或者名称"
					size="small"
					clearable
					@change="search"
					style="width: 250px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="podStore.state.type" placeholder="输入标签或者名称" style="width: 60px" size="small">
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
				<el-button type="danger" size="small" class="ml10" :disabled="podStore.state.selectData.length == 0" @click="deletePods">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="podStore.state.pods"
				style="width: 100%"
				@selection-change="handleSelectionChange"
				v-loading="podStore.state.loading"
				max-height="100vh - 235px"
			>
				<el-table-column type="selection" width="55" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column prop="metadata.name" label="名称" width="300px" show-overflow-tooltip>
					<template #default="scope">
						<el-button link type="primary" @click="jumpPodDetail(scope.row)">{{ scope.row.metadata.name }}</el-button>
						<!-- <div v-if="scope.row.status.phase != 'Running'" style="color: red">
							<div v-if="scope.row.status.containerStatuses">
								{{ scope.row.status.containerStatuses[0].state }}
							</div>
							<div v-else>{{ scope.row.status.conditions[0].reason }}:{{ scope.row.status.conditions[0].message }}</div>
						</div> -->
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
			<Pagination :total="podStore.state.total" @handlePageChange="handlePageChange" />
		</el-card>
		<YamlDialog ref="yamlRef" :resourceType="'pod'" :update-resource="updatePod" />
	</div>
</template>

<script setup lang="ts" name="k8sPod">
import { defineAsyncComponent, h, onBeforeUnmount, onMounted, reactive, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import router from '@/router';
import { podInfo } from '@/stores/pod';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ContainerStatus, Pod, PodCondition, PodStatus } from 'kubernetes-types/core/v1';
import { useWebsocketApi } from '@/api/kubernetes/websocket';
import { PageInfo } from '@/types/kubernetes/common';
import YAML from 'js-yaml';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { deepClone } from '@/utils/other';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));

const route = useRoute();
const yamlRef = ref();
const k8sStore = kubernetesInfo();
const podStore = podInfo();
const websocketApi = useWebsocketApi();

const search = () => {
	podStore.listPod();
};

const handleSelectionChange = (value: any) => {
	podStore.state.selectData = value;
};

const handleChange = () => {
	podStore.state.loading = true;
	podStore.listPod();
	podStore.state.loading = false;
};

const podRestart = (status: PodStatus) => {
	let count = 0;
	status.containerStatuses!.forEach((item) => {
		count += item.restartCount;
	});
	return count;
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
// FIXME

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
const handlePageChange = (pageInfo: PageInfo) => {
	podStore.state.query.page = pageInfo.page;
	podStore.state.query.limit = pageInfo.limit;
	podStore.state.loading = true;
	podStore.listPod();
	podStore.state.loading = false;
};

const updatePod = () => {
	const updateData = deepClone(YAML.load(yamlRef.value.code) as Pod);
	delete updateData.status;
	yamlRef.value.handleClose();
};
const editPod = (pod: Pod) => {
	delete pod.metadata?.managedFields;
	yamlRef.value.openDialog(pod);
};
const jumpPodExec = (p: Pod) => {
	podStore.state.podShell = p;
	router.push({
		name: 'podShell',
	});
};
const jumpPodDetail = (pod: Pod) => {
	podStore.state.podDetail = pod;
	router.push({
		name: 'podDetail',
	});
};
const jumpPodLog = (p: Pod) => {
	podStore.state.podShell = p;
	router.push({
		name: 'podLog',
	});
};

const deletePods = async () => {
	podStore.state.loading = true;
	podStore.state.selectData.forEach((pod: Pod) => {
		podStore.deletePod(pod);
	});
	podStore.state.loading = false;
};
const deletePod = async (p: Pod) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${p.metadata?.name}`),
			h('span', null, ' 容器. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			podStore
				.deletePod(p)
				.then((res) => {
					ElMessage({
						type: 'success',
						message: `${p.metadata?.name} 已删`,
					});
				})
				.catch((e) => {
					ElMessage.error(e.message);
				});
			podStore.listPod();
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};

const filterPod = (pods: Array<Pod>) => {
	const podList = [] as Pod[];
	if (podStore.state.type === '1') {
		pods.forEach((pod: Pod) => {
			if (pod.metadata?.name?.includes(podStore.state.inputValue)) {
				podList.push(pod);
			}
		});
	} else {
		pods.forEach((pod: Pod) => {
			if (pod.metadata?.labels) {
				for (let k in pod.metadata.labels) {
					if (k.includes(podStore.state.inputValue) || pod.metadata.labels[k].includes(podStore.state.inputValue)) {
						podList.push(pod);
						break;
					}
				}
			}
		});
	}
	podStore.state.pods = podList;
};
const ws = websocketApi.createWebsocket('pod');
ws.onmessage = (e) => {
	if (e.data === 'ping') {
		return;
	} else {
		const object = JSON.parse(e.data);
		if (object.type === 'pod' && object.result.namespace === k8sStore.state.activeNamespace && object.cluster == k8sStore.state.activeCluster) {
			// podStore.state.pods = object.result.data;
			filterPod(object.result.data);
		}
	}
};

onMounted(() => {
	podStore.state.loading = true;
	podStore.listPod();
	podStore.state.loading = false;
});
onBeforeUnmount(() => {
	ws.close();
});
</script>

<style scoped lang="scss">
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
.ellipsis {
	height: 60px;
	white-space: nowrap;
	overflow: hidden;
	overflow-y: auto;
	// text-overflow: ellipsis;
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
