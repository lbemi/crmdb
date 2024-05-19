<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="50%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div>
			<el-form v-model="data" :rules="formRules" label-width="90px">
				<el-form-item label="命名空间" prop="namespace">
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
							:value="item.metadata!.name!"
						/>
					</el-select>
				</el-form-item>
				<el-form-item label="不可修改:"
					><el-checkbox v-model="data.secrets.immutable" size="default" label="创建后不可修改" :disabled="data.isUpdate"></el-checkbox>
				</el-form-item>
				<el-form-item label="Secret名称" v-if="data.secrets.metadata"
					><el-input :disabled="data.isUpdate" size="small" v-model="data.secrets.metadata.name" style="width: 400px"></el-input>
				</el-form-item>
				<!-- <el-form-item> -->
				<Label :name="'标签'" :labelData="data.labels" @updateLabels="getLabels" :label-width="'90px'" />
				<!-- </el-form-item> -->

				<Label :name="'注解'" :labelData="data.annotations" @updateLabels="getAnnotations" :label-width="'90px'" />

				<el-form-item label="类型">
					<el-radio-group v-model="data.secrets.type" size="small" :disabled="data.isUpdate">
						<el-radio-button v-for="(item, index) in secretType" :label="item.value" :key="index">{{ item.key }}</el-radio-button>
					</el-radio-group>
				</el-form-item>
				<el-form-item>
					<div v-if="data.secrets.type === 'Opaque'">
						<el-table :data="data.keyValues" style="width: 100%">
							<el-table-column label="名称" width="180">
								<template #default="scope">
									<el-input v-model="scope.row.key" size="small" />
								</template>
							</el-table-column>
							<el-table-column label="值" width="380">
								<template #default="scope">
									<el-input type="textarea" v-model="scope.row.value" size="small" :rows="4" />
								</template>
							</el-table-column>
							<el-table-column width="30px">
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
											<div class="el-upload__tip">file with a size less than 500KB.</div>
										</template>
										<template #trigger>
											<el-button type="primary" size="small" text>从文件上传</el-button>
										</template>
									</el-upload>
								</template>
							</el-table-column>
						</el-table>
						<div class="flex-center">
							<el-button size="small" @click="addKey()" type="primary" plain
								><el-icon><Plus /></el-icon>添加</el-button
							>
						</div>
					</div>
					<div v-else-if="data.secrets.type === 'kubernetes.io/tls'">
						<!--					<el-form-item label="证书公钥名称">-->
						<!--						<el-input size="small" v-model="data.tlsValue[0].key" style="width: 400px"></el-input>-->
						<!--					</el-form-item>-->
						<div>
							<el-form-item label="证书公钥内容" label-width="auto">
								<el-input
									placeholder="-----BEGIN CERTIFICATE-----"
									type="textarea"
									:rows="5"
									v-model="data.tlsValue[0].value"
									style="width: 430px"
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

						<!--					<el-form-item label="证书私钥名称">-->
						<!--						<el-input size="small" v-model="data.tlsValue[1].key" style="width: 400px"></el-input>-->
						<!--					</el-form-item>-->
						<div>
							<el-form-item label="证书私钥内容" label-width="auto">
								<el-input
									placeholder="-----BEGIN PRIVATE KEY-----"
									type="textarea"
									:rows="5"
									v-model="data.tlsValue[1].value"
									style="width: 430px"
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
					<div v-else-if="data.secrets.type === 'kubernetes.io/dockerconfigjson'">
						<el-form label-width="auto">
							<el-form-item label="仓库地址:">
								<el-input v-model="data.registerInfo.register" size="small" style="width: 350px"></el-input>
							</el-form-item>
							<el-form-item label="用户名:">
								<el-input v-model="data.registerInfo.username" size="small"></el-input>
							</el-form-item>
							<el-form-item label="密码:">
								<el-input size="small" type="password" show-password v-model="data.registerInfo.password"></el-input>
							</el-form-item>
							<el-form-item label="邮箱:">
								<el-input size="small" v-model="data.registerInfo.email"></el-input>
							</el-form-item>
						</el-form>
					</div>
					<div v-else-if="data.secrets.type === 'kubernetes.io/basic-auth'">
						<el-form label-width="auto">
							<el-form-item label="用户名:" style="width: 350px">
								<el-input v-model="data.basicAuth.username" size="small"></el-input>
							</el-form-item>
							<el-form-item label="密码:" style="width: 350px">
								<el-input type="password" show-password size="small" v-model="data.basicAuth.password"></el-input>
							</el-form-item>
						</el-form>
					</div>
					<div v-else>
						<el-table :data="data.keyValues" style="width: 100%">
							<el-table-column label="名称" width="180">
								<template #default="scope">
									<el-input v-model="scope.row.key" size="small" />
								</template>
							</el-table-column>
							<el-table-column label="值" width="380">
								<template #default="scope">
									<el-input type="textarea" v-model="scope.row.value" size="small" :rows="4" />
								</template>
							</el-table-column>
							<el-table-column width="30px">
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
										<template #trigger>
											<el-button type="primary" size="small" text>从文件上传</el-button>
										</template>
									</el-upload>
								</template>
							</el-table-column>
						</el-table>
						<div>
							<el-button size="small" @click="addKey()" style="width: 90%" type="primary" plain
								><el-icon><Plus /></el-icon>添加</el-button
							>
						</div>
					</div>
				</el-form-item>
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

import { Secret } from 'kubernetes-models/v1';
import { defineAsyncComponent, onMounted, reactive, ref } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { Plus, RemoveFilled } from '@element-plus/icons-vue';
import { useSecretApi } from '@/api/kubernetes/secret';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { deepClone } from '@/utils/other';
import { MirrorRepository } from '@/types/kubernetes/common';

const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));

const k8sStore = kubernetesInfo();
const secretApi = useSecretApi();

const data = reactive({
	isUpdate: false,
	visible: false,
	labels: [] as Array<{ key: string; value: string }>,
	annotations: [] as Array<{ key: string; value: string }>,
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
	basicAuth: {
		username: '',
		password: '',
	},
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
		immutable: false,
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
			if (typeof fileReader.result === 'string') {
				data.tlsValue[0].value = fileReader.result;
			}
		};
		fileReader.readAsText(uploadFile.raw);
	}
};
const uploadKey = (uploadFile: UploadFile | undefined) => {
	if (uploadFile?.raw) {
		let fileReader = new FileReader();
		fileReader.onload = async () => {
			if (typeof fileReader.result === 'string') {
				data.tlsValue[1].value = fileReader.result;
			}
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
	visible: boolean;
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
	data.secrets.stringData = {} as { [key: string]: string };
	kvs.forEach((item) => {
		let obj = {} as { [key: string]: string };
		obj[item.key] = item.value;
		data.secrets.stringData = { ...data.secrets.stringData, ...obj };
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
				[data.registerInfo.register]: {
					username: data.registerInfo.username,
					password: data.registerInfo.password,
					email: data.registerInfo.email,
					auth: authSecret,
				},
			},
		}),
	};
};

const convertBasicAuth = () => {
	data.secrets.stringData = data.basicAuth;
};
const confirm = async () => {
	switch (data.secrets.type) {
		case 'Opaque':
			convertSecret(data.keyValues);
			break;
		case 'kubernetes.io/tls':
			convertSecret(data.tlsValue);
			break;
		case 'kubernetes.io/dockerconfigjson':
			convertRegister();
			break;
		case 'kubernetes.io/basic-auth':
			convertBasicAuth();
			break;
		default:
			convertSecret(data.keyValues);
			break;
	}
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
};

const handleClose = () => {
	emit('update:visible', false);
};

const convertSecretTo = () => {
	switch (data.secrets.type) {
		case 'kubernetes.io/tls':
			if (data.secrets.data) {
				data.tlsValue[0].value = atob(data.secrets.data['tls.crt']);
				data.tlsValue[1].value = atob(data.secrets.data['tls.key']);
			}
			break;
		case 'Opaque':
			let kvs = [] as Array<{ key: string; value: string }>;
			if (data.secrets.data) {
				Object.keys(data.secrets.data).forEach((k) => {
					kvs.push({
						key: k,
						value: atob(data.secrets.data![k]),
					});
				});
				data.keyValues = kvs;
			}
			break;
		case 'kubernetes.io/dockerconfigjson':
			if (data.secrets.data) {
				const obj = JSON.parse(decodeURI(atob(data.secrets.data['.dockerconfigjson']))) as MirrorRepository<string>;
				let register = '';
				for (let key in obj.auths) {
					register = key;
				}
				if (obj.auths[register]) {
					data.registerInfo = {
						register: register,
						username: obj.auths[register].username,
						password: obj.auths[register].password,
						email: obj.auths[register].email,
					};
				}
			}
			break;
		case 'kubernetes.io/basic-auth':
			if (data.secrets.data) {
				// const obj = JSON.parse(data.secrets.data);
				data.basicAuth.username = atob(data.secrets.data.username);
				data.basicAuth.password = atob(data.secrets.data.password);
			}
			break;
		default:
			let kvss = [] as Array<{ key: string; value: string }>;
			if (data.secrets.data && data.secrets.type) {
				Object.keys(data.secrets.data).forEach((k) => {
					kvss.push({
						key: k,
						value: atob(data.secrets.data![k]),
					});
				});
				data.keyValues = kvss;
				secretType.push({
					key: data.secrets.type,
					value: data.secrets.type,
				});
			}
			break;
	}

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
let secretType = [
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
	{
		key: '基础认证',
		value: 'kubernetes.io/basic-auth',
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
