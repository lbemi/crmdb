<template>
	<div class="layout-pd">
		<div>
			<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">
				<el-tab-pane v-for="(item, index) in data.containers" :key="index" :name="index" :closable="index != 0">
					<template #label>
						<span v-if="data.isInitContainer" class="custom-tabs-label">
							<SvgIcon name="iconfont icon-container-" class="svg" />{{ ' init容器 ' + (index + 1) }}
						</span>
						<span v-else class="custom-tabs-label"> <SvgIcon name="iconfont icon-container-" class="svg" />{{ ' 容器 ' + (index + 1) }} </span>
					</template>
					<container :container="item" :index="index" @updateContainer="getContainer" />
				</el-tab-pane>
				<el-tab-pane
					v-if="data.initContainers.length != 0"
					v-for="(item, index) in data.initContainers"
					:key="index"
					:name="index"
					:closable="index != 0"
				>
					<template #label>
						<span class="custom-tabs-label"> <SvgIcon name="iconfont icon-container-" class="svg" />{{ ' init容器 ' + (index + 1) }} </span>
					</template>
					<container :container="item" :index="index" @updateContainer="getContainer" />
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, watch } from 'vue';
import { ref } from 'vue-demi';
import { Container } from 'kubernetes-types/core/v1';
import type { TabPaneName } from 'element-plus';
import { isObjectValueEqual } from '/@/utils/arrayOperation';
import { deepClone } from '/@/utils/other';

const Container = defineAsyncComponent(() => import('./container.vue'));

const editableTabsValue = ref(0);

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
	if (action === 'add') {
		const newTabName = data.containers.length;

		data.containers.push(data.container);
		editableTabsValue.value = newTabName;
	} else if (action === 'remove') {
		if (targetName === 0) {
			return;
		}
		const tabs = data.containers;
		let activeName = editableTabsValue.value;
		if (activeName === targetName) {
			activeName = activeName - 1;
		}
		tabs.forEach((tab, index) => {
			if (index === targetName) {
				tabs.splice(index, 1);
			}
		});
		data.containers = tabs;
		editableTabsValue.value = activeName;
	}
};

const data = reactive({
	isInitContainer: false,
	loadFromParent: false,
	initContainers: [] as Array<Container>,
	containers: [
		{
			name: '',
			image: '',
			imagePullPolicy: 'IfNotPresent',
			securityContext: {
				privileged: false,
			},
			ports: [],
			env: [],
			resources: {},
			livenessProbe: {},
			readinessProbe: {},
			startupProbe: {},
			lifecycle: {},
		},
	] as Array<Container>,
	container: <Container>{
		name: '',
		image: '',
		imagePullPolicy: 'IfNotPresent',
		securityContext: {
			privileged: false,
		},
		ports: [],
		env: [],
		resources: {},
		livenessProbe: {},
		readinessProbe: {},
		startupProbe: {},
		lifecycle: {},
	},
});

const getContainer = (index: number, isInitContainer: boolean, container: Container) => {
	if (index === editableTabsValue.value) {
		// // FIXME  初始化container name
		if (isInitContainer) {
			data.isInitContainer = isInitContainer;
			data.initContainers[index] = deepClone(container) as Container;
			delete data.containers[index];
		} else {
			delete data.initContainers[index];
			data.containers[index] = deepClone(container) as Container;
		}
	}
};

const getInitContainer = (index: number, isInitContainer: boolean, container: Container) => {
	if (index === editableTabsValue.value) {
		if (isInitContainer) {
			data.isInitContainer = isInitContainer;
			data.initContainers[index] = deepClone(container) as Container;
		} else {
			delete data.initContainers[index];
			data.containers[index] = deepClone(container) as Container;
		}
	}
};

const props = defineProps({
	containers: Array<Container>,
	initContainers: Array<Container>,
});

watch(
	() => props,
	() => {
		if (props.containers && props.containers.length != 0 && !isObjectValueEqual(data.containers, props.containers)) {
			data.loadFromParent = true;
			data.containers = deepClone(props.containers) as Container[];
			setTimeout(() => {
				// 延迟一下，不然会触发循环更新
				data.loadFromParent = false;
			}, 10);
		}
		if (props.initContainers && props.initContainers.length != 0 && !isObjectValueEqual(data.initContainers, props.initContainers)) {
			data.loadFromParent = true;
			data.initContainers = deepClone(props.initContainers) as Container[];
			setTimeout(() => {
				// 延迟一下，不然会触发循环更新
				data.loadFromParent = false;
			}, 10);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
const emit = defineEmits(['updateContainers']);

watch(
	() => [data.containers, data.initContainers],
	() => {
		if (!data.loadFromParent) {
			emit('updateContainers', deepClone(data.containers), deepClone(data.initContainers));
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
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
