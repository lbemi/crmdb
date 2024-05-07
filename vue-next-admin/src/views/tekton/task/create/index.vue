<template>
	<div class="dynamic-form-container layout-pd">
		<el-card shadow="hover">
			<template #header>
				<div class="flex-between">
					<div class="card-header">
						<span>{{ state.title }}Task</span>
					</div>
					<div class="mr15">
						<el-button plain type="primary" @click="onOpenYaml()">yaml</el-button>
					</div>
				</div>
			</template>
			<el-form ref="formRef" :model="state.task" label-width="120px">
				<Meta ref="metaRef" :meta-data="{ metadata: state.task.metadata }" :isUpdate="state.update" :resourceType="'task'" :label-width="'120px'" />
				<el-form-item label="描述" prop="description">
					<el-input v-model="state.task.spec!.description" placeholder="请输入描述信息" style="width: 250px" />
				</el-form-item>
				<Params v-if="state.task.spec" ref="paramRef" :params="state.task.spec.params" :name="'参数'" />
				<Result v-if="state.task.spec" ref="resultRef" :results="state.task.spec.results" />
				<Workspace v-if="state.task.spec" ref="workspaceRef" :workspaces="state.task.spec.workspaces" />
				<Step v-if="state.task.spec" ref="stepRef" :steps="state.task.spec.steps" />
			</el-form>
		</el-card>
		<YamlDialog
			:dialogVisible="state.dialogVisible"
			:code-data="state.task"
			@update="updateTaskYaml"
			v-if="state.dialogVisible"
			:disabledUpdate="true"
		/>
		<el-row class="flex mt15">
			<div class="flex-margin">
				<el-button size="default" @click="onCancel()">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					取消
				</el-button>
				<el-button size="default" type="primary" @click="onSubmitForm()">
					<SvgIcon name="iconfont icon-shuxing" />
					确认
				</el-button>
			</div>
		</el-row>
	</div>
</template>

<script setup lang="ts" name="test">
import { defineAsyncComponent, onBeforeMount, reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { useThemeConfig } from '@/stores/themeConfig';
import { Task } from '@/types/tekton/api';
import { useTektonTasksApi } from '@/api/tekton/tasks';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useRouter, useRoute } from 'vue-router';

const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
const Params = defineAsyncComponent(() => import('@/components/tekton/params.vue'));
const Result = defineAsyncComponent(() => import('@/components/tekton/results.vue'));
const Workspace = defineAsyncComponent(() => import('@/components/tekton/workspaces.vue'));
const Step = defineAsyncComponent(() => import('@/components/tekton/steps.vue'));
// 定义变量内容
const formRef = ref<FormInstance>();
const metaRef = ref();
const paramRef = ref();
const resultRef = ref();
const workspaceRef = ref();
const stepRef = ref();

const route = useRoute();
const router = useRouter();
const api = useTektonTasksApi();
const k8sStore = kubernetesInfo();
const theme = useThemeConfig();

const state = reactive({
	update: false,
	title: '创建',
	dialogVisible: false,
	size: theme.themeConfig.globalComponentSize,
	task: <Task>{
		metadata: {
			name: '',
			namespace: 'default',
			annotations: {},
		},
		spec: {
			description: '',
			params: [],
			workspaces: [],
			results: [],
			steps: [],
		},
	},
	validateRef: <Array<FormInstance>>[],
});

const validate = async () => {
	if (!formRef.value) return;
	state.validateRef.push(formRef.value);
	const res = metaRef.value.getMeta();
	state.task.metadata = res.meta;
	state.validateRef.push(...res.validateRefs);

	state.task.spec!.params = paramRef.value.getParams();
	state.validateRef.push(...paramRef.value.ruleFormRef);

	//校验results
	state.task.spec!.results = resultRef.value.getResults();
	state.validateRef.push(...resultRef.value.ruleFormRef);

	//校验workspaces
	state.task.spec!.workspaces = workspaceRef.value.getWorkspaces();
	state.validateRef.push(...workspaceRef.value.ruleFormRef);

	//校验steps
	state.task.spec!.steps = stepRef.value.getSteps();
	if (stepRef.value.ruleFormRef && stepRef.value.ruleFormRef.length > 0) {
		state.validateRef.push(...stepRef.value.ruleFormRef);
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
onBeforeMount(() => {
	if (route.query.update == 'true') {
		state.task = k8sStore.state.tekton.updateTask;
		state.title = '更新';
		state.update = true;
	}
});
const onOpenYaml = async () => {
	await validate();
	state.dialogVisible = true;
};
const updateTaskYaml = async (task: Task) => {
	try {
		await api.updateTask({ cloud: k8sStore.state.activeCluster }, task);
		ElMessage.success('更新成功');
	} catch (error) {
		ElMessage.error('更新失败');
	}
};
// 表单验证并提交
const onSubmitForm = async () => {
	if (!(await validate())) return;
	try {
		if (state.update) {
			await api.updateTask({ cloud: k8sStore.state.activeCluster }, state.task);
		} else {
			await api.createTask({ cloud: k8sStore.state.activeCluster }, state.task);
		}
		ElMessage.success('创建成功');
		await router.push('/tekton/task');
	} catch (error) {
		ElMessage.error('创建失败');
	}
};

// 返回上一次路由地址
const onCancel = () => {
	window.history.go(-1);
};
</script>
