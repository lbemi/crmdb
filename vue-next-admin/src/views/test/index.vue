<template>
	<div class="dynamic-form-container layout-pd">
		<el-card shadow="hover" header="创建Task">
			<el-form ref="formRef" :model="state.task" label-width="80px">
				<Meta ref="metaRef" :meta="state.task.metadata" :isUpdate="false" :resourceType="'task'" />
				<el-form-item label="描述" prop="description">
					<el-input v-model="state.task.spec!.description" placeholder="请输入描述信息" style="width: 250px" />
				</el-form-item>
				<Params v-if="state.task.spec" ref="paramRef" :params="state.task.spec.params" :name="'参数'" />
				<el-form-item label="results">
					<div>
						<div>
							<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddResult">新增</el-button>
						</div>
						<div v-for="(item,index) in state.task.spec!.results" :key="index" class="mb10">
							<el-form :model="state" inline>
								<el-form-item
									label="名称"
									:key="index"
									prop="'results.' + index + '.name'"
									:rules="[{ required: true, message: '请输入名称', trigger: 'blur' }]"
								>
									<el-input v-model="item.name" placeholder="请输入名称" style="width: 250px" size="small" />
								</el-form-item>
								<el-form-item label="描述" :key="index" class="mr1">
									<el-input v-model="item.description" placeholder="描述，可选" style="width: 250px" size="small" />
								</el-form-item>
								<el-form-item>
									<el-button :icon="RemoveFilled" type="primary" size="small" text @click="onRemoveResult(index)"></el-button>
								</el-form-item>
							</el-form>
						</div>
					</div>
				</el-form-item>
				<el-form-item label="步骤">
					<el-input placeholder="请输入步骤" />
				</el-form-item>
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
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';

const Meta = defineAsyncComponent(() => import('@/components/kubernetes/meta.vue'));
const Params = defineAsyncComponent(() => import('@/components/tekton/params.vue'));
// 定义变量内容
const formRef = ref<FormInstance>();
const metaRef = ref();
const paramRef = ref();

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
		},
	},
	validateRef: <Array<FormInstance>>[],
});
const onAddResult = () => {
	state.task.spec!.results!.push({
		name: '',
		description: '',
	});
};
const onRemoveResult = (index: number) => {
	state.task.spec!.results!.splice(index, 1);
};
const validate = async () => {
	if (!formRef.value) return;
	state.validateRef.push(formRef.value);
	const res = metaRef.value.getMeta();
	state.task.metadata = res.meta;
	state.validateRef.push(...res.validateRefs);

	state.task.spec!.params = paramRef.value.getParams();
	state.validateRef.push(...paramRef.value.ruleFormRef);

	try {
		for (const item of state.validateRef) {
			// 使用 Promise.all 来等待所有表单验证完成
			await Promise.all([item.validate()]);
		}

		ElMessage.success('验证成功');
	} catch (error) {
		// 如果有表单验证不通过，则返回 false
		return false;
	}
};

// 表单验证
const onSubmitForm = async () => {
	await validate();
	console.log(state.task);
};
// 重置表单
const onResetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
</script>
