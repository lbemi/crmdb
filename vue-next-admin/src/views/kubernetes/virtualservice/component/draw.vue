<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="45%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div v-if="data.virtualService.metadata">
			<el-form v-model="data.virtualService" ref="formRulesOneRef" label-width="100px" class="mt35">
				<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
					<el-form-item label="命名空间：" prop="data.virtualService.metadata.namespace">
						<el-select
							v-model="data.virtualService.metadata.namespace"
							style="max-width: 180px"
							size="small"
							class="m-2"
							placeholder="Select"
							:disabled="data.isUpdate"
						>
							<el-option
								v-for="item in k8sStore.state.namespace"
								:key="item.metadata?.name"
								:label="item.metadata?.name"
								:value="item.metadata!.name!"
							/>
						</el-select>
					</el-form-item>
				</el-col>
				<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
					<el-form-item label="虚拟服务名称:"
						><el-input :disabled="data.isUpdate" size="small" v-model="data.virtualService.metadata.name"></el-input>
					</el-form-item>
				</el-col>
				<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
					<Label class="label" :labelData="data.labels" :name="'标签'" @updateLabels="getLabels" />
				</el-col>
				<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
					<Label class="label" :labelData="data.annotations" :name="'注解'" @updateLabels="getAnnotations" />
				</el-col>
				<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
					<Hosts :hosts="data.virtualService.spec?.hosts" :name="'Hosts'" />
				</el-col>
			</el-form>
			<div class="footer">
				<el-button size="small" @click="handleClose">取消</el-button>
				<el-button type="primary" size="small" @click="confirm">确认</el-button>
			</div>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer, ElMessage, FormInstance, FormRules, UploadFile } from 'element-plus';

import { VirtualService } from '@kubernetes-models/istio/networking.istio.io/v1beta1/VirtualService';
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import { ref } from 'vue';
import { genFileId } from 'element-plus';
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useVirtualServiceApi } from '@/api/kubernetes/virtualService';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { deepClone } from '@/utils/other';
const formRulesOneRef = ref<FormInstance>();
const Label = defineAsyncComponent(() => import('@/components/istio/kubernetes/label.vue'));
const Hosts = defineAsyncComponent(() => import('@/components/istio/kubernetes/hosts.vue'));
const k8sStore = kubernetesInfo();
const virtualServiceApi = useVirtualServiceApi();

const data = reactive({
	isBinaryData: false,
	isUpdate: false,
	visible: false,
	labels: [],
	annotations: [],
	host: '',
	virtualService: {
		metadata: {
			name: '',
			labels: {},
		},
		spec: {
			hosts: [''],
		},
	} as VirtualService,
	keyValues: [] as Array<{ key: string; value: string }>,
});

const onAddRow = () => {
	data.virtualService.spec?.hosts?.push('');
};
const onDelRow = (k: number) => {
	data.virtualService.spec?.hosts?.splice(k, 1);
};

const formRules = reactive<FormRules>({});

const upload = ref<UploadInstance>();

const handleExceed: UploadProps['onExceed'] = (files) => {
	upload.value!.clearFiles();
	const file = files[0] as UploadRawFile;
	file.uid = genFileId();
	upload.value!.handleStart(file);
};

const submitUpload = (uploadFile: UploadFile | undefined, d: any) => {
	if (uploadFile?.raw) {
		let fileReader = new FileReader();
		fileReader.onload = async () => {
			d.value = fileReader.result;
		};
		fileReader.readAsText(uploadFile.raw);
		d.key = uploadFile.name;
	}
};

const addKey = () => {
	data.keyValues.push({
		key: '',
		value: '',
	});
};
const emit = defineEmits(['update:visible', 'refresh']);

const props = defineProps({
	visible: Boolean,
	virtualService: {
		type: Object as () => VirtualService,
	},
	title: String,
});

const handleLabels = (label: { [key: string]: string }) => {
	const labels = deepClone(label);
	const labelsTup = [];
	for (let key in labels) {
		const l = {
			key: key,
			value: labels![key],
		};
		labelsTup.push(l);
	}
	if (labelsTup != data.labels) {
		data.labels = labelsTup;
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
	if (labelsTup != data.annotations) {
		data.annotations = labelsTup;
	}
};

const convertConfigMap = () => {
	const res = {};
	data.keyValues.forEach((item) => {
		let obj = {};
		obj[item.key] = item.value;
		Object.assign(res, obj);
	});
};

const getAnnotations = (labels: any) => {
	data.virtualService.metadata!.annotations = labels;
};

const getLabels = (labels: any) => {
	data.virtualService.metadata!.labels = labels;
};

const confirm = async () => {
	convertConfigMap();
	if (!data.isUpdate) {
		console.log(data.virtualService);
		// await virtualServiceApi
		// 	.createVirtualService({ cloud: k8sStore.state.activeCluster }, data.virtualService)
		// 	.then(() => {
		// 		ElMessage.success('创建成功');
		// 		handleClose();
		// 		emit('refresh');
		// 	})
		// 	.catch((e: any) => {
		// 		ElMessage.error(e.message);
		// 	});
	} else {
		await virtualServiceApi
			.updateVirtualService({ cloud: k8sStore.state.activeCluster }, data.virtualService)
			.then((res: any) => {
				ElMessage.success(res.message);
				handleClose();
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
	}
};

const handleClose = () => {
	emit('update:visible', false);
};

const convertConfigMapTo = (obj: { [name: string]: string }) => {
	let kvs = [] as Array<{ key: string; value: string }>;
	Object.keys(obj).forEach((k) => {
		kvs.push({
			key: k,
			value: obj[k],
		});
	});
	data.keyValues = kvs;
	handleLabels(data.virtualService.metadata!.labels!);
	handAnnotations(data.virtualService.metadata!.annotations!);
};

onMounted(() => {
	data.isUpdate = false;
	data.visible = props.visible;
	if (props.virtualService && !isObjectValueEqual(props.virtualService, Object.create({}))) {
		data.isUpdate = true;
		data.virtualService = props.virtualService;
	}
});
</script>
<style scoped>
.el-form {
	margin-left: 20px;
	margin-top: 10px;
}
.footer {
	display: flex;
	margin-top: 50px;
	/*margin-left: 80px;*/
	justify-content: center;
}
.label {
	margin-top: 10px;
}
</style>
