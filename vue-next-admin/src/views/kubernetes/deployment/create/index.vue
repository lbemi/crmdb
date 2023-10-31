<template>
	<div class="layout-padding div-container">
		<el-card shadow="hover" class="layout-padding-auto">
			<el-backtop :right="100" :bottom="100" />

			<div>
				<el-button type="primary" @click="edit">编辑</el-button>
				<el-steps :active="data.active" finish-status="success" simple>
					<el-step title="基本信息" description="Some description" />
					<el-step title="容器配置" description="Some description" />
					<el-step title="高级配置" description="Some description" />
				</el-steps>
			</div>
			<el-row>
				<el-col :span="15">
					<el-card style="padding-left: 20px; margin-top: 15px">
						<div>
							<div style="margin-top: 10px" id="0" v-show="data.active === 0">
								<Meta :bindData="data.bindMetaData" @updateData="getMeta" />
							</div>
							<div style="margin-top: 10px" id="1" v-show="data.active === 1">
								<Containers
									:containers="deepClone(data.deployments.spec!.template.spec!.containers) as Array< Container>"
									@updateContainers="getContainers"
								/>
							</div>
							<div style="margin-top: 10px" id="2" v-show="data.active === 2"></div>
						</div>
					</el-card>
				</el-col>
				<!--				<el-col :span="1" />-->
				<el-col :span="7">
					<el-card style="margin-top: 15px; height: 99%">
						<codemirror v-model="data.code" style="margin-top: 15px" :autofocus="true" :tabSize="2" :extensions="extensions" />
					</el-card>
				</el-col>
				<el-col :span="2" style="margin-left: 20px">
					<div class="btn">
						<div>
							<el-link type="primary" :underline="false" @click="jumpTo(0)" class="men">基础信息</el-link>
						</div>
						<div>
							<el-link type="primary" :underline="false" @click="jumpTo(1)" class="men">容器配置</el-link>
						</div>
						<div>
							<el-link type="primary" :underline="false" @click="jumpTo(2)" class="men">高级配置</el-link>
						</div>
						<el-button @click="next" style="margin-top: 5px" size="small">下一步</el-button>
						<el-button @click="confirm" style="margin-top: 5px" size="small">确认</el-button>
					</div>
				</el-col>
			</el-row>
		</el-card>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onBeforeMount, onUnmounted, reactive, ref, watch } from 'vue';
import { Codemirror } from 'vue-codemirror';
import { oneDark } from '@codemirror/theme-one-dark';
import { Container } from 'kubernetes-types/core/v1';
import { Deployment } from 'kubernetes-types/apps/v1';
import yamlJs from 'js-yaml';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage } from 'element-plus';
import router from '@/router';
import { useRoute } from 'vue-router';
import mittBus from '@/utils/mitt';
import { deepClone } from '@/utils/other';
import { CreateK8SBindData, CreateK8SMetaData } from '@/types/kubernetes/custom';
import type { FormInstance } from 'element-plus';
import { StreamLanguage } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';

const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
const Containers = defineAsyncComponent(() => import('@/components/kubernetes/containers.vue'));

const kubeInfo = kubernetesInfo();
const deployApi = useDeploymentApi();

const edit = () => {
	const dep = deepClone(kubeInfo.state.activeDeployment);
	delete dep.metadata?.resourceVersion;
	delete dep.metadata?.managedFields;
	delete dep.status;
	data.code = yamlJs.dump(dep);
};
const metaRef = ref<FormInstance>();
const data = reactive({
	loadCode: false,
	active: 0,
	//初始化deployment
	deployments: <Deployment>{
		apiVersion: 'apps/v1',
		kind: 'Deployment',
		metadata: {
			// name: '',
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
	bindMetaData: <CreateK8SBindData>{
		resourceType: 'deployment',
	},
});
const extensions = [oneDark, StreamLanguage.define(yaml)];
const getContainers = (containers: Array<Container>) => {
	data.deployments.spec!.template.spec!.containers = containers;
	updateCodeMirror();
};

const getMeta = (newData: CreateK8SMetaData, metaRefs: FormInstance) => {
	metaRef.value = metaRefs;
	const dep = deepClone(newData);
	const metaLabels = deepClone(newData);
	data.deployments.metadata = newData.meta;
	//更新labels
	if (dep.meta.name) data.deployments.metadata!.labels!.app = dep.meta.name;
	//更新selector.matchLabels
	data.deployments.spec!.selector.matchLabels = dep.meta.labels;
	data.deployments.spec!.template.metadata!.labels = metaLabels.meta.labels;
	data.deployments.spec!.replicas = newData.replicas;
	updateCodeMirror();
};
const nextStep = (formEl: FormInstance | undefined) => {
	if (!formEl) {
		ElMessage.error('请输入必填项');
		return;
	}
	formEl.validate((valid) => {
		if (valid) {
			if (data.active++ > 2) data.active = 0;
		} else {
			ElMessage.error('请输检查字段');
		}
	});
};
const jumpTo = (id: number) => {
	data.active = id;

	document.getElementById(id + '')!.scrollIntoView(true);
};
const next = () => {
	// data.deployment.metadata = metaRef.value.data.meta;
	// data.deployment.spec!.replicas = metaRef.value.data.replicas;
	// data.code = yaml.dump(data.deployment);
	nextStep(metaRef.value);
	// if (data.active === 0) {
	// 	if (nextStep(metaRef.value)) {
	// 		data.active += 1;
	// if (data.active++ > 2) data.active = 0;
	// 	}
	// }

	// if (data.active++ > 2) data.active = 0;
};

// 定义变量内容
const route = useRoute();
mittBus.on('updateVolumes', (res) => {
	data.deployments.spec!.template.spec!.volumes = res;
	updateCodeMirror();
});

const confirm = () => {
	// data.code = yaml.dump(data.deployment);
	deployApi
		.createDeployment({ cloud: kubeInfo.state.activeCluster }, data.deployments)
		.then(() => {
			router.push({
				name: 'k8sDeployment',
			});
			mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));

			ElMessage.success('创建成功');
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
};
const updateCodeMirror = () => {
	data.loadCode = true;
	data.code = yamlJs.dump(data.deployments);
	setTimeout(() => {
		data.loadCode = false;
	}, 1);
};
watch(
	() => data.code,
	(newValue, oldValue) => {
		if (newValue && !data.loadCode) {
			if (newValue != oldValue) {
				const newData = yamlJs.load(newValue) as Deployment;
				if (typeof newData === 'object' && newData != null) {
					data.bindMetaData.metadata = newData.metadata!;
					data.bindMetaData.replicas = newData.spec?.replicas!;
					data.deployments = newData;
					mittBus.emit('updateDeploymentVolumes', newData.spec!.template.spec!.volumes);
					//重新更新一下关联字段，并更新cide
					data.deployments.metadata!.labels!.app = newData.metadata!.name!;
					data.deployments.spec!.selector.matchLabels = deepClone(newData).metadata!.labels;
					data.deployments.spec!.template.metadata!.labels = deepClone(newData).metadata!.labels;
					updateCodeMirror();
				}
			}
		}
	},
	{
		deep: true,
	}
);

onBeforeMount(() => {
	updateCodeMirror();
});
onUnmounted(() => {
	//卸载
	mittBus.off('updateVolumes', () => {});
});
</script>

<style scoped lang="scss">
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
.div-container {
	:deep(.el-card__body) {
		display: flex;
		flex-direction: column;
		flex: 1;
		overflow: auto;
		.el-table {
			flex: 1;
		}
	}
}
</style>
