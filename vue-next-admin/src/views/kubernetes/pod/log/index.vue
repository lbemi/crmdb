<template>
	<div class="layout-pd">
		<el-card shadow="hover" class="mb15">
			<span>容器组名: {{ pod?.metadata.name }}</span>
			<el-select v-model="selectContainer" class="m-2" placeholder="选择容器" size="large">
				<el-option v-for="item in pod?.spec.containers" :key="item.name" :label="item.name" :value="item.name" />
			</el-select>
			<el-button type="primary" @click="showLogs">显示日志</el-button>
			<el-button @click="logs = ''">清空</el-button>
			<div id="logs" class="logs">
				{{ logs }}
			</div>
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podLog">
import { ref } from 'vue';
import request from '/@/utils/request';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { podInfo } from '/@/stores/pod';

const k8sStore = kubernetesInfo();
const selectContainer = ref();
const pod = podInfo().state.podShell;
const logs = ref();

const getLog = async (namespace: string, podName: string, container: string) => {
	request({
		url: '/pod/log/' + namespace + '/' + podName + '/' + container + '?cloud=' + k8sStore.state.activeCluster,
		method: 'GET',
		onDownloadProgress: (e) => {
			console.log(e);
			const dataChunk = e.currentTarget.response;
			logs.value.log += dataChunk;
			console.log('++++++++++++++++++++', logs);
			const logDiv = document.getElementById('logs');
			if (logDiv !== undefined && logDiv?.scrollTop !== undefined) {
				logDiv.scrollTop = logDiv.scrollHeight;
			}
		},
	});
};

const showLogs = () => {
	request({
		url: '/pod/log/' + pod?.metadata.namespace! + '/' + pod?.metadata.name! + '/' + selectContainer.value + '?cloud=' + k8sStore.state.activeCluster,
		method: 'GET',
		onDownloadProgress: (e) => {

			const dataChunk = e.event.currentTarget.response;
			logs.value += dataChunk;
			const logDiv = document.getElementById('logs');
			if (logDiv  && logDiv?.scrollTop != undefined) {
				logDiv.scrollTop = logDiv.scrollHeight;
			}
		},
	}).catch((e) => {
		// router.push({name: 'pod'})
	});
};
</script>
<style lang="scss">
.logs {
	overflow: auto;

	margin: 10px auto;
	min-height: 200px;
	max-height: 100%;
	border: solid 1px black;
	background-color: #454545;
	padding: 10px;

	color: #27aa5e;
	line-height: 21pt;
	white-space: pre;
	width: 100%;
}
</style>
