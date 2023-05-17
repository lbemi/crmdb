<template>
	<div class="system-user-dialog-container">
		<el-dialog v-model="dialogVisible" width="780px" :title="title" @close="handleClose()">
			<el-form ref="ruleFormRef" :model="data.service" label-width="120px" class="demo-ruleForm" status-icon>
				<el-form-item label="服务名称" prop="name">
					<el-input v-model="data.service.metadata!.name" style="width: 190px" />
					<el-form-item label="命名空间">
						<el-select v-model="data.service.metadata!.namespace" placeholder="指定命名空间">
							<el-option
								v-for="item in k8sStore.state.namespace"
								:key="item.metadata!.name"
								:label="item.metadata!.name"
								:value="item.metadata!.name!"
							/>
						</el-select>
					</el-form-item>
				</el-form-item>
				<el-form-item v-if="data.service.spec?.type === 'ClusterIP'" style="margin-top: 5px">
					<div style="display: flex; justify-items: center">
						<el-checkbox v-model="data.headless" style="margin-right: 5px" />
						<el-tooltip
							class="box-item"
							effect="light"
							content="<div>不为服务分配IP地址,可以通过集群DNS机制机制在集群内部访问服务</div>"
							placement="top-start"
							raw-content
						>
							无头服务(用户实例间服务发现)
						</el-tooltip>
					</div>
				</el-form-item>

				<el-form-item label="类型">
					<el-select v-model="data.service.spec!.type">
						<el-option v-for="item in data.serviceType" :key="item.key" :label="item.key" :value="item.value" />
					</el-select>

					<el-form-item label="负载均衡供应商" v-if="data.service.spec?.type === 'LoadBalance'">
						<el-select v-model="data.loadBalancerProvider">
							<el-option v-for="item in data.loadBalancerProviders" :key="item.key" :label="item.key" :value="item.value" />
						</el-select>
					</el-form-item>
				</el-form-item>
				<!-- <el-form-item label="负载均衡供应商" v-if="data.service.spec?.type === 'LoadBalance'">
					<el-select v-model="data.loadBalancerProvider">
						<el-option v-for="item in data.loadBalancerProviders" :key="item.key" :label="item.key" :value="item.value" />
					</el-select>
				</el-form-item> -->
				<el-form-item label="标签选择器">
					<Label :labelData="data.service.spec!.selector!" @update-labels="getLabels" />
				</el-form-item>

				<el-form-item label="关联后端">
					<el-popover :visible="data.visible" placement="right" :width="400">
						<el-text class="mx-1" type="info">使用工作负载的标签作为选择器</el-text>
						<el-tabs
							v-model="data.activeName"
							type="border-card"
							style="margin-top: 8px; margin-bottom: 8px"
							@tab-change="(TabPaneName) => chanageTab(TabPaneName)"
						>
							<el-tab-pane label="无状态" name="deployment">
								请选择:
								<el-select class="m-2" placeholder="Select" v-model="data.selectWorkLoad">
									<el-option v-for="item in data.deployments" :key="item.metadata!.name" :label="item.metadata!.name" :value="item.metadata!.name!" />
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
				<!-- ServicePort 端口设置 -->
				<ServicePort :ports="data.service.spec?.ports" @updatePort="getServicePorts" />
				<el-form-item label="标签">
					<Label :labelData="data.service.metadata!.labels!" @update-labels="getMetaLabels" />
				</el-form-item>
				<el-form-item label="注解">
					<Label :labelData="data.service.metadata!.annotations!" @update-labels="getAnnotations" />
				</el-form-item>
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
import { V1DaemonSet, V1Deployment, V1ObjectMeta, V1Service, V1StatefulSet } from '@kubernetes/client-node';
import { defineAsyncComponent, reactive, ref, watch } from 'vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { ElMessage } from 'element-plus';

const code = ref('');
const dialogVisible = ref(false);
const k8sStore = kubernetesInfo();
const deploymentApi = useDeploymentApi();

const Label = defineAsyncComponent(() => import('/@/components/kubernetes/labels.vue'));
const ServicePort = defineAsyncComponent(() => import('/@/components/kubernetes/servicePort.vue'));
const handleClose = () => {
	emit('update:dialogVisible', false);
};

const emit = defineEmits(['update', 'update:dialogVisible']);

const update = () => {
	emit('update', code.value);
};

const data = reactive({
	headless: false,
	visible: false,
	deployments: [] as V1Deployment[],
	daemonSets: [] as V1DaemonSet[],
	statefulSets: [] as V1StatefulSet[],
	activeName: 'deployment',
	selectWorkLoad: '',
	service: <V1Service>{
		metadata: {
			name: '',
			namespace: 'default',
			labels: {},
			annotations: {},
		},
		spec: {
			selector: {},
			type: 'ClusterIP',
			ports: [],
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

const selectWorkLoad = () => {
	data.visible = true;
	getDeployments();
};

const getDeployments = () => {
	deploymentApi
		.listDeployment(data.service.metadata!.namespace!, { cloud: k8sStore.state.activeCluster })
		.then((res) => {
			if (res.code === 200) {
				data.deployments = res.data.data;
			}
		})
		.catch(() => {
			ElMessage.error('获取deployment失败');
		});
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
const chanageTab = (chanageTab: string | number) => {
	console.log(chanageTab);
};

const confirm = () => {
	if (data.service.spec?.type === 'ClusterIP' && data.headless) {
		data.service.spec!.clusterIP = 'None';
	}
	console.log(data.service);
};

const props = defineProps({
	title: String,
	codeData: Object,
	dialogVisible: Boolean,
});

watch(
	() => props,
	() => {
		dialogVisible.value = props.dialogVisible;
	},
	{
		immediate: true,
	}
);
</script>

<style scoped></style>
