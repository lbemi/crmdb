<template>
	<div class="layout-padding">
		<el-card shadow="hover">
			<template #header>
				<div style="align-items: center; display: flex">
					<el-form-item label="Pod名称:" style="margin-bottom: 0px">
						<el-text type="primary" style="margin-right: 10px">{{ pod.metadata?.name }}</el-text>
					</el-form-item>
					<el-form-item label="容器:" style="margin-bottom: 0px">
						<el-select v-model="selectContainer" class="m-2" placeholder="选择容器" size="small">
							<el-option v-for="item in pod.spec?.containers" :key="item.name" :label="item.name"
								:value="item.name" />
						</el-select>
					</el-form-item>
					<el-button type="primary" size="small" class="ml10" @click="getLog">显示日志</el-button>
					<el-button type="primary" size="small" class="ml10" @click="logs = ''">清空</el-button>

					<el-button plain size="small" type="primary" class="ml10" @click="download">下载日志</el-button>
					<VideoPlay v-if="stop" style="width: 1.9em; height: 1.9em; color: #2c9505; margin-left: 10px"
						@click="stop = !stop" />
					<VideoPause v-else style="width: 1.9em; height: 1.9em; color: #ff0007; margin-left: 10px"
						@click="stop = !stop" />
				</div>
			</template>

			<Codemirror ref="codemirrorRef" v-model="logs" style="height: 730px; margin-bottom: 10px;"
				class="layout-padding" @ready="handleReady" :autofocus="true" :tabSize="2" :extensions="extensions"
				disabled />
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podLog">
import { onBeforeUnmount, ref, shallowRef } from 'vue';
import { podInfo } from '@/stores/pod';
import { useWebsocketApi } from '@/api/kubernetes/websocket';
import { VideoPause, VideoPlay } from '@element-plus/icons-vue';
import { oneDark } from '@codemirror/theme-one-dark';
import { StreamLanguage } from '@codemirror/language';
import { liveScript } from '@codemirror/legacy-modes/mode/livescript';
import { Codemirror } from 'vue-codemirror';
import { EditorView } from '@codemirror/view';
import { saveAs } from 'file-saver';

const codemirrorRef = ref();

const extensions = [oneDark, StreamLanguage.define(liveScript)];
const selectContainer = ref();
const pod = podInfo().state.podShell;
const logs = ref('');
const websocketApi = useWebsocketApi();
const stop = ref(false);
const ws = ref();

const cmView = shallowRef<EditorView>();
const handleReady = ({ view }: any) => {
	cmView.value = view;
};
const getLog = async () => {
	logs.value = ''; //清空数据
	if (ws.value) {
		ws.value.close();
	}
	ws.value = websocketApi.createLogWebsocket(pod.metadata?.namespace!, pod.metadata?.name!, selectContainer.value);

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

				cmView.value?.scrollDOM.scrollTo(0, cmView.value?.scrollDOM.scrollHeight + 40);
			} else {
				cacheLog.value += e.data; //暂停的时候保存日志，否则会丢失这部分日志
			}
		}
	};
};

const download = () => {
	let str = new Blob([logs.value], { type: 'text/plain;charset=utf-8' });
	saveAs(str, pod.metadata?.name! + '.log');
};
onBeforeUnmount(() => {
	ws.value.close();
});
</script>
<style lang="scss"></style>
