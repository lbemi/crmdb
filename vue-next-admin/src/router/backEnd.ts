import { RouteRecordRaw } from 'vue-router';
import pinia from '@/stores/index';
import { useUserInfo } from '@/stores/userInfo';
import { useRequestOldRoutes } from '@/stores/requestOldRoutes';
import { Session } from '@/utils/storage';
import { NextLoading } from '@/utils/loading';
import { dynamicRoutes, notFoundAndNoPower } from '@/router/route';
import { formatTwoStageRoutes, formatFlatteningRoutes, router } from '@/router/index';
import { useRoutesList } from '@/stores/routesList';
import { useTagsViewRoutes } from '@/stores/tagsViewRoutes';
import { useMenuApi } from '@/api/system/menu';

// 后端控制路由

// 引入 api 请求接口
const menuApi = useMenuApi();

/**
 * 获取目录下的 .vue、.tsx 全部文件
 * @method import.meta.glob
 * @link 参考：https://cn.vitejs.dev/guide/features.html#json
 */
const layouModules: any = import.meta.glob('../layout/routerView/*.{vue,tsx}');
const viewsModules: any = import.meta.glob('../views/**/*.{vue,tsx}');
const dynamicViewsModules: Record<string, Function> = Object.assign({}, { ...layouModules }, { ...viewsModules });

/**
 * 后端控制路由：初始化方法，防止刷新时路由丢失
 * @method NextLoading 界面 loading 动画开始执行
 * @method useUserInfo().setUserInfos() 触发初始化用户信息 pinia
 * @method useRequestOldRoutes().setRequestOldRoutes() 存储接口原始路由（未处理component），根据需求选择使用
 * @method setAddRoute 添加动态路由
 * @method setFilterMenuAndCacheTagsViewRoutes 设置路由到 pinia routesList 中（已处理成多级嵌套路由）及缓存多级嵌套数组处理后的一维数组
 */
export async function initBackEndControlRoutes() {
	const storesRoutesList = useRoutesList(pinia);
	// 界面 loading 动画开始执行
	if (window.nextLoading === undefined) NextLoading.start();
	// 无 token 停止执行下一步
	if (!Session.get('token')) return false;
	// 触发初始化用户信息 pinia
	// https://gitee.com/lyt-top/vue-next-admin/issues/I5F1HP
	await useUserInfo().setUserInfos();

	// 获取路由菜单数据
	// const res = await getBackEndControlRoutes();
	const { kubernetesMenus: kubernetesMenus, permissions: permission, newMenus: menus } = await parseMenus();
	// 无登录权限时，添加判断
	// https://gitee.com/lyt-top/vue-next-admin/issues/I64HVO
	if (menus.length <= 0) return Promise.resolve(true);
	// 设置按钮权限
	await useUserInfo().setUserAuthButton(permission);
	console.log('>>>>>', storesRoutesList.isKubernetes);
	if (!storesRoutesList.isKubernetes) {
		console.log('-----:', storesRoutesList.isKubernetes);
	}

	// 处理路由（page），替换 dynamicRoutes（@/router/route）第一个顶级 children 的路由
	if (storesRoutesList.isKubernetes) {
		// 存储接口原始路由（未处理component），根据需求选择使用
		console.log('-----<<<<', kubernetesMenus);
		await useRequestOldRoutes().setRequestOldRoutes(JSON.parse(JSON.stringify(kubernetesMenus)));
		dynamicRoutes[0].children = await backEndComponent(kubernetesMenus);
	} else {
		await useRequestOldRoutes().setRequestOldRoutes(JSON.parse(JSON.stringify(menus)));
		dynamicRoutes[0].children = await backEndComponent(menus);
	}
	// dynamicRoutes[0].children = await backEndComponent(menus);
	// dynamicRoutes[1].children = await backEndComponent(kubernetesMenus);
	// 添加动态路由
	await setAddRoute();
	// 设置路由到 pinia routesList 中（已处理成多级嵌套路由）及缓存多级嵌套数组处理后的一维数组
	await setFilterMenuAndCacheTagsViewRoutes();
}

const parseMenus = async () => {
	const menuList = await getBackEndControlRoutes();
	const menus = menuList.data.menus;
	const permissions = menuList.data.permission;
	let kubernetesMenus: any[] = [];

	const newMenus = menus
		.map((item: any) => {
			// 检查当前项的path属性
			if (item.path === '/kubernetes') {
				// 如果是kubernetes条目，处理其children
				const newChildren = item.children.map((child: any) => {
					if (child.meta && child.meta.isK8s === true) {
						// 如果是K8s条目，则添加到filteredK8sEntries数组中
						kubernetesMenus.push(child);
						// 返回null以从原始数据中删除该项
						return null;
					}
					// 否则，保持不变
					return child;
				});
				// 更新当前项的children
				item.children = newChildren.filter((child: any) => child !== null);
			}
			// 如果当前项不是kubernetes条目，保持不变
			return item;
		})
		.filter((item: any) => item !== null); // 过滤掉null值

	//返回kubernetes 和 permission,newMenus
	return { kubernetesMenus, permissions, newMenus };
};
/**
 * 设置路由到 pinia routesList 中（已处理成多级嵌套路由）及缓存多级嵌套数组处理后的一维数组
 * @description 用于左侧菜单、横向菜单的显示
 * @description 用于 tagsView、菜单搜索中：未过滤隐藏的(isHide)
 */
export async function setFilterMenuAndCacheTagsViewRoutes() {
	const storesRoutesList = useRoutesList(pinia);
	await storesRoutesList.setRoutesList(dynamicRoutes[0].children as any);
	setCacheTagsViewRoutes();
}

/**
 * 缓存多级嵌套数组处理后的一维数组
 * @description 用于 tagsView、菜单搜索中：未过滤隐藏的(isHide)
 */
export function setCacheTagsViewRoutes() {
	const storesTagsView = useTagsViewRoutes(pinia);
	storesTagsView.setTagsViewRoutes(formatTwoStageRoutes(formatFlatteningRoutes(dynamicRoutes))[0].children);
}

/**
 * 处理路由格式及添加捕获所有路由或 404 Not found 路由
 * @description 替换 dynamicRoutes（@/router/route）第一个顶级 children 的路由
 * @returns 返回替换后的路由数组
 */
export function setFilterRouteEnd() {
	let filterRouteEnd: any = formatTwoStageRoutes(formatFlatteningRoutes(dynamicRoutes));
	// notFoundAndNoPower 防止 404、401 不在 layout 布局中，不设置的话，404、401 界面将全屏显示
	// 关联问题 No match found for location with path 'xxx'
	filterRouteEnd[0].children = [...filterRouteEnd[0].children, ...notFoundAndNoPower];
	return filterRouteEnd;
}

/**
 * 添加动态路由
 * @method router.addRoute
 * @description 此处循环为 dynamicRoutes（@/router/route）第一个顶级 children 的路由一维数组，非多级嵌套
 * @link 参考：https://next.router.vuejs.org/zh/api/#addroute
 */
export async function setAddRoute() {
	await setFilterRouteEnd().forEach((route: RouteRecordRaw) => {
		router.addRoute(route);
	});
}

/**
 * 请求后端路由菜单接口
 * @description isRequestRoutes 为 true，则开启后端控制路由
 * @returns 返回后端路由菜单数据
 */
export function getBackEndControlRoutes() {
	return menuApi.getUserMenu();
}

/**
 * 重新请求后端路由菜单接口
 * @description 用于菜单管理界面刷新菜单（未进行测试）
 * @description 路径：/src/views/system/menu/page/addMenu.vue
 */
export async function setBackEndControlRefreshRoutes() {
	await getBackEndControlRoutes();
}

/**
 * 后端路由 page 转换
 * @param routes 后端返回的路由表数组
 * @returns 返回处理成函数后的 page
 */
export function backEndComponent(routes: any) {
	if (!routes) return;
	return routes.map((item: any) => {
		if (item.component) item.component = dynamicImport(dynamicViewsModules, item.component as string);
		item.children && backEndComponent(item.children);
		return item;
	});
}

/**
 * 后端路由 page 转换函数
 * @param dynamicViewsModules 获取目录下的 .vue、.tsx 全部文件
 * @param component 当前要处理项 page
 * @returns 返回处理成函数后的 page
 */
export function dynamicImport(dynamicViewsModules: Record<string, Function>, component: string) {
	const keys = Object.keys(dynamicViewsModules);
	const matchKeys = keys.filter((key) => {
		const k = key.replace(/..\/views|../, '');
		return k.startsWith(`${component}`) || k.startsWith(`/${component}`);
	});
	if (matchKeys?.length === 1) {
		const matchKey = matchKeys[0];
		return dynamicViewsModules[matchKey];
	}
	if (matchKeys?.length > 1) {
		return false;
	}
}
