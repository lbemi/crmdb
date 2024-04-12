<template>
	<div class="dynamic-form-container layout-pd">
		<el-card shadow="hover" header="创建Task">
			<el-form :model="state.task" ref="formRulesOneRef" size="default" label-width="100px" class="mt35">
				<el-row :gutter="35" v-for="(v, k) in state.form.list" :key="k">
					<el-col :xs="24" :sm="12" :md="8" :lg="8" :xl="6" class="mb20">
						<el-form-item label="名称" :prop="state.task.metadata.name" :rules="[{ required: true, message: '名称不能为空', trigger: 'blur' }]">
							<el-input v-model="state.task.metadata.name" placeholder="名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="8" :lg="8" :xl="6" class="mb20">
						<el-form-item label="命名空间">
							<el-select v-model="state.task.metadata.namespace" style="max-width: 280px" class="m-2" placeholder="Select" :size="state.size"
								><el-option key="all" label="所有命名空间" value="all"></el-option>
								<el-option
									v-for="item in k8sStore.state.namespace"
									:key="item.metadata?.name"
									:label="item.metadata?.name"
									:value="item.metadata!.name!"
								/>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="8" :lg="8" :xl="6" class="mb20">
						<Label :name="'注解'" />
					</el-col>
				</el-row>
			</el-form>
		</el-card>
		<el-row class="flex mt15">
			<div class="flex-margin">
				<el-button size="default" @click="onResetForm(formRulesOneRef)">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					重置表单
				</el-button>
				<el-button size="default" type="primary" @click="onSubmitForm(formRulesOneRef)">
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
import { Task, annotations } from '@/types/tekton/task';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));
// 定义变量内容
const formRulesOneRef = ref<FormInstance>();
const theme = useThemeConfig();
const k8sStore = kubernetesInfo();

const state = reactive({
	size: theme.themeConfig.globalComponentSize,
	task: {
		metadata: {
			name: '',
			namespace: 'default',
			annotations: {} as annotations,
		},
	} as Task,
	form: {
		namespace: 'default',
		name: '',
		email: '',
		autograph: '',
		occupation: '',
		list: [
			{
				year: '',
				month: '',
				day: '',
			},
		],
		remarks: '',
	},
});

// 新增行
const onAddRow = () => {
	state.form.list.push({
		year: '',
		month: '',
		day: '',
	});
};
// 删除行
const onDelRow = (k: number) => {
	state.form.list.splice(k, 1);
};
// 表单验证
const onSubmitForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate((valid: boolean) => {
		if (valid) {
			ElMessage.success('验证成功');
		} else {
			return false;
		}
	});
};
// 重置表单
const onResetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
</script>
