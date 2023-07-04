<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div style="margin-bottom: 8px">
				<a>容器组名: {{ pod?.metadata?.name }}：</a>
				<el-select v-model="selectContainer" placeholder="选择容器" size="small" @change="containerChange">
					<el-option v-for="item in pod?.spec?.containers" :key="item.name" :label="item.name" :value="item.name" />
				</el-select>
			</div>

			<div ref="terminal" element-loading-text="拼命连接中" id="xterm" class="xterm" :style="{ height: state.height }" />
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podShell">
import { reactive, onMounted, onBeforeUnmount, nextTick, ref } from 'vue';
import 'xterm/css/xterm.css';
import { ITerminalOptions, ITheme, Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { ElMessage } from 'element-plus';
import { useWebsocketApi } from '@/api/kubernetes/websocket';
import { podInfo } from '@/stores/pod';

const selectContainer = ref();
const pod = podInfo().state.podShell;

const state = reactive({
	machineId: 0,
	cmd: '',
	height: '600px',
	term: null as any,
	socket: null as any,
});

const webSocketApi = useWebsocketApi();
const containerChange = () => {
	state.term.clear();
	initSocket();
};

onMounted(() => {
	state.height = window.innerHeight - 200 + 'px';
});

onBeforeUnmount(() => {
	closeAll();
});

nextTick(() => {
	initXterm();
});

function initXterm() {
	const term = new Terminal({
		fontSize: 15,
		fontWeight: 'normal',
		fontFamily: 'Consolas, Lucida Console,monospace,JetBrainsMono, monaco',
		cursorBlink: true,
		disableStdin: false,
		theme: {
			foreground: '#ffffff', //字体
			background: '#060101', //背景色
			cursor: 'help', //设置光标
		} as ITheme,
	} as ITerminalOptions);

	const fitAddon = new FitAddon();
	term.loadAddon(fitAddon);
	term.open(document.getElementById('xterm')!);
	term.writeln('\n 正在初始化终端,请选择容器...');
	fitAddon.fit();
	term.focus();
	state.term = term;
	// 监听窗口resize
	window.addEventListener('resize', () => {
		try {
			// 窗口大小改变时，触发xterm的resize方法使自适应
			fitAddon.fit();
			if (state.term) {
				state.term!.focus();
			}
		} catch (e) {
			console.log(e);
		}
	});

	term.onData((key: any) => {
		state.socket.send(key);
	});
}

function initSocket() {
	state.socket = webSocketApi.createShellWebsocket(pod.metadata?.namespace!, pod.metadata?.name!, selectContainer.value);
	// state.socket = webSocketApi.createTestShellWebsocket();

	// 监听socket错误信息
	state.socket.onerror = () => {
		ElMessage.error('连接失败');
	};

	state.socket.onclose = () => {
		if (state.term) {
			state.term.writeln('\r\n\x1b[31m提示: 连接已关闭...');
		}
	};
	state.socket.onmessage = (e: any) => {
		state.term.write(e.data);
	};
}

function closeAll() {
	if (state.socket) {
		state.socket.close();
	}
	if (state.term) {
		state.term.dispose();
		state.term = null;
	}
}
</script>
<style lang="scss">
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
