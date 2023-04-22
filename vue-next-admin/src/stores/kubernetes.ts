import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { V1Deployment, V1Namespace } from '@kubernetes/client-node';

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
			activeDeployment: {} as V1Deployment,
			clusters: [],
			namespace: [] as V1Namespace[],
			creatDeployment: {
				namespace: '',
				name: '',
			},
		});
		return { state };
	},
	{
		persist: true,
	}
);
