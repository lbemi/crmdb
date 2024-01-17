import { defineStore } from 'pinia';
import { onMounted, reactive } from 'vue';
import { DaemonSet, Deployment, StatefulSet } from 'kubernetes-types/apps/v1';
import { Namespace, Service } from 'kubernetes-types/core/v1';
import { Node } from '../types/kubernetes/cluster';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { ElMessage } from 'element-plus';
import { useServiceApi } from '@/api/kubernetes/service';
import { useDaemonsetApi } from '@/api/kubernetes/daemonset';
import { CronJob, Job } from 'kubernetes-types/batch/v1';
import { useJobApi } from '@/api/kubernetes/job';
import { useCronJobApi } from '@/api/kubernetes/cronjob';
import { useNamespaceApi } from '@/api/kubernetes/namespace';

/**
 * k8s集群信息
 * @methods kubernetesInfo 设置k8s集群信息
 */
const deploymentApi = useDeploymentApi();
const daemonSetApi = useDaemonsetApi();
const serviceApi = useServiceApi();
const jobApi = useJobApi();
const cronJobApi = useCronJobApi();
const namespaceApi = useNamespaceApi();
export const kubernetesInfo = defineStore(
	'kubernetesInfo',
	() => {
		const state = reactive({
			activeNode: {} as Node,
			activeCluster: '',
			activeNamespace: 'default',
			activeDeployment: {} as Deployment,
			activeDaemonSet: {} as DaemonSet,
			activeStatefulSet: {} as StatefulSet,
			activeCronJob: {} as CronJob,
			activeJob: {} as Job,
			activeService: {} as Service,
			search: {
				type: '1',
				value: '',
			},
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
			createStatefulSet: {
				namespace: '',
				name: '',
			},
			createCronJob: {
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
		const refreshActiveJob = async () => {
			if (!isObjectValueEqual(state.activeJob, {}))
				await jobApi
					.getJob(state.activeJob.metadata!.namespace!, state.activeJob.metadata!.name!, {
						cloud: state.activeCluster,
					})
					.then((res: any) => {
						state.activeJob = res.data;
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
		};

		const refreshActiveCronJob = async () => {
			if (!isObjectValueEqual(state.activeJob, {}))
				await cronJobApi
					.getCronJob(state.activeCronJob.metadata!.namespace!, state.activeCronJob.metadata!.name!, {
						cloud: state.activeCluster,
					})
					.then((res: any) => {
						state.activeCronJob = res.data;
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
		};
		const refreshActiveDaemonset = async () => {
			if (!isObjectValueEqual(state.activeDeployment, {}))
				await daemonSetApi
					.getDaemonset(state.activeDaemonSet.metadata!.namespace!, state.activeDaemonSet.metadata!.name!, {
						cloud: state.activeCluster,
					})
					.then((res: any) => {
						state.activeDaemonSet = res.data;
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

		const listNamespace = () => {
			namespaceApi.listNamespace({ cloud: state.activeCluster }).then((res) => {
				state.namespace = res.data.data;
				state.namespaceTotal = res.data.total;
			});
		};
		return { state, refreshActiveDeployment, refreshActiveService, refreshActiveDaemonset, refreshActiveJob, refreshActiveCronJob, listNamespace };
	},
	{
		persist: true,
	}
);
