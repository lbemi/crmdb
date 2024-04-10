<template>
	<el-menu
		router
		:default-active="state.defaultActive"
		background-color="transparent"
		:collapse="state.isCollapse"
		:unique-opened="getThemeConfig.isUniqueOpened"
		:collapse-transition="false"
	>
		<el-menu-item v-if="k8sStore.state.isKubernetesRoutes">
			<template #title>
				<SvgIcon :name="'ele-ArrowLeftBold'" @click="backKubernetesList()" />集群：

				<el-select
					v-model="k8sStore.state.activeCluster"
					placeholder="Select"
					size="small"
					popper-class="my-select-dropdown layout-navbars-breadcrumb-span"
					@change="handleChange"
				>
					<el-option v-for="item in k8sStore.state.clusterList" :key="item.name" :label="item.name" :value="item.name!" />
				</el-select>
			</template>
		</el-menu-item>

		<template v-for="val in menuLists">
			<el-sub-menu :index="val.path" v-if="val.children && val.children.length > 0" :key="val.path">
				<template #title>
					<SvgIcon :name="val.meta.icon" />
					<span>{{ $t(val.meta.title) }}</span>
				</template>
				<SubItem :chil="val.children" />
			</el-sub-menu>
			<template v-else>
				<el-menu-item :index="val.path" :key="val.path">
					<SvgIcon :name="val.meta.icon" />
					<template #title v-if="!val.meta.isLink || (val.meta.isLink && val.meta.isIframe)">
						<span>{{ $t(val.meta.title) }}</span>
					</template>
					<template #title v-else>
						<a class="w100" @click.prevent="onALinkClick(val)">{{ $t(val.meta.title) }}</a>
					</template>
				</el-menu-item>
			</template>
		</template>
	</el-menu>
</template>

<script setup lang="ts" name="navMenuVertical">
import { defineAsyncComponent, reactive, computed, onMounted, watch } from 'vue';
import { useRoute, onBeforeRouteUpdate, RouteRecordRaw } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '@/stores/themeConfig';
import other from '@/utils/other';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useRouter } from 'vue-router';
import { initBackEndControlRoutes } from '@/router/backEnd';
import mittBus from '@/utils/mitt';

const k8sStore = kubernetesInfo();
const router = useRouter();
// 引入组件
const SubItem = defineAsyncComponent(() => import('@/layout/navMenu/subItem.vue'));

// 定义父组件传过来的值
const props = defineProps({
	// 菜单列表
	menuList: {
		type: Array<RouteRecordRaw>,
		default: () => [],
	},
});

// 定义变量内容
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const route = useRoute();
const state = reactive({
	// 修复：https://gitee.com/lyt-top/vue-next-admin/issues/I3YX6G
	defaultActive: route.meta.isDynamic ? route.meta.isDynamicPath : route.path,
	isCollapse: false,
});
const backKubernetesList = async () => {
	k8sStore.setKubernetesRoutes(false);
	await initBackEndControlRoutes();
	mittBus.emit('getBreadcrumbIndexSetFilterRoutes');
	await router.push({ path: '/kubernetes' });
};
const handleChange = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
// 获取父级菜单数据
const menuLists = computed(() => {
	return <RouteItems>props.menuList;
});
// 获取布局配置信息
const getThemeConfig = computed(() => {
	return themeConfig.value;
});
// 菜单高亮（详情时，父级高亮）
const setParentHighlight = (currentRoute: RouteToFrom) => {
	const { path, meta } = currentRoute;
	const pathSplit = meta?.isDynamic ? meta.isDynamicPath!.split('/') : path!.split('/');
	if (pathSplit.length >= 4 && meta?.isHide) return pathSplit.splice(0, 3).join('/');
	else return path;
};
// 打开外部链接
const onALinkClick = (val: RouteItem) => {
	other.handleOpenLink(val);
};
// 页面加载时
onMounted(() => {
	state.defaultActive = setParentHighlight(route);
});
// 路由更新时
onBeforeRouteUpdate((to) => {
	// 修复：https://gitee.com/lyt-top/vue-next-admin/issues/I3YX6G
	state.defaultActive = setParentHighlight(to);
	const clientWidth = document.body.clientWidth;
	if (clientWidth < 1000) themeConfig.value.isCollapse = false;
});
// 设置菜单的收起/展开
watch(
	() => themeConfig.value.isCollapse,
	(isCollapse) => {
		document.body.clientWidth <= 1000 ? (state.isCollapse = false) : (state.isCollapse = isCollapse);
	},
	{
		immediate: true,
	}
);
</script>
<style lang="scss" scoped>
:deep(.el-input__wrapper) {
	background-color: #ffffff00;
}
:deep(.el-input__inner) {
	font-size: 16px;
	color: #626aef;
}
/* 全局样式文件或组件的 <style> 标签中 */
:deep(.el-input) {
	width: 100px;
	--el-input-focus-border: #ffffff00;
	--el-input-transparent-border: 0 0 0 0px;
	--el-input-border-color: #ffffff00;
	--el-input-hover-border: 0px !important;
	--el-input-hover-border-color: #ffffff00;
	--el-input-focus-border-color: #ffffff00;
	--el-input-clear-hover-color: #ffffff00;
	box-shadow: 0 0 0 0px !important;
	--el-input-border: 0px;
	border-bottom: 1px solid #626aef;
}
:deep(.el-select .el-input__wrapper.is-focus) {
	box-shadow: 0 0 0 0px !important;
	background-color: #ffffff00;
}
:deep(.el-select .el-input.is-focus .el-input__wrapper) {
	box-shadow: 0 0 0 0px !important;
}
:deep(.el-select) {
	--el-select-border-color-hover: #ffffff00;
}
// 自定义el-select的下拉箭头
// :deep(.el-select__caret) {
// 	/*很关键：将默认的select选择框样式清除*/
// 	appearance: none;
// 	-moz-appearance: none;
// 	-webkit-appearance: none;
// 	/*为下拉小箭头留出一点位置，避免被文字覆盖*/
// 	padding-right: 15px;
// 	/*自定义图片*/
// 	// background: url('http://ourjs.github.io/static/2015/arrow.png') no-repeat scroll right center transparent;
// 	/*自定义图片的大小*/
// 	background-size: 14px 12px;
// }
/*将小箭头的样式去去掉*/
:deep(.el-icon-arrow-up:before) {
	content: '';
}
</style>
