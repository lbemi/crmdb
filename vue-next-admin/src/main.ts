import { createApp } from 'vue';
import pinia from '@/stores/index';
import App from '@/App.vue';
import router from '@/router';
import { directive } from '@/directive/index';
import { i18n } from '@/i18n/index';
import other from '@/utils/other';
import {addDateRange,parseTime,handleTree,selectDictLabel,download} from '@/utils'
import { dateStrFormat } from "@/utils/formatTime"
import mitt from 'mitt';

import ElementPlus from 'element-plus';
import '@/theme/index.scss';
import VueGridLayout from 'vue-grid-layout';

const app = createApp(App);

directive(app);
other.elSvg(app);

app.use(pinia).use(router).use(ElementPlus).use(i18n).use(VueGridLayout).mount('#app');

// 全局方法挂载
app.config.globalProperties.mittBus = mitt();
app.config.globalProperties.addDateRange = addDateRange;
app.config.globalProperties.parseTime = parseTime
app.config.globalProperties.handleTree = handleTree
app.config.globalProperties.selectDictLabel = selectDictLabel
app.config.globalProperties.download = download
app.config.globalProperties.dateStrFormat = dateStrFormat