<template>
		<el-form :key="portIndex" v-for="(item, portIndex) in data.labels"  >
			<div v-if="item.key !== 'app'" style="display: flex">
				<el-form-item label="key" >
					<el-input placeholder="key" v-model="item.key" size="small" style="width: 120px" />
				</el-form-item>
				<el-form-item label="value" style="margin-left: 10px" >
					<el-input placeholder="value" v-model="item.value" size="small" />
				</el-form-item>
				<el-form-item>
					<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.labels.splice(portIndex, 1)"></el-button>
				</el-form-item>
			</div>
		</el-form>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, watch } from 'vue';

interface label {
	key: string;
	value: string;
}
const data = reactive({
	labels: [

	] as label[],
});
const props = defineProps({
	labelData: Array,
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

const emit = defineEmits(['updateLabels']);

// 监听父组件传递来的数据
watch(
	() => props.labelData,
	() => {
		if ( props.labelData) {
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
		// const labels = handleLabels();
    // if(!isObjectValueEqual(labels,{'':''})) {
    //   emit('updateLabels', labels);
    // }
		emit('updateLabels', data.labels);
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>

<style scoped>

</style>
