import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { Pod } from 'kubernetes-types/core/v1';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { usePodApi } from '/@/api/kubernetes/pod';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */

export const podInfo = defineStore(
	'podInfo',
	() => {
		const k8sStore = kubernetesInfo();
		const podApi = usePodApi();
		const state = reactive({
			podDetail: {} as Pod,
			pods: [] as Pod[],
			query: {
				cloud: k8sStore.state.activeCluster,
				page: 1,
				limit: 10,
				key: '',
				type: '1',
			},
			total: 0,
			loading: false,
			selectData: [],
			podShell: {} as Pod,
		});
		const listPod = async () => {
			state.query.cloud = k8sStore.state.activeCluster;
			await podApi.listPods(k8sStore.state.activeNamespace, state.query).then((res) => {
				state.pods = res.data.data;
				state.total = res.data.total;
			});
		};
		const deletePod = async (pod: Pod) => {
			state.query.cloud = k8sStore.state.activeCluster;
			await podApi.deletePod(pod.metadata?.namespace, pod.metadata?.name, { cloud: k8sStore.state.activeCluster });
		};
		const searchPods = async () => {
			state.loading = true;
			if (state.query.key != '') {
				state.query.cloud = k8sStore.state.activeCluster;
				await podApi.searchPods(k8sStore.state.activeNamespace, state.query).then((res) => {
					if (res.code == 200) {
						state.pods = res.data.data;
						state.total = res.data.total;
					}
				});
			} else {
				listPod();
			}
			state.loading = false;
		};
		return { state, listPod, deletePod, searchPods };
	},
	{
		persist: {
			key: 'state',
		},
	}
);
