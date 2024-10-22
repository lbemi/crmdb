import { defineStore } from 'pinia';
import { reactive } from 'vue';
import { DaemonSet, Deployment, StatefulSet } from 'kubernetes-models/apps/v1';
import { Namespace, Service } from 'kubernetes-models/v1';
import { ClusterInfo, Node } from '@/types/kubernetes/cluster';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { ElMessage } from 'element-plus';
import { useServiceApi } from '@/api/kubernetes/service';
import { useDaemonsetApi } from '@/api/kubernetes/daemonset';
import { CronJob, Job } from 'kubernetes-models/batch/v1';
import { useJobApi } from '@/api/kubernetes/job';
import { useCronJobApi } from '@/api/kubernetes/cronjob';
import { useNamespaceApi } from '@/api/kubernetes/namespace';
import jsPlumb from 'jsplumb';
import isEmpty = jsPlumb.jsPlumbUtil.isEmpty;
import { TaskProps } from '@/types/cdk8s-pipelines/lib';

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
			isKubernetesRoutes: false,
			activeNode: {} as Node,
			clusterList: [] as ClusterInfo[],
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
			tekton:{
				updateTask:{} as TaskProps
			}
		});
		const refreshActiveDeployment = async () => {
			if (!isEmpty(state.activeDeployment))
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
			if (!isEmpty(state.activeJob))
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
			if (!isEmpty(state.activeJob))
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
		const refreshActiveDaemonSet = async () => {
			if (!isEmpty(state.activeDeployment))
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
			if (!isEmpty(state.activeService))
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
		const setKubernetesRoutes = (enable: boolean) => {
			state.isKubernetesRoutes = enable;
		};
		return {
			state,
			refreshActiveDeployment,
			refreshActiveService,
			refreshActiveDaemonset: refreshActiveDaemonSet,
			refreshActiveJob,
			refreshActiveCronJob,
			listNamespace,
			setKubernetesRoutes,
		};
	},
	{
		persist: true,
	}
);
