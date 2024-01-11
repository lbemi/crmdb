<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="60%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div v-if="data.virtualService.metadata">
			<el-form :model="data.virtualService.metadata" ref="formRulesOneRef" label-width="100px" :rules="rules">
				<div>
					<el-collapse v-model="data.activeCollapse" @change="aggregate()">
						<el-collapse-item name="1">
							<template #title>
								基础设置
								<el-icon size="medium" color="#529b2e" class="no-inherit" v-if="data.baseStepStatus">
									<CircleCheck />
								</el-icon>
								<el-icon size="medium" color="#c45656" class="no-inherit" v-else>
									<CircleClose />
								</el-icon>
							</template>
							<el-form-item label="命名空间：" prop="namespace">
								<el-select
									v-model="data.virtualService.metadata.namespace"
									style="max-width: 200px"
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
							<el-form-item label="名称:" prop="name"
								><el-input :disabled="data.isUpdate" size="small" v-model="data.virtualService.metadata.name" style="max-width: 160px"></el-input>
							</el-form-item>
							<Label class="label" ref="labelRef" :labelData="data.labels" :name="'标签'" @updateLabels="getLabels" />
							<Label class="label" ref="annoRef" :labelData="data.annotations" :name="'注解'" @updateLabels="getAnnotations" />
						</el-collapse-item>
						<el-collapse-item name="2">
							<template #title>
								主机设置
								<el-icon size="medium" color="#529b2e" class="no-inherit" v-if="data.hostStepStatus">
									<CircleCheck />
								</el-icon>
								<el-icon size="medium" color="#c45656" class="no-inherit" v-else>
									<CircleClose />
								</el-icon>
							</template>
							<Hosts ref="hostRef" :hosts="data.virtualService.spec?.hosts" :name="'Hosts'" />
						</el-collapse-item>
						<el-collapse-item name="3">
							<template #title>
								HTTP设置
								<el-icon size="medium" color="#529b2e" class="no-inherit" v-if="data.httpStepStatus">
									<CircleCheck />
								</el-icon>
								<el-icon size="medium" color="#c45656" class="no-inherit" v-else>
									<CircleClose />
								</el-icon>
							</template>
							<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
								<VSHTTP ref="httpRef"></VSHTTP>
							</el-col>
						</el-collapse-item>
						<el-collapse-item title="Controllability" name="4">
							<div>Decision making: giving advices about operations is acceptable, but do not make decisions for the users;</div>
							<div>
								Controlled consequences: users should be granted the freedom to operate, including canceling, aborting or terminating current
								operation.
							</div>
						</el-collapse-item>
					</el-collapse>
				</div>
			</el-form>
			<div class="footer">
				<el-button size="small" @click="handleClose">取消</el-button>
				<el-button type="primary" size="small" @click="confirm">确认</el-button>
			</div>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer, ElMessage, FormInstance, FormRules } from 'element-plus';

import { VirtualService } from '@kubernetes-models/istio/networking.istio.io/v1beta1/VirtualService';
import { defineAsyncComponent, onMounted, reactive, ref } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useVirtualServiceApi } from '@/api/istio/virtualService';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { deepClone } from '@/utils/other';
import { CircleCheck, CircleClose, InfoFilled } from '@element-plus/icons-vue';
import yamlJs from 'js-yaml';

const formRulesOneRef = ref<FormInstance>();
const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));
const Hosts = defineAsyncComponent(() => import('@/components/istio/hosts.vue'));
const VSHTTP = defineAsyncComponent(() => import('@/components/istio/vshttp.vue'));

const k8sStore = kubernetesInfo();
const virtualServiceApi = useVirtualServiceApi();

const httpRef = ref();
const hostRef = ref();
const labelRef = ref();
const annoRef = ref();

const data = reactive({
	baseStepStatus: false,
	hostStepStatus: false,
	httpStepStatus: false,

	activeCollapse: '1',
	isBinaryData: false,
	isUpdate: false,
	visible: false,
	labels: [],
	annotations: [],
	host: '',
	virtualService2: {
		metadata: {
			name: '',
			namespace: k8sStore.state.activeNamespace,
			labels: {},
		},
		spec: {
			hosts: [''],
		},
	},
	virtualService: new VirtualService({
		metadata: {
			name: '',
			namespace: k8sStore.state.activeNamespace,
			labels: {},
		},
		spec: {
			hosts: [],
		},
	}),
	keyValues: [] as Array<{ key: string; value: string }>,
});

const rules = reactive<FormRules>({
	name: [
		{ required: true, message: '请输入名字', trigger: 'blur' },
		{ min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
	],
});

const emit = defineEmits(['update:visible', 'refresh']);

const props = defineProps({
	visible: Boolean,
	virtualService: {
		type: Object as () => VirtualService,
	},
	title: String,
});

const getAnnotations = (labels: any) => {
	data.virtualService.metadata!.annotations = labels;
};

const getLabels = (labels: any) => {
	data.virtualService.metadata!.labels = labels;
};

const handleHosts = async () => {
	data.hostStepStatus = await hostRef.value.validateHandler();
	data.virtualService.spec!.hosts! = hostRef.value.returnHosts();
};
const handleHttps = async () => {
	data.httpStepStatus = await httpRef.value.validateHandler();
	data.virtualService.spec!.http = httpRef.value.returnHttps();
};
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

const handleName = async () => {
	if (!formRulesOneRef.value) return false;
	await formRulesOneRef.value.validate((valid: boolean) => {
		data.baseStepStatus = valid;
	});
};
const aggregate = () => {
	//处理hosts字段
	handleHosts();
	handleHttps();
	handleName();
};
const confirm = async () => {
	aggregate();
	console.log(data.virtualService.validate());
	if (!data.isUpdate) {
		console.log(yamlJs.dump(data.virtualService));
		if (data.httpStepStatus && data.hostStepStatus && data.baseStepStatus) {
			await virtualServiceApi
				.createVirtualService({ cloud: k8sStore.state.activeCluster }, data.virtualService)
				.then(() => {
					ElMessage.success('创建成功');
					handleClose();
					emit('refresh');
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
		} else {
			ElMessage.error('请检查配置');
		}
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
<style scoped lang="scss">
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

.el-icon {
	margin-left: 10px;
}
:deep .el-collapse-item__header {
	background-color: #f1f1f1 !important;
	//position: relative;
	padding-left: 20px;
	margin-bottom: 5px;
	border: none;
	height: 35px;
}

//:deep .el-collapse-item__content {
//	text-align: left;
//	color: #fff;
//	background-color: #313743;
//	padding-bottom: 0;
//	div {
//		height: 0.96rem;
//		line-height: 0.96rem;
//	}
//}
</style>
