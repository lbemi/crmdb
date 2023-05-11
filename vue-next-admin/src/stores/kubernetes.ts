import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { V1Deployment, V1Namespace, V1Node, V1Service } from '@kubernetes/client-node';
import { Node } from '../types/kubernetes/cluster';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
export const kubernetesInfo = defineStore(
	'kubernetesInfo',
	() => {
		const state = reactive({
			activeNode: {} as Node,
			activeCluster: '',
			activeNamespace: 'default',
			activeDeployment: {} as V1Deployment,
			activeService: {} as V1Service,
			clusters: [],
			namespace: [] as V1Namespace[],
			namespaceTotal: 0,
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
