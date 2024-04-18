<template>
	<div class="dynamic-form-container layout-pd">
		<el-card shadow="hover" header="创建Task">
			<el-form ref="formRef" :model="state.task" label-width="120px">
				<Meta ref="metaRef" :meta="state.task.metadata" :isUpdate="false" :resourceType="'task'" :label-width="'120px'" />
				<el-form-item label="描述" prop="description">
					<el-input v-model="state.task.spec!.description" placeholder="请输入描述信息" style="width: 250px" />
				</el-form-item>
				<Params v-if="state.task.spec" ref="paramRef" :params="state.task.spec.params" :name="'参数'" />
				<Result v-if="state.task.spec" ref="resultRef" :results="state.task.spec.results" />
				<Workspace v-if="state.task.spec" ref="workspaceRef" :workspaces="state.task.spec.workspaces" />
				<Step v-if="state.task.spec" ref="stepRef" :steps="state.task.spec.steps" />
			</el-form>
		</el-card>
		<el-row class="flex mt15">
			<div class="flex-margin">
				<el-button size="default" @click="onResetForm(formRef)">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					重置表单
				</el-button>
				<el-button size="default" type="primary" @click="onSubmitForm()">
					<SvgIcon name="iconfont icon-shuxing" />
					验证表单
				</el-button>
			</div>
		</el-row>
	</div>
</template>

<script setup lang="ts" name="test">
import { defineAsyncComponent, reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { useThemeConfig } from '@/stores/themeConfig';
import { TaskProps } from '@/types/cdk8s-pipelines/lib/tasks';
import { useTektonTasksApi } from '@/api/tekton/tasks';
import { kubernetesInfo } from '@/stores/kubernetes';

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

const api = useTektonTasksApi();
const k8sStore = kubernetesInfo();
const theme = useThemeConfig();

const state = reactive({
	size: theme.themeConfig.globalComponentSize,
	task: <TaskProps>{
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
	state.validateRef.push(...stepRef.value.ruleFormRef);

	try {
		for (const item of state.validateRef) {
			// 使用 Promise.all 来等待所有表单验证完成
			await Promise.all([item.validate()]);
		}
		ElMessage.success('验证成功');
		return true;
	} catch (error) {
		// 如果有表单验证不通过，则返回 false
		return false;
	}
};

// 表单验证
const onSubmitForm = async () => {
	if (!(await validate())) return;
	console.log(state.task);
	try {
		await api.createTask({ cloud: k8sStore.state.activeCluster }, state.task);
		ElMessage.success('创建成功');
	} catch (error) {
		console.log(error);
		ElMessage.error('创建失败');
	}
};
// 重置表单
const onResetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
</script>
