<template>
	<div class="layout-pd">
		<el-card shadow="hover">
			<el-backtop :right="100" :bottom="100" />
			<div style="height: 100%">
				<el-steps :active="data.active" finish-status="success" simple>
					<el-step title="基本信息" description="Some description" />
					<el-step title="容器配置" description="Some description" />
					<el-step title="高级配置" description="Some description" />
				</el-steps>
			</div>
			<el-card style="height: 100%">
				<el-row>
					<el-col :span="1" />
					<el-col :span="15">
						<div>
							<div style="margin-top: 10px" id="0" v-show="data.active === 0">
								<Meta ref="metaRef" v-bind:metadata="data.tableData" v-bind:resourceType="'deployment'" />
							</div>
							<div style="margin-top: 10px" id="1" v-show="data.active === 1">
								<Containers ref="containersRef" />
							</div>
							<div style="margin-top: 10px" id="2" v-show="data.active === 2">
								<h1>asdj</h1>
							</div>
						</div>
					</el-col>
					<el-col :span="1" />
					<el-col :span="6">
						<codemirror v-model="data.code" :style="{ height: '100%' }" :autofocus="true" :tabSize="2" :extensions="extensions" />
					</el-col>
					<el-col :span="3" style="margin-left: 20px">
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
		</el-card>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive } from 'vue';

import { ref } from 'vue-demi';
import { Codemirror } from 'vue-codemirror';
import { javascript } from '@codemirror/lang-javascript';
import { oneDark } from '@codemirror/theme-one-dark';
import { V1Container, V1Deployment, V1DeploymentSpec, V1EnvVar, V1PodTemplate } from '@kubernetes/client-node';
import yaml from 'js-yaml';

const Meta = defineAsyncComponent(() => import('/@/components/kubernetes/meta.vue'));
const Containers = defineAsyncComponent(() => import('/@/components/kubernetes/containers.vue'));
const metaRef = ref<InstanceType<typeof Meta>>();
const containersRef = ref<InstanceType<typeof Containers>>();

// 格式化 env
const data = reactive({
	active: 0,
	deployment: {
		apiVersion: 'apps/v1',
		kind: 'Deployment',
		metadata: {
			namespace: 'default',
		},
		spec: {
			replicas: 1,
			template: {
				spec: {
					container: [],
				},
			},
		},
	} as V1Deployment,
	code: '',
	tableData: [],
});

const extensions = [javascript(), oneDark];
const jumpTo = (id) => {
	data.active = id;
	document.getElementById(id).scrollIntoView(true);
};
const next = () => {
	data.deployment.metadata = metaRef.value.data.meta;
	data.deployment.spec!.replicas = metaRef.value.data.replicas;
	data.code = yaml.dump(data.deployment);
	if (data.active++ > 2) data.active = 0;
};
const confirm = () => {
	data.deployment.metadata = metaRef.value.data.meta;
	data.deployment.spec!.template.spec!.container = containersRef.value.getContainers();
	data.deployment.spec!.replicas = metaRef.value.data.replicas;
	console.log('获取到的deployment数据：', data.deployment);
	data.code = yaml.dump(data.deployment);
};
onMounted(() => {
	data.code = yaml.dump(data.deployment);
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
</style>
