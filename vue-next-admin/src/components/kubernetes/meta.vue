<template>
	<div style="margin-top: 10px">
		<el-form ref="metaRef" :model="data.meta" v-if="data.meta" label-width="120px" label-position="left" :rules="metaRules" size="default">
			<el-form-item label="命名空间">
				<el-select v-model="data.meta.namespace" style="max-width: 220px" class="m-2" placeholder="Select" :disabled="props.isUpdate">
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata!.name" :label="item.metadata!.name" :value="item.metadata!.name!" />
				</el-select>
			</el-form-item>
			<el-form-item label="应用名称" prop="name">
				<el-input v-model="data.meta.name" style="width: 220px" :disabled="props.isUpdate" />
			</el-form-item>
			<el-form-item label="类型">
				<el-select v-model="data.resourceType" class="m-2" placeholder="Select" style="width: 220px" :disabled="enableEdit">
					<el-option v-for="item in types" :key="item.name" :label="item.name" :value="item.value" />
				</el-select>
			</el-form-item>
			<el-form-item label="副本数量" v-if="resourceType">
				<el-input-number v-model="data.replicas" :min="1" :max="100" />
			</el-form-item>
			<el-form-item label="标签">
				<Label :labelData="data.labelData" @updateLabels="getLabels" />
			</el-form-item>
			<el-form-item label="注解">
				<Label :labelData="data.annotationsData" @updateLabels="getAnnotations" />
			</el-form-item>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { computed, defineAsyncComponent, reactive, watch } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ref } from 'vue-demi';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { deepClone } from '@/utils/other';
import { CreateK8SBindData, CreateK8SMetaData, ResourceType } from '@/types/kubernetes/custom';
import type { FormInstance, FormRules } from 'element-plus';
import { ObjectMeta } from 'kubernetes-types/meta/v1';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));

const metaRef = ref<FormInstance>();
const k8sStore = kubernetesInfo();
const enableEdit = ref(false);

const emit = defineEmits(['updateData']);
// 当父组件未传递值是使用下面的初始化值
const data = reactive<CreateK8SMetaData>({
	loadFromParent: false,
	labelData: [{ key: 'app', value: '' }],
	annotationsData: [],
	replicas: 1,
	resourceType: 'deployment' as ResourceType,
	meta: {
		name: '',
		namespace: 'default',
		labels: { app: '' },
	},
});
const metaRules = reactive<FormRules>({
	name: [
		{ required: true, message: '请输入名称', trigger: 'blur' },
		{ min: 1, max: 50, message: '长度大于1个字符小于50个字符', trigger: 'blur' },
	],
});

const resourceType = computed(() => {
	switch (data.resourceType) {
		case 'deployment': {
			return true;
		}
		case 'daemonSet': {
			return true;
		}
		case 'statefulSet': {
			return true;
		}
		default:
			return false;
	}
});

const getLabels = (labels: any) => {
	if (!isObjectValueEqual(data.meta!.labels, labels)) {
		data.meta!.labels = labels;
	}
};

const getAnnotations = (labels: any) => {
	if (!isObjectValueEqual(data.meta!.annotations, labels)) {
		data.meta!.annotations = labels;
	}
	if (Object.keys(data.meta!.annotations!).length === 0) {
		delete data.meta?.annotations;
	}
};

const handleLabels = (label: { [key: string]: string }) => {
	const labels = JSON.parse(JSON.stringify(label));
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
	bindData?: CreateK8SBindData;
	isUpdate: boolean;
}>();

watch(
	() => props.bindData,
	() => {
		if (props.bindData) {
			data.loadFromParent = true;
			if (props.bindData.metadata) {
				data.meta = deepClone(props.bindData.metadata) as ObjectMeta;
			}
			//处理labels标签
			if (props.bindData.metadata) {
				if (props.bindData.metadata.labels) {
					handleLabels(props.bindData.metadata.labels);
				}
				if (props.bindData.metadata.annotations) {
					handAnnotations(props.bindData.metadata.annotations!);
				}
			}
			if (props.bindData.replicas) data.replicas = props.bindData.replicas;
			if (props.bindData.resourceType) {
				data.resourceType = props.bindData.resourceType;
				enableEdit.value = true;
			}

			setTimeout(() => {
				//延迟一下，不然会触发循环更新
				data.loadFromParent = false;
			}, 1);
		}
	},
	{ immediate: true, deep: true }
);

watch(
	() => [data.meta, data.replicas],
	(value, oldValue) => {
		// 监听变化，父组件传值时不更新
		if (!data.loadFromParent) {
			if (data.meta) {
				// 缓存namespace
				if (data.meta.namespace) {
					k8sStore.state.creatDeployment.namespace = data.meta.namespace;
				}
				// FIXME 优化，这个会处罚两次更新
				// if (data.meta.name && data.meta.name != '') {
				// 	data.meta.labels!.app = data.meta.name;
				// }
			}
			emit('updateData', data, metaRef.value); //触发更新数据事件
		}
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

<style scoped>
/* .el-form-item {
	margin-bottom: 4px;
} */
</style>
