<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="45%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div>
			<el-form :inline="true" class="demo-form-inline" v-model="data" :rules="formRules">
				<el-form-item label="命名空间：" prop="data.configMap.metadata.namespace">
					<el-select
						v-model="data.configMap.metadata.namespace"
						style="max-width: 180px"
						size="small"
						class="m-2"
						placeholder="Select"
						:disabled="data.isUpdate"
						><el-option key="all" label="所有命名空间" value="all"></el-option>
						<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata.name" :value="item.metadata.name" />
					</el-select>
				</el-form-item>
				<div>
					<el-form-item label="配置项名称:"
						><el-input :disabled="data.isUpdate" size="small" v-model="data.configMap.metadata.name"></el-input>
					</el-form-item>
				</div>
				<div>
					<el-form-item label="标签">
						<Label class="label" :labelData="data.labels" @updateLabels="getLabels" />
					</el-form-item>
				</div>

				<el-form-item label="注解">
					<Label class="label" :labelData="data.annotations" @updateLabels="getAnnotations" />
				</el-form-item>
				<el-form-item label="数据:">
					<div>
						<el-table :data="data.keyValues" style="width: 100%">
							<el-table-column label="名称" width="180">
								<template #default="scope">
									<el-input v-model="scope.row.key" size="small" />
								</template>
							</el-table-column>
							<el-table-column label="值" width="380">
								<template #default="scope">
									<el-input type="textarea" v-model="scope.row.value" size="small" />
								</template>
							</el-table-column>
							<el-table-column>
								<template #default="scope">
									<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.keyValues.splice(scope.$index, 1)"></el-button>
								</template>
							</el-table-column>
							<el-table-column width="200px">
								<template #default="scope">
									<el-upload
										ref="upload"
										class="upload-demo"
										:limit="1"
										:on-change="($evnet) => submitUpload($evnet, scope.row)"
										:on-exceed="handleExceed"
										:auto-upload="false"
									>
										<template #tip>
											<div class="el-upload__tip text-red">limit 1 file</div>
										</template>
										<template #trigger>
											<el-button type="primary" size="small" text>从文件上传</el-button>
										</template>
										<!--									<el-button class="ml-3" type="success" @click="submitUpload"> upload to server </el-button>-->
									</el-upload>
								</template>
							</el-table-column>
						</el-table>
					</div>
				</el-form-item>
				<div>
					<el-button size="small" @click="addKey()" style="width: 90%" type="primary" plain
						><el-icon><Plus /></el-icon>添加</el-button
					>
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
import { ElDrawer, ElMessage, FormRules, UploadFile } from 'element-plus';

import { ConfigMap } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import { ref } from 'vue';
import { genFileId } from 'element-plus';
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus';
import { kubernetesInfo } from '@/stores/kubernetes';
import { RemoveFilled, Plus } from '@element-plus/icons-vue';
import { useConfigMapApi } from '@/api/kubernetes/configMap';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { deepClone } from '@/utils/other';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));

const k8sStore = kubernetesInfo();
const configMapApi = useConfigMapApi();

const data = reactive({
	isUpdate: false,
	visible: false,
	labels: [],
	annotations: [],
	configMap: {
		metadata: {
			name: '',
			namespace: '',
		},
		data: {},
	} as ConfigMap,
	keyValues: [] as Array<{ key: string; value: string }>,
});

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
	configMap: {
		type: Object as () => ConfigMap,
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
	data.keyValues.forEach((item) => {
		let obj = {};
		obj[item.key] = item.value;
		Object.assign(data.configMap.data, obj);
	});
};

const getAnnotations = (labels: any) => {
	data.configMap.metadata!.annotations = labels;
};

const getLabels = (labels: any) => {
	data.configMap.metadata!.labels = labels;
};

const confirm = async () => {
	convertConfigMap();
	if (!data.isUpdate) {
		await configMapApi
			.createConfigMap({ cloud: k8sStore.state.activeCluster }, data.configMap)
			.then(() => {
				ElMessage.success('创建成功');
				handleClose();
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
	} else {
		await configMapApi
			.updateConfigMap({ cloud: k8sStore.state.activeCluster }, data.configMap)
			.then((res: any) => {
				ElMessage.success(res.message);
				handleClose();
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
	}
	console.log(data.configMap);
};

const handleClose = () => {
	emit('update:visible', false);
};

const convertConfigMapTo = () => {
	let kvs = [] as Array<{ key: string; value: string }>;
	Object.keys(data.configMap.data).forEach((k) => {
		kvs.push({
			key: k,
			value: data.configMap.data![k],
		});
	});
	data.keyValues = kvs;
	handleLabels(data.configMap.metadata!.labels!);
	handAnnotations(data.configMap.metadata!.annotations!);
};

onMounted(() => {
	data.isUpdate = false;
	data.visible = props.visible;
	if (!isObjectValueEqual(props.configMap, {})) {
		data.isUpdate = true;
		data.configMap = props.configMap;
		convertConfigMapTo();
	}
});

// watch(
// 	() => props,
// 	() => {
// 		console.log('----', props);
// 		data.isUpdate = false;
// 		data.visible = props.visible;
// 		if (!isObjectValueEqual(props.configMap, {})) {
// 			data.isUpdate = true;
// 			data.configMap = props.configMap;
// 			convertConfigMapTo();
// 		}
// 	},
// 	{
// 		deep: true,
// 		immediate: true,
// 	}
// );
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
