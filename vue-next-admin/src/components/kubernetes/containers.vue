<template>
	<div class="layout-pd">
		<div>
			<el-tabs v-model="editableTabsValue" @tab-remove="tabRemove" class="demo-tabs" :before-leave="tabChange">
				<el-tab-pane v-for="(item, index) in data.containers" :key="index" :name="index" :closable="data.containers.length != 1">
					<template #label>
						<span v-if="item.isInitContainer" class="custom-tabs-label">
							<SvgIcon name="iconfont icon-container-" class="svg" />{{ ' init容器 ' }}
						</span>
						<span v-else class="custom-tabs-label"> <SvgIcon name="iconfont icon-container-" class="svg" />{{ ' 容器 ' }} </span>
					</template>
					<ContainerDiv :ref="(el: refItem) => setItemRef(el, index)" :container="item" :index="index" :volumes="props.volumes" />
				</el-tab-pane>
				<el-tab-pane key="CustomBtn" name="CustomBtn" :closable="false">
					<template #label>
						<el-link type="primary" :underline="false" :icon="Plus">添加容器</el-link>
					</template>
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import { ComponentPublicInstance, ref } from 'vue-demi';
import { Container, Volume } from 'kubernetes-models/v1';
import type { FormInstance, TabPaneName } from 'element-plus';
import { deepClone } from '@/utils/other';
import { Plus } from '@element-plus/icons-vue';
import { CustomizeContainer } from '@/types/kubernetes/container';

const ContainerDiv = defineAsyncComponent(() => import('./container.vue'));

type refItem = Element | ComponentPublicInstance | null;
const editableTabsValue = ref(0);
const itemRefs = ref<(Element | ComponentPublicInstance | null)[]>([]);

const setItemRef = (el: refItem, index: number) => {
	if (el) {
		itemRefs.value[index] = el;
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
	data.containers = tabs.filter((_tab: any, index: number) => index !== TabPaneName);
	editableTabsValue.value = data.containers.length - 1;
};

const data = reactive({
	validateRefs: <Array<FormInstance>>[],
	loadFromParent: false,
	currentIndex: 1,
	addIndex: 1,
	tabIndex: 1,
	editableTabsValue: '1',
	volumes: [] as Volume[],
	containers: <Array<CustomizeContainer>>[],
	container: new CustomizeContainer({
		isInitContainer: false,
		name: '',
		image: '',
		imagePullPolicy: 'IfNotPresent',
		securityContext: {
			privileged: false,
		},
	}),
});

const getContainers = () => {
	data.validateRefs = [];
	const vs = [];
	const containers = data.containers;
	for (let i = 0, len = itemRefs.value.length; i < len; i++) {
		const refValue = itemRefs.value[i];
		if (!refValue) continue;
		const { index, container, volumes, validateRefs } = refValue.returnContainer();
		data.validateRefs.push(...validateRefs);
		vs.push(...volumes);
		containers[index] = { ...container };
	}

	data.volumes = vs;
};

type propsType = {
	containers: Array<Container>;
	initContainers: Array<Container>;
	volumes: Array<Volume>|undefined;
};
const props = defineProps<propsType>();
onMounted(() => {
	const containers =
		props.containers.length > 0
			? props.containers.map(
					(item) =>
						new CustomizeContainer({
							...item,
							isInitContainer: false,
						})
			  )
			: [data.container];

	if (props.initContainers.length > 0) {
		const initContainers = props.initContainers.map(
			(item) =>
				new CustomizeContainer({
					...item,
					isInitContainer: true,
				})
		);
		data.containers = [...containers, ...initContainers];
	} else {
		data.containers = containers;
	}
});

const returnContainers = () => {
	getContainers();
	const containers = [] as Container[];
	const initContainers = [] as Container[];

	let resIndex = 0;
	const resLength = data.containers.length;
	while (resIndex < resLength) {
		const item = data.containers[resIndex];
		if (item.isInitContainer) {
			initContainers.push({
				...item,
				isInitContainer: undefined,
				lifecycle: undefined,
				livenessProbe: undefined,
				readinessProbe: undefined,
				startupProbe: undefined,
			} as Omit<CustomizeContainer, 'isInitContainer' | 'lifecycle' | 'livenessProbe' | 'readinessProbe' | 'startupProbe'>);
		} else {
			containers.push(item);
		}
		resIndex++;
	}

	return { containers, initContainers, volumes: data.volumes, validateRefs: data.validateRefs };
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
</style>
