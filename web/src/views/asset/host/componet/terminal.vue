<template>
  <!-- xterm虚拟终端 -->
  <el-card>
    <h1>SSH....{{ query.ip }}</h1>
    <br />
    <div
      id="xterm"
      v-loading="loading"
      ref="terminal"
      element-loading-text="拼命连接中"
    ></div>
  </el-card>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus'
import { onMounted, onUnmounted, reactive, ref } from 'vue'
//引入xterm终端依赖
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import 'xterm/css/xterm.css'
import 'xterm/lib/xterm.js'
import { useRoute } from 'vue-router'

const { query } = useRoute()
const loading = ref(false)

interface Data {
  term: Terminal | null
  socket: WebSocket | null
}

const data = reactive<Data>({
  term: null,
  socket: null
})

const props = defineProps({
  webshelloptions: {
    type: Object,
    default() {
      return {}
    }
  }
})
console.log('参数:', query)
onMounted(() => {
  // 优化体验
  initTerm()
  initSocket()
})

onUnmounted(() => {
  closeSocket()
})

const initTerm = () => {
  var containerWidth = window.screen.height
  var containerHeight = window.screen.width
  var cols = Math.floor((containerWidth - 30) / 9)
  var rows = Math.floor(window.innerHeight / 17) - 2

  //初始化xterm实例
  data.term = new Terminal({
    rows: cols, //行数
    cols: rows,
    convertEol: false, //启用时，光标将设置为下一行的开头
    scrollback: 10, //终端中的回滚量
    disableStdin: false, //是否应禁用输入
    cursorStyle: 'underline', //光标样式
    cursorBlink: true, //光标闪烁
    theme: {
      foreground: 'white', //字体
      background: '#060101', //背景色
      cursor: 'help' //设置光标
    }
  })
  //绑定dom
  data.term.open(document.getElementById('xterm')!)
  data.term?.write('Hello from \x1B[1;3;31mxterm.js\x1B[0m $ ')
  //终端适应父元素大小
  const fitAddon = new FitAddon()
  data.term.loadAddon(fitAddon)
  fitAddon.fit()
  //获取终端的焦点
  data.term.focus()
  let _data = data //一定要重新定义一个this，不然this指向会出问题
  //onData方法用于定义输入的动作
  data.term.onData(function (key) {
    // 这里key值是输入的值，数据格式就是后端定义的 {"operation":"stdin","data":"ls"}
    let msgOrder = {
      operation: 'stdin',
      data: key
    }
    console.log('******', msgOrder)

    //发送数据
    _data.socket?.send(JSON.stringify(msgOrder))
  })
}
const initSocket = () => {
  if (data.socket !== null) {
    return
  }
  // const websocketAddr = import.meta.env.VITE_BASE_API.replace("http", "ws");
  // //定义websocket连接地址
  // let terminalWsUrl =
  //   websocketAddr +
  //   `/clouds/webshell/ws?cloud=${props.webshelloptions.cloud}&namespace=${props.webshelloptions.namespace}&pod=${props.webshelloptions.pod}&container=${props.webshelloptions.container}`;

  let url =
    (location.protocol === 'http:' ? 'ws' : 'wss') +
    '://' +
    location.hostname +
    ':8080' +
    '/api/v1/host/' +
    `${query.id}` +
    '/ws' +
    '?' +
    '&rows=' +
    data.term?.rows +
    '&cols=' +
    data.term?.cols

  //实例化
  data.socket = new WebSocket(url, [localStorage.getItem('token')!])
  console.log(data.socket, '....')
  //关闭连接时的方法
  // socketOnClose();
  //接收消息的方法
  socketOnMessage()
  //报错时的方法
  socketOnError()
}

const socketOnMessage = () => {
  data.socket!.onmessage = (msg) => {
    //接收到消息后将字符串转为对象，输出data内容
    let content = JSON.parse(msg.data)
    console.log('发送消息--:', msg)

    data.term!.write(content.data)
  }
}

const socketOnClose = () => {
  data.socket!.onclose = () => {
    //关闭连接后打印在终端里
    console.log('执行关闭....')
    data.socket = null
  }
}
const socketOnError = () => {
  data.socket!.onerror = (e) => {
    console.log('err: socket-->', data.socket)
    console.log('errr---: ', e)

    ElMessage.error('连接失败.')
  }
}
const closeSocket = () => {
  //若没有实例化，则不需要关闭
  if (data.socket === null) {
    return
  }
  console.log('执行关闭....!!!')

  data.socket.close()
}
</script>

<style scoped="scoped">
#terminal {
  width: 100%;
  height: 100%;
}
</style>
