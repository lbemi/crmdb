<template>
	<div class="layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<!-- <div>Cluster-集群</div> -->
			<el-tabs>
				<el-tab-pane lazy label="概览"><Dashboard /></el-tab-pane>
				<el-tab-pane lazy label="基本信息"><Info /></el-tab-pane>
				<el-tab-pane label="连接信息">连接信息</el-tab-pane>
				<el-tab-pane label="集群资源">集群资源</el-tab-pane>
				<el-tab-pane lazy label="集群日志"><Log /></el-tab-pane>
			</el-tabs>
		</el-card>
	</div>
</template>

<script setup lang="ts" name="kubernetesDashboard">
import Dashboard from './component/dashboard.vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { defineAsyncComponent, onMounted } from 'vue';

const Info = defineAsyncComponent(() => import('./component/info.vue'));
const Log = defineAsyncComponent(() => import('./component/log.vue'));

const k8sStore = kubernetesInfo();

onMounted(() => {
	k8sStore.listNamespace();
});
</script>

<style scoped lang="scss">
.body {
	margin-left: 10px;
	width: 100%;
	// height: 100;
}
.el-tabs--border-card {
	// border-bottom: 0px;
	// border-left: 0px;
	// border-right: 0px;
	border: 0;
}
</style>
