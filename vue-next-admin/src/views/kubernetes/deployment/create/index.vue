<template>
	<div class="layout-pd">
		<el-card shadow="hover">
			<template #header>
				<div class="card-header">
					<span>创建deployment</span>
					<el-button type="primary" :icon="View" size="small" @click="showYaml">预览YAML</el-button>
				</div>
			</template>
			<el-row :gutter="30">
				<el-col :span="3" :offset="1" style="margin-top: 15px">
					<el-steps :active="data.active" finish-status="success" direction="vertical">
						<el-step title="基本信息" description="基础配置信息" />
						<el-step title="容器配置" description="容器相关配置信息" />
						<el-step title="高级配置" description="service高级配置信息" />
					</el-steps>
				</el-col>
				<el-col :span="20">
					<div id="0" v-show="data.active === 0">
						<Meta ref="metaRef" :bindData="data.bindMetaData" :isUpdate="data.isUpdate" :label-width="'90px'" />
					</div>
					<div id="1" v-show="data.active === 1">
						<Containers
							ref="containersRef"
							:containers="data.containers"
							:initContainers="data.initContainers"
							:volumes="data.deployment.spec!.template.spec!.volumes"
						/>
					</div>
					<div id="2" v-show="data.active === 2">
						<el-checkbox v-model="data.enableService" label="配置service" />
					</div>
				</el-col>
			</el-row>
			<div class="footer">
				<el-button @click="up" size="small">上一步</el-button>
				<el-button @click="next" size="small">下一步</el-button>
				<el-button @click="confirm" type="primary" size="small">确认</el-button>
			</div>
		</el-card>
		<YamlDialog v-model:dialogVisible="data.yamlDialogVisible" :code-data="data.deployment" v-if="data.yamlDialogVisible" />
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onBeforeMount, reactive, ref } from 'vue';
import { Container } from 'kubernetes-models/v1';
import { Deployment } from 'kubernetes-models/apps/v1';
import yamlJs from 'js-yaml';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage } from 'element-plus';
import { View } from '@element-plus/icons-vue';
import { deepClone } from '@/utils/other';
import { CreateK8SMeta } from '@/types/kubernetes/custom';
import type { FormInstance } from 'element-plus';
import { useDeploymentApi } from '@/api/kubernetes/deployment';

const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
const Containers = defineAsyncComponent(() => import('@/components/kubernetes/containers.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));

const containersRef = ref();
const kubeInfo = kubernetesInfo();
const deploymentApi = useDeploymentApi();
const metaRef = ref();

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
					initContainers: [],
					containers: [],
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
	bindMetaData: <CreateK8SMeta>{
		resourceType: 'deployment',
	},
	validateRef: <Array<FormInstance>>[],
});

const showYaml = async () => {
	getContainers();
	data.yamlDialogVisible = true;
};

const getContainers = () => {
	// delete data.deployment.spec!.template.spec!.initContainers;
	const { containers, initContainers, volumes, validateRefs } = containersRef.value.returnContainers();

	if (validateRefs.length > 0) {
		data.validateRef.push(...validateRefs);
	}

	if (volumes.length > 0) {
		data.deployment.spec!.template.spec!.volumes = volumes;
	}
	if (containers.length > 0) {
		data.deployment.spec!.template.spec!.containers = containers;
	}
	if (initContainers.length > 0) {
		data.deployment.spec!.template.spec!.initContainers = initContainers;
	}
};

const getMeta = () => {
	const metaData = deepClone(metaRef.value.getMeta());

	if (metaData.validateRefs.length > 0) {
		data.validateRef.push(...metaData.validateRefs);
	}
	data.deployment.metadata = deepClone(metaData.meta);
	//更新labels
	data.deployment.metadata!.labels!.app = metaData.meta.name;
	//更新selector.matchLabels
	data.deployment.spec!.selector.matchLabels = deepClone(metaData.meta.labels);
	data.deployment.spec!.template.metadata!.labels = deepClone(metaData.meta.labels);
	data.deployment.spec!.replicas = metaData.replicas;
	updateCodeMirror();
};
const validate = async () => {
	data.validateRef = [];
	if (data.active == 0) {
		getMeta();
	}

	if (data.active == 1) {
		getContainers();
	}

	try {
		for (const item of data.validateRef) {
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
	if (data.active == 0) {
		if (!(await validate())) return;
	}
	if (data.active == 1) {
		if (!(await validate())) return;
	}
	if (data.active++ > 2) data.active = 0;
};

const up = () => {
	if (data.active-- == 0) data.active = 0;
};
const next = () => {
	nextStep();
};

const confirm = async () => {
	if (!(await validate())) return;
	// 校验deployment是否有问题
	try {
		data.deployment.validate();
		deploymentApi
			.createDeployment({ cloud: kubeInfo.state.activeCluster }, data.deployment)
			.then(() => {
				ElMessage.success('创建成功');
				// handleClose();
			})
			.catch((e) => {
				ElMessage.error(e.message);
				// handleClose();
			});
	} catch (error) {
		ElMessage.error((error as Error).message);
	}
	// } else {
	// 	await deploymentApi
	// 		.updateDeployment(data.deployments, { cloud: kubeInfo.state.activeCluster })
	// 		.then(() => {
	// 			ElMessage.success('更新成功');
	// 			handleClose();
	// 		})
	// 		.catch((e) => {
	// 			ElMessage.error(e.message);
	// 		});
	// }
};
const updateCodeMirror = () => {
	data.loadCode = true;
	data.code = yamlJs.dump(data.deployment);
	setTimeout(() => {
		data.loadCode = false;
	}, 1);
};

onBeforeMount(() => {
	updateCodeMirror();
});
</script>

<style scoped>
.footer {
	margin-top: 30px;
	display: flex;
	justify-content: center;
}

.card-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.text {
	font-size: 14px;
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
</style>
