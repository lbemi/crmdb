<template>
	<div>
		<el-form v-model="data.container" label-width="100px" label-position="left">
			<el-card>
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
						<el-input placeholder="如：0.5" v-model="data.container.resources.requests.cpu" size="small" style="width: 80px" />
						<span> Core</span>
						<el-divider direction="vertical" />
						<a>内存</a>
						<el-input placeholder="如：300Mi" v-model="data.container.resources.requests.memory" size="small" style="width: 80px" /><span>
							(单位：MiB)</span
						>
					</div>
					<div style="font-size: 6px; color: #00bb00">
						<el-icon size="12px" color="#00bb00"><InfoFilled /></el-icon>建议根据实际使用情况设置，防止由于资源约束而无法调度或引发内存不足(OOM)错误
					</div>
				</el-form-item>
				<el-form-item v-if="data.container.resources?.limits && data.resourceSet">
					<div style="height: 28px">
						<span>资源限制： CPU</span>
						<el-input placeholder="如：0.5" v-model="data.container.resources.limits.cpu" size="small" style="width: 80px" />
						<span> Core</span>
						<el-divider direction="vertical" />
						<a>内存</a>
						<el-input placeholder="如：500Mi" v-model="data.container.resources.limits.memory" size="small" style="width: 80px" /><span>
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
					<el-checkbox v-model="data.container.securityContext.privileged" />
				</el-form-item>
			</el-card>
			<el-card>
				<el-form-item label="端口设置：">
					<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="pushPort">新增</el-button>
				</el-form-item>
				<el-form-item>
					<div>
						<el-form :key="portIndex" v-for="(item, portIndex) in data.ports" style="display: flex">
							<el-form-item label="名称">
								<el-input v-model="item.name" size="small" style="width: 120px" />
							</el-form-item>
							<el-form-item label="容器端口" style="margin-left: 10px">
								<el-input-number v-model="item.containerPort" size="small" />
							</el-form-item>
							<el-form-item label="协议" style="margin-left: 10px">
								<el-select v-model="item.protocol" size="small" style="width: 80px">
									<el-option v-for="item in protocolType" :key="item.type" :label="item.type" :value="item.value" />
								</el-select>
							</el-form-item>
							<el-form-item>
								<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.ports.splice(portIndex, 1)"></el-button>
							</el-form-item>
						</el-form>
					</div>
				</el-form-item>
			</el-card>
			<el-card>
				<el-form-item label="环境变量：">
					<el-button
						:icon="CirclePlusFilled"
						type="primary"
						size="small"
						text
						style="padding-left: 0"
						@click="data.env.push({ type: 'custom', name: '', value: '', otherValue: '' })"
						>新增</el-button
					>
				</el-form-item>
				<el-form-item label-width="45">
					<el-table
						:data="data.env"
						style="width: 100%; font-size: 10px"
						v-show="data.env.length != 0"
						:cell-style="{ padding: '5px' }"
						:header-cell-style="{ padding: '5px' }"
					>
						<el-table-column prop="name" label="类型" width="130">
							<template #default="scope">
								<el-select v-model="scope.row.type" size="small">
									<el-option v-for="item in envType" :key="item.type" :label="item.type" :value="item.value" />
								</el-select>
							</template>
						</el-table-column>

						<el-table-column prop="name" label="变量名称" width="180">
							<template #default="scope">
								<el-input v-model="scope.row.name" size="small" />
							</template>
						</el-table-column>
						<el-table-column prop="name" label="变量/变量引用" width="290">
							<template #default="scope">
								<el-input v-model="scope.row.value" size="small" style="width: 120px" />
								<el-input v-model="scope.row.otherValue" size="small" v-if="scope.row.type != 'custom'" style="width: 120px; margin-left: 5px" />
							</template>
						</el-table-column>
						<el-table-column>
							<template #default="scope">
								<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.env.splice(scope.$index, 1)"></el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-form-item>
			</el-card>
			<el-card>
				<el-form-item label="存活检查">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							存活检查：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.liveCheck" label="开启" size="small" />
					<el-button
						v-if="data.showLiveCheck"
						type="info"
						v-show="data.liveCheck"
						text
						:icon="CaretTop"
						@click="data.showLiveCheck = !data.showLiveCheck"
						size="small"
						style="margin-left: 30px"
						>隐藏</el-button
					>
					<el-button
						v-else
						type="info"
						v-show="data.liveCheck"
						text
						:icon="CaretBottom"
						@click="data.showLiveCheck = !data.showLiveCheck"
						size="small"
						style="margin-left: 30px"
						>展开</el-button
					>
				</el-form-item>
				<el-form-item v-show="data.liveCheck && data.showLiveCheck">
					<HealthCheck :checkData="data.container.livenessProbe" @updateCheckData="getLivenessData" />
				</el-form-item>
				<el-form-item label="就绪检查">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							就绪检查：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.readyCheck" label="开启" size="small" />
					<el-button
						v-if="data.showReadyCheck"
						type="info"
						v-show="data.readyCheck"
						text
						:icon="CaretTop"
						@click="data.showReadyCheck = !data.showReadyCheck"
						size="small"
						style="margin-left: 30px"
						>隐藏</el-button
					>
					<el-button
						v-else
						type="info"
						v-show="data.readyCheck"
						text
						:icon="CaretBottom"
						@click="data.showReadyCheck = !data.showReadyCheck"
						size="small"
						style="margin-left: 30px"
						>展开</el-button
					>
				</el-form-item>
				<el-form-item v-show="data.readyCheck && data.showReadyCheck">
					<HealthCheck :checkData="data.container.readinessProbe" @updateCheckData="getReadinessData" />
				</el-form-item>
				<el-form-item label="启动探测">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							启动探测：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.startCheck" label="开启" size="small" />
					<el-button
						v-if="data.showStartCheck"
						type="info"
						v-show="data.startCheck"
						text
						:icon="CaretTop"
						@click="data.showStartCheck = !data.showStartCheck"
						size="small"
						style="margin-left: 30px"
						>隐藏</el-button
					>
					<el-button
						v-else="data.showStartCheck"
						type="info"
						v-show="data.startCheck"
						text
						:icon="CaretBottom"
						@click="data.showStartCheck = !data.showStartCheck"
						size="small"
						style="margin-left: 30px"
						>展开</el-button
					>
				</el-form-item>
				<el-form-item v-show="data.startCheck && data.showStartCheck">
					<HealthCheck :checkData="data.container.startupProbe" @updateCheckData="getStartupData" />
				</el-form-item>
			</el-card>
			<el-card>
				<el-form-item label="生命周期配置" />
				<el-form-item label="启动前：">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							启动前：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.lifePostStartSet" label="开启" size="small" />
					<el-button
						v-if="data.lifeShow"
						type="info"
						v-show="data.lifePostStartSet"
						text
						:icon="CaretTop"
						@click="data.lifeShow = !data.lifeShow"
						size="small"
						style="margin-left: 30px"
						>隐藏</el-button
					>
					<el-button
						v-else
						type="info"
						v-show="data.lifePostStartSet"
						text
						:icon="CaretBottom"
						@click="data.lifeShow = !data.lifeShow"
						size="small"
						style="margin-left: 30px"
						>展开</el-button
					>
				</el-form-item>
				<el-form-item>
					<LifeSet v-show="data.lifePostStartSet && data.lifeShow" :lifeData="data.container.lifecycle?.postStart" @updateLifeData="getPostStart" />
				</el-form-item>
				<el-form-item label="停止前：">
					<template #label>
						<el-tooltip class="box-item" effect="light" content="用来检查容器是否正常，不正常则重启容器" placement="top-start" raw-content>
							停止前：
						</el-tooltip>
					</template>
					<el-checkbox v-model="data.lifePreStopSet" label="开启" size="small" />
					<el-button
						v-if="data.lifePreShow"
						type="info"
						v-show="data.lifePreStopSet"
						text
						:icon="CaretTop"
						@click="data.lifePreShow = !data.lifePreShow"
						size="small"
						style="margin-left: 30px"
						>隐藏</el-button
					>
					<el-button
						v-else
						type="info"
						v-show="data.lifePreStopSet"
						text
						:icon="CaretBottom"
						@click="data.lifePreShow = !data.lifePreShow"
						size="small"
						style="margin-left: 30px"
						>展开</el-button
					>
				</el-form-item>
				<el-form-item>
					<LifeSet v-show="data.lifePreStopSet && data.lifePreShow" :lifeData="data.container.lifecycle?.preStop" @updateLifeData="getPreStop" />
				</el-form-item>
			</el-card>
			<el-card>
				<CommandSet :args="data.container.args" :commands="data.container.command" @updateCommand="getCommand" />
			</el-card>
			<el-card>
				<VolumeMount :volumeMounts="data.container.volumeMounts" @updateVolume="getVolumeMount" />
			</el-card>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, watch } from 'vue';
import { V1Container, V1ContainerPort, V1EnvVar, V1SecurityContext, V1Volume, V1VolumeMount } from '@kubernetes/client-node';
import { CaretBottom, CaretTop, CirclePlusFilled, Delete, Edit, InfoFilled, RemoveFilled } from '@element-plus/icons-vue';
import { isObjectValueEqual } from '/@/utils/arrayOperation';
import { V1LifecycleHandler } from '@kubernetes/client-node/dist/gen/model/v1LifecycleHandler';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;
import { deepClone } from '/@/utils/other';

const HealthCheck = defineAsyncComponent(() => import('./check.vue'));
const LifeSet = defineAsyncComponent(() => import('./life.vue'));
const CommandSet = defineAsyncComponent(() => import('./startCommand.vue'));
const VolumeMount = defineAsyncComponent(() => import('./volume.vue'));

interface envImp {
	name: string;
	value: string;
	otherValue: string;
	type: string;
}

const data = reactive({
	lifePostStartSet: false,
	lifePreStopSet: false,
	lifePreShow: true,
	lifeShow: true,
	resourceHasSet: false,
	liveCheck: false,
	showLiveCheck: true,
	readyCheck: false,
	showReadyCheck: true,
	startCheck: false,
	showStartCheck: true,
	resourceSet: false,
	containers: [] as V1Container[],
	container: {
		name: '',
		imagePullPolicy: 'ifNotPresent',
		securityContext: {
			privileged: false,
		} as V1SecurityContext,
		// livenessProbe: {},
		// readinessProbe: {},
		// startupProbe: {},
		// env: [] as V1EnvVar[],
		// ports: [] as V1ContainerPort,
		// resources: {
		// 	limits: {
		// 		cpu: '',
		// 		memory: '',
		// 	},
		// 	requests: {
		// 		cpu: '',
		// 		memory: '',
		// 	},
		// },
	} as V1Container,
	limit: {
		cpu: '',
		memory: 0,
	},
	require: {
		cpu: 0.5,
		memory: 500,
	},

	ports: [] as V1ContainerPort[],
	env: [] as envImp[],
});

const getVolumeMount = (volumeMounts: any) => {
	data.container.volumeMounts = volumeMounts;
	console.log('>>>>>>>>>>volumeMounts', volumeMounts, data.container.volumeMounts);
};
const getCommand = (c: any) => {
	data.container.command = c.commands;
	data.container.args = c.args;
};
const getPostStart = (postStart: V1LifecycleHandler) => {
	if (data.lifePostStartSet && data.container.lifecycle) {
        console.log("kdsjhfjhkasgdfjkhsdjkfhsdjfsdhjfashjdfasjfgsajhfgsdjhfgsjfdgk");
        
		data.container.lifecycle.postStart = postStart;
	} else {
		delete data.container.lifecycle?.postStart;
	}
};
const getPreStop = (preStop: V1LifecycleHandler) => {
	if (data.lifePreStopSet) {
		data.container.lifecycle!.preStop = preStop;
	} else {
		delete data.container.lifecycle?.preStop;
	}
};
// 更新存活检查数据
const getLivenessData = (liveData: {}) => {
	if (data.liveCheck) {
		data.container.livenessProbe = liveData;
	}
};
const getReadinessData = (readData: {}) => {
	if (data.readyCheck) {
		data.container.readinessProbe = readData;
	}
};
const getStartupData = (startData: {}) => {
	if (data.startCheck) {
		data.container.startupProbe = startData;
	}
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
const pushPort = () => {
	const name = uuid().toString().split('-')[1];
	data.ports.push({ name: 'p-' + name, containerPort: 80, protocol: 'TCP' });
};

const cancelResource = () => {
	data.resourceSet = false;
	delete data.container.resources?.limits;
	delete data.container.resources?.requests;
	data.resourceHasSet = false;
};
const buildEnv = () => {
	const envData = [] as V1EnvVar[];
	const envTup = JSON.parse(JSON.stringify(data.env));
	envTup.forEach((item, index) => {
		if (item.type === 'custom') {
			//自定义变量
			const envVar: V1EnvVar = {
				name: item.name,
				value: item.value,
			};
			envData[index] = envVar;
		} else if (item.type === 'fieldRef') {
			const envVar: V1EnvVar = {
				name: item.name,
				valueFrom: {
					fieldRef: {
						fieldPath: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		} else if (item.type === 'resourceFieldRef') {
			const envVar: V1EnvVar = {
				name: item.name,
				valueFrom: {
					resourceFieldRef: {
						containerName: item.value,
						resource: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		} else if (item.type === 'configMapKeyRef') {
			const envVar: V1EnvVar = {
				name: item.name,
				valueFrom: {
					configMapKeyRef: {
						name: item.value,
						key: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		} else if (item.type === 'secretKeyRef') {
			const envVar: V1EnvVar = {
				name: item.name,
				valueFrom: {
					secretKeyRef: {
						name: item.value,
						key: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		}
	});

	if (!isObjectValueEqual(data.env, envData)) {
		data.container.env = envData;
	}
};
const parseEnv = (envs: Array<V1EnvVar>) => {
	const envData = [] as envImp[];
	envs.forEach((env, index) => {
		envData.push({
			name: '',
			value: '',
			type: '',
			otherValue: '',
		});
		envData[index].name = env.name;
		if (env.valueFrom) {
			if (env.valueFrom.fieldRef) envData[index].type = 'fieldRef';
			if (env.valueFrom.fieldRef?.fieldPath) envData[index].value = env.valueFrom.fieldRef.fieldPath;

			if (env.valueFrom.secretKeyRef) envData[index].type = 'secretKeyRef';
			if (env.valueFrom.secretKeyRef?.key) envData[index].otherValue = env.valueFrom.secretKeyRef.key;
			if (env.valueFrom.secretKeyRef?.name) envData[index].value = env.valueFrom.secretKeyRef.name;

			if (env.valueFrom.configMapKeyRef) envData[index].type = 'configMapKeyRef';
			if (env.valueFrom.configMapKeyRef?.key) envData[index].otherValue = env.valueFrom.configMapKeyRef.key;
			if (env.valueFrom.configMapKeyRef?.name) envData[index].value = env.valueFrom.configMapKeyRef.name;

			if (env.valueFrom.resourceFieldRef) envData[index].type = 'resourceFieldRef';
			if (env.valueFrom.resourceFieldRef?.resource) envData[index].otherValue = env.valueFrom.resourceFieldRef.resource;
			if (env.valueFrom.resourceFieldRef?.containerName) envData[index].value = env.valueFrom.resourceFieldRef.containerName;
		}
		if (env.value) {
			envData[index].value = env.value;
			envData[index].type = 'custom';
		}
	});
	console.log('解析从COde传递来的的ENV数据', envData);
	data.env = envData;
};
const props = defineProps({
	container: Object,
	index: Number,
});

const emit = defineEmits(['updateContainer']);

watch(
	() => props.container,
	() => {
		if (props.container && !isObjectValueEqual(data.container, props.container)) {
			console.log('YYYYYYYYYYYYY', props.container);
			const copyData = deepClone(props.container) as V1Container;
			if (!data.container.volumeMounts) {
				data.container.volumeMounts = [] as V1VolumeMount[];
			}

			if (!data.container.lifecycle) {
				data.container.lifecycle = {};
			}

			if (copyData.env && copyData.env.length != 0) {
				parseEnv(copyData.env);
			}
			// }
			data.container = copyData;
			if (copyData.ports) data.ports = copyData.ports;
		}
	},
	{
		deep: true,
		immediate: true,
	}
);

watch(
	() => [data.container, data.ports, data.liveCheck, data.readyCheck, data.startCheck, data.resourceSet, data.lifePostStartSet],
	() => {
		console.log('1.触发updateContainer，>>>>>', data.container);
		// if (data.container.name != k8sStore.state.creatDeployment.name) {
		// 	data.container.name = k8sStore.state.creatDeployment.name;
		// }
		if (data.ports && data.ports.length != 0) {
			data.container.ports = data.ports;
		} else {
			delete data.container.ports;
		}

		if (data.container.securityContext?.privileged && !data.container.securityContext?.privileged) {
			delete data.container.securityContext;
		}

		if (!data.lifePostStartSet) {
			delete data.container.lifecycle?.postStart;
		}
		if (!data.lifePreStopSet) {
			delete data.container.lifecycle?.preStop;
		}
		if (!data.liveCheck) {
			delete data.container.livenessProbe;
		}
		if (!data.startCheck) {
			delete data.container.startupProbe;
		}
		if (!data.readyCheck) {
			delete data.container.readinessProbe;
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
	},
	{
		deep: true,
		immediate: true,
	}
);
watch(
	() => data.env,
	() => {
		if (data.env && data.env.length != 0) {
			buildEnv();
		} else {
			delete data.container.env;
		}
		emit('updateContainer', props.index, data.container);
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
const envType = [
	{
		type: '配置项',
		value: 'configMapKeyRef',
	},
	{
		type: '资源引用',
		value: 'fieldRef',
	},
	{
		type: '资源引用2',
		value: 'resourceFieldRef',
	},
	{
		type: '加密字典',
		value: 'secretKeyRef',
	},
	{
		type: '自定义',
		value: 'custom',
	},
];
const protocolType = [
	{
		type: 'tcp',
		value: 'TCP',
	},
	{
		type: 'udp',
		value: 'UDP',
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
