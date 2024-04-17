<template>
	<el-form-item label="工作空间">
		<div>
			<div>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddResult">新增</el-button>
				<a class="k-description">用于共享数据,默认是/workspace</a>
				<a class="k-description" target="_blank" href="https://tekton.dev/docs/pipelines/tasks/#emitting-results">详细信息</a>
			</div>
			<div v-for="(item, index) in state.workspaces" :key="index" class="mb10">
				<el-form ref="ruleFormRef" :model="state" inline>
					<el-form-item label="名称" :key="index" :prop="'workspaces.' + index + '.name'" :rules="formRules.name">
						<el-input v-model="item.name" placeholder="请输入名称" style="width: 150px" size="small" clearable />
					</el-form-item>
					<el-form-item label="指定工作目录">
						<el-input v-model="item.mountPath" placeholder="指定工作目录" style="width: 150px" size="small" clearable />
					</el-form-item>
					<el-form-item label="描述" :key="index" class="mr1">
						<el-input v-model="item.description" placeholder="描述，可选" size="small" />
					</el-form-item>
					<el-form-item>
						<el-button :icon="RemoveFilled" type="primary" size="small" text @click="onRemoveResult(index)"></el-button>
					</el-form-item>
				</el-form>
			</div>
		</div>
	</el-form-item>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, ref, onMounted } from 'vue';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { FormInstance, FormRules } from 'element-plus';
import { TaskWorkspace } from '@/types/cdk8s-pipelines';

const ruleFormRef = ref<Array<FormInstance>>([]);

const state = reactive({
	workspaces: [] as TaskWorkspace[],
});

const onAddResult = async () => {
	state.workspaces.push({
		name: '',
		description: '',
		mountPath: '',
	});
};

const onRemoveResult = (index: number) => {
	state.workspaces.splice(index, 1);
};

//校验Name，不能重复
const validateName = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('请输入名称'));
	} else {
		let count = 0;
		state.workspaces.forEach((item: TaskWorkspace) => {
			if (item.name === value) {
				count++;
			}
		});
		if (count > 1) {
			callback(new Error('key已存在'));
		} else {
			callback();
		}
	}
};

// 校验value
const formRules = reactive<FormRules>({
	name: [{ required: true, validator: validateName, trigger: 'blur' }],
});

//指定接收值
const props = defineProps<{
	workspaces: Array<TaskWorkspace> | undefined;
}>();

const getWorkspaces = () => {
	return state.workspaces;
};

onMounted(() => {
	if (props.workspaces) {
		state.workspaces = props.workspaces;
	}
});

defineExpose({
	ruleFormRef,
	getWorkspaces,
});
</script>
