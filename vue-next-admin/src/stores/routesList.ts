import { defineStore } from 'pinia';
import { reactive } from 'vue';

/**
 * 路由列表
 * @methods setRoutesList 设置路由数据
 * @methods setColumnsMenuHover 设置分栏布局菜单鼠标移入 boolean
 * @methods setColumnsNavHover 设置分栏布局最左侧导航鼠标移入 boolean
 */
export const useRoutesList = defineStore('routesList', {
	state: (): RoutesListState =>
		reactive({
			routesList: [],
			isKubernetes: false,
			kubernetesRoutesList: [],
			isColumnsMenuHover: false,
			isColumnsNavHover: false,
		}),
	actions: {
		async setRoutesList(data: Array<string>) {
			this.routesList = data;
		},
		async setKubernetesList(data: Array<string>) {
			this.kubernetesRoutesList = data;
		},
		async setColumnsMenuHover(bool: Boolean) {
			this.isColumnsMenuHover = bool;
		},
		async setColumnsNavHover(bool: Boolean) {
			this.isColumnsNavHover = bool;
		},
	},
	// persist: true,
});
