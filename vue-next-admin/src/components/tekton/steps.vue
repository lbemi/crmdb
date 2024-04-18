<template>
	<el-form-item label="步骤">
		<div>
			<div>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddStep">添加步骤</el-button>
			</div>
			<div v-for="(step, index) in state.steps" :key="index" class="flex mt20">
				<el-card class="card" shadow="hover">
					<el-form ref="ruleFormRef" :model="state" label-width="90px" :rules="formRules">
						<el-form-item label="名称" :prop="'steps.' + index + '.name'">
							<el-input v-model="step.name" placeholder="步骤名称" clearable />
						</el-form-item>
						<el-form-item label="镜像" :prop="'steps.' + index + '.image'" :rules="formRules.image">
							<el-input v-model="step.image" placeholder="镜像:tag" clearable />
						</el-form-item>
						<el-form-item label="脚本" :prop="'steps.' + index + '.script'">
							<el-input
								type="textarea"
								rows="3"
								v-model="step.script"
								placeholder="!#/bin/bash
echo hello"
								clearable
							/>
						</el-form-item>
						<el-form-item label="命令" :prop="'steps.' + index + '.command'">
							<el-input
								type="textarea"
								rows="2"
								v-model="state.tmpCommand[index]"
								placeholder="例如：sh -c,echo test"
								clearable
								@blur="updateCommand(index)"
							/>
							<div class="k-description">多个命令以逗号隔开</div>
						</el-form-item>
						<el-form-item label="参数" :prop="'steps.' + index + '.args'">
							<el-input
								type="textarea"
								rows="2"
								v-model="state.tmpArgs[index]"
								placeholder="例如：sh -c,echo test"
								clearable
								@blur="updateArgs(index)"
							/>
							<div class="k-description">多个参数以逗号隔开</div>
						</el-form-item>
						<Env ref="envsRef" :env="step.env" />
						<el-form-item label="工作路径" :prop="'steps.' + index + '.workingDir'">
							<el-input v-model="step.workingDir" placeholder="例如： $(workspaces.source.path)" />
						</el-form-item>
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

const Env = defineAsyncComponent(() => import('./env.vue'));

const ruleFormRef = ref<FormInstance>();
const envsRef = ref<FormInstance>();
const state = reactive({
	tmpCommand: [] as string[],
	tmpArgs: [] as string[],
	steps: <Array<TaskStep>>[],
});

defineProps<{
	steps: Array<TaskStep> | undefined;
}>();

const formRules = {
	image: [{ required: true, message: '请输入镜像', trigger: 'blur' }],
};
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
const updateCommand = (index: number) => {
	const tmpCommand = state.tmpCommand[index];
	if (tmpCommand) {
		const command = [];
		const parts = tmpCommand.split(',');
		let i = parts.length;
		while (i--) {
			const part = parts[i].trim();
			if (part) {
				command.unshift(part);
			}
		}
		state.steps[index].command = command;
	}
};
const updateArgs = (index: number) => {
	const tmpArgs = state.tmpArgs[index];
	if (tmpArgs) {
		const args = [];
		const parts = tmpArgs.split(',');
		let i = parts.length;
		while (i--) {
			const part = parts[i].trim();
			if (part) {
				args.unshift(part);
			}
		}
		state.steps[index].args = args;
	}
};
const getSteps = () => {
	return state.steps;
};

defineExpose({
	ruleFormRef,
	getSteps,
});
</script>

<style scoped lang="scss">
.card {
	width: 50vw;
}
.el-form-item {
	margin-bottom: 18px;
}

.custom-header {
	display: flex;
	align-items: center;
	justify-content: flex-start;
	/* 您可以根据需要调整头部的样式 */
	width: 100px; /* 或者您想要的宽度 */
	height: 100%;
	padding: 10px;
	box-sizing: border-box;
}

.header-title {
	writing-mode: vertical-rl;
	transform: rotate(180deg);
	/* 根据需要调整文字样式 */
	font-size: 16px;
	font-weight: bold;
}
</style>
