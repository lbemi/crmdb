<template>
	<el-form-item label="results">
		<div>
			<div>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddResult">新增</el-button>
				<a class="k-description">一般用于共享少了字符数据</a>
				<a class="k-description" target="_blank" href="https://tekton.dev/docs/pipelines/tasks/#emitting-results">详细信息</a>
			</div>
			<div v-for="(item, index) in state.results" :key="index" class="mb10">
				<el-form ref="ruleFormRef" :model="state" inline>
					<el-form-item label="名称" :key="index" :prop="'results.' + index + '.name'" :rules="formRules.name">
						<el-input v-model="item.name" placeholder="请输入名称" style="width: 150px" size="small" clearable />
					</el-form-item>
					<el-form-item label="类型" :key="index" :prop="'results.' + index + '.type'" style="width: 170px" :rules="formRules.type">
						<el-select v-model="item.type" placeholder="请选择" size="small">
							<el-option v-for="item in resourceType" :key="item.value" :label="item.name" :value="item.value"> </el-option>
						</el-select>
					</el-form-item>
					<KV ref="kvRef" v-show="item.type === 'object'" v-model="item.properties" :name="'对象属性'" @kv-update="updateKV(index, $event)" />
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
import { reactive, ref, onMounted, defineAsyncComponent } from 'vue';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { FormInstance, FormRules } from 'element-plus';
import { TaskSpecResult } from '@/types/cdk8s-pipelines';

const KV = defineAsyncComponent(() => import('./kv.vue'));

const ruleFormRef = ref<Array<FormInstance>>([]);
const kvRef = ref();

const state = reactive({
	results: [] as TaskSpecResult[],
});

const onAddResult = async () => {
	state.results.push({
		name: '',
		description: '',
		type: 'string',
	});
};

const onRemoveResult = (index: number) => {
	state.results.splice(index, 1);
};

//校验Name，不能重复
const validateName = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('请输入名称'));
	} else {
		let count = 0;
		state.results.forEach((item: TaskSpecResult) => {
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

const updateKV = (index: number, result: any) => {
	state.results[index].properties = result;
};

// 校验value
const formRules = reactive<FormRules>({
	name: [{ required: true, validator: validateName, trigger: 'blur' }],
	type: [{ required: true, message: '请选择入类型', trigger: 'blur' }],
});

//指定接收值
const props = defineProps<{
	results: Array<TaskSpecResult> | undefined;
}>();

const getResults = () => {
	return state.results;
};

onMounted(() => {
	if (props.results) {
		state.results = props.results;
	}
});

defineExpose({
	ruleFormRef,
	getResults,
});

const resourceType = [
	{
		name: '字符类型(string)',
		value: 'string',
	},
	{
		name: '数组类型(array)',
		value: 'array',
	},
	{
		name: '对象类型(object)',
		value: 'object',
	},
];
</script>
