<template>
	<div>
		<el-form v-model="data.container" label-width="100px" label-position="left">
			<el-card shadow="hover">
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
					<el-select v-model="data.container.imagePullPolicy" class="m-2" placeholder="Select" size="default"
						style="width: 296px">
						<el-option v-for="item in imagePullPolicy" :key="item.name" :label="item.name"
							:value="item.value" />
					</el-select>
				</el-form-item>
				<el-tooltip class="box-item" effect="light"
					content="<div>即为该应用预留资源额度，包括CPU和内存两种资源，即容器独占该资源，</div><div> 防止因资源不足而被其他服务或进程争夺资源，导致应用不可用</div>"
					placement="top-start" raw-content>
					<el-form-item label="资源配置：">
						<el-button :icon="Edit" type="primary" size="small" text style="padding-left: 0"
							v-if="data.resourceSet" @click="setResource">配置</el-button>
						<el-button :icon="Delete" type="primary" size="small" text v-else style="padding-left: 0"
							@click="cancelResource">取消配置</el-button>
					</el-form-item>
				</el-tooltip>
				<div v-if="!data.resourceSet">
					<el-form-item v-if="data.container.resources">
						<div v-if="data.container.resources.requests">
							<span>所需资源： CPU</span>
							<el-input placeholder="如：0.5" v-model="data.container.resources.requests.cpu" size="small"
								class="limit" />
							<span class="limit"> Core</span>
							<el-divider direction="vertical" />
							<a>内存</a>
							<el-input placeholder="如：300Mi" v-model="data.container.resources.requests.memory" size="small"
								class="limit" /><span>
								(单位：MiB)</span>
						</div>
						<div style="font-size: 12px; color: #181a18">
							<el-icon size="13px" color="#181a18">
								<InfoFilled />
							</el-icon>建议根据实际使用情况设置，防止由于资源约束而无法调度或引发内存不足(OOM)错误
						</div>

						<div v-if="data.container.resources.limits">
							<span>资源限制： CPU</span>
							<el-input placeholder="如：0.5" v-model="data.container.resources.limits.cpu" size="small"
								class="limit" />
							<span class="limit"> Core</span>
							<el-divider direction="vertical" />
							<a>内存</a>
							<el-input placeholder="如：500Mi" v-model="data.container.resources.limits.memory" size="small"
								class="limit" /><span>
								(单位：MiB)</span>
						</div>
						<div style="font-size: 12px; color: #181a18">
							<el-icon size="13px" color="#181a18">
								<InfoFilled />
							</el-icon> 建议根据实际使用情况设置，防止因资源不足导致应用不可用
						</div>
					</el-form-item>
				</div>

				<el-form-item label="容器启动项：">
					<template #label>
						<el-tooltip class="box-item" effect="light"
							content="<div>默认情况下，容器是不可以访问宿主上的任何设备；特权容器则</div><div> 被授权访问宿主上所有设备，享有宿主上运行的进程的所有访问权限</div>"
							placement="top-start" raw-content>
							容器启动项：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.container.stdin" label="stdin" />
					<el-checkbox v-model="data.container.tty" label="tty" />
				</el-form-item>
				<el-form-item label="特权容器：">
					<template #label>
						<el-tooltip class="box-item" effect="light"
							content="<div>默认情况下，容器是不可以访问宿主上的任何设备；特权容器则</div><div> 被授权访问宿主上所有设备，享有宿主上运行的进程的所有访问权限</div>"
							placement="top-start" raw-content>
							特权容器：
						</el-tooltip>
					</template>
					<el-checkbox v-if="data.container.securityContext"
						v-model="data.container.securityContext.privileged" />
				</el-form-item>
			</el-card>
			<el-card shadow="hover">
				<Ports ref="portsRef" :ports="data.container.ports" />
			</el-card>
			<el-card shadow="hover">
				<Env ref="envsRef" :env="data.container.env" />
			</el-card>
			<el-card shadow="hover" v-show="!data.container.isIntiContainer">
				<el-form-item label="存活检查">
					<template #label>
						<el-tooltip effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							存活检查：
						</el-tooltip>
					</template>
					<HealthCheck ref="lifeCheckRef" :checkData="data.container.livenessProbe" />
				</el-form-item>
				<el-form-item label="就绪检查">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start"
							raw-content>
							就绪检查：
						</el-tooltip>
					</template>
					<HealthCheck ref="readyCheckRef" :checkData="data.container.readinessProbe" />
				</el-form-item>
				<el-form-item label="启动探测">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start"
							raw-content>
							启动探测：
						</el-tooltip>
					</template>
					<HealthCheck ref="startCheckRef" :checkData="data.container.startupProbe" />
				</el-form-item>
			</el-card>
			<el-card shadow="hover" v-show="!data.container.isIntiContainer">
				<el-form-item label="生命周期配置" />
				<el-form-item label="启动前：">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start"
							raw-content>
							启动前：
						</el-tooltip>
					</template>
					<LifeSet ref="preLifeRef" :lifeData="data.container.lifecycle?.postStart" />
				</el-form-item>
				<el-form-item label="停止前：">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start"
							raw-content>
							停止前：
						</el-tooltip>
					</template>
					<LifeSet ref="postLifeRef" :lifeData="data.container.lifecycle?.preStop" />
				</el-form-item>
			</el-card>
			<el-card shadow="hover">
				<CommandSet ref="startCmdRef" :args="data.container.args" :commands="data.container.command" />
			</el-card>
			<el-card shadow="hover">
				<VolumeMountDiv ref="volumeMountRef" :volumeMounts="data.container.volumeMounts" :volumes="props.volumes" />
			</el-card>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive, ref } from 'vue';
import { Container, ContainerPort, EnvVar, Volume } from 'kubernetes-types/core/v1';
import { Delete, Edit, InfoFilled } from '@element-plus/icons-vue';
import { deepClone } from '@/utils/other';
import { ContainerType } from '@/types/kubernetes/common';
import { isObjectValueEqual } from '@/utils/arrayOperation';

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
const portsRef = ref();
const envsRef = ref();

const data = reactive({
	loadFromParent: false,
	lifePostStartSet: false,
	lifePreStopSet: false,
	lifePreShow: true,
	lifeShow: true,
	resourceHasSet: false,
	resourceSet: true,
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

const getPorts = () => {
	const ports = portsRef.value.returnPorts();
	if (ports && ports.length != 0) {
		data.container.ports = deepClone(ports) as Array<ContainerPort>;
	} else {
		delete data.container.ports;
	}
};

const getEnvs = () => {
	const envs = envsRef.value.returnEnvs();
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
	data.resourceSet = !data.resourceSet;
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
};

const cancelResource = () => {
	data.resourceSet = !data.resourceSet;
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

onMounted(() => {
	const copyData = deepClone(props.container) as ContainerType;
	if (!copyData.securityContext) {
		copyData.securityContext = {
			privileged: false,
		};
	}
	if (!isObjectValueEqual(copyData.resources, {})) data.resourceSet = true;
	data.container = copyData;
});

const returnContainer = () => {
	getPorts();
	getLivenessData();
	getStartupData();
	getReadinessData();
	getPreStop();
	getPostStart();
	getCommands();
	getVolumeMounts();
	getEnvs();

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
.limit {
	margin-left: 10px;
	width: 90px;
}

// .btn {
// 	position: fixed;
// 	right: 50px;
// 	text-align: center;
// 	top: 50%;
// }

// .men {
// 	font-size: 13px;
// 	letter-spacing: 3px;
// }



.el-card {
	margin-bottom: 10px;
}
</style>
