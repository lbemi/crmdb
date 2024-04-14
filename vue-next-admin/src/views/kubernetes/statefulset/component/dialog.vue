<template>
	<div class="system-user-dialog-container">
		<el-dialog class="dialog" v-model="dialogVisible" width="1000px" @close="handleClose()">
			<template #header>
				<div style="display: flex; justify-content: space-between">
					{{ title }}
					<el-button type="primary" @click="showYaml" size="small" :icon="View" style="margin-right: 20px">YAML</el-button>
				</div>
				<el-divider style="margin: 8px 0" />
			</template>
			<el-backtop :right="100" :bottom="100" />
			<div>
				<el-steps :active="data.active" finish-status="success" align-center>
					<el-step title="基本信息" />
					<el-step title="容器配置" />
					<el-step title="高级配置" />
				</el-steps>
			</div>
			<div>
				<div style="margin-top: 10px" id="0" v-show="data.active === 0">
					<el-card>
						<Meta :bindData="data.bindMetaData" :isUpdate="data.isUpdate" @updateData="getMeta" />
					</el-card>
				</div>
				<div style="margin-top: 10px" id="1" v-show="data.active === 1">
					<Containers
						ref="containersRef"
						:containers="data.containers"
						:initContainers="data.initContainers"
						:volumes="data.deployments.spec?.template.spec?.volumes"
					/>
				</div>
				<div style="margin-top: 10px" id="2" v-show="data.active === 2">
					<el-checkbox v-model="data.enableService" label="配置service" />
				</div>
			</div>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="up" size="small">上一步</el-button>
					<el-button @click="next" size="small">下一步</el-button>
					<el-button @click="confirm" type="success" size="small">确认</el-button>
					<el-button size="small" type="primary" @click="handleClose">关闭</el-button>
				</span>
			</template>
		</el-dialog>
		<YamlDialog v-model:dialogVisible="data.yamlDialogVisible" :code-data="data.deployments" v-if="data.yamlDialogVisible" />
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onBeforeMount, onMounted, reactive, ref } from 'vue';
import { Container } from 'kubernetes-models/v1';
import { Deployment } from 'kubernetes-models/apps/v1';
import yamlJs from 'js-yaml';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage } from 'element-plus';
import { View } from '@element-plus/icons-vue';
import { deepClone } from '@/utils/other';
import { CreateK8SMeta, CreateK8SMetaData } from '@/types/kubernetes/custom';
import type { FormInstance } from 'element-plus';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { isObjectValueEqual } from '@/utils/arrayOperation';

const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
const Containers = defineAsyncComponent(() => import('@/components/kubernetes/containers.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));

const dialogVisible = ref(false);
const containersRef = ref();
const kubeInfo = kubernetesInfo();
const deploymentApi = useDeploymentApi();
const metaRef = ref<FormInstance>();

const data = reactive({
	isUpdate: false,
	enableService: false,
	yamlDialogVisible: false,
	dialogVisible: false,
	codeData: {} as Deployment,
	loadCode: false,
	active: 0,
	containers: [] as Container[],
	initContainers: [] as Container[],
	//初始化deployment
	deployments: <Deployment>{
		apiVersion: 'apps/v1',
		kind: 'Deployment',
		metadata: {
			namespace: 'default',
		},
		spec: {
			replicas: 1,
			selector: {
				matchLabels: {},
			},
			template: {
				metadata: {
					labels: {},
				},
				spec: {
					serviceAccount: 'default',
					initContainers: [] as Container[],
					containers: [],
					// volumes: [],
				},
			},
			strategy: {
				type: 'RollingUpdate',
				rollingUpdate: {
					maxUnavailable: '25%',
					maxSurge: '25%',
				},
			},
		},
	},
	code: '',
	// 绑定初始值
	bindMetaData: <CreateK8SMeta>{
		resourceType: 'deployment',
	},
});

const showYaml = async () => {
	getContainers();
	data.yamlDialogVisible = true;
};

const getContainers = () => {
	delete data.deployments.spec!.template.spec!.containers;
	delete data.deployments.spec!.template.spec!.initContainers;
	const { containers, initContainers, volumes } = containersRef.value.returnContainers();
	if (volumes.length > 0) {
		data.deployments.spec!.template.spec!.volumes = volumes;
	}
	if (containers.length > 0) {
		data.deployments.spec!.template.spec!.containers = containers;
	}
	if (initContainers.length > 0) {
		data.deployments.spec!.template.spec!.initContainers = initContainers;
	}
};

const getMeta = (newData: CreateK8SMetaData, metaRefs: FormInstance) => {
	metaRef.value = metaRefs;
	const dep = deepClone(newData);
	const metaLabels = deepClone(newData);
	data.deployments.metadata = newData.meta;
	//更新labels
	if (!data.isUpdate) {
		if (dep.meta.name) data.deployments.metadata!.labels!.app = dep.meta.name;
	}
	//更新selector.matchLabels
	data.deployments.spec!.selector.matchLabels = dep.meta.labels;
	data.deployments.spec!.template.metadata!.labels = metaLabels.meta.labels;
	data.deployments.spec!.replicas = newData.replicas;
	updateCodeMirror();
};
const nextStep = () => {
	if (data.active++ > 2) data.active = 0;
};
const up = () => {
	if (data.active-- == 0) data.active = 0;
};
const next = () => {
	nextStep();
};

const confirm = async () => {
	// data.code = yaml.dump(data.deployment);
	getContainers();
	if (props.title === '创建deployment') {
		deploymentApi
			.createDeployment({ cloud: kubeInfo.state.activeCluster }, data.deployments)
			.then(() => {
				ElMessage.success('创建成功');
				handleClose();
			})
			.catch((e) => {
				ElMessage.error(e.message);
				// handleClose();
			});
	} else {
		await deploymentApi
			.updateDeployment(data.deployments, { cloud: kubeInfo.state.activeCluster })
			.then(() => {
				ElMessage.success('更新成功');
				handleClose();
			})
			.catch((e) => {
				ElMessage.error(e.message);
			});
	}
};
const updateCodeMirror = () => {
	data.loadCode = true;
	data.code = yamlJs.dump(data.deployments);
	setTimeout(() => {
		data.loadCode = false;
	}, 1);
};

onBeforeMount(() => {
	updateCodeMirror();
});

const emit = defineEmits(['update:dialogVisible', 'refresh']);

const handleClose = () => {
	emit('update:dialogVisible', false);
	emit('refresh');
};

const props = defineProps({
	title: String,
	dialogVisible: Boolean,
	deployment: Object,
});

onMounted(() => {
	dialogVisible.value = props.dialogVisible;
	if (!isObjectValueEqual(props.deployment, {})) {
		data.isUpdate = true;
		data.deployments = props.deployment as Deployment;
		data.bindMetaData.metadata = data.deployments.metadata;
		data.bindMetaData.replicas = data.deployments.spec?.replicas;

		if (data.deployments.spec?.template.spec?.initContainers) {
			data.initContainers = data.deployments.spec!.template.spec!.initContainers!;
		}
		if (data.deployments.spec?.template.spec?.containers) {
			data.containers = data.deployments.spec!.template.spec!.containers!;
		}
	}
});
</script>

<style scoped>
.card {
	margin-bottom: 10px;
}

.d2 {
	min-width: 100%;
	height: 100%;
	position: relative;
	display: flex;
	justify-content: flex-end;
}

.btn {
	position: fixed;
	right: 50px;
	text-align: center;
	top: 50%;
}

.men {
	font-size: 13px;
	letter-spacing: 3px;
}

.el-form-item {
	margin-bottom: 2px;
}

.el-table-column {
	padding-top: 2px;
	padding-bottom: 2px;
}

.dialog .el-dialog__body {
	--el-dialog-padding-primary: -10px;
}
</style>
