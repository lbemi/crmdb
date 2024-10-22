import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { Pod } from 'kubernetes-models/v1';
import { kubernetesInfo } from '@/stores/kubernetes';
import { usePodApi } from '@/api/kubernetes/pod';
import { ElMessage } from 'element-plus';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
export const podInfo = defineStore(
	'podInfo',
	() => {
		const k8sStore = kubernetesInfo();
		const podApi = usePodApi();
		const state = reactive({
			podDetail: {} as Pod,
			pods: [] as Array<Pod>,
			query: <queryType>{
				cloud: k8sStore.state.activeCluster,
				page: 1,
				limit: 10,
			},
			total: 0,
			loading: false,
			selectData: [],
			podShell: {} as Pod,
			type: '1',
			inputValue: '',
		});
		const listPod = async () => {
			state.loading = true;
			if (state.type == '1') {
				state.query.name = k8sStore.state.search.value;
				delete state.query.label;
			} else if (state.type == '0') {
				state.query.label = k8sStore.state.search.value;
				delete state.query.name;
			}
			if (state.inputValue === '') {
				delete state.query.label;
				delete state.query.name;
			}
			state.query.cloud = k8sStore.state.activeCluster;

			await podApi
				.listPods(k8sStore.state.activeNamespace, state.query)
				.then((res) => {
					state.pods = res.data.data;
					state.total = res.data.total;
				})
				.catch((e: any) => {
					if (e.code != 5003) ElMessage.error(e.message);
				});
			state.loading = false;
		};
		const refreshPodDetail = async () => {
			await podApi
				.getPod(state.podDetail.metadata?.namespace!, state.podDetail.metadata?.name!, { cloud: k8sStore.state.activeCluster })
				.then((res: any) => {
					state.podDetail = res.data;
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
		};

		const deletePod = async (pod: Pod) => {
			state.query.cloud = k8sStore.state.activeCluster;
			await podApi.deletePod(pod.metadata?.namespace, pod.metadata?.name, { cloud: k8sStore.state.activeCluster });
		};

		return { state, listPod, deletePod, refreshPodDetail };
	},
	{
		persist: {
			key: 'state',
		},
	}
);
