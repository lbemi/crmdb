<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-input size="default" placeholder="请输入集群名称" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button @click="loadCluster">ceshi </el-button>
				<el-button size="default" type="success" class="ml10" @click="loadCluster()"> 导入集群 </el-button>
			</div>

			<el-table :data="data.clusters" style="width: 100%">
				<el-table-column prop="id" label="ID" width="100" />
				<el-table-column prop="name" label="Name" width="120">
					<template #default="scope">
						<el-button link type="primary" @click="handleCluster(scope.row)">{{ scope.row.name }}</el-button>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="状态" width="120">
					<template #default="scope">
						<div v-if="scope.row.status == true">
							<div style="display: inline-block; width: 12px; height: 12px; background: #67c23a; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #67c23a">运行中 </span>
						</div>
						<div v-else>
							<div style="display: inline-block; width: 12px; height: 12px; background: #f56c6c; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #f56c6c">故障 </span>
						</div>
					</template>
				</el-table-column>
				<el-table-column prop="nodes" label="节点数量" width="120" />
				<el-table-column prop="version" label="Version" width="120" />
				<el-table-column prop="runtime" label="运行时" width="160" />
				<el-table-column prop="pod_cidr" label="Pod_CIDR" width="130" />
				<el-table-column prop="service_cidr" label="Service_CIDR" width="130" />
				<el-table-column prop="cpu" label="CPU" width="120" />
				<el-table-column prop="memory" label="内存" width="160">
					<template #default="scope"> {{ Math.round(scope.row.memory / 1024 / 1024) }}G </template>
				</el-table-column>
				<el-table-column fixed="right" label="操作" width="170">
					<template #default="scope">
						<el-button link type="danger" size="small" @click="deleteCluster(scope.row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
		<CreateCluster v-model:dialogVisible="state.dialogVisible" @value-change="getCluster()" :title="state.title" v-if="state.dialogVisible" />
	</div>
</template>

<script setup lang="ts" name="kubernetesCluster">
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import router from '/@/router';
import { useClusterApi } from '/@/api/kubernetes/cluster';
import { ElMessage, ElMessageBox } from 'element-plus';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { ClusterInfo } from '/@/types/kubernetes/cluster';

const CreateCluster = defineAsyncComponent(() => import('./component/create.vue'));

const k8sStore = kubernetesInfo();
const clusterAPI = useClusterApi();
const state = reactive({
	dialogVisible: false,
	title: '创建集群',
});

onMounted(() => {
	getCluster();
});

const loadCluster = () => {
	state.dialogVisible = true;
};

const data = reactive({
	clusters: [] as Array<ClusterInfo>,
});

const getCluster = async () => {
	const res = await clusterAPI.listCluster();
	data.clusters = res.data;
};

const deleteCluster = async (cluster: any) => {
	ElMessageBox.confirm(`此操作将删除[ ${cluster.name} ]集群 . 是否继续?`, '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			clusterAPI
				.deleteCluster(cluster.id)
				.then(() => {
					getCluster();
					ElMessage.success('操作成功');
				})
				.catch(() => {
					ElMessage.error('操作失败');
				});
		})
		.catch(); // 取消
};

const handleCluster = (cluster: any) => {
	k8sStore.state.activeCluster = cluster.name;

	router.push({
		name: 'kubernetesDashboard',
	});
};
</script>

<style scoped lang="scss">
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
