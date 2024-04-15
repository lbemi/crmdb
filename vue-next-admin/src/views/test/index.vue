<template>
	<div class="dynamic-form-container layout-pd">
		<el-card shadow="hover" header="创建Task">
			<el-form ref="formRef" :model="state.task.metadata" label-width="80px">
				<Meta ref="metaRef" :meta="state.task.metadata" :isUpdate="false" :resourceType="'task'" />
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
import { useThemeConfig } from '@/stores/themeConfig';
import { TaskProps } from '@/types/cdk8s-pipelines/lib/tasks';

const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
// 定义变量内容
const formRef = ref<FormInstance>();
const metaRef = ref();
const theme = useThemeConfig();

const state = reactive({
	size: theme.themeConfig.globalComponentSize,
	task: <TaskProps>{
		metadata: {
			name: '',
			namespace: 'default',
			annotations: {},
		},
	},
	validateRef: <Array<FormInstance>>[],
});

// 表单验证
const onSubmitForm = async (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	state.validateRef.push(formEl);
	const res = metaRef.value.getMeta();
	state.task.metadata = res.meta;
	state.validateRef.push(...res.validateRefs);

	try {
		for (const item of state.validateRef) {
			// 使用 Promise.all 来等待所有表单验证完成
			await Promise.all([item.validate()]);
		}

		console.log(state.task);
		ElMessage.success('验证成功');
	} catch (error) {
		// 如果有表单验证不通过，则返回 false
		return false;
	}
};
// 重置表单
const onResetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
</script>
