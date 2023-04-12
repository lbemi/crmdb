<template>
	<div>
		<el-form :model="data.container" label-width="120px" label-position="left">
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
				<el-form-item label="所需资源：">
					<div style="height: 28px">
						<span>CPU</span>
						<el-input placeholder="如：0.5" v-model.number="data.limit.cpu" size="small" style="width: 80px" />
						<span> Core</span>
						<el-divider direction="vertical" />
						<a>内存</a>
						<el-input placeholder="如：0.5" v-model.number="data.limit.memory" size="small" style="width: 80px" /><span> MiB</span>
					</div>
					<div style="font-size: 6px; color: #00bb00">
						<el-icon size="12px" color="#00bb00"><InfoFilled /></el-icon>建议根据实际使用情况设置，防止由于资源约束而无法调度或引发内存不足(OOM)错误
					</div>
				</el-form-item>
				<el-form-item label="资源限制：">
					<div style="height: 28px">
						<span>CPU</span>
						<el-input placeholder="如：0.5" v-model.number="data.require.cpu" size="small" style="width: 80px" />
						<span> Core</span>
						<el-divider direction="vertical" />
						<a>内存</a>
						<el-input placeholder="如：0.5" v-model.number="data.require.memory" size="small" style="width: 80px" /><span> MiB</span>
					</div>
				</el-form-item>
				<el-form-item>
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
				<el-form-item label="端口设置：">
					<el-button
						:icon="CirclePlusFilled"
						type="primary"
						size="small"
						text
						style="padding-left: 0"
						@click="data.ports.push({ name: '', containerPort: 0 })"
						>新增</el-button
					>
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
				<el-form-item>
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
						v-else="data.showLiveCheck"
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
				<el-form-item v-if="data.liveCheck && data.showLiveCheck">
					<HealthCheck ref="livenessRef" />
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
						v-else="data.showReadyCheck"
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
				<el-form-item v-if="data.readyCheck && data.showReadyCheck">
					<HealthCheck ref="readyRef" />
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
				<el-form-item v-if="data.startCheck && data.showStartCheck">
					<HealthCheck ref="startRef" />
				</el-form-item>
			</el-card>
			<el-card>
				<el-form-item label="生命周期：">
					<el-input v-model="data.container.image" size="default" style="width: 296px" />
				</el-form-item>
			</el-card>
		</el-form>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, ref, toRefs, watch } from 'vue';
import { V1Container, V1ContainerPort, V1EnvVar, V1SecurityContext } from '@kubernetes/client-node';
import { CaretBottom, CaretTop, CirclePlusFilled, InfoFilled, RemoveFilled } from '@element-plus/icons-vue';

const HealthCheck = defineAsyncComponent(() => import('./check.vue'));
const readyRef = ref<InstanceType<typeof HealthCheck>>();
const livenessRef = ref<InstanceType<typeof HealthCheck>>();
const startRef = ref<InstanceType<typeof HealthCheck>>();

interface envImp {
	name: string;
	value: string;
	otherValue: string;
	type: string;
}

// 格式化 env
const getContainer = () => {
	//先置空
	data.container.env = [];
	data.container.livenessProbe = {};
	data.container.startupProbe = {};
	data.container.readinessProbe = {};
	data.container.ports = data.ports;

	if (data.liveCheck) {
		data.container.livenessProbe = livenessRef.value.getData();
	}
	if (data.startCheck) {
		data.container.startupProbe = startRef.value.getData();
	}
	if (data.readyCheck) {
		data.container.readinessProbe = readyRef.value.getData();
	}
	data.env.forEach((item, index) => {
		if (item.type === 'custom') {
			//自定义变量
			const envVar: V1EnvVar = {
				name: item.name,
				value: item.value,
			};
			data.container.env![index] = envVar;
		} else if (item.type === 'fieldRef') {
			const envVar: V1EnvVar = {
				name: item.name,
				valueFrom: {
					fieldRef: {
						fieldPath: item.otherValue,
					},
				},
			};
			data.container.env![index] = envVar;
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
			data.container.env![index] = envVar;
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
			data.container.env![index] = envVar;
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
			data.container.env![index] = envVar;
		}
	});

	return data.container;
};

const data = reactive({
	liveCheck: false,
	showLiveCheck: true,
	readyCheck: false,
	showReadyCheck: true,
	startCheck: false,
	showStartCheck: true,
	containers: [] as V1Container[],
	container: {
		securityContext: {
			privileged: false,
		} as V1SecurityContext,
		livenessProbe: {},
		readinessProbe: {},
		startupProbe: {},
		env: [] as V1EnvVar[],
		ports: [] as V1ContainerPort,
	} as V1Container,
	limit: {
		cpu: 0,
		memory: 0,
	},
	require: {
		cpu: 0.5,
		memory: 500,
	},

	ports: [] as V1ContainerPort[],
	env: [] as envImp[],
});

const props = defineProps({
	container: V1Container,
});

const emit = defineEmits(['updateContainer']);

// defineExpose({
// 	getContainer,
// 	data,
// });
watch(
	() => props.container,
	() => {
		if (props.container) {
			data.container = props.container;
		}
		console.log('container:接受的容器：', props.container);
	},
	{
		deep: true,
		immediate: true,
	}
);
watch(
	() => data.container,
	() => {
		emit('updateContainer', data.container);
	},
	{
		deep: true,
		immediate: true,
	}
);

const imagePullPolicy = [
	{
		name: '优先使用本地镜像(ifNotPresent)',
		value: 'ifNotPresent',
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
		value: 'tcp',
	},
	{
		type: 'udp',
		value: 'udp',
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
