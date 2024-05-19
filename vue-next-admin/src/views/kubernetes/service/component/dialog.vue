<template>
	<div class="system-user-dialog-container">
		<el-dialog v-model="dialogVisible" width="780px" :title="title" @close="handleClose()">
			<el-form ref="ruleFormRef" :model="data.service" label-width="90px" class="demo-ruleForm" status-icon>
				<el-card class="card">
					<el-form-item label="类型" prop="type">
						<el-select v-if="data.service.spec" v-model="data.service.spec.type" size="small">
							<el-option v-for="item in data.serviceType" :key="item.key" :label="item.key" :value="item.value" />
						</el-select>

						<el-form-item label="负载均衡供应商" v-if="data.service.spec?.type === 'LoadBalance'" prop="loadBalancerProvider">
							<el-select v-model="data.loadBalancerProvider">
								<el-option v-for="item in data.loadBalancerProviders" :key="item.key" :label="item.key" :value="item.value" />
							</el-select>
						</el-form-item>
					</el-form-item>
					<el-form-item label="服务名称" prop="name" v-if="data.service.metadata">
						<el-input v-model="data.service.metadata.name" style="width: 190px" :disabled="data.updateFlag" size="small" />
						<el-form-item label="命名空间">
							<el-select
								v-model="data.service.metadata.namespace"
								v-if="k8sStore.state.namespace.length > 0"
								placeholder="指定命名空间"
								size="small"
								:disabled="data.updateFlag"
							>
								<el-option
									v-for="item in k8sStore.state.namespace"
									:key="item.metadata && item.metadata.name"
									:label="item.metadata && item.metadata.name"
									:value="(item.metadata && item.metadata.name) || ''"
								/>
							</el-select>
						</el-form-item>
					</el-form-item>
					<el-form-item v-if="data.service.spec?.type === 'ClusterIP'" style="margin-top: 5px" prop="headless">
						<div style="display: flex; justify-items: center">
							<el-checkbox v-model="data.headless" style="margin-right: 5px" />
							<el-tooltip
								class="box-item"
								effect="light"
								content="<div>不为服务分配IP地址,可以通过集群DNS机制机制在集群内部访问服务</div>"
								placement="right"
								raw-content
							>
								无头服务(用于实例间服务发现)
							</el-tooltip>
						</div>
					</el-form-item>
				</el-card>

				<!-- <el-form-item label="负载均衡供应商" v-if="data.service.spec?.type === 'LoadBalance'">
					<el-select v-model="data.loadBalancerProvider">
						<el-option v-for="item in data.loadBalancerProviders" :key="item.key" :label="item.key" :value="item.value" />
					</el-select>
				</el-form-item> -->

				<el-card class="card">
					<el-form-item label="关联后端">
						<el-popover :visible="data.visible" placement="right" :width="400">
							<el-text class="k-description" type="info">使用工作负载的标签作为选择器</el-text>
							<el-tabs v-model="data.activeName" type="border-card" style="margin-top: 8px; margin-bottom: 8px">
								<el-tab-pane label="无状态" name="deployment">
									请选择:
									<el-select class="m-2" placeholder="Select" v-model="data.selectWorkLoad" size="small">
										<el-option
											v-for="item in data.deployments"
											:key="item.metadata && item.metadata.name"
											:label="item.metadata && item.metadata.name"
											:value="(item.metadata && item.metadata.name) || ''"
										/>
									</el-select>
								</el-tab-pane>
								<el-tab-pane label="守护进程" name="daemonSet">daemonSet</el-tab-pane>
								<el-tab-pane label="有状态" name="statefulSets">statefulSets</el-tab-pane>
							</el-tabs>
							<div style="text-align: right; margin: 0">
								<el-button size="small" text @click="data.visible = false">取消</el-button>
								<el-button size="small" type="primary" @click="confirmSelect">确定</el-button>
							</div>
							<template #reference>
								<el-button type="primary" plain size="small" @click="selectWorkLoad">选择后端</el-button>
							</template>
						</el-popover>
					</el-form-item>
					<!--					<el-form-item label="标签选择器" prop="selector">-->
					<Label v-if="data.service.spec" :label="'标签选择器'" :labels="data.service.spec.selector" @update-labels="getLabels" />
					<!--					</el-form-item>-->
				</el-card>

				<el-card class="card">
					<el-form-item label="会话保持" style="margin-top: 5px" prop="keepAlive">
						<el-checkbox v-model="data.keepAlive" style="margin-right: 150px" />
					</el-form-item>
					<el-form-item label="最长回话保持时间" label-width="150" v-if="data.keepAlive" prop="keepAliveTime">
						<el-input v-model.number="data.keepAliveTime" style="width: 190px" />
					</el-form-item>
					<div v-if="data.keepAlive" style="font-size: 8px; margin-left: 150px">设置最大会话保持时间。取值范围为 0 到 86400，默认值 10800。</div>
				</el-card>

				<el-card class="card">
					<!-- ServicePort 端口设置 -->
					<ServicePort :ports="data.service.spec?.ports" @updatePort="getServicePorts" />
				</el-card>
				<el-card class="card" v-if="data.service.metadata">
					<Label :label="'标签'" :labels="data.service.metadata.labels" :name="'标签'" @update-labels="getMetaLabels" />
					<Label label="'注解'" :labels="data.service.metadata.annotations" @update-labels="getAnnotations" />
				</el-card>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button size="small" @click="handleClose">关闭</el-button>
					<el-button type="primary" size="small" @click="confirm">确定</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { Service } from 'kubernetes-models/v1';
import { DaemonSet, Deployment, StatefulSet } from 'kubernetes-models/apps/v1';
import { defineAsyncComponent, reactive, ref, watch } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { ElMessage } from 'element-plus';
import { useServiceApi } from '@/api/kubernetes/service';
import { isObjectValueEqual } from '@/utils/arrayOperation';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));
const ServicePort = defineAsyncComponent(() => import('@/components/kubernetes/servicePort.vue'));

const dialogVisible = ref(false);
const k8sStore = kubernetesInfo();
const deploymentApi = useDeploymentApi();
const serviceApi = useServiceApi();
const ruleFormRef = ref();

const emit = defineEmits(['update', 'update:dialogVisible', 'refresh']);

const data = reactive({
	updateFlag: false,
	keepAliveTime: 10800,
	keepAlive: false,
	headless: false,
	visible: false,
	deployments: [] as Deployment[],
	daemonSets: [] as DaemonSet[],
	statefulSets: [] as StatefulSet[],
	activeName: 'deployment',
	selectWorkLoad: '',
	service: <Service>{
		metadata: {
			name: '',
			namespace: 'default',
			params: {},
			annotations: {},
		},
		spec: {
			selector: {},
			type: 'ClusterIP',
			ports: [],
			sessionAffinity: 'None',
		},
	},
	loadBalancerProvider: 'aliyun',
	loadBalancerProviders: [
		{ key: '阿里云', value: 'aliyun' },
		{ key: '其他', value: 'other' },
	],
	serviceType: [
		{ key: '虚拟IP', value: 'ClusterIP' },
		{ key: '节点端口', value: 'NodePort' },
		{ key: '负载均衡', value: 'LoadBalance' },
	],
});

const handleClose = () => {
	emit('update:dialogVisible', false);
	ruleFormRef.value.resetFields();
};

const selectWorkLoad = () => {
	data.visible = true;
	getDeployments();
};

const getDeployments = () => {
	deploymentApi
		.listDeployment(data.service.metadata!.namespace!, { cloud: k8sStore.state.activeCluster })
		.then((res: any) => {
			data.deployments = res.data.data;
		})
		.catch(() => {
			ElMessage.error('获取deployment失败');
		});
};

const createService = () => {
	serviceApi
		.createService({ cloud: k8sStore.state.activeCluster }, data.service)
		.then((res: any) => {
			if (res.code == 200) {
				ElMessage.success('创建成功');
			}
		})
		.catch((e: any) => {
			ElMessage.error(e);
		});
	handleClose();
	emit('refresh');
};
const updateService = () => {
	serviceApi
		.updateService({ cloud: k8sStore.state.activeCluster }, data.service)
		.then((res: any) => {
			ElMessage.success(res.message);
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	handleClose();
	emit('refresh');
};
const confirmSelect = () => {
	if (data.activeName === 'deployment') {
		let res = data.deployments.filter((item) => item.metadata!.name === data.selectWorkLoad);

		if (res.length === 1) data.service.spec!.selector = res[0].metadata?.labels;
	}

	data.visible = false;
};

const getLabels = (labels: any) => {
	data.service.spec!.selector = labels;
};

const getMetaLabels = (labels: any) => {
	data.service.metadata!.labels = labels;
};
const getAnnotations = (labels: any) => {
	data.service.metadata!.annotations = labels;
};
const getServicePorts = (ports: any) => {
	data.service.spec!.ports = ports;
};

const confirm = () => {
	if (data.service.spec?.type === 'ClusterIP' && data.headless) {
		data.service.spec!.clusterIP = 'None';
	} else {
		delete data.service.spec?.clusterIP;
	}
	if (data.keepAlive) {
		data.service.spec!.sessionAffinity = 'ClientIP';
		data.service.spec!.sessionAffinityConfig = {
			clientIP: {
				timeoutSeconds: data.keepAliveTime,
			},
		};
	} else {
		delete data.service.spec?.sessionAffinity;
		delete data.service.spec?.sessionAffinityConfig;
	}
	if (data.updateFlag) {
		updateService();
	} else {
		createService();
	}
};

const props = defineProps({
	title: String,
	codeData: Object,
	dialogVisible: Boolean,
	service: Object,
});

watch(
	() => props,
	() => {
		dialogVisible.value = props.dialogVisible;

		if (!isObjectValueEqual(props.service, {})) {
			data.service = props.service as Service;
			data.updateFlag = true;
		}
	},
	{
		immediate: true,
	}
);
</script>

<style scoped>
.card {
	margin-bottom: 10px;
}
</style>
