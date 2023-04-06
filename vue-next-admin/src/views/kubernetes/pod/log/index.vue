<template>
	<div class="layout-pd">
		<el-card shadow="hover" class="mb15">
			<span class="mb15">容器组名: {{ pod?.metadata.name }}</span> :
			<el-select v-model="selectContainer" class="m-2" placeholder="选择容器" size="default">
				<el-option v-for="item in pod?.spec.containers" :key="item.name" :label="item.name" :value="item.name" />
			</el-select>
			<el-button type="primary" size="default" class="ml10" @click="getLog">显示日志</el-button>

			<el-button type="primary" size="default" class="ml10" @click="logs = ''">清空</el-button>
			<el-button type="primary" size="default" class="ml10" @click="stop = !stop">{{ stop ? '继续' : '暂停' }}</el-button>
        <el-scrollbar ref="scrollbarRef" height="800px" >
          <div ref="innerRef" id="logs" class="logs">
            {{ logs }}
          </div>
        </el-scrollbar>
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podLog">
import { onBeforeUnmount, ref } from 'vue';
import { ElScrollbar } from 'element-plus'
import { podInfo } from '/@/stores/pod';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';

const selectContainer = ref();
const pod = podInfo().state.podShell;
const logs = ref();
const websocketApi = useWebsocketApi();
const stop = ref(false);
const ws = ref();
const innerRef = ref<HTMLDivElement>()
const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>()

const getLog = async () => {
	logs.value = ''; //清空数据
	if (ws.value) {
		ws.value.close();
	}
	ws.value = websocketApi.createLogWebsocket(pod.metadata?.namespace, pod?.metadata.name!, selectContainer.value);
	const logDiv = document.getElementById('logs');
	const cacheLog = ref('');
	ws.value.onmessage = (e) => {
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
          scrollbarRef.value!.setScrollTop(innerRef.value!.clientHeight+10)
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
</style>
