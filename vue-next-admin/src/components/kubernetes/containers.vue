<template>
	<div class="layout-pd">
		<div>
			<el-tabs v-model="editableTabsValue" type="border-card" @tab-remove="tabRemove" class="demo-tabs" :before-leave="tabChange">
				<el-tab-pane v-for="(item, index) in data.containers" :key="index" :name="index" :closable="data.containers.length != 1">
					<template #label>
						<span v-if="item.isIntiContainer" class="custom-tabs-label">
							<SvgIcon name="iconfont icon-container-" class="svg" />{{ ' init容器 ' }}
						</span>
						<span v-else class="custom-tabs-label"> <SvgIcon name="iconfont icon-container-" class="svg" />{{ ' 容器 ' }} </span>
					</template>
					<ContainerDiv :ref="(el:refItem)=>setItemRef(el)" :container="item" :index="index" :volumes="props.volumes" />
				</el-tab-pane>
				<el-tab-pane key="CustomBtn" name="CustomBtn" :closable="false">
					<template #label>
						<el-link type="primary" :underline="false" :icon="Plus"></el-link>
					</template>
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import { ComponentPublicInstance, ref } from 'vue-demi';
import { Container, Volume } from 'kubernetes-types/core/v1';
import type { TabPaneName } from 'element-plus';
import { deepClone } from '@/utils/other';
import { Plus } from '@element-plus/icons-vue';
import { ContainerType } from '@/types/kubernetes/common';

const ContainerDiv = defineAsyncComponent(() => import('./container.vue'));

type refItem = Element | ComponentPublicInstance | null;
const editableTabsValue = ref(0);
const itemRefs = ref([]);

const setItemRef = (el: refItem) => {
	if (el) {
		itemRefs.value.push(el);
	}
};
const tabChange = (currentName: TabPaneName) => {
	if (currentName === 'CustomBtn') {
		tabAdd();
		return false;
	}
};

const tabAdd = () => {
	const newTabName = data.containers.length;
	data.containers.push(data.container);
	editableTabsValue.value = newTabName;
};

const tabRemove = (TabPaneName: TabPaneName) => {
	if (data.containers.length <= 1) {
		return false;
	}
	const tabs = deepClone(data.containers);
	data.containers = tabs.filter((tab, index) => index !== TabPaneName);
	editableTabsValue.value = data.containers.length - 1;
};

const data = reactive({
	isInitContainer: false,
	loadFromParent: false,
	currentIndex: 1,
	addIndex: 1,
	tabIndex: 1,
	editableTabsValue: '1',
	volumes: [] as Volume[],
	containers: [] as Array<ContainerType>,
	container: <ContainerType>{
		isIntiContainer: false,
		name: '',
		image: '',
		imagePullPolicy: 'IfNotPresent',
		securityContext: {
			privileged: false,
		},
		// ports: [],
		// env: [],
		// resources: {},
		// livenessProbe: {},
		// readinessProbe: {},
		// startupProbe: {},
		// lifecycle: {},
	},
});

const getContainers = () => {
	itemRefs.value.forEach((refValue) => {
		const { index, container, volumes } = refValue.returnContainer();
		data.volumes = volumes;
		data.containers[index] = deepClone(container) as ContainerType;
	});
};

type propsType = {
	containers: Array<Container>;
	initContainers: Array<Container>;
	volumes: Array<Volume> | undefined;
};
const props = defineProps<propsType>();

onMounted(() => {
	const containers = [] as ContainerType[];
	if (props.containers && props.containers.length > 0) {
		props.containers.forEach((item: Container) => {
			const container: ContainerType = {
				args: item.args,
				command: item.command,
				env: item.env,
				envFrom: item.envFrom,
				image: item.image,
				imagePullPolicy: item.imagePullPolicy,
				isIntiContainer: false,
				lifecycle: item.lifecycle,
				livenessProbe: item.livenessProbe,
				name: item.name,
				ports: item.ports,
				readinessProbe: item.readinessProbe,
				resources: item.resources,
				securityContext: item.securityContext,
				startupProbe: item.startupProbe,
				stdin: item.stdin,
				stdinOnce: item.stdinOnce,
				terminationMessagePath: item.terminationMessagePath,
				terminationMessagePolicy: item.terminationMessagePolicy,
				tty: item.tty,
				volumeDevices: item.volumeDevices,
				volumeMounts: item.volumeMounts,
				workingDir: item.workingDir,
			};
			containers.push(container);
		});
	} else {
		containers.push(data.container);
	}
	if (props.initContainers && props.initContainers.length > 0) {
		props.initContainers.forEach((item: Container) => {
			const container: ContainerType = {
				args: item.args,
				command: item.command,
				env: item.env,
				envFrom: item.envFrom,
				image: item.image,
				imagePullPolicy: item.imagePullPolicy,
				isIntiContainer: true,
				lifecycle: item.lifecycle,
				livenessProbe: item.livenessProbe,
				name: item.name,
				ports: item.ports,
				readinessProbe: item.readinessProbe,
				resources: item.resources,
				securityContext: item.securityContext,
				startupProbe: item.startupProbe,
				stdin: item.stdin,
				stdinOnce: item.stdinOnce,
				terminationMessagePath: item.terminationMessagePath,
				terminationMessagePolicy: item.terminationMessagePolicy,
				tty: item.tty,
				volumeDevices: item.volumeDevices,
				volumeMounts: item.volumeMounts,
				workingDir: item.workingDir,
			};
			containers.push(container);
		});
		data.containers = containers;
	}
	data.containers = containers;
});

const returnContainers = () => {
	getContainers();
	const containers = [] as Container[];
	const initContainers = [] as Container[];
	const res = deepClone(data.containers) as ContainerType[];
	res.forEach((item) => {
		if (item.isIntiContainer) {
			delete item.isIntiContainer;
			delete item.lifecycle;
			delete item.livenessProbe;
			delete item.readinessProbe;
			delete item.startupProbe;
			initContainers.push(item);
		} else {
			delete item.isIntiContainer;
			containers.push(item);
		}
	});
	return { containers: containers, initContainers: initContainers, volumes: data.volumes };
};

defineExpose({
	returnContainers,
});
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
</style>
