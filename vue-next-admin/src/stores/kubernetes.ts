import { defineStore } from 'pinia';
import { reactive, } from 'vue';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
export const kubernetesInfo = defineStore(
	'kubernetesInfo',
	() => {
		const state = reactive({
			activeCluster: '',
			activeNamespace: 'default',
			activeDeployment:  '',
			clusters: [],
			namespace: [],
		})
		return {state };
	},
	{
		persist: true,
	}
);
