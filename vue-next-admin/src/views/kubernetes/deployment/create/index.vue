<template>
	<div class="layout-pd">
		<el-card shadow="hover">
			<el-backtop :right="100" :bottom="100" />
			<el-row>
				<el-col :span="3">
					<div style="height: 100%">
						<el-steps :active="data.active" finish-status="success" direction="vertical" style="margin-top: 5px; height: 100%" align-center>
							<el-step title="基本信息" description="Some description" />
							<el-step title="容器配置" description="Some description" />
							<el-step title="高级配置" description="Some description" />
						</el-steps>
					</div>
				</el-col>
				<el-col :span="11">
					<div>
						<div style="margin-top: 10px" id="0" v-show="data.active === 0">
							<Meta ref="metaRef" v-bind:metadata="data.tableData" v-bind:resourceType="'deployment'" />
						</div>
						<div style="margin-top: 10px" id="1" v-show="data.active === 1">
							<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">
								<el-tab-pane label="容器 1" key="1" name="1">
									<el-form :model="data.containers" label-width="120px" label-position="left">
										<el-form-item label="镜像名称：">
											<el-input v-model="data.containers.image" size="default" />
										</el-form-item>
										<el-form-item label="拉取策略：">
											<el-select v-model="data.containers.imagePullPolicy" class="m-2" placeholder="Select" size="default">
												<el-option v-for="item in imagePullPolicy" :key="item.name" :label="item.name" :value="item.value" />
											</el-select>
										</el-form-item>
										<el-form-item label="所需资源：">
											<div style="height: 28px">
												<span>CPU</span>
												<el-input placeholder="如：0.5" v-model.number="data.limit.cpu" size="small" style="width: 80px" />
												<span> Core</span>
												<el-divider direction="vertical" />
												<a>内存</a>
												<el-input placeholder="如：0.5" v-model.number="data.limit.memory" size="small" style="width: 80px" /><span> MiB</span>
											</div>
											<div style="font-size: 6px; color: #00bb00">
												<el-icon size="12px" color="#00bb00"><InfoFilled /></el-icon
												>建议根据实际使用情况设置，防止由于资源约束而无法调度或引发内存不足(OOM)错误
											</div>
										</el-form-item>
										<el-form-item label="资源限制：">
											<div style="height: 28px">
												<span>CPU</span>
												<el-input placeholder="如：0.5" v-model.number="data.require.cpu" size="small" style="width: 80px" />
												<span> Core</span>
												<el-divider direction="vertical" />
												<a>内存</a>
												<el-input placeholder="如：0.5" v-model.number="data.require.memory" size="small" style="width: 80px" /><span> MiB</span>
											</div>
											<el-tooltip
												class="box-item"
												effect="light"
												content="<div>即为该应用预留资源额度，包括CPU和内存两种资源，即容器独占该资源，</div><div> 防止因资源不足而被其他服务或进程争夺资源，导致应用不可用</div>"
												placement="top-start"
												raw-content
											>
												<div style="font-size: 6px; color: #00bb00">
													<el-icon size="12px" color="#00bb00"><InfoFilled /></el-icon> 建议根据实际使用情况设置，防止因资源不足导致应用不可用
												</div>
											</el-tooltip>
										</el-form-item>
										<el-form-item label="特权容器：">
											<template #label>
											<el-tooltip
												class="box-item"
												effect="light"
												content="<div>默认情况下，容器是不可以访问宿主上的任何设备；特权容器则</div><div> 被授权访问宿主上所有设备，享有宿主上运行的进程的所有访问权限</div>"
												placement="top-start"
												raw-content
											>
												特权容器
											</el-tooltip>
											</template>
											<el-switch
												v-model="data.containers.securityContext!.privileged"
												class="mt-2"
												style="margin-left: 24px"
												inline-prompt
												:active-icon="Check"
												:inactive-icon="Close"
												size="small"
											/>
										</el-form-item>
										<el-form-item label="特权容器：">
											<template #label>
												<el-tooltip
														class="box-item"
														effect="light"
														content="<div>默认情况下，容器是不可以访问宿主上的任何设备；特权容器则</div><div> 被授权访问宿主上所有设备，享有宿主上运行的进程的所有访问权限</div>"
														placement="top-start"
														raw-content
												>
													特权容器
												</el-tooltip>
											</template>
											<el-switch
													v-model="data.containers.securityContext!.privileged"
													class="mt-2"
													style="margin-left: 24px"
													inline-prompt
													:active-icon="Check"
													:inactive-icon="Close"
													size="small"
											/>
										</el-form-item>
									</el-form>
								</el-tab-pane>
								<el-tab-pane v-for="item in editableTabs" :key="item.name" :label="item.title" :name="item.name">
									{{ item.content }}
								</el-tab-pane>
							</el-tabs>
						</div>
						<div style="margin-top: 10px" id="2" v-show="data.active === 2">
							<h1>asdj</h1>
						</div>
					</div>
				</el-col>
				<el-col :span="8">
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
						<el-button @click="next" style="margin-top: 5px" size="small">确认</el-button>
					</div>
				</el-col>
			</el-row>
		</el-card>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive } from 'vue';

import { ref } from 'vue-demi';
import { Codemirror } from 'vue-codemirror';
import { javascript } from '@codemirror/lang-javascript';
import { oneDark } from '@codemirror/theme-one-dark';
import { V1Container, V1Deployment, V1SecurityContext } from '@kubernetes/client-node';
import yaml from 'js-yaml';
import type { TabPaneName } from 'element-plus';
import { Check, Close, InfoFilled } from '@element-plus/icons-vue';

let tabIndex = 1;
const editableTabsValue = ref('1');
const editableTabs = ref([]);

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
	if (action === 'add') {
		const newTabName = `${++tabIndex}`;
		editableTabs.value.push({
			title: 'New Tab',
			name: newTabName,
			content: 'New Tab content',
		});
		editableTabsValue.value = newTabName;
	} else if (action === 'remove') {
		const tabs = editableTabs.value;
		let activeName = editableTabsValue.value;
		if (activeName === targetName) {
			tabs.forEach((tab, index) => {
				if (tab.name === targetName) {
					const nextTab = tabs[index + 1] || tabs[index - 1];
					if (nextTab) {
						activeName = nextTab.name;
					}
				}
			});
		}

		editableTabsValue.value = activeName;
		editableTabs.value = tabs.filter((tab) => tab.name !== targetName);
	}
};
const imagePullPolicy = [
	{
		name: '优先使用本地镜像(ifNotPresent)',
		value: 'ifNotPresent',
	},
	{
		name: '总是拉取镜像(Always)',
		value: 'Always',
	},
	{
		name: '仅使用本地镜像(Never)',
		value: 'Never',
	},
];
const Meta = defineAsyncComponent(() => import('/@/components/kubernetes/meta.vue'));

const metaRef = ref<InstanceType<typeof Meta>>();

const data = reactive({
	containers: {
		securityContext: {
			privileged: false,
		} as V1SecurityContext,
	} as V1Container,
	// 上面测试
	limit: {
		cpu: 0,
		memory: 0,
	},
	require: {
		cpu: 0.5,
		memory: 500,
	},
	active: 0,
	deployment: {
		apiVersion: 'apps/v1',
		kind: 'Deployment',
		metadata: {
			namespace: 'default',
		},
		spec: {},
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
	console.log(metaRef.value.data.meta);
	data.deployment.metadata = metaRef.value.data.meta;
	data.deployment.spec!.replicas = metaRef.value.data.replicas;
	data.code = yaml.dump(data.deployment);
	const htmlElement = document.getElementById(data.active + '');
	if (htmlElement) htmlElement.scrollIntoView(true);
	if (data.active++ > 2) data.active = 0;
};

onMounted(() => {
	data.code = yaml.dump(data.deployment);
});

// watch((data),()=>{
// 	console.log("******",data.code)
// 	// data.code = yaml.dump(metaRef.value.data.meta)
// 	// data.tableData = YAML.load(data.code).metadata
// })
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
</style>
