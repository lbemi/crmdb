import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { DaemonSet, Deployment } from 'kubernetes-types/apps/v1';
import { Namespace, Service } from 'kubernetes-types/core/v1';
import { Node } from '../types/kubernetes/cluster';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { ElMessage } from 'element-plus';
import { useServiceApi } from '@/api/kubernetes/service';
import { useDaemonsetApi } from '@/api/kubernetes/daemonset';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
const deploymentApi = useDeploymentApi();
const daemonSetApi = useDaemonsetApi();
const serviceApi = useServiceApi();
export const kubernetesInfo = defineStore(
	'kubernetesInfo',
	() => {
		const state = reactive({
			activeNode: {} as Node,
			activeCluster: '',
			activeNamespace: 'default',
			activeDeployment: {} as Deployment,
			activeDaemonSet: {} as DaemonSet,
			activeService: {} as Service,
			clusters: [],
			namespace: [] as Namespace[],
			namespaceTotal: 0,
			creatDeployment: {
				namespace: '',
				name: '',
			},
			creatDaemonSet: {
				namespace: '',
				name: '',
			},
		});
		const refreshActiveDeployment = async () => {
			if (!isObjectValueEqual(state.activeDeployment, {}))
				await deploymentApi
					.getDeployment(state.activeDeployment.metadata!.namespace!, state.activeDeployment.metadata!.name!, {
						cloud: state.activeCluster,
					})
					.then((res: any) => {
						state.activeDeployment = res.data;
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
		};

		const refreshActiveDaemonset = async () => {
			if (!isObjectValueEqual(state.activeDeployment, {}))
				await daemonSetApi
					.getDaemonset(state.activeDeployment.metadata!.namespace!, state.activeDeployment.metadata!.name!, {
						cloud: state.activeCluster,
					})
					.then((res: any) => {
						state.activeDeployment = res.data;
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
		};
		const refreshActiveService = async () => {
			if (!isObjectValueEqual(state.activeService, {}))
				await serviceApi
					.getService(state.activeService.metadata!.namespace!, state.activeService.metadata!.name!, {
						cloud: state.activeCluster,
					})
					.then((res: any) => {
						state.activeService = res.data;
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
		};

		return { state, refreshActiveDeployment, refreshActiveService, refreshActiveDaemonset };
	},
	{
		persist: true,
	}
);
