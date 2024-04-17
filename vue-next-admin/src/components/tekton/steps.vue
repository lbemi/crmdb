<template>
	<el-form-item label="步骤">
		<div>
			<div>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddStep">添加步骤</el-button>
			</div>
			<div v-for="(step, index) in state.steps" :key="index" class="flex mt20">
				<el-card class="card" shadow="hover">
					<el-form ref="ruleFormRef" :model="state.steps" label-width="60px">
						<el-form-item label="名称" :prop="'steps.' + index + '.name'">
							<el-input v-model="step.name" placeholder="步骤名称" clearable />
						</el-form-item>
						<el-form-item label="镜像" :prop="'steps.' + index + '.image'" :rules="[{ required: true, message: '请输入镜像', trigger: 'blur' }]">
							<el-input v-model="step.image" placeholder="镜像:tag" clearable />
						</el-form-item>
						<el-form-item label="script" :prop="'steps.' + index + '.script'">
							<el-input type="textarea" rows="3" v-model="step.script" placeholder="脚本" clearable />
						</el-form-item>
						<Container ref="containerRef" v-model="step.container" />
					</el-form>
				</el-card>
				<el-form-item>
					<el-button :icon="RemoveFilled" type="danger" size="small" text @click="onRemoveStep(index)">删除</el-button>
				</el-form-item>
			</div>
		</div>
	</el-form-item>
</template>
<script setup lang="ts">
import { TaskStep } from '@/types/cdk8s-pipelines';
import { defineAsyncComponent, reactive, ref } from 'vue';
import { FormInstance } from 'element-plus';
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
const Container = defineAsyncComponent(() => import('@/components/kubernetes/container.vue'));

const ruleFormRef = ref<FormInstance>();

const state = reactive({
	steps: <Array<TaskStep>>[],
});

defineProps<{
	steps: Array<TaskStep> | undefined;
}>();

const onAddStep = () => {
	state.steps.push({
		name: '',
		image: '',
		command: Array<string>(),
		args: Array<string>(),
		workingDir: '',
		env: [],
	});
};
const onRemoveStep = (index: number) => {
	state.steps.splice(index, 1);
};
</script>
<style scoped lang="scss">
.card {
	width: 50vw;
}
.el-form-item {
	margin-bottom: 18px;
}
</style>
