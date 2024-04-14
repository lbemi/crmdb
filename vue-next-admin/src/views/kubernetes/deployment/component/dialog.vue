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
				<el-steps :active="state.active" finish-status="success" align-center>
					<el-step title="基本信息" />
					<el-step title="容器配置" />
					<el-step title="高级配置" />
				</el-steps>
			</div>
			<div>
				<div style="margin-top: 10px" id="0" v-show="state.active === 0">
					<el-card>
						<Meta ref="metaRef" :metaData="state.metaData" :isUpdate="state.isUpdate" />
					</el-card>
				</div>
				<div style="margin-top: 10px" id="1" v-show="state.active === 1">
					<Containers
						ref="containersRef"
						:containers="state.containers"
						:initContainers="state.initContainers"
						:volumes="state.deployment.spec?.template.spec?.volumes"
					/>
				</div>
				<div style="margin-top: 10px" id="2" v-show="state.active === 2">
					<el-checkbox v-model="state.enableService" label="配置service" />
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
		<YamlDialog v-model:dialogVisible="state.yamlDialogVisible" :code-data="state.deployment" v-if="state.yamlDialogVisible" />
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

const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
const Containers = defineAsyncComponent(() => import('@/components/kubernetes/containers.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));

const dialogVisible = ref(false);
const containersRef = ref();
const kubeInfo = kubernetesInfo();
const deploymentApi = useDeploymentApi();
const metaRef = ref<FormInstance>();

const state = reactive({
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
	deployment: new Deployment({
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
					containers: [] as Container[],
					volumes: [],
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
	}),
	code: '',
	// 绑定初始值
	metaData: <CreateK8SMeta>{
		resourceType: 'deployment',
	},
	validateRef: <Array<FormInstance>>[],
});

const showYaml = async () => {
	getContainers();
	state.yamlDialogVisible = true;
};

const getContainers = () => {
	const { containers, initContainers, volumes, validateRefs } = containersRef.value.returnContainers();

	if (validateRefs.length > 0) {
		state.validateRef.push(...validateRefs);
	}

	if (volumes.length > 0) {
		state.deployment.spec!.template.spec!.volumes = volumes;
	}
	if (containers.length > 0) {
		state.deployment.spec!.template.spec!.containers = containers;
	}
	if (initContainers.length > 0) {
		state.deployment.spec!.template.spec!.initContainers = initContainers;
	}
};

const getMeta = () => {
	const metaData = deepClone(metaRef.value.getMeta());

	if (metaData.validateRefs.length > 0) {
		state.validateRef.push(...metaData.validateRefs);
	}
	state.deployment.metadata = deepClone(metaData.meta);
	//更新labels
	state.deployment.metadata!.labels!.app = metaData.meta.name;
	//更新selector.matchLabels
	state.deployment.spec!.selector.matchLabels = deepClone(metaData.meta.labels);
	state.deployment.spec!.template.metadata!.labels = deepClone(metaData.meta.labels);
	state.deployment.spec!.replicas = metaData.replicas;
	updateCodeMirror();
};
const validate = async () => {
	state.validateRef = [];
	if (state.active == 0) {
		getMeta();
	}

	if (state.active == 1) {
		getContainers();
	}

	try {
		for (const item of state.validateRef) {
			// 使用 Promise.all 来等待所有表单验证完成
			await Promise.all([item.validate()]);
		}

		return true;
	} catch (error) {
		// 如果有表单验证不通过，则返回 false
		return false;
	}
};
const nextStep = async () => {
	if (state.active == 0) {
		if (!(await validate())) return;
	}
	if (state.active == 1) {
		if (!(await validate())) return;
	}
	if (state.active++ > 2) state.active = 0;
};
const up = () => {
	if (state.active-- == 0) state.active = 0;
};
const next = async () => {
	if (!(await validate())) return;
	nextStep();
};

const confirm = async () => {
	// data.code = yaml.dump(data.deployment);
	getContainers();
	if (props.title === '创建deployment') {
		deploymentApi
			.createDeployment({ cloud: kubeInfo.state.activeCluster }, state.deployment)
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
			.updateDeployment(state.deployment, { cloud: kubeInfo.state.activeCluster })
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
	state.loadCode = true;
	state.code = yamlJs.dump(state.deployment);
	setTimeout(() => {
		state.loadCode = false;
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
	if (props.deployment) {
		state.isUpdate = true;
		state.deployment = deepClone(props.deployment) as Deployment;
		state.metaData.metadata = state.deployment.metadata;
		state.metaData.replicas = state.deployment.spec?.replicas;

		if (state.deployment.spec?.template.spec?.initContainers) {
			state.initContainers = state.deployment.spec!.template.spec!.initContainers!;
		}
		if (state.deployment.spec?.template.spec?.containers) {
			state.containers = state.deployment.spec!.template.spec!.containers!;
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
