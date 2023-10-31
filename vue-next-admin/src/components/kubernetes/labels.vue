<template>
	<div>
		<el-form-item>
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="data.labels!.push({ key: '', value: '' })"
				>新增</el-button
			>
		</el-form-item>

		<el-form ref="labelRef" :key="index" :model="data" v-for="(item, index) in data.labels">
			<div style="display: flex">
				<el-form-item label="key" :prop="'labels.' + index + '.key'" :rules="labelRules.key">
					<el-input placeholder="key" v-model="item.key" size="small" style="width: 120px" />
				</el-form-item>
				<el-form-item label="value" :prop="'labels.' + index + '.value'" style="margin-left: 10px" :rules="labelRules.value">
					<el-input placeholder="value" v-model="item.value" size="small" />
				</el-form-item>
				<el-form-item>
					<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.labels.splice(index, 1)"></el-button>
				</el-form-item>
			</div>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, watch, ref } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { FormInstance } from 'element-plus';

const labelRef = ref<Array<FormInstance>>();
interface label {
	key: string;
	value: string;
}
const data = reactive({
	labels: [] as label[],
});

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
		}
		callback();
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
const labelRules = {
	key: [{ validator: validateKey, required: true, trigger: 'blur' }],
	value: [{ required: true, validator: validateValue, trigger: 'blur' }],
};
//指定接收值
const props = defineProps({
	labelData: {} as { [key: string]: string },
});

const handleLabels = () => {
	const labelsTup = {};
	for (const k in data.labels) {
		// if (data.labels[k].key != '' && data.labels[k].value != '') {
		labelsTup[data.labels[k].key] = data.labels[k].value;
		// }
	}
	return labelsTup;
};

const parseLabels = (label: { [key: string]: string }) => {
	const labels = JSON.parse(JSON.stringify(label));
	const labelsTup = [];
	for (let key in labels) {
		const l = {
			key: key,
			value: labels![key],
		};
		labelsTup.push(l);
	}
	if (!isObjectValueEqual(data.labels, labelsTup)) {
		data.labels = labelsTup;
	}
};

const emit = defineEmits(['updateLabels']);
// FIXME 在父组件校验
const nextStep = (formEl: Array<FormInstance> | undefined, labels: Object) => {
	formEl?.forEach((item) => {
		item.validate((valid, fields) => {
			if (valid) {
				emit('updateLabels', labels);
			}
		});
	});
};
// 监听父组件传递来的数据
watch(
	() => props.labelData,
	() => {
		if (props.labelData && !isObjectValueEqual(props.labelData, { '': '' })) {
			let labels = JSON.parse(JSON.stringify(props.labelData));
			parseLabels(labels);
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
			nextStep(labelRef.value, labels);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>

<style scoped></style>
