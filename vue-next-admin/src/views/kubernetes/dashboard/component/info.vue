/** * Created by lei on 2022/12/08 */
<template>
	<div>
				<table class="info">
					<tr>
						<td colspan="3" class="title">基本信息</td>
					</tr>
					<tr>
						<td>集群ID: {{ state.cluster.id }}</td>
						<td>集群名称: {{ state.cluster.name }}</td>
						<td colspan="2">
							<div v-if="state.cluster.status == true">
								<div style="display: inline-block; width: 12px; height: 12px; background: #67c23a; border-radius: 50%"></div>
								<span style="margin-left: 5px; font-size: 12px; color: #67c23a">运行中 </span>
							</div>
							<div v-else>
								<div style="display: inline-block; width: 12px; height: 12px; background: #f56c6c; border-radius: 50%"></div>
								<span style="margin-left: 5px; font-size: 12px; color: #f56c6c">故障 </span>
							</div>
						</td>
					</tr>
				</table>

			<div style="margin-top: 15px">
				<table class="info">
					<tr>
						<td colspan="3" class="title">详细信息</td>
					</tr>
					<tr>
						<td>API Server 内网连接端点</td>
						<td colspan="2">http://{{ state.cluster.internal_ip }}</td>
					</tr>
					<tr>
						<td>Pod 网络 CIDR</td>
						<td colspan="2">{{ state.cluster.pod_cidr }}</td>
					</tr>
					<tr>
						<td>Service CIDR</td>
						<td colspan="2">10.244.0.0/24</td>
					</tr>
					<tr>
						<td>运行时</td>
						<td colspan="2">{{ state.cluster.runtime }}</td>
					</tr>
					<tr>
						<td>节点 IP 数量</td>
						<td colspan="2">{{ state.cluster.nodes }}</td>
					</tr>
					<tr>
						<td>网络插件</td>
						<td colspan="2">Calico</td>
					</tr>
					<tr>
						<td>创建时间</td>
						<td colspan="2">{{dateStrFormat (state.cluster.created_at )}}</td>
					</tr>
				</table>
			</div>
	</div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useClusterApi } from '/@/api/kubernetes/cluster';
import { ClusterInfo } from '/@/types/cluster';

const k8sStore = kubernetesInfo();
const clusterApi = useClusterApi();
const state = reactive({
	cluster: {} as ClusterInfo,
	query: {
		name: '',
	},
});

onMounted(() => {
	getCluster();
});
const getCluster = async () => {
	await clusterApi.getCluster(k8sStore.activeCluster).then((res) => {
		state.cluster = res.data;
	});
};
</script>

<style scoped lang="scss">
table,
table tr th,
table tr td {
	table-layout: fixed;
	border: 1px solid #dbdbdbe5;
	height: 45px;
	width: 100%;
	font-size: 12px;
	color: rgb(145, 143, 143);
	padding-left: 16px;
}
table {
	table-layout: fixed;
	width: 100%;
	min-height: 25px;
	line-height: 25px;
	text-align: left;
	border-collapse: collapse;
}
.title {
	background-color: #f5f7fa;
	color: #333333;
	font-size: 14px;
}
.info {
	max-height: calc(100vh - 200px);
}
</style>
