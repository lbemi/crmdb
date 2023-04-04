<template>
	<div class="layout-pd">
		<el-card shadow="hover">
			<div class="mb15">
				命名空间:
				<el-select v-model="kubeStore.activeNamespace" style="max-width: 180px" class="m-2" placeholder="Select" size="default" @change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in namespace" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name" />
				</el-select>
				<el-button type="primary" size="default" class="ml10">创建Deployment</el-button>
				<el-button type="danger" size="default" class="ml10" :disabled="data.selectData.length == 0" @click="deleteDeployments(data.selectData)"
					>批量删除</el-button
				>
			</div>

			<el-table
				:data="data.deployments"
				style="width: 100%"
				@selection-change="handleSelectionChange"
				v-loading="data.loading"
				max-height="100vh - 235px"
			>
				<el-table-column type="selection" width="55" />

				<el-table-column prop="metadata.name" label="名称" width="220px">
					<template #default="scope">
						<el-button link type="primary" @click="deployDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
					</template>
				</el-table-column>
				<el-table-column label="状态" width="220px">
					<template #default="scope">
						<el-button v-if="scope.row.status.conditions[0].status === 'True'" type="success" :icon="Check" size="small" circle />
						<el-button v-else type="danger" :icon="Close" size="small" circle />
					</template>
				</el-table-column>

				<el-table-column label="镜像" width="540px">
					<template #default="scope">
						<el-tag type="success" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{
							item.image.split('@')[0]
						}}</el-tag>
					</template>
				</el-table-column>

				<el-table-column label="标签" width="280px" show-overflow-tooltip="true">
					<template #default="scope">
						<el-tag type="info" v-for="(item, key, index) in scope.row.metadata.labels" :key="index"> {{ key }}:{{ item }} </el-tag>
					</template>
				</el-table-column>

				<el-table-column prop="spec.replicas" label="Pods" width="80px">
					<template #default="scope">
						<a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a
						>/ <a style="color: green">{{ scope.row.status.replicas }}</a
						>/
						<a style="color: red">{{ scope.row.status.unavailableReplicas || '0' }}</a>
					</template>
				</el-table-column>
				<el-table-column label="创建时间" width="180px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<Pagination :total="data.total" @handlePageChange="handlePageChange" />
		</el-card>
	</div>
</template>

<script setup lang="ts" name="k8sDeployment">
import { reactive, onMounted, onBeforeUnmount, defineAsyncComponent } from 'vue';
import { Check, Close } from '@element-plus/icons-vue';

import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1Deployment } from '@kubernetes/client-node';
import { PageInfo } from '/@/types/kubernetes/common';
import { kubernetesInfo } from '/@/stores/kubernetes';
import router from '/@/router';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));
const deploymentApi = useDeploymentApi();
const { activeCluster, activeNamespace, setActiveDeployment, setActiveNamespace, namespace } = kubernetesInfo();
const kubeStore = kubernetesInfo();
onMounted(() => {
	listDeployment();
});
const deleteDeployments = (data: any) => {};
//
// var dns = webSocketURL + activeCluster+ '/deployment'
// var ws = new WebSocket(dns)
// ws.onopen = () => {
//   console.log('ws connected.')
// }
// ws.onmessage = (e) => {
//   if (e.data === 'ping') {
//     return
//   } else {
//     const object = JSON.parse(e.data)
//     if (
//       object.type === 'deployment' &&
//       object.result.namespace === activeNamespace &&
//       object.cluster == activeCluster
//     ) {
//       data.deployments = object.result.data
//     }
//   }
// }
// ws.onclose = () => {
//   console.log('close')
// }

// const data = reactive(new Data())
const data = reactive({
	query: {
		cloud: '',
		page: 1,
		limit: 10,
	},
	namespace: '',
	loading: false,
	deployments: [] as V1Deployment[],
	selectData: [] as V1Deployment[],
	total: 0,
});
const handleSelectionChange = (value: any) => {
	data.selectData = value;
	console.log(data.selectData);
};

onBeforeUnmount(() => {
	console.log('关闭....');
	// ws.close()
});
const listDeployment = async () => {
	data.namespace = activeNamespace;
	data.query.cloud = activeCluster;
	try {
		data.loading = true;
		await deploymentApi.listDeployment(kubeStore.activeNamespace, data.query).then((res) => {
			data.deployments = res.data.data;
			data.total = res.data.total;
		});
	} catch (error) {
		console.log(error);
	}
	data.loading = false;
};
const handleChange = () => {
	// data.namespace = activeNamespace
	// setActiveNamespace(data.namespace)
	//   setActiveNamespace(data.namespace)
	listDeployment();
};
const handlePageChange = (pageInfo: PageInfo) => {
	data.query.page = pageInfo.page;
	data.query.limit = pageInfo.limit;
	listDeployment();
};
const deployDetail = async (dep: V1Deployment) => {
	await setActiveDeployment(dep);
	router.push({
		name: 'detail',
	});
};
</script>

<style scoped lang="less"></style>
