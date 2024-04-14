<template>
	<div class="dynamic-form-container layout-pd">
		<el-card shadow="hover" header="创建Task">
			<el-form :model="state.task" ref="formRef" size="default" label-width="100px">
				<el-form-item label="名称" prop="metadata.name" :rules="[{ required: true, message: '名称不能为空', trigger: 'blur' }]">
					<el-input v-model="state.task.metadata.name" placeholder="名称" clearable style="width: 225px"></el-input>
				</el-form-item>
				<el-form-item label="命名空间">
					<el-select v-model="state.task.metadata.namespace" style="max-width: 280px" class="m-2" placeholder="Select" :size="state.size">
						<el-option
							v-for="item in k8sStore.state.namespace"
							:key="item.metadata?.name"
							:label="item.metadata?.name"
							:value="item.metadata!.name!"
						/>
					</el-select>
				</el-form-item>
				<KeyValue ref="labelRef" :labels="state.task.metadata.annotations" :name="'注解'" />
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
				<el-button size="default" type="primary" @click="onSubmitForm(formRef)">
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
import { kubernetesInfo } from '@/stores/kubernetes';
import { useThemeConfig } from '@/stores/themeConfig';
import { Task, Annotations } from '@/types/tekton/task';
import { Deployment } from 'kubernetes-models/apps/v1';

const KeyValue = defineAsyncComponent(() => import('@/components/kubernetes/KeyValue.vue'));
// 定义变量内容
const formRef = ref<FormInstance>();
const labelRef = ref();
const theme = useThemeConfig();
const k8sStore = kubernetesInfo();

const state = reactive({
	size: theme.themeConfig.globalComponentSize,
	task: <Task>{
		metadata: {
			name: '',
			namespace: 'default',
			annotations: <Annotations>{},
		},
	},
	validateRef: <Array<FormInstance>>[],
});

// 表单验证
const onSubmitForm = async (formEl: FormInstance | undefined) => {
	console.log(labelRef.value.ruleFormRef);
	if (!formEl) return;
	state.validateRef.push(formEl);
	if (!labelRef.value.ruleFormRef) return;
	state.validateRef.push(...labelRef.value.ruleFormRef);
	try {
		for (const item of state.validateRef) {
			// 使用 Promise.all 来等待所有表单验证完成
			await Promise.all([item.validate()]);
		}
		state.task.metadata.annotations = labelRef.value.getLabels();
		console.log(state.task);
		const dep = new Deployment();
		dep.metadata = {
			name: [state.task.metadata.name],
			test: 'xxx',
		};
		dep.validate();
		console.log(dep);
		ElMessage.success('验证成功');
	} catch (error) {
		console.log(error);
		// 如果有表单验证不通过，则返回 false
		ElMessage.error(error.message);
	}
};
// 重置表单
const onResetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
</script>
