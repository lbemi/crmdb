import { createApp } from 'vue'
import App from './App.vue'
import { initRouter } from './router'
import { initStore } from './store'

import { ElMessage, ElMessageBox } from 'element-plus';
import 'element-plus/theme-chalk/el-message.css';
import 'element-plus/theme-chalk/el-message-box.css';


const app = createApp(App)
//初始化路由
initRouter(app)
//初始化存储
initStore(app)

app.mount('#app')

