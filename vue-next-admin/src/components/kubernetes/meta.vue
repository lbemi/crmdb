<template>
	<div>
		<el-form ref="metaRef" :model="state.meta" v-if="state.meta" :label-width="labelWidth" :rules="metaRules" size="default">
			<el-form-item label="命名空间">
				<el-select v-model="state.meta.namespace" style="max-width: 220px" placeholder="Select" :disabled="props.isUpdate">
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata!.name" :label="item.metadata!.name" :value="item.metadata!.name!" />
				</el-select>
			</el-form-item>
			<el-form-item label="应用名称" prop="name">
				<div>
					<el-input v-model="state.meta.name" style="width: 220px" :disabled="props.isUpdate" />
					<div class="k-description">最长63个字符，只能包含小写字母、数字及分隔符("-")</div>
				</div>
			</el-form-item>
			<el-form-item label="类型">
				<el-select v-model="state.resourceType" placeholder="Select" style="width: 220px" :disabled="enableEdit">
					<el-option v-for="item in resourceTypes" :key="item.name" :label="item.name" :value="item.value" />
				</el-select>
			</el-form-item>
			<el-form-item label="副本数量" v-if="resourceType !== 'task'">
				<el-input-number v-model="state.replicas" :min="1" :max="100" />
			</el-form-item>
			<Label ref="labelsRef" :labels="state.meta.labels" :name="'标签'" :label-width="labelWidth" />
			<Label ref="annotationRef" :labels="state.meta.annotations" :name="'注解'" :label-width="labelWidth" />
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { computed, defineAsyncComponent, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ref } from 'vue-demi';
import { deepClone } from '@/utils/other';
import { CreateK8SMeta, KubernetesResourceType } from '@/types/kubernetes/custom';
import type { FormInstance, FormRules } from 'element-plus';
import { IObjectMeta } from '@kubernetes-models/apimachinery/apis/meta/v1/ObjectMeta';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));
const labelsRef = ref();
const annotationRef = ref();
const metaRef = ref<FormInstance>();
const k8sStore = kubernetesInfo();
const enableEdit = ref(false);

// 当父组件未传递值是使用下面的初始化值
const state = reactive({
	replicas: 1,
	resourceType: <KubernetesResourceType>'deployment',
	meta: <IObjectMeta>{
		name: '',
		namespace: 'default',
		params: { app: '' },
		annotations: {},
	},
	validateRefs: <Array<FormInstance>>[],
});
const metaRules = reactive<FormRules>({
	name: [
		{ required: true, message: '请输入名称', trigger: 'blur' },
		{ min: 1, max: 50, message: '长度大于1个字符小于50个字符', trigger: 'blur' },
	],
});
const resourceType = computed(() => {
	return state.resourceType as KubernetesResourceType;
});

const props = defineProps<{
	metaData?: CreateK8SMeta;
	isUpdate: boolean;
	resourceType?: KubernetesResourceType;
	labelWidth?: string;
}>();

onMounted(() => {
	if (props.resourceType) {
		state.resourceType = props.resourceType;
	}
	if (props.metaData) {
		if (props.metaData.metadata) {
			state.meta = deepClone(props.metaData.metadata) as IObjectMeta;
		}

		if (props.metaData.replicas) state.replicas = props.metaData.replicas;
		if (props.metaData.resourceType) {
			state.resourceType = props.metaData.resourceType;
			enableEdit.value = true;
		}
	}
});

const getMeta = () => {
	state.validateRefs = [];
	if (labelsRef.value.ruleFormRef) {
		state.validateRefs.push(...labelsRef.value.ruleFormRef);
		state.meta.labels = labelsRef.value.getLabels();
	}
	if (annotationRef.value.ruleFormRef) {
		state.validateRefs.push(...annotationRef.value.ruleFormRef);
		state.meta.annotations = annotationRef.value.getLabels();
	}
	if (metaRef.value) {
		state.validateRefs.push(metaRef.value);
	}
	// 缓存namespace
	if (state.meta.namespace) {
		k8sStore.state.creatDeployment.namespace = state.meta.namespace;
	}
	return state;
};
defineExpose({
	getMeta,
});

const resourceTypes = [
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
	{
		value: 'task',
		name: 'Task(CI任务)',
	},
];
</script>
