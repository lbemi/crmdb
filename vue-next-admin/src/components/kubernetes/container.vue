<template>
	<div>
		<el-form v-model="data.container" label-width="100px" label-position="left">
			<el-card>
				<el-form-item label="初始化容器:">
					<el-checkbox v-model="data.container.isIntiContainer">设置为初始化容器</el-checkbox>
				</el-form-item>
				<el-form-item label="容器名称：">
					<el-input v-model="data.container.name" size="default" style="width: 296px" />
				</el-form-item>
				<el-form-item label="镜像名称：">
					<el-input v-model="data.container.image" size="default" style="width: 296px" />
				</el-form-item>
				<el-form-item label="拉取策略：">
					<el-select v-model="data.container.imagePullPolicy" class="m-2" placeholder="Select" size="default" style="width: 296px">
						<el-option v-for="item in imagePullPolicy" :key="item.name" :label="item.name" :value="item.value" />
					</el-select>
				</el-form-item>
				<el-form-item label="资源配置：">
					<el-button :icon="Edit" type="primary" size="small" text style="padding-left: 0" v-if="!data.resourceSet" @click="setResource"
						>配置</el-button
					>
					<el-button :icon="Delete" type="primary" size="small" text v-else style="padding-left: 0" @click="cancelResource">取消配置</el-button>
				</el-form-item>
				<el-form-item v-if="data.container.resources?.requests && data.resourceSet">
					<div style="height: 28px">
						<span>所需资源： CPU</span>
						<el-input placeholder="如：0.5" v-model="data.container.resources!.requests!.cpu" size="small" style="width: 80px" />
						<span> Core</span>
						<el-divider direction="vertical" />
						<a>内存</a>
						<el-input placeholder="如：300Mi" v-model="data.container.resources!.requests!.memory" size="small" style="width: 80px" /><span>
							(单位：MiB)</span
						>
					</div>
					<div style="font-size: 6px; color: #00bb00">
						<el-icon size="12px" color="#00bb00"><InfoFilled /></el-icon>建议根据实际使用情况设置，防止由于资源约束而无法调度或引发内存不足(OOM)错误
					</div>
				</el-form-item>
				<el-form-item v-if="data.resourceSet">
					<div style="height: 28px">
						<span>资源限制： CPU</span>
						<el-input placeholder="如：0.5" v-model="data.container.resources!.limits!.cpu" size="small" style="width: 80px" />
						<span> Core</span>
						<el-divider direction="vertical" />
						<a>内存</a>
						<el-input placeholder="如：500Mi" v-model="data.container.resources!.limits!.memory" size="small" style="width: 80px" /><span>
							(单位：MiB)</span
						>
					</div>
				</el-form-item>
				<el-form-item v-if="data.resourceSet">
					<el-tooltip
						class="box-item"
						effect="light"
						content="<div>即为该应用预留资源额度，包括CPU和内存两种资源，即容器独占该资源，</div><div> 防止因资源不足而被其他服务或进程争夺资源，导致应用不可用</div>"
						placement="top-start"
						raw-content
					>
						<div style="font-size: 6px; color: #00bb00">
							<el-icon size="12px" color="#00bb00"><InfoFilled /></el-icon> 建议根据实际使用情况设置，防止因资源不足导致应用不可用
						</div>
					</el-tooltip>
				</el-form-item>
				<el-form-item label="容器启动项：">
					<template #label>
						<el-tooltip
							class="box-item"
							effect="light"
							content="<div>默认情况下，容器是不可以访问宿主上的任何设备；特权容器则</div><div> 被授权访问宿主上所有设备，享有宿主上运行的进程的所有访问权限</div>"
							placement="top-start"
							raw-content
						>
							容器启动项：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.container.stdin" label="stdin" />
					<el-checkbox v-model="data.container.tty" label="tty" />
				</el-form-item>
				<el-form-item label="特权容器：">
					<template #label>
						<el-tooltip
							class="box-item"
							effect="light"
							content="<div>默认情况下，容器是不可以访问宿主上的任何设备；特权容器则</div><div> 被授权访问宿主上所有设备，享有宿主上运行的进程的所有访问权限</div>"
							placement="top-start"
							raw-content
						>
							特权容器：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.container.securityContext!.privileged" />
				</el-form-item>
			</el-card>
			<el-card>
				<Ports :ports="data.container.ports" @updatePort="getPorts" />
			</el-card>
			<el-card>
				<Env :env="data.container.env" @updateEnv="getEnvs" />
			</el-card>
			<el-card v-show="!data.container.isIntiContainer">
				<el-form-item label="存活检查">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							存活检查：
						</el-tooltip>
					</template>
					<HealthCheck ref="lifeCheckRef" :checkData="data.container.livenessProbe" />
				</el-form-item>
				<el-form-item label="就绪检查">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							就绪检查：
						</el-tooltip>
					</template>
					<HealthCheck ref="readyCheckRef" :checkData="data.container.readinessProbe" />
				</el-form-item>
				<el-form-item label="启动探测">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							启动探测：
						</el-tooltip>
					</template>
					<HealthCheck ref="startCheckRef" :checkData="data.container.startupProbe" />
				</el-form-item>
			</el-card>
			<el-card v-show="!data.container.isIntiContainer">
				<el-form-item label="生命周期配置" />
				<el-form-item label="启动前：">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							启动前：
						</el-tooltip>
					</template>
					<LifeSet ref="preLifeRef" :lifeData="data.container.lifecycle?.postStart" />
				</el-form-item>
				<el-form-item label="停止前：">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							停止前：
						</el-tooltip>
					</template>
					<LifeSet ref="postLifeRef" :lifeData="data.container.lifecycle?.preStop" />
				</el-form-item>
			</el-card>
			<el-card>
				<CommandSet ref="startCmdRef" :args="data.container.args" :commands="data.container.command" />
			</el-card>
			<el-card>
				<VolumeMountDiv ref="volumeMountRef" :volumeMounts="data.container.volumeMounts" :volumes="props.volumes" />
			</el-card>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive, ref, watch } from 'vue';
import { Container, ContainerPort, EnvVar, Volume } from 'kubernetes-types/core/v1';
import { Delete, Edit, InfoFilled } from '@element-plus/icons-vue';
import { deepClone } from '@/utils/other';
import { ContainerType } from '@/types/kubernetes/common';

//子组件引用
const HealthCheck = defineAsyncComponent(() => import('./check.vue'));
const LifeSet = defineAsyncComponent(() => import('./life.vue'));
const CommandSet = defineAsyncComponent(() => import('./startCommand.vue'));
const VolumeMountDiv = defineAsyncComponent(() => import('./volume.vue'));
const Ports = defineAsyncComponent(() => import('./port.vue'));
const Env = defineAsyncComponent(() => import('./env.vue'));

const lifeCheckRef = ref();
const readyCheckRef = ref();
const startCheckRef = ref();
const preLifeRef = ref();
const postLifeRef = ref();
const startCmdRef = ref();
const volumeMountRef = ref();
const data = reactive({
	loadFromParent: false,
	lifePostStartSet: false,
	lifePreStopSet: false,
	lifePreShow: true,
	lifeShow: true,
	resourceHasSet: false,
	resourceSet: false,
	containers: [] as Container[],
	container: <ContainerType>{
		name: '',
		imagePullPolicy: 'ifNotPresent',
		securityContext: {
			privileged: false,
		},
		// resources: {
		// 	limits: {
		// 		cpu: '',
		// 		memory: '0',
		// 	},
		// 	requests: {
		// 		cpu: '0.5',
		// 		memory: '500',
		// 	},
		// },
	},
});

const getPorts = (ports: Array<ContainerPort>) => {
	if (ports && ports.length != 0) {
		data.container.ports = deepClone(ports) as Array<ContainerPort>;
	} else {
		delete data.container.ports;
	}
};

const getEnvs = (envs: Array<EnvVar>) => {
	if (envs && envs.length != 0) {
		data.container.env = deepClone(envs) as Array<EnvVar>;
	} else {
		delete data.container.env;
	}
};
// 更新存活检查数据
const getLivenessData = () => {
	const { set, probe } = lifeCheckRef.value.returnHealthCheck();
	if (set) {
		data.container.livenessProbe = probe;
	} else {
		delete data.container.livenessProbe;
	}
};

const getReadinessData = () => {
	const { set, probe } = readyCheckRef.value.returnHealthCheck();
	if (set) {
		data.container.readinessProbe = probe;
	} else {
		delete data.container.readinessProbe;
	}
};

const getStartupData = () => {
	const { set, probe } = startCheckRef.value.returnHealthCheck();
	if (set) {
		data.container.startupProbe = probe;
	} else {
		delete data.container.startupProbe;
	}
};

const getPostStart = () => {
	const { set, lifeProbe } = postLifeRef.value.returnLife();
	if (set) {
		data.container.lifecycle = {
			postStart: lifeProbe,
		};
	} else {
		delete data.container.lifecycle?.postStart;
	}
};

const getPreStop = () => {
	const { set, lifeProbe } = preLifeRef.value.returnLife();
	if (set) {
		data.container.lifecycle = {
			preStop: lifeProbe,
		};
	} else {
		delete data.container.lifecycle?.preStop;
	}
};

const getCommands = () => {
	const { set, commands, args } = startCmdRef.value.returnStartCommand();
	if (set) {
		data.container.command = commands;
		data.container.args = args;
	} else {
		delete data.container.command;
		delete data.container.args;
	}
};

const getVolumeMounts = () => {
	data.container.volumeMounts = volumeMountRef.value.returnVolumeMounts();
};

const setResource = () => {
	data.resourceSet = true;
	data.container.resources = {
		limits: {
			cpu: '',
			memory: '',
		},
		requests: {
			cpu: '',
			memory: '',
		},
	};
	data.resourceHasSet = true;
};

const cancelResource = () => {
	data.resourceSet = false;
	delete data.container.resources?.limits;
	delete data.container.resources?.requests;
	data.resourceHasSet = false;
};

type propsType = {
	container: ContainerType;
	volumes: Array<Volume> | undefined;
	index: Number;
};
const props = defineProps<propsType>();

const emit = defineEmits(['updateContainer']);
onMounted(() => {
	const copyData = deepClone(props.container) as ContainerType;
	console.log(')))>>>', copyData);
	// if (!copyData.volumeMounts) {
	// 	copyData.volumeMounts = [] as VolumeMount[];
	// }
	//
	// if (copyData.resources?.limits || copyData.resources?.requests) {
	// 	data.resourceSet = true;
	// 	console.log(data.container.name, '************', data.container.resources?.requests.cpu, '-----', props.container.resources);
	// }

	if (!copyData.securityContext) {
		copyData.securityContext = {
			privileged: false,
		};
	}
	data.container = copyData;
});
// watch(
// 	() => props.container,
// 	() => {
//
// 		if (props.container && !isObjectValueEqual(data.container, props.container)) {
// 			data.loadFromParent = true;
//
// 			const copyData = deepClone(props.container) as ContainerType;
// 			if (!copyData.volumeMounts) {
// 				copyData.volumeMounts = [] as VolumeMount[];
// 			}
//
// 			if (copyData.resources?.limits || copyData.resources?.requests) {
// 				data.resourceSet = true;
// 				console.log(data.container.name, '************', data.container.resources?.requests.cpu, '-----', props.container.resources);
// 			}
//
// 			if (!copyData.securityContext) {
// 				copyData.securityContext = {
// 					privileged: false,
// 				};
// 			}
// 			if (!copyData.lifecycle) {
// 				copyData.lifecycle = {};
// 			}
//
// 			data.container = copyData;
//
// 			setTimeout(() => {
// 				//延迟一下，不然会触发循环更新
// 				data.loadFromParent = false;
// 			}, 1);
// 		}
// 	},
// 	{
// 		deep: true,
// 		immediate: true,
// 	}
// );
const returnContainer = () => {
	getLivenessData();
	getStartupData();
	getReadinessData();
	getPreStop();
	getPostStart();
	getCommands();
	getVolumeMounts();

	if (data.container.securityContext?.privileged && !data.container.securityContext?.privileged) {
		delete data.container.securityContext;
	}

	if (data.resourceSet && !data.resourceHasSet) {
		data.container.resources = {
			limits: {
				cpu: '',
				memory: '',
			},
			requests: {
				cpu: '',
				memory: '',
			},
		};
		data.resourceHasSet = true;
	}
	return { index: props.index, container: data.container, volumes: volumeMountRef.value.returnVolumes() };
};

defineExpose({
	returnContainer,
});

watch(
	() => [data.container, data.resourceSet],
	() => {
		// 父组件传值直接渲染，不触发循环更新
		if (!data.loadFromParent) {
			if (data.container.securityContext?.privileged && !data.container.securityContext?.privileged) {
				delete data.container.securityContext;
			}

			if (data.resourceSet && !data.resourceHasSet) {
				data.container.resources = {
					limits: {
						cpu: '',
						memory: '',
					},
					requests: {
						cpu: '',
						memory: '',
					},
				};
				data.resourceHasSet = true;
			}
			emit('updateContainer', props.index, data.container);
		}
	},
	{
		deep: true,
		immediate: true,
	}
);

const imagePullPolicy = [
	{
		name: '优先使用本地镜像(ifNotPresent)',
		value: 'IfNotPresent',
	},
	{
		name: '总是拉取镜像(Always)',
		value: 'Always',
	},
	{
		name: '仅使用本地镜像(Never)',
		value: 'Never',
	},
];
</script>

<style scoped lang="scss">
.btn {
	position: fixed;
	right: 50px;
	text-align: center;
	top: 50%;
}
.men {
	font-size: 13px;
	letter-spacing: 3px;
}
.el-form-item {
	margin-bottom: 2px;
}
.el-table-column {
	padding-top: 2px;
	padding-bottom: 2px;
}
.el-card {
	margin-bottom: 3px;
}
</style>
