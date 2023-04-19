<template>
	<div class="layout-padding div-container">
		<el-card shadow="hover" class="layout-padding-auto">
			<el-backtop :right="100" :bottom="100" />

			<div>
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
								<Containers :containers="data.deployment.spec.template.spec.containers" @updateContainers="getContainers" />
							</div>
							<div style="margin-top: 10px" id="2" v-show="data.active === 2">
								<h1>asdj</h1>
							</div>
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
import { defineAsyncComponent, onMounted, reactive, watch } from 'vue';

import { ref } from 'vue-demi';
import { Codemirror } from 'vue-codemirror';
import { javascript } from '@codemirror/lang-javascript';
import { oneDark } from '@codemirror/theme-one-dark';
import { V1Container, V1Deployment, V1DeploymentSpec, V1LabelSelector } from '@kubernetes/client-node';
import yaml from 'js-yaml';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { ElMessage } from 'element-plus';
import router from '/@/router';
import { useRoute } from 'vue-router';
import mittBus from '/@/utils/mitt';
import { deepClone } from '/@/utils/other';
const kubeInfo = kubernetesInfo();
const deployApi = useDeploymentApi();

const Meta = defineAsyncComponent(() => import('/@/components/kubernetes/meta.vue'));
const Containers = defineAsyncComponent(() => import('/@/components/kubernetes/containers.vue'));
// 格式化 env
const data = reactive({
	active: 1,
	deployment: {
		apiVersion: 'apps/v1',
		kind: 'Deployment',
		metadata: {
			name: '',
			namespace: 'default',
		},
		spec: {
			selector: {
				matchLabels: {},
			},
			replicas: 1,
			template: {
				metadata: {
					labels: {},
				},
				spec: {
					containers: [],
				},
			},
		},
	},
	code: '',
	// 绑定初始值
	bindMetaData: {
		metadata: {
			namespace: 'default',
			labels: { app: '' },
		} as V1DeploymentSpec,
		replicas: 1,
		resourceType: 'deployment',
	},
});
const extensions = [javascript(), oneDark];
const getContainers = (containers: Array<V1Container>) => {
	data.deployment.spec.template.spec.containers = containers;
	data.code = yaml.dump(data.deployment);
};
const getMeta = (newData) => {
	// console.log('获取到的deployment数据:', newData, data, isObjectValueEqual(data.deployment.metadata, newData.meta));
	// if (!isObjectValueEqual(data.deployment.metadata,newData.meta )  || data.deployment.spec!.replicas != newData.replicas) {
	const dep = deepClone(newData);
	const metaLabels = deepClone(newData);
	data.deployment.metadata = newData.meta;
	data.deployment.spec.selector.matchLabels = dep.meta.labels;
	data.deployment.spec.template.metadata.labels = metaLabels.meta.labels;
	data.deployment.spec!.replicas = newData.replicas;
	data.code = yaml.dump(data.deployment);
	// }
};
const jumpTo = (id) => {
	data.active = id;
	document.getElementById(id).scrollIntoView(true);
};
const next = () => {
	// data.deployment.metadata = metaRef.value.data.meta;
	// data.deployment.spec!.replicas = metaRef.value.data.replicas;
	// data.code = yaml.dump(data.deployment);
	if (data.active++ > 2) data.active = 0;
};

// 定义变量内容
const route = useRoute();

const confirm = () => {
	// data.code = yaml.dump(data.deployment);
	deployApi
		.createDeployment({ cloud: kubeInfo.state.activeCluster }, data.deployment)
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

watch(
	() => data.code,
	(newValue, oldValue) => {
		// console.log("Code ----新的：",newValue, "老的:",oldValue)
		if (newValue) {
			if (newValue != oldValue) {
				const newData = yaml.load(newValue);
				console.log('code变化了，回填数据', newData, 'oldCPde:', oldValue);
				data.bindMetaData.metadata = newData.metadata;
				data.bindMetaData.replicas = newData.spec?.replicas!;
				data.deployment.spec.template.spec.containers = newData.spec.template.spec.containers;
			}
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
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
