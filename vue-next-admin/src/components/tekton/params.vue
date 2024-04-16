<template>
	<el-form-item :label="name">
		<div>
			<div>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="addParam">新增</el-button>
			</div>
			<div :key="index" v-for="(item, index) in state.params" class="mb10">
				<template v-if="item">
					<el-form ref="ruleFormRef" :model="state" :inline="true">
						<el-form-item label="名称" :prop="'params.' + index + '.name'" :rules="labelRules.name">
							<el-input placeholder="名称" v-model="item.name" size="small" style="width: 120px" clearable />
						</el-form-item>
						<el-form-item label="类型" :prop="'params.' + index + '.type'" :rules="labelRules.type" style="width: 180px">
							<el-select v-model="item.type" placeholder="请选择" size="small">
								<el-option v-for="item in paramType" :key="item.value" :label="item.name" :value="item.value"> </el-option>
							</el-select>
						</el-form-item>
						<KV v-if="item.type === 'object'" ref="kvRef" v-model="item.properties" :name="'对象属性'" @kv-update="updateKV(index, $event)" />
						<el-form-item label="默认值" :prop="'params.' + index + '.default'" style="width: 200px">
							<div v-if="item.type === 'object'">
								<el-select v-model="state.default[index].key" placeholder="请选择" size="small">
									<el-option v-for="(k, v) in item.properties" :key="v" :label="v" :value="v"> </el-option>
								</el-select>
								<el-input
									@blur="updateDefault(index)"
									placeholder="指定属性默认值，可不填写"
									v-model="state.default[index].value"
									size="small"
									clearable
									:disabled="state.default[index].key === ''"
								/>
							</div>
							<div v-else>
								<el-input placeholder="默认值，可不填写" v-model="item.default" size="small" clearable />
							</div>
						</el-form-item>
						<el-form-item label="描述" :prop="'params.' + index + '.description'" class="mr5">
							<el-input placeholder="描述，可不填写" v-model="item.description" size="small" clearable />
						</el-form-item>
						<el-form-item>
							<el-button :icon="RemoveFilled" type="primary" size="small" text @click="removeParam(index)"></el-button>
						</el-form-item>
					</el-form>
				</template>
			</div>
			<a class="k-description">object类型暂不支持</a>
		</div>
	</el-form-item>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, ref, onMounted, defineAsyncComponent } from 'vue';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { FormInstance, FormRules } from 'element-plus';
import { TaskSpecParam } from '@/types/cdk8s-pipelines';

const KV = defineAsyncComponent(() => import('./kv.vue'));
const ruleFormRef = ref<Array<FormInstance>>([]);
const kvRef = ref();

const state = reactive({
	default: <Array<{ key: string; value: string }>>[],
	params: <Array<TaskSpecParam>>[],
	validate: false,
});

const addParam = async () => {
	state.params.push({
		name: '',
		type: '',
		default: '',
		description: '',
	});
	state.default.push({
		key: '',
		value: '',
	});
};
const removeParam = (index: number) => {
	state.params.splice(index, 1);
};
//校验key，不能重复
const validateName = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('请输入名称'));
	} else {
		let count = 0;
		state.params.forEach((item: TaskSpecParam) => {
			if (item.name === value) {
				count++;
			}
		});
		if (count > 1) {
			callback(new Error('名称已存在'));
		} else {
			callback();
		}
	}
};

const labelRules = reactive<FormRules>({
	name: [{ required: true, validator: validateName, trigger: 'blur' }],
	type: [{ required: true, message: '请输入类型', trigger: 'blur' }],
});

const updateKV = (index: number, result: any) => {
	state.params[index].properties = result;
};

const updateDefault = (index: number) => {
	state.params[index].default = {
		[`${state.default[index].key}`]: state.default[index].value,
	};
};
//指定接收值
const props = defineProps<{
	params: Array<TaskSpecParam> | undefined;
	name: string;
}>();

onMounted(() => {
	if (props.params) {
		state.params = props.params;
	}
});

const getParams = () => {
	return state.params;
};
defineExpose({
	ruleFormRef,
	getParams,
});

const paramType = [
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
