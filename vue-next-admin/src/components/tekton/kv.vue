<template>
	<el-form-item :label="name">
		<div>
			<div>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="addLabel">新增</el-button
				><a class="k-description">默认类型为： { type: string }</a>
			</div>
			<div :key="index" v-for="(item, index) in data.labels" class="mb10">
				<template v-if="item">
					<el-form ref="ruleFormRef" :model="data" :inline="true" v-if="item.key !== 'app'">
						<el-form-item label="键" :prop="'labels.' + index + '.key'" :rules="labelRules.key" class="mr1">
							<el-input placeholder="key" v-model="item.key" size="small" style="width: 120px" />
						</el-form-item>
						<el-form-item>
							<el-button :icon="RemoveFilled" type="primary" size="small" text @click="removeLabel(index)"></el-button>
						</el-form-item>
					</el-form>
				</template>
			</div>
			<!--			<a class="k-description">标签键值以字母、数字开头和结尾, 且只能包含字母、数字及分隔符。</a>-->
		</div>
	</el-form-item>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, ref, onMounted, watch } from 'vue';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { FormInstance, FormRules } from 'element-plus';

const ruleFormRef = ref<Array<FormInstance>>([]);
interface label {
	key: string;
	value: object;
}
const data = reactive({
	labels: [] as label[],
	validate: false,
});

const addLabel = async () => {
	data.labels.push({ key: '', value: { type: 'string' } });
};
const removeLabel = (index: number) => {
	data.labels.splice(index, 1);
};
//校验key，不能重复
const validateKey = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('请输入key'));
	} else {
		let count = 0;
		data.labels.forEach((item: label) => {
			if (item.key === value) {
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
const labelRules = reactive<FormRules>({
	key: [{ required: true, validator: validateKey, trigger: 'blur' }],
});

//指定接收值
const props = defineProps({
	labels: Object,
	name: {
		type: String,
		default: '标签',
	},
});

const getLabels = () => {
	if (!validate()) return;
	const labelsTup: { [key: string]: object } = {};
	for (const k in data.labels) {
		if (data.labels[k].key != '') {
			labelsTup[data.labels[k].key] = {
				type: 'string',
			};
		}
	}
	emits('kv-update', labelsTup);
};
const emits = defineEmits(['kv-update']);
const parseLabels = (label: { [key: string]: object }) => {
	const labels = JSON.parse(JSON.stringify(label));
	const labelsTup = [];
	for (let key in labels) {
		const l = {
			key: key,
			value: labels![key],
		};
		labelsTup.push(l);
	}
	return labelsTup;
};
const validate = async () => {
	try {
		for (const item in ruleFormRef) {
			await Promise.all([item.validate()]);
		}
		return true;
	} catch (error) {
		return false;
	}
};
onMounted(() => {
	if (props.labels) {
		data.labels = parseLabels(props.labels);
	}
});

watch(
	() => data.labels,
	() => {
		getLabels();
	},
	{
		deep: true,
		immediate: true,
	}
);
//
// defineExpose({
// 	ruleFormRef,
// 	getLabels,
// });
</script>
