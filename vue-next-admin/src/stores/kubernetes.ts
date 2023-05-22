import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { Deployment } from 'kubernetes-types/apps/v1';
import { Namespace, Service } from 'kubernetes-types/core/v1';
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
			activeDeployment: {} as Deployment,
			activeService: {} as Service,
			clusters: [],
			namespace: [] as Namespace[],
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
