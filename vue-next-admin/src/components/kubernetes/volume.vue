<template>
	<div>
		<el-form-item label-width="90px" style="margin-bottom: 0">
			<template #label>
				<el-tooltip
					class="box-item"
					effect="light"
					content="自定义容器启动时运行的命令; 默认情况下，容器启动时将运行镜像默认命令"
					placement="top-start"
					raw-content
				>
					数据卷：
				</el-tooltip>
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
						<el-select v-model="scope.row.type" size="small">
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
						{{ scope.row.type }}
						<el-input
							v-if="scope.row.type === 'hostPath' && scope.row.hostPath"
							v-model="scope.row.hostPath.path"
							size="small"
							placeholder="主机路径：/tmp"
						/>
						<div v-if="scope.row.type === 'persistentVolumeClaim'" style="display: flex">
							<el-select v-model="scope.row.persistentVolumeClaim.name" size="small" :loading="data.loading" @click="getPvc" show-overflow-tooltip>
								<el-option v-for="item in data.pvcdata" :key="item.metadata!.name" :label="item.metadata!.name" :value="item.metadata!.name" />
							</el-select>
						</div>
						<div v-if="scope.row.type === 'configMap'" style="display: flex">
							<el-select
								v-model="scope.row.configMap.name"
								size="small"
								:loading="data.loading"
								@click="getConfigMap(scope.row)"
								show-overflow-tooltip
							>
								<el-option v-for="item in data.configMapData" :key="item.metadata!.name" :label="item.metadata!.name" :value="item.metadata!.name" />
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
								<el-option v-for="item in data.secretData" :key="item.metadata!.uid" :label="item.metadata!.name" :value="item.metadata!.name" />
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
						<!-- <el-input v-model="scope.row.hostPath.path" size="small" placeholder="主机路径：/tmp" v-if="scope.row.type === 'tmp'" /> -->
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
import { onUnmounted, reactive, ref, watch } from 'vue';
import { V1ConfigMap, V1PersistentVolumeClaim, V1Secret, V1Volume, V1VolumeMount } from '@kubernetes/client-node';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useConfigMapApi } from '/@/api/kubernetes/configMap';
import { useSecretApi } from '/@/api/kubernetes/secret';
import { V1KeyToPath } from '@kubernetes/client-node/dist/gen/model/v1KeyToPath';
import { usePVCApi } from '/@/api/kubernetes/persitentVolumeClaim';
import mittBus from '/@/utils/mitt';
import { isObjectValueEqual } from '/@/utils/arrayOperation';
import { CreateK8SVolumentData } from '/@/types/kubernetes/custom';

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
	volumeData: [] as Array<CreateK8SVolumentData>,
	pvcdata: [] as V1PersistentVolumeClaim[],
	tmpVolumes: [] as V1Volume[],
	volumes: [] as V1Volume[],
	volumeMount: [] as V1VolumeMount[],
	configMapData: [] as V1ConfigMap[],
	loading: false,
	secretData: [] as V1Secret[],
	form: {
		key: '',
		path: '',
	},
	items: [] as Array<V1KeyToPath>,
});

const addKey = () => {
	data.items.push({
		key: '',
		path: '',
	});
};
const getConfigMap = (config: any) => {
	configMapApi.listConfigMap(k8sStore.state.creatDeployment.namespace, { cloud: k8sStore.state.activeCluster }).then((res) => {
		data.configMapData = res.data.data;
		config.keySetShow = true;
	});
};
const getSecret = () => {
	secretApi.listSecret(k8sStore.state.creatDeployment.namespace, { cloud: k8sStore.state.activeCluster }).then((res) => {
		data.secretData = res.data.data;
	});
};
const getPvc = () => {
	pvcApi.listPVC(k8sStore.state.creatDeployment.namespace, { cloud: k8sStore.state.activeCluster }).then((res) => {
		data.pvcdata = res.data.data;
	});
};
const openDialog = (config: any, index: number) => {
	if (config.type === 'configMap') {
		if (config.configMap.items != undefined) {
			data.items = config.configMap.items;
		}
		dialogFormVisible.value = true;
		if (dialogFormVisible.value) {
			data.configMapData.forEach((item: V1ConfigMap) => {
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
const handleClose = () => {
	dialogFormVisible.value = false;
	//置空数据
	data.tmpData = {};
	data.index = 0;
	data.items = [];
};
const handleConfirm = () => {
	if (data.tmpData.type === 'configMap') {
		data.volumeData[data.index].configMap!.items = data.items;
	} else if (data.tmpData.type === 'secret') {
		data.volumeData[data.index].secret!.items = data.items;
	}
	handleClose();
};

const handleVolumeData = () => {
	const tmpVolume = [] as V1Volume[];
	const tempVolumeMount = [] as V1VolumeMount[];

	data.volumeData.forEach((item: any, index: number) => {
		if (tempVolumeMount.length === index) {
			tempVolumeMount.push({} as V1VolumeMount);
		}
		if (tmpVolume.length === index) {
			tmpVolume.push({} as V1Volume);
		}
		tmpVolume[index].name = item.name;
		tempVolumeMount[index].name = item.name;
		tempVolumeMount[index].mountPath = item.volumeMountData.mountPath;
		tempVolumeMount[index].subPath = item.volumeMountData.subPath;
		switch (item.type) {
			case 'configMap':
				tmpVolume[index].configMap = item.configMap;
				break;
			case 'secret':
				tmpVolume[index].secret = item.secret;
				break;
			case 'hostPath':
				tmpVolume[index].hostPath = item.hostPath;
				break;
			case 'persistentVolumeClaim':
				tmpVolume[index].persistentVolumeClaim = item.persistentVolumeClaim;
				break;
			case 'tmp':
				tmpVolume[index].emptyDir = item.emptyDir;
				break;
		}
	});
	data.volumeMount = tempVolumeMount;
	data.tmpVolumes = tmpVolume;
	data.volumes = tmpVolume;
};
const handleSet = () => {
	const name = 'volume-' + uuid().toString().split('-')[1];
	data.volumeData.push({
		keySet: false,
		keySetShow: false,
		type: 'hostPath',
		name: name,
		hostPath: {
			path: '',
		},
		secret: {},
		configMap: {},
		persistentVolumeClaim: {
			claimName: '',
		},
		emptyDir: {},
		volumeMountData: {
			name: name,
			mountPath: '',
			subPath: '',
		},
	} as CreateK8SVolumentData);
};

mittBus.on('updateDeploymentVolumes', (volumes: any) => {
	// if (!isObjectValueEqual(volumes, data.volumes)) {
	// 	data.loadFromParent = true;
	// 	console.log('--------------------', volumes);
	// 	data.tmpVolumes = volumes;
	// 	parseVolumeMount(data.volumeMount);
	// 	setTimeout(() => {
	// 		data.loadFromParent = false;
	// 	}, 100);
	// }
});

onUnmounted(() => {
	//卸载
	mittBus.off('updateDeploymentVolumes', () => {});
});

const props = defineProps({
	volumeMounts: Array<V1VolumeMount>,
});

const parseVolumeMount = (volumeMount: Array<V1VolumeMount>) => {
	const tmpVolumeMount = [] as Array<CreateK8SVolumentData>;
	volumeMount.forEach((item: V1VolumeMount) => {
		data.tmpVolumes.forEach((v: V1Volume) => {
			if (item.name === v.name) {
				tmpVolumeMount.push({
					name: item.name,
					type: 'hostPath',
					emptyDir: v.emptyDir,
					secret: v.secret,
					configMap: v.configMap,
					persistentVolumeClaim: v.persistentVolumeClaim,
					hostPath: v.hostPath,
					volumeMountData: {
						name: item.name,
						mountPath: item.mountPath,
						subPath: item.subPath,
					} as V1VolumeMount,
				});
			}
		});
	});
	console.log('#$^^^^^%%%%%%%%%%%%%%%%%', tmpVolumeMount, volumeMount, data.tmpVolumes);
	if (!isObjectValueEqual(data.volumeData, tmpVolumeMount)) data.volumeData = tmpVolumeMount;
};
watch(
	() => [props.volumeMounts, data.tmpVolumes],
	() => {
		if (props.volumeMounts && Object.keys(props.volumeMounts).length > 0) {
			data.loadFromParent = true;
			console.log('更新。。。。。。', props.volumeMounts);

			parseVolumeMount(props.volumeMounts);
			setTimeout(() => {
				data.loadFromParent = false;
			}, 100);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);

const emit = defineEmits(['updateVolumeMount']);
watch(
	() => [data.volumeData],
	() => {
		if (!data.loadFromParent) {
			console.log('触发更新volume--->', data.volumeData);
			handleVolumeData();
			emit('updateVolumeMount', data.volumeMount);
			mittBus.emit('updateVolumes', data.volumes);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>

<style scoped>
.el-card {
	border: none;
	box-shadow: none;
}
</style>
