<template>
	<div>
		<el-form-item label-width="90px" style="margin-bottom: 0">
			<template #label>
				<el-tooltip class="box-item" effect="light" content="用于挂载到容器中使用" placement="top-start" raw-content> 数据卷： </el-tooltip>
			</template>
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="handleSet">新增</el-button>
			<el-button
				v-if="data.show"
				type="info"
				v-show="data.volumeData.length != 0"
				text
				:icon="CaretTop"
				@click="data.show = !data.show"
				size="small"
				style="margin-left: 30px"
				>隐藏</el-button
			>
			<el-button v-else type="info" text :icon="CaretBottom" @click="data.show = !data.show" size="small" style="margin-left: 30px">展开</el-button>
		</el-form-item>
		<el-form-item label-width="0">
			<el-table
				:data="data.volumeData"
				v-show="data.volumeData.length != 0 && data.show"
				style="width: 100%; font-size: 10px"
				:cell-style="{ padding: '0,5px' }"
				:row-style="{ padding: '2px' }"
				:header-cell-style="{ padding: '5px' }"
				:header-row-style="{ padding: '5px' }"
			>
				<el-table-column prop="" label="类型" width="130">
					<template #default="scope">
						<el-select v-model="scope.row.type" size="small" @change="(val:string) => handleTypeChange(val, scope.$index)">
							<el-option v-for="item in data.typeList" :key="item.value" :label="item.label" :value="item.value" />
						</el-select>
					</template>
				</el-table-column>

				<el-table-column prop="name" label="名称" width="150">
					<template #default="scope">
						<el-input v-model="scope.row.name" size="small" />
					</template>
				</el-table-column>
				<el-table-column prop="" label="挂载源" width="200">
					<template #default="scope">
						<el-input
							v-if="scope.row.type === 'hostPath' && scope.row.hostPath"
							v-model="scope.row.hostPath.path"
							size="small"
							placeholder="主机路径：/tmp"
						/>
						<div v-if="scope.row.type === 'persistentVolumeClaim'" style="display: flex">
							<el-select
								v-model="scope.row.persistentVolumeClaim.claimName"
								size="small"
								:loading="data.loading"
								@click="getPvc"
								show-overflow-tooltip
							>
								<el-option v-for="item in data.pvcdata" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name" />
							</el-select>
						</div>
						<div v-if="scope.row.type === 'configMap'" style="display: flex">
							<el-select v-model="scope.row.configMap.name" size="small" :loading="data.loading" @click="getConfigMap()" show-overflow-tooltip>
								<el-option v-for="item in data.configMapData" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name" />
							</el-select>
							<el-button text type="primary" @click="openDialog(scope.row, scope.$index)" size="small" style="margin-left: 3px">
								<el-tooltip placement="top" effect="light">
									<template #content>
										<div
											v-for="(item, index) in scope.row.configMap.items"
											:key="index"
											style="display: flex; justify-content: space-between; width: 280px"
										>
											<span>Key : {{ item.key }}</span>
											<span>Path: {{ item.path }}</span>
										</div>
									</template>
									高级
								</el-tooltip>
							</el-button>
						</div>
						<div v-if="scope.row.type === 'secret'" style="display: flex">
							<el-select v-model="scope.row.secret.secretName" size="small" :loading="data.loading" @click="getSecret" show-overflow-tooltip>
								<el-option v-for="item in data.secretData" :key="item.metadata.uid" :label="item.metadata.name" :value="item.metadata.name" />
							</el-select>
							<el-button text type="primary" @click="openDialog(scope.row, scope.$index)" size="small" style="margin-left: 3px">
								<el-tooltip placement="top" effect="dark">
									<template #content>
										<div
											v-for="(item, index) in scope.row.secret.items"
											:key="index"
											style="display: flex; justify-content: space-between; width: 280px"
										>
											<span>Key : {{ item.key }}</span>
											<span>Path: {{ item.path }}</span>
										</div>
									</template>
									高级
								</el-tooltip>
							</el-button>
						</div>
						<span v-if="scope.row.type === 'tmp'">临时目录</span>
					</template>
				</el-table-column>
				<el-table-column prop="mountPath" label="容器挂载路径" width="150">
					<template #default="scope">
						<el-input v-model="scope.row.volumeMountData.mountPath" size="small" placeholder="容器路径：/app" />
					</template>
				</el-table-column>
				<el-table-column prop="subPath" label="容器子路径" width="120">
					<template #default="scope">
						<el-input v-model="scope.row.volumeMountData.subPath" size="small" placeholder="默认为空" />
					</template>
				</el-table-column>
				<el-table-column>
					<template #default="scope">
						<el-button :icon="RemoveFilled" type="danger" size="small" text @click="data.volumeData.splice(scope.$index, 1)"></el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-form-item>
		<el-dialog ref="dialogRef" v-model="dialogFormVisible" title="指定键：" width="400px" v-if="dialogFormVisible">
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="addKey">新增</el-button>
			<el-table
				:data="data.items"
				size="small"
				style="width: 100%; font-size: 10px"
				:cell-style="{ padding: '0,5px' }"
				:row-style="{ padding: '2px' }"
				:header-cell-style="{ padding: '5px' }"
				:header-row-style="{ padding: '5px' }"
			>
				<el-table-column label="键名">
					<template #default="scope">
						<el-select v-model="scope.row.key" placeholder="选择key" size="small">
							<el-option v-for="(item, key, index) in data.keyValData" :value="key" :key="index" :label="key"> {{ key }} </el-option>
						</el-select>
					</template>
				</el-table-column>
				<el-table-column label="挂载路径">
					<template #default="scope">
						<el-input v-model="scope.row.path" autocomplete="off" size="small" placeholder="请使用相对路径：tmp" />
					</template>
				</el-table-column>
				<el-table-column width="30">
					<template #default="scope">
						<el-button :icon="RemoveFilled" type="danger" size="small" text @click="data.items.splice(scope.$index, 1)"></el-button>
					</template>
				</el-table-column>
			</el-table>

			<template #footer>
				<span class="dialog-footer">
					<el-button @click="handleClose" size="small">Cancel</el-button>
					<el-button type="primary" @click="handleConfirm" size="small"> Confirm </el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { CaretBottom, CaretTop, CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { onMounted, onUnmounted, reactive, ref } from 'vue';
import { ConfigMap, PersistentVolumeClaim, Secret, Volume, VolumeMount, KeyToPath } from '@/types/kubernetes-types/core/v1';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;
import { kubernetesInfo } from '@/stores/kubernetes';
import { useConfigMapApi } from '@/api/kubernetes/configMap';
import { useSecretApi } from '@/api/kubernetes/secret';
import { usePVCApi } from '@/api/kubernetes/persitentVolumeClaim';
import mittBus from '@/utils/mitt';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { CreateK8SVolumeData } from '@/types/kubernetes/custom';

const k8sStore = kubernetesInfo();
const configMapApi = useConfigMapApi();
const pvcApi = usePVCApi();
const secretApi = useSecretApi();
const dialogFormVisible = ref(false);
const dialogRef = ref();

const data = reactive({
	loadFromParent: false,
	tmpData: {} as any,
	index: 0,
	keyValData: {} as { [key: string]: string } | undefined,
	set: false,
	show: true,
	typeList: [
		{
			label: '主机目录',
			value: 'hostPath',
		},
		{
			label: '加密字典',
			value: 'secret',
		},
		{
			label: '配置项',
			value: 'configMap',
		},
		{
			label: '临时目录',
			value: 'tmp',
		},
		{
			label: 'PVC',
			value: 'persistentVolumeClaim',
		},
	],
	volumeData: [] as Array<CreateK8SVolumeData>,
	pvcdata: [] as PersistentVolumeClaim[],
	tmpVolumes: [] as Volume[],
	volumes: [] as Volume[],
	volumeMount: [] as VolumeMount[],
	configMapData: [] as ConfigMap[],
	loading: false,
	secretData: [] as Secret[],
	form: {
		key: '',
		path: '',
	},
	items: [] as Array<KeyToPath>,
});

const addKey = () => {
	data.items.push({
		key: '',
		path: '',
	});
};

// 从接口获取configMap数据
const getConfigMap = async () => {
	await configMapApi.listConfigMap(k8sStore.state.creatDeployment.namespace, { cloud: k8sStore.state.activeCluster }).then((res) => {
		data.configMapData = res.data.data;
		// config.keySetShow = true;
	});
};

// 从接口获取secret数据
const getSecret = async () => {
	await secretApi.listSecret(k8sStore.state.creatDeployment.namespace, { cloud: k8sStore.state.activeCluster }).then((res) => {
		data.secretData = res.data.data;
	});
};
// 从接口获取pvc数据
const getPvc = async () => {
	await pvcApi.listPVC(k8sStore.state.creatDeployment.namespace, { cloud: k8sStore.state.activeCluster }).then((res) => {
		data.pvcdata = res.data.data;
	});
};
// 打开dialog,并处理初始化选项
const openDialog = (config: any, index: number) => {
	if (config.type === 'configMap') {
		if (config.configMap.items != undefined) {
			data.items = config.configMap.items;
		}
		dialogFormVisible.value = true;
		if (dialogFormVisible.value) {
			data.configMapData.forEach((item: ConfigMap) => {
				if (item.metadata?.name === config.configMap.name) {
					data.keyValData = item.data;
				}
			});
			data.index = index;
			data.tmpData = config;
		}
	} else if (config.type === 'secret') {
		if (config.secret.items != undefined) {
			data.items = config.secret.items;
		}
		dialogFormVisible.value = true;
		if (dialogFormVisible.value) {
			data.secretData.forEach((item) => {
				if (item.metadata?.name === config.secret.secretName) {
					data.keyValData = item.data;
				}
			});
			data.index = index;
			data.tmpData = config;
		}
	}
};
// 关闭dialog
const handleClose = () => {
	dialogFormVisible.value = false;
	//置空数据
	data.tmpData = {};
	data.index = 0;
	data.items = [];
};
// 确认指定特定的键值
const handleConfirm = () => {
	if (data.tmpData.type === 'configMap') {
		data.volumeData[data.index].configMap!.items = data.items;
	} else if (data.tmpData.type === 'secret') {
		data.volumeData[data.index].secret!.items = data.items;
	}
	handleClose();
};

// 转换volumeData 为k8s所需类型的数据
const handleVolumeData = () => {
	const tmpVolume = [] as Volume[];
	const tempVolumeMount = [] as VolumeMount[];

	data.volumeData.forEach((item: any, index: number) => {
		if (tempVolumeMount.length === index) {
			tempVolumeMount.push({} as VolumeMount);
		}
		tempVolumeMount[index].name = item.name;
		tempVolumeMount[index].mountPath = item.volumeMountData.mountPath;
		tempVolumeMount[index].subPath = item.volumeMountData.subPath;
		let flag = false;
		if (tmpVolume.length > 0) {
			tmpVolume.filter((i) => {
				if (i.name === item.name) {
					flag = true;
					return;
				}
			});
		}
		if (!flag) {
			let volume: Volume = {
				name: item.name,
			};
			switch (item.type) {
				case 'configMap':
					volume.configMap = item.configMap;
					break;
				case 'secret':
					volume.secret = item.secret;
					break;
				case 'hostPath':
					volume.hostPath = item.hostPath;
					break;
				case 'persistentVolumeClaim':
					volume.persistentVolumeClaim = item.persistentVolumeClaim;
					break;
				case 'tmp':
					volume.emptyDir = item.emptyDir;
					break;
			}
			tmpVolume.push(volume);
		}
	});
	data.volumeMount = tempVolumeMount;
	data.tmpVolumes = tmpVolume;
	data.volumes = tmpVolume;
};

// 添加volumeData数据
const handleSet = () => {
	const name = 'volume-' + uuid().toString().split('-')[1];
	data.volumeData.push({
		type: 'hostPath',
		name: name,
		hostPath: {
			path: '',
		},
		volumeMountData: {
			name: name,
			mountPath: '',
		},
	} as CreateK8SVolumeData);
};

// 根据不同的type初始化volumeData的值
const handleTypeChange = (type: string, index: number) => {
	// 切换type类型时初始化不同的值
	switch (type) {
		case 'hostPth':
			data.volumeData[index].hostPath = {
				path: '',
			};
			break;
		case 'secret':
			data.volumeData[index].secret = {};
			break;
		case 'configMap':
			data.volumeData[index].configMap = {};
			break;
		case 'tmp':
			data.volumeData[index].emptyDir = {};
			break;
		case 'persistentVolumeClaim':
			data.volumeData[index].persistentVolumeClaim = {
				claimName: '',
			};
			break;
	}
};

type propsType = {
	volumeMounts: Array<VolumeMount> | undefined;
	volumes: Array<Volume> | undefined;
};
//接受父组件传递的值
const props = defineProps<propsType>();

//解析volumeMount为所需要的CreateK8SVolumeData 类型
const parseVolumeMount = (volumeMount: Array<VolumeMount>) => {
	const tmpVolumeData = [] as Array<CreateK8SVolumeData>;
	volumeMount.forEach((item: VolumeMount) => {
		data.tmpVolumes.forEach((v: Volume) => {
			let volumeType = '';

			if (v.hostPath) {
				volumeType = 'hostPath';
			}
			if (v.secret) {
				volumeType = 'secret';
				getSecret();
			}
			if (v.configMap) {
				volumeType = 'configMap';
				getConfigMap();
			}
			if (v.emptyDir) {
				volumeType = 'tmp';
			}
			if (v.persistentVolumeClaim) {
				volumeType = 'persistentVolumeClaim';
				getPvc();
			}

			if (item.name === v.name) {
				tmpVolumeData.push({
					name: item.name,
					type: volumeType,
					emptyDir: v.emptyDir,
					secret: v.secret,
					configMap: v.configMap,
					persistentVolumeClaim: v.persistentVolumeClaim,
					hostPath: v.hostPath,
					volumeMountData: {
						name: item.name,
						mountPath: item.mountPath,
						subPath: item.subPath,
					} as VolumeMount,
				});
			}
		});
	});
	data.volumeData = tmpVolumeData;
};

onMounted(() => {
	if (!isObjectValueEqual(props.volumes, data.volumes) && props.volumes != undefined) {
		data.tmpVolumes = props.volumes;
		// parseVolumeMount(data.volumeMount);
	}

	if (props.volumeMounts && Object.keys(props.volumeMounts).length > 0) {
		parseVolumeMount(props.volumeMounts);
	}
});

const returnVolumeMounts = () => {
	handleVolumeData();
	return data.volumeMount;
};

const returnVolumes = () => {
	// handleVolumeData();
	console.log('返回volumes：：', data.volumes);
	return data.volumes;
};
defineExpose({
	returnVolumeMounts,
	returnVolumes,
});
</script>

<style scoped>
.el-card {
	border: none;
	box-shadow: none;
}
</style>
