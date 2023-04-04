import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';
import { V1Namespace, V1NamespaceList } from '@kubernetes/client-node';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
export const kubernetesInfo = defineStore(
	'kubernetesInfo',
	() => {
		const activeCluster = ref<string>('');
		const activeNamespace = ref('default');
		const activeDeployment = ref();
		const clusters = ref();
		const namespace = ref<[]>();
		const setActiveDeployment = async (deploy: any) => {
			activeDeployment.value = deploy;
		};
		const setActiveNamespace = async (ns: string) => {
			activeNamespace.value = ns;
		};
		const setNamespace = async (data: any) => {
			namespace.value = data;
		};
		return { activeCluster, activeNamespace, activeDeployment, clusters,namespace, setActiveDeployment,setActiveNamespace, setNamespace };
	},
	{
		persist: true,
	}
);
