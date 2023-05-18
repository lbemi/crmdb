<template>
	<div>
		<el-table :data="data.pods" style="width: 100%" @selection-change="handleSelectionChange" v-loading="data.loading" max-height="100vh - 235px">
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
</template>

<script setup lang="ts">
import { defineAsyncComponent, onBeforeUnmount, onMounted, reactive, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import router from '/@/router';
import { podInfo } from '/@/stores/pod';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { V1ContainerStatus, V1Pod, V1PodCondition, V1PodStatus } from '@kubernetes/client-node';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';
import { PageInfo } from '/@/types/kubernetes/common';
import YAML from 'js-yaml';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));

const yamlRef = ref();
const k8sStore = kubernetesInfo();

const data = reactive({
	pods: [] as V1Pod[],
	query: {
		cloud: k8sStore.state.activeCluster,
		page: 1,
		limit: 10,
	},
	total: 0,
	loading: false,
	selectData: [],
});

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
const handleSelectionChange = () => {};
const handlePageChange = (pageInfo: PageInfo) => {
	data.query.page = pageInfo.page;
	data.query.limit = pageInfo.limit;
	data.loading = true;
	podStore.listPod();
	data.loading = false;
};
const editPod = (pod: V1Pod) => {
	delete pod.metadata?.managedFields;
	yamlRef.value.openDialog(pod);
};
const jumpPodExec = (p: V1Pod) => {
	data.podShell = p;
	router.push({
		name: 'podShell',
	});
};
const jumpPodDetail = (pod: V1Pod) => {
	data.podDetail = pod;
	router.push({
		name: 'podDetail',
	});
};
const jumpPodLog = (p: V1Pod) => {
	data.podShell = p;
	router.push({
		name: 'podLog',
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
</script>

<style scoped></style>
