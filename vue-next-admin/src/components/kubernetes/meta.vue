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
				<Label :labelData="data.labelData" @on-click="getLabels" v-bind:tableData="tableData" />
			</el-form-item>
			<el-form-item label="注解">
				<Label :labelData="data.annotationsData" @on-click="getAnnotations" />
			</el-form-item>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, watch } from 'vue';
import { V1ObjectMeta } from '@kubernetes/client-node';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { ref } from 'vue-demi';

const Label = defineAsyncComponent(() => import('/@/components/label/index.vue'));

const k8sStore = kubernetesInfo();
const enableEdit = ref(false);
const tableData = ref([]);
const data = reactive({
	labelData: [],
  annotationsData: [],
	replicas: 1,
	resourceType: 'deployment',
	meta: {
		namespace: 'default',
	} as V1ObjectMeta,
});

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
const getLabels = (labels: any) => {
	data.meta.labels = labels;
};
const getAnnotations = (labels: any) => {
	data.meta.annotations = labels;
};

const handleLables = (lables: { [key: string]: string }) => {
  const labelsTup = []
  Object.keys(lables).forEach((key)=>{
    labelsTup.push({key:lables[key]})
  })
  return labelsTup
};
const props = defineProps<{
	bindData;
}>();

const emit = defineEmits(['updateData']);

defineExpose({
	data,
});

watch(
	props,
	(newValue,oldValue) => {
		console.log('监测到原始code发生变化 了：', props.bindData.metadata.labels);
		data.resourceType = props.bindData.resourceType;
		data.meta = props.bindData.metadata;
		data.replicas = props.bindData.replicas;
    //处理labels标签
    if(props.bindData.metadata.labels) {
      data.labelData = handleLables(props.bindData.metadata.labels)
    }
    if(props.bindData.metadata.annotations) {
      data.annotationsData = handleLables(props.bindData.metadata.annotations)
    }
		enableEdit.value = true;
	},
	{ immediate: true, deep: true }
);

watch(
	data,
	() => {
		if (data) {
			console.log('子组件的data发生改变了。。。', data);
			emit('updateData', data); //触发更新数据事件
		}
	},
	{ immediate: true, deep: true }
);
</script>

<style scoped></style>
