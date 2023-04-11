<template>
	<div style="margin-top: 10px">
		<el-form :model="data.meta" label-width="120px" label-position="left">
			<el-form-item label="命名空间">
				<el-select v-model="data.meta.namespace" style="max-width: 220px" class="m-2" placeholder="Select" size="default">
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name" />
				</el-select>
			</el-form-item>
			<el-form-item label="应用名称">
				<el-input v-model="data.meta.name" size="default" style="width: 220px" />
			</el-form-item>
			<el-form-item label="类型">
				<el-select v-model="data.resourceType" class="m-2" placeholder="Select" size="default" style="width: 220px" :disabled="enableEdit">
					<el-option v-for="item in types" :key="item.name" :label="item.name" :value="item.value" />
				</el-select>
			</el-form-item>
			<el-form-item
				label="副本数量"
				v-if="data.resourceType === 'deployment' || data.resourceType === 'daemonSet' || data.resourceType === 'statefulSet'"
				size="default"
			>
				<el-input-number v-model="data.replicas" :min="1" :max="100" />
			</el-form-item>
			<el-form-item label="标签">
				<el-button
					:icon="CirclePlusFilled"
					type="primary"
					size="small"
					text
					style="padding-left: 0"
					@click="data.labelData.push({ key: '', value: '' })"
					>新增</el-button
				>
			</el-form-item>
			<el-form-item v-if="data.labelData.length != 0">
				<Label :labelData="data.labelData" @updateLabels="getLabels" />
			</el-form-item>
			<el-form-item label="注解">
				<el-button
					:icon="CirclePlusFilled"
					type="primary"
					size="small"
					text
					style="padding-left: 0"
					@click="data.annotationsData.push({ key: '', value: '' })"
					>新增</el-button
				>
			</el-form-item>
			<el-form-item v-if="data.annotationsData.length != 0">
				<Label :labelData="data.annotationsData" @updateLabels="getAnnotations" />
			</el-form-item>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, watch } from 'vue';
import { V1ObjectMeta } from '@kubernetes/client-node';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { ref } from 'vue-demi';
import { CirclePlusFilled } from '@element-plus/icons-vue';
import { isObjectValueEqual } from '/@/utils/arrayOperation';
interface label {
  key: string;
  value: string;
}
const Label = defineAsyncComponent(() => import('/@/components/kubernetes/label.vue'));

const k8sStore = kubernetesInfo();
const enableEdit = ref(false);
// const tableData = ref([]);
const data = reactive({
	labelData: [] as  label[],
	annotationsData: [],
	replicas: 1,
	resourceType: 'deployment',
	meta: {
		namespace: 'default',
	} as V1ObjectMeta,
});

const getLabels = (labels: any) => {
	if (!isObjectValueEqual(data.meta.labels, labels)) {
		data.meta.labels = labels;
	}
};

const getAnnotations = (labels: any) => {
	if (!isObjectValueEqual(data.meta.annotations, labels)) {
		data.meta.annotations = labels;
	}
};

const handleLabels = (labels: { [key: string]: string }) => {
	const labelsTup = [];
	for (let key in labels) {
		const l = {
			key: key,
			value: labels![key],
		};
		labelsTup.push(l);
	}
	if (labelsTup != data.labelData) {
		data.labelData = labelsTup;
	}
};
const handAnnotations = (labels: { [key: string]: string }) => {
	const labelsTup = [];
	for (let key in labels) {
		const l = {
			key: key,
			value: labels![key],
		};
		labelsTup.push(l);
	}
	if (labelsTup != data.annotationsData) {
		data.annotationsData = labelsTup;
	}
};
const props = defineProps<{
	bindData;
}>();

const emit = defineEmits(['updateData']);

defineExpose({
	data,
});

watch(
	() => props.bindData,
	(newValue, oldValue) => {
		// console.log('反向更新数据，', props.bindData.metadata.labels);
		data.resourceType = props.bindData.resourceType;
		data.meta = props.bindData.metadata;
		//处理labels标签
		// console.log("反向更新数据，",props.bindData.metadata.labels)
		handleLabels(props.bindData.metadata.labels);
		handAnnotations(props.bindData.metadata.annotations);
		// if (props.bindData.metadata.annotations) {
		// 	data.annotationsData = handleLabels(props.bindData.metadata.annotations);
		// }
		data.replicas = props.bindData.replicas;
		enableEdit.value = true;
	},
	{ immediate: true, deep: true }
);

watch(
	() => [data.meta, data.replicas],
	() => {
		emit('updateData', data); //触发更新数据事件
	},
	{ immediate: true, deep: true }
);
const types = [
	{
		value: 'deployment',
		name: 'Deployment(无状态)',
	},
	{
		value: 'daemonSet',
		name: 'DaemonSet(守护进程集)',
	},
	{
		value: 'statefulSet',
		name: 'StatefulSet(有状态)',
	},
	{
		value: 'job',
		name: 'Job(任务)',
	},
	{
		value: 'cronJob',
		name: 'CronJob(定时任务)',
	},
];
</script>

<style scoped></style>
