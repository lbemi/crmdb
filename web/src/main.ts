import { createApp } from 'vue'
import App from './App.vue'
import { initRouter } from './router'
import pinia from '@/store/index'
import 'element-plus/theme-chalk/el-message.css'
import 'element-plus/theme-chalk/el-message-box.css'
import 'virtual:windi.css'
import { directive } from '@/utils/directive'
import { dateStrFormat } from '@/utils/date'
import '@/assets/iconfont/iconfont.js'
import '@/assets/iconfont/iconfont.css'
import SvgIcon from '@/component/svgIcon/svgIcon.vue'

const app = createApp(App)
//初始化存储
app.use(pinia)
//初始化路由
initRouter(app)
//自定义指令
directive(app)

// 全局时间格式化
type Filter = {
  dateFormat: <T extends any>(str: T) => T
}

declare module '@vue/runtime-core' {
  export interface ComponentCustomProperties {
    $filters: Filter
  }
}
app.config.globalProperties.$filters = {
  dateFormat(value: any) {
    if (!value) {
      return ''
    }
    return dateStrFormat('yyyy-MM-dd HH:mm:ss', value)
  }
}

// 加载全局icon组件
app.component('SvgIcon', SvgIcon)
app.mount('#app')
