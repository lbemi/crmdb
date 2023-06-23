<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="45%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div>
			<el-form :inline="true" class="demo-form-inline" v-model="data" :rules="formRules" label-position="right" label-width="auto">
				<el-form-item label="命名空间：" prop="namespace">
					<el-select
						v-if="data.secrets.metadata"
						v-model="data.secrets.metadata.namespace"
						style="max-width: 180px"
						size="small"
						class="m-2"
						placeholder="Select"
						:disabled="data.isUpdate"
						><el-option key="all" label="所有命名空间" value="all"></el-option>
						<el-option
							v-for="item in k8sStore.state.namespace"
							:key="item.metadata!.name"
							:label="item.metadata!.name!"
							:value="item.metadata!.name"
						/>
					</el-select>
				</el-form-item>
				<div>
					<el-form-item label="Secret名称:" v-if="data.secrets.metadata"
						><el-input :disabled="data.isUpdate" size="small" v-model="data.secrets.metadata.name"></el-input>
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
				<div>
					<el-form-item label="类型">
						<el-radio-group v-model="data.secrets.type" size="small">
							<el-radio-button v-for="(item, index) in secretType" :label="item.value" :key="index">{{ item.key }}</el-radio-button>
						</el-radio-group>
					</el-form-item>
				</div>

				<!--				<el-form-item label="数据:">-->
				<div v-if="data.secrets.type === 'Opaque'">
					<el-table :data="data.keyValues" style="width: 100%">
						<el-table-column label="名称" width="180">
							<template #default="scope">
								<el-input v-model="scope.row.key" size="small" />
							</template>
						</el-table-column>
						<el-table-column label="值" width="280">
							<template #default="scope">
								<el-input type="textarea" v-model="scope.row.value" size="small" :rows="1" />
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
									<!--										<template #tip>-->
									<!--											<div class="el-upload__tip text-red">limit 1 file</div>-->
									<!--										</template>-->
									<template #trigger>
										<el-button type="primary" size="small" text>从文件上传</el-button>
									</template>
								</el-upload>
							</template>
						</el-table-column>
					</el-table>
				</div>
				<div v-if="data.secrets.type === 'kubernetes.io/tls'">
					<el-form-item label="证书公钥名称">
						<el-input size="small" v-model="data.tlsValue[0].key" style="width: 400px"></el-input>
					</el-form-item>
					<div>
						<el-form-item label="证书公钥内容">
							<el-input
								placeholder="-----BEGIN CERTIFICATE-----"
								type="textarea"
								:rows="5"
								v-model="data.tlsValue[0].value"
								style="width: 400px"
							></el-input>
							<el-upload
								ref="upload"
								class="upload-demo"
								:limit="1"
								:on-change="($evnet) => uploadCert($evnet)"
								:on-exceed="handleExceed"
								:auto-upload="false"
							>
								<template #trigger>
									<el-button type="primary" size="small" text>从文件上传</el-button>
								</template>
							</el-upload>
						</el-form-item>
					</div>

					<el-form-item label="证书私钥名称">
						<el-input size="small" v-model="data.tlsValue[1].key" style="width: 400px"></el-input>
					</el-form-item>
					<div>
						<el-form-item label="证书私钥内容">
							<el-input
								placeholder="-----BEGIN PRIVATE KEY-----"
								type="textarea"
								:rows="5"
								v-model="data.tlsValue[1].value"
								style="width: 400px"
							></el-input>
							<el-upload
								ref="upload"
								class="upload-demo"
								:limit="1"
								:on-change="($evnet) => uploadKey($evnet)"
								:on-exceed="handleExceed"
								:auto-upload="false"
							>
								<template #trigger>
									<el-button type="primary" size="small" text>从文件上传</el-button>
								</template>
							</el-upload>
						</el-form-item>
					</div>
				</div>
				<div v-if="data.secrets.type === 'kubernetes.io/dockerconfigjson'">
					<el-form-item label="仓库地址:">
						<el-input v-model="data.registerInfo.register" size="small"></el-input>
					</el-form-item>
					<div>
						<el-form-item label="用户名:">
							<el-input v-model="data.registerInfo.username" size="small"></el-input>
						</el-form-item>
					</div>
					<el-form-item label="密码:">
						<el-input size="small" v-model="data.registerInfo.password"></el-input>
					</el-form-item>
					<el-form-item label="邮箱:">
						<el-input size="small" v-model="data.registerInfo.email"></el-input>
					</el-form-item>
				</div>
				<!--				</el-form-item>-->
				<div v-if="data.secrets.type === 'Opaque'">
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
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus';
import { ElDrawer, ElMessage, FormRules, genFileId, UploadFile } from 'element-plus';

import { Secret } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, onMounted, reactive, ref } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { Plus, RemoveFilled } from '@element-plus/icons-vue';
import { useSecretApi } from '@/api/kubernetes/secret';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { deepClone } from '@/utils/other';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));

const k8sStore = kubernetesInfo();
const secretApi = useSecretApi();

const data = reactive({
	isUpdate: false,
	visible: false,
	labels: [],
	annotations: [],
	keyValues: [] as Array<{ key: string; value: string }>,
	tlsValue: [
		{
			key: 'tls.crt',
			value: '',
		},
		{
			key: 'tls.key',
			value: '',
		},
	],
	registerInfo: {
		register: '',
		username: '',
		password: '',
		email: '',
	},
	secrets: <Secret>{
		metadata: {
			namespace: '',
			name: '',
		},
		type: 'Opaque',
		stringData: {},
	},
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
const uploadCert = (uploadFile: UploadFile | undefined) => {
	if (uploadFile?.raw) {
		let fileReader = new FileReader();
		fileReader.onload = async () => {
			data.tlsValue[0].value = JSON.stringify(fileReader.result);
		};
		fileReader.readAsText(uploadFile.raw);
	}
};
const uploadKey = (uploadFile: UploadFile | undefined) => {
	if (uploadFile?.raw) {
		let fileReader = new FileReader();
		fileReader.onload = async () => {
			data.tlsValue[1].value = JSON.stringify(fileReader.result);
		};
		fileReader.readAsText(uploadFile.raw);
	}
};
const addKey = () => {
	data.keyValues.push({
		key: '',
		value: '',
	});
};
const emit = defineEmits(['update:visible', 'refresh']);

type propsType = {
	visible: Boolean;
	secret: Secret | undefined;
	title: String;
};
const props = defineProps<propsType>();

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

const convertSecret = (kvs: Array<{ key: string; value: string }>) => {
	data.secrets.stringData = {};
	kvs.forEach((item) => {
		let obj = {};
		obj[item.key] = item.value;
		Object.assign(data.secrets.stringData, obj);
	});
};

const getAnnotations = (labels: any) => {
	data.secrets.metadata!.annotations = labels;
};

const getLabels = (labels: any) => {
	data.secrets.metadata!.labels = labels;
};

const convertRegister = () => {
	const authSecret = window.btoa(data.registerInfo.username + ':' + data.registerInfo.password);
	data.secrets.stringData = {
		'.dockerconfigjson': JSON.stringify({
			auths: {
				d: {
					username: data.registerInfo.username,
					password: data.registerInfo.password,
					email: data.registerInfo.email,
					auth: authSecret,
				},
			},
		}),
	};
};
const confirm = async () => {
	if (data.secrets.type === 'Opaque') {
		convertSecret(data.keyValues);
	} else if (data.secrets.type === 'kubernetes.io/tls') {
		convertSecret(data.tlsValue);
	} else if (data.secrets.type === 'kubernetes.io/dockerconfigjson') {
		convertRegister();
	}

	if (!data.isUpdate) {
		await secretApi
			.createSecret({ cloud: k8sStore.state.activeCluster }, data.secrets)
			.then(() => {
				ElMessage.success('创建成功');
				handleClose();
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
	} else {
		await secretApi
			.updateSecret({ cloud: k8sStore.state.activeCluster }, data.secrets)
			.then((res: any) => {
				ElMessage.success(res.message);
				handleClose();
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
	}
	console.log(data.secrets);
};

const handleClose = () => {
	emit('update:visible', false);
};

const convertSecretTo = () => {
	let kvs = [] as Array<{ key: string; value: string }>;
	Object.keys(data.secrets.data).forEach((k) => {
		kvs.push({
			key: k,
			value: data.secrets.data![k],
		});
	});
	data.keyValues = kvs;
	handleLabels(data.secrets.metadata!.labels!);
	handAnnotations(data.secrets.metadata!.annotations!);
};

onMounted(() => {
	data.isUpdate = false;
	data.visible = props.visible;
	if (!isObjectValueEqual(props.secret, {}) && props.secret != undefined) {
		data.isUpdate = true;
		data.secrets = props.secret;
		convertSecretTo();
	}
});
const secretType = [
	{
		key: 'Opaque',
		value: 'Opaque',
	},
	{
		key: '仓库登录密钥',
		value: 'kubernetes.io/dockerconfigjson',
	},
	{
		key: 'TLS证书',
		value: 'kubernetes.io/tls',
	},
];
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
.footer {
	margin-top: 20px;
}
.el-form-item {
	width: 90%;
}
</style>
