import { defineStore } from 'pinia';
import { ref } from 'vue';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
export const kubernetesInfo = defineStore('kubernetesInfo',
	() => {
		const activeCluster = ref<string>('');
		const clusters = ref();
		return { activeCluster, clusters };
	},
	{
		persist: true,
	}
);
