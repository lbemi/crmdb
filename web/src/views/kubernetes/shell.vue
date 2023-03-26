<template >
    <span>容器组名: {{ pod?.metadata.name }}</span>
    <el-select v-model="selectContiner" class="m-2" placeholder="选择容器" size="large" @change="containerChange">
        <el-option v-for="item in pod?.spec.containers" :key="item.name" :label="item.name" :value="item.name" />
    </el-select>
    <div ref="terminal" element-loading-text="拼命连接中" id="xterm" class="xterm" :style="{ height: state.height }" />
</template>
<script setup lang="ts">
import { reactive, onMounted, onBeforeUnmount, nextTick, ref } from 'vue'
import { podStore } from '@/store/kubernetes/pods'
import 'xterm/css/xterm.css'
import { ITerminalOptions, ITheme, Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { ElMessage } from 'element-plus'

const cloudName = kubeStore().activeCluster
const selectContiner = ref()
const pod = podStore().podShell

const state = reactive({
    machineId: 0,
    cmd: '',
    height: '600px',
    term: null as any,
    socket: null as any
})

const containerChange = () => {
    state.term.clear()
    initSocket()
}


onMounted(() => {
    state.height = window.innerHeight - 200 + 'px'
})

onBeforeUnmount(() => {
    closeAll()
})

nextTick(() => {
    initXterm()
})

function initXterm() {

    const term = new Terminal({
        fontSize: 15,
        fontWeight: 'normal',
        fontFamily: 'Consolas, Lucida Console,monospace,JetBrainsMono, monaco',
        cursorBlink: true,
        disableStdin: false,
        theme: {
            foreground: '#7e9192', //字体
            background: '#060101', //背景色
            cursor: 'help' //设置光标
        } as ITheme
    } as ITerminalOptions)

    const fitAddon = new FitAddon()
    term.loadAddon(fitAddon)
    term.open(document.getElementById('xterm')!)
    term.writeln("\n\n 正在初始化终端...")
    fitAddon.fit()
    term.focus()
    state.term = term
    // 监听窗口resize
    window.addEventListener('resize', () => {
        try {
            // 窗口大小改变时，触发xterm的resize方法使自适应
            fitAddon.fit()
            if (state.term) {
                state.term!.focus()
                // send({
                //     type: resize,
                //     Cols: parseInt(state.term.cols),
                //     Rows: parseInt(state.term.rows)
                // })
            }
        } catch (e) {
            console.log(e)
        }
    })

    term.onData((key: any) => {
        state.socket.send(key)
    })
}


function initSocket() {
    let url =
        (location.protocol === 'http:' ? 'ws' : 'wss') +
        '://' +
        location.hostname +
        ':8080' +
        '/api/v1/pod/' +
        `${pod!.metadata.namespace}/${pod!.metadata.name}/` + selectContiner.value +
        '?cloud=' + cloudName
    state.socket = new WebSocket(url)
    // 监听socket连接
    state.socket.onopen = () => {
        console.log("open");
    }

    // 监听socket错误信息
    state.socket.onerror = (e: any) => {
        ElMessage.error('连接失败')
    }

    state.socket.onclose = () => {
        if (state.term) {
            state.term.writeln('\r\n\x1b[31m提示: 连接已关闭...')
        }
       
    }
    state.socket.onmessage = (e: any) => {
        state.term.write(e.data)
    }
}


function closeAll() {
    if (state.socket) {
        state.socket.close()
    }
    if (state.term) {
        state.term.dispose()
        state.term = null
    }
}



</script>
<style lang="less"></style>