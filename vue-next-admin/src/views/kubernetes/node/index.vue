<template>
	<div class="layout-pd">
		<el-card shadow="hover">
			<el-table :data="data.nodes" style="width: 100%" max-height="100vh - 200px">
				<el-table-column label="名称/IP地址/UID" width="140" align="center" show-overflow-tooltip>
					<template #default="scope">
						<div>{{ scope.row.metadata.name }}</div>
						<el-button link type="primary">{{ scope.row.status.addresses[0].address }}</el-button>
						<div>{{ scope.row.metadata.uid }}</div>
					</template>
				</el-table-column>
				<el-table-column label="角色/状态" width="120" align="center">
					<template #default="scope">
						<div v-if="scope.row.metadata.labels['kubernetes.io/role']">Master</div>
						<div v-else>Worker</div>
						<div
							v-if="scope.row.status.conditions.slice(-1)[0]['status'] == 'True'"
							style="display: flex; align-items: center; justify-content: center"
						>
							<div style="width: 12px; height: 12px; background: #67c23a; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #67c23a">运行中 </span>
							<el-tooltip placement="right" effect="light">
								<template #content>
									<div
										v-for="(item, index) in scope.row.status.conditions"
										:key="index"
										style="display: flex; justify-content: space-between; width: 180px"
									>
										<span>{{ item.type }}</span>
										<span> {{ item.status }}</span>
									</div>
								</template>
								<el-icon size="17px" color="#909399" style="margin-left: 3px"><InfoFilled /></el-icon>
							</el-tooltip>
						</div>
						<div v-else>
							<div style="display: inline-block; width: 12px; height: 12px; background: #f56c6c; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #f56c6c">故障 </span>
							<el-tooltip content="Right center" placement="right" effect="light">
								<el-icon size="17px" color="#909399" style="margin-left: 3px"><InfoFilled /></el-icon>
							</el-tooltip>
						</div>
						<div v-if="scope.row.spec.unschedulable">
							<span style="color: red">不可调度</span>
						</div>
						<div v-else>可调度</div>
					</template>
				</el-table-column>
				<el-table-column width="250" align="center">
					<template #header> <span>配置</span><br /><span style="font-size: 10px; font-weight: 50">(系统版本/内核版本)</span> </template>
					<template #default="scope">
						<div>
							{{ scope.row.status.capacity.cpu }}vCPU
							{{ parseInt(scope.row.status.capacity.memory.split('Ki')[0] / 1000000 + '') + 'GiB' }}
						</div>
						<span>{{ scope.row.status.nodeInfo.osImage }}</span>
						<div>{{ scope.row.status.nodeInfo.kernelVersion }}</div>
					</template>
				</el-table-column>
				<el-table-column width="110" align="center">
					<template #header> <span>容器组</span><br /><span style="font-size: 10px; font-weight: 50">(已分配/总量)</span> </template>
					<template #default="scope">
						<div>
							{{ scope.row.usage.pod }}/
							{{ scope.row.status.capacity.pods }}
						</div>
					</template>
				</el-table-column>
				<el-table-column width="80" align="center">
					<template #header> <span>CPU</span><br /><span style="font-size: 10px; font-weight: 50">(使用率)</span> </template>
					<template #default="scope">
						<div>{{ Math.round(scope.row.usage.cpu * 100) }}%</div>
					</template>
				</el-table-column>
				<el-table-column width="120" align="center">
					<template #header> <span>内存</span><br /><span style="font-size: 10px; font-weight: 50">(使用率)</span> </template>
					<template #default="scope">
						<div>{{ Math.round(scope.row.usage.memory * 100) }}%</div>
					</template>
				</el-table-column>

				<el-table-column width="180" align="center">
					<template #header> <span>Kubelet版本</span><br /><span style="font-size: 10px; font-weight: 50">(Runtime版本/系统类型)</span> </template>
					<template #default="scope">
						<div>
							{{ scope.row.status.nodeInfo.kubeletVersion }}
						</div>
						<span>{{ scope.row.status.nodeInfo.containerRuntimeVersion }}</span>
						<div>{{ scope.row.status.nodeInfo.operatingSystem }}/{{ scope.row.status.nodeInfo.architecture }}</div>
					</template>
				</el-table-column>
				<el-table-column label="标签" width="60" align="center">
					<template #default="scope">
						<el-tooltip placement="right" effect="light">
							<template #content>
								<div style="display: flex; flex-direction: column">
									<el-tag class="label" type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index">
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<SvgIcon iconName="icon-biaoqian" className="icon-1-4em" />
						</el-tooltip>
					</template>
				</el-table-column>
				<el-table-column label="污点" width="280" align="center">
					<template #default="scope">
						<el-tag type="success" v-for="(item, index) in scope.row.spec.taints" :key="index"> {{ item.key }}: {{ item.effect }} </el-tag>
					</template>
				</el-table-column>
				<el-table-column label="创建时间" width="180" align="center">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>

				<el-table-column fixed="right" label="操作" width="130">
					<template #default="scope">
						<div style="display: flex; align-items: center">
							<el-button link type="primary" size="default">详情</el-button>
							<el-divider direction="vertical" />
							<el-dropdown>
								<span class="el-dropdown-link">
									更多<el-icon class="el-icon--right"><CaretBottom /></el-icon>
								</span>
								<template #dropdown>
									<el-dropdown-menu>
										<el-dropdown-item @click="updateLabels(scope.row)">修改标签</el-dropdown-item>
										<el-dropdown-item>添加污点</el-dropdown-item>
										<el-dropdown-item>是否可调度</el-dropdown-item>
										<el-dropdown-item>删除节点</el-dropdown-item>
									</el-dropdown-menu>
								</template>
							</el-dropdown>
						</div>
					</template>
				</el-table-column>
			</el-table>
			<Pagination :total="data.total" @handlePageChange="handlePageChange" />
		</el-card>
		<Labels
			v-model:visible="data.visible"
			:title="data.title"
			:data="data.data"
			:cloud="clusterStore.activeCluster"
			vi-if="data.visible"
			@valuechange="listNodes()"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sNode">
import { reactive, onMounted, defineAsyncComponent } from 'vue';
import { InfoFilled, CaretBottom } from '@element-plus/icons-vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { V1Node } from '@kubernetes/client-node';
import { useNodeApi } from '/@/api/kubernetes/node';
import { PageInfo } from '/@/types/kubernetes/common';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));
const Labels = defineAsyncComponent(() => import('./component/labels.vue'));

const nodeApi = useNodeApi();
const data = reactive({
	query: {
		cloud: '',
		page: 1,
		limit: 10,
	},
	nodes: [] as V1Node[],
	total: 1,
	visible: false,
	title: '修改标签',
	data: {} as Node,
});
const clusterStore = kubernetesInfo();

onMounted(() => {
	listNodes();
});
const listNodes = () => {
	data.query.cloud = clusterStore.activeCluster;
	nodeApi
		.listNode(data.query)
		.then((result) => {
			data.nodes = result.data.data;
			data.total = result.data.total;
		})
		.catch((err) => {
			console.log(err);
		});
};

const handlePageChange = (pageInfo: PageInfo) => {
	data.query.page = pageInfo.page;
	data.query.limit = pageInfo.limit;
	listNodes();
};
const updateLabels = (node: Node) => {
	data.visible = true;
	data.data = node;
};
</script>

<style scoped lang="scss">
//// .tooltip-base-box {
////   width: 600px;
//// }
//.el-popper is-dark {
//	display: flex;
//	align-items: center;
//	flex-direction: column;
//	justify-content: space-between;
//}
//.label {
//	margin-top: 3px;
//}
//.example-showcase .el-dropdown + .el-dropdown {
//	margin-left: 15px;
//}
//.example-showcase .el-dropdown-link {
//	cursor: pointer;
//	color: var(--el-color-primary);
//	display: flex;
//	align-items: center;
//}
</style>