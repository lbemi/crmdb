/** * Created by lei on 2023/03/21 */
<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				命名空间:
				<el-select
					v-model="k8sStore.state.activeNamespace"
					style="max-width: 180px"
					size="default"
					class="m-2"
					placeholder="Select"
					@change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata?.name" />
				</el-select>
				<el-button type="danger" size="default" class="ml10" :disabled="podStore.state.selectData.length == 0">批量删除</el-button>
			</div>
			<el-table
				:data="podStore.state.pods"
				style="width: 100%"
				@selection-change="handleSelectionChange"
				v-loading="podStore.state.loading"
				max-height="100vh - 235px"
			>
				<el-table-column type="selection" width="55" />

				<el-table-column prop="metadata.name" label="名称" width="300px" show-overflow-tooltip>
					<template #default="scope">
						<el-button link type="primary">{{ scope.row.metadata.name }}</el-button>
						<div v-if="scope.row.status.phase != 'Running'" style="color: red">
							<div v-if="scope.row.status.containerStatuses">
								{{ scope.row.status.containerStatuses[0].state }}
							</div>
							<div v-else>{{ scope.row.status.conditions[0].reason }}:{{ scope.row.status.conditions[0].message }}</div>
						</div>
					</template>
				</el-table-column>
				<el-table-column label="状态" width="90px">
					<template #default="scope">
						<span v-if="scope.row.status.phase == 'Running'" style="color: green"> {{ scope.row.status.phase }}</span>
						<span v-else style="color: red"> {{ scope.row.status.phase }}</span>
					</template>
				</el-table-column>
				<el-table-column label="重启次数" width="100px">
					<template #default="scope">
						<div v-if="scope.row.status.containerStatuses">
							{{ scope.row.status.containerStatuses[0].restartCount }}
						</div>
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
						<el-button link type="primary" size="small">详情</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small">编辑</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="deletePod(scope.row)">删除</el-button>
						<el-button link type="primary" size="small" @click="jumpPodExec(scope.row)">终端</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="jumpPodLog(scope.row)">日志</el-button>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<Pagination :total="podStore.state.total" @handlePageChange="handlePageChange" />
		</el-card>
	</div>
</template>

<script setup lang="ts" name="k8sPod">
import { defineAsyncComponent, onBeforeUnmount, onMounted } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import router from '/@/router';
import { podInfo } from '/@/stores/pod';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { V1Pod } from '@kubernetes/client-node';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';
import { PageInfo } from '/@/types/kubernetes/common';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));

const k8sStore = kubernetesInfo();
const podStore = podInfo();
const websocketApi = useWebsocketApi();

const handleSelectionChange = () => {};

const handleChange = () => {
	podStore.state.loading = true;
	podStore.listPod();
	podStore.state.loading = false;
};

const handlePageChange = (pageInfo: PageInfo) => {
	podStore.state.query.page = pageInfo.page;
	podStore.state.query.limit = pageInfo.limit;
	podStore.state.loading = true;
	podStore.listPod();
	podStore.state.loading = false;
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

const ws = websocketApi.createWebsocket('pod');
ws.onmessage = (e) => {
	if (e.data === 'ping') {
		return;
	} else {
		const object = JSON.parse(e.data);
		if (object.type === 'pod' && object.result.namespace === k8sStore.state.activeNamespace && object.cluster == k8sStore.state.activeCluster) {
			podStore.state.pods = object.result.data;
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
