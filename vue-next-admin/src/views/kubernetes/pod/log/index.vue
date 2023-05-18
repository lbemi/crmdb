<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<span class="mb15">容器组名: {{ pod.metadata?.name }}</span> :
			<el-select v-model="selectContainer" class="m-2" placeholder="选择容器" size="small">
				<el-option v-for="item in pod.spec?.containers" :key="item.name" :label="item.name" :value="item.name" />
			</el-select>
			<el-button type="primary" size="small" class="ml10" @click="getLog">显示日志</el-button>

			<el-button type="primary" size="small" class="ml10" @click="logs = ''">清空</el-button>
			<el-button type="success" size="small" class="ml10" v-if="stop" @click="stop = !stop" :icon="VideoPlay">{{ stop ? '继续' : '暂停' }}</el-button>
			<el-button type="danger" size="small" class="ml10" v-else @click="stop = !stop" :icon="VideoPause">{{ stop ? '继续' : '暂停' }}</el-button>
			<el-scrollbar ref="scrollbarRef" height="800px" class="logs">
				<div ref="innerRef" id="logs">
					{{ logs }}
				</div>
			</el-scrollbar>
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podLog">
import { onBeforeUnmount, ref } from 'vue';
import { ElScrollbar } from 'element-plus';
import { podInfo } from '/@/stores/pod';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';
import { VideoPause, VideoPlay } from '@element-plus/icons-vue';

const selectContainer = ref();
const pod = podInfo().state.podShell;
const logs = ref();
const websocketApi = useWebsocketApi();
const stop = ref(false);
const ws = ref();
const innerRef = ref<HTMLDivElement>();
const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>();

const getLog = async () => {
	logs.value = ''; //清空数据
	if (ws.value) {
		ws.value.close();
	}
	ws.value = websocketApi.createLogWebsocket(pod.metadata?.namespace!, pod.metadata?.name!, selectContainer.value);
	const logDiv = document.getElementById('logs');
	const cacheLog = ref('');
	ws.value.onmessage = (e: any) => {
		if (e.data === 'ping') {
			return;
		} else {
			// const object = JSON.parse(e.data);
			if (!stop.value) {
				logs.value += cacheLog.value;
				cacheLog.value = ''; //置空
				logs.value += e.data;
				if (logDiv && logDiv?.scrollTop != undefined) {
					// logDiv.scrollTop = logDiv.scrollHeight;
					scrollbarRef.value!.setScrollTop(innerRef.value!.clientHeight + 10);
				}
			} else {
				cacheLog.value += e.data; //暂停的时候保存日志，否则会丢失这部分日志
			}
		}
	};
};

onBeforeUnmount(() => {
	ws.value.close();
});
</script>
<style lang="scss">
.logs {
	margin-top: 10px;
	color: #27aa5e;
	line-height: 18pt;
	width: 100%;
	//height: 800px;
	//overflow-y: scroll;
	background-color: #414141;
	border: 1px solid black;
	padding: 10px;
	white-space: pre-line;
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
