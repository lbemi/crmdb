<template>
	<el-form-item :label="name" :label-width="labelWidth">
		<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="addLabel">新增</el-button>
	</el-form-item>
	<el-form-item :key="index" v-for="(item, index) in data.labels" :label-width="labelWidth">
		<template v-if="item">
			<el-form ref="labelRef" :model="data" :inline="true" v-if="item.key !== 'app'">
				<el-form-item label="键" :prop="'labels.' + index + '.key'" :rules="labelRules.key">
					<el-input placeholder="key" v-model="item.key" size="small" style="width: 120px" />
				</el-form-item>
				<el-form-item label="值" :prop="'labels.' + index + '.value'" :rules="labelRules.value">
					<el-input placeholder="value" v-model="item.value" size="small" />
				</el-form-item>
				<el-form-item>
					<el-button :icon="RemoveFilled" type="primary" size="small" text @click="removeLabel(index)"></el-button>
				</el-form-item>
			</el-form>
		</template>
	</el-form-item>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, watch, ref } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { FormInstance, FormRules } from 'element-plus';

const labelRef = ref<Array<FormInstance>>([]);
interface label {
	key: string;
	value: string;
}
const data = reactive({
	labels: [] as label[],
});

const addLabel = () => {
	data.labels.push({ key: '', value: '' });
	// labelRef.value.push(null); // Push null initially, it will be populated when the form is rendered
};
const removeLabel = (index: number) => {
	data.labels.splice(index, 1);
	// labelRef.value.splice(index, 1);
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
const validateValue = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('请输入value'));
	} else {
		callback();
	}
};
const labelRules = reactive<FormRules>({
	key: [{ required: true, validator: validateKey, trigger: 'blur' }],
	value: [{ required: true, validator: validateValue, trigger: 'blur' }],
});

//指定接收值
const props = defineProps({
	labelData: Array,
	name: {
		type: String,
		default: '标签',
	},
	labelWidth: {
		type: String,
		default: '90px',
	},
});

const handleLabels = () => {
	const labelsTup: { [key: string]: string } = {};
	for (const k in data.labels) {
		if (data.labels[k].key != '' && data.labels[k].value != '') {
			labelsTup[data.labels[k].key] = data.labels[k].value;
		}
	}
	return labelsTup;
};

const emit = defineEmits(['updateLabels']);
// FIXME 在父组件校验
const validateHandler = (formEl: Array<FormInstance> | undefined, labels: Object) => {
	let status = false;
	formEl?.forEach((item) => {
		item.validate((valid) => {
			status = valid;
			if (valid) {
				emit('updateLabels', labels);
			}
		});
	});
	return status;
};
// 监听父组件传递来的数据
watch(
	() => props.labelData,
	() => {
		if (props.labelData) {
			data.labels = JSON.parse(JSON.stringify(props.labelData));
		}
	},
	{
		immediate: true,
		deep: true,
	}
);

// 监听表单数据，如果发生变化则传递到父组件
watch(
	() => data.labels,
	() => {
		const labels = handleLabels();
		if (!isObjectValueEqual(labels, { '': '' })) {
			validateHandler(labelRef.value, labels);
			// emit('updateLabels', labels);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);

defineExpose({
	validateHandler,
});
</script>

<style scoped></style>
