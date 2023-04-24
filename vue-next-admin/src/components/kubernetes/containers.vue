<template>
	<div class="layout-pd">
		<div>
			<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">
				<el-tab-pane v-for="(item, index) in data.containers" :key="index" :label="'容器' + (index + 1)" :name="index" :closable="index != 0">
					<container :container="item" :index="index" @updateContainer="getContainer" />
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, watch } from 'vue';
import { ref } from 'vue-demi';
import { V1Container } from '@kubernetes/client-node';
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
	loadFromParent: false,
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
	] as Array<V1Container>,
	container: <V1Container>{
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

const getContainer = (index: number, container: V1Container) => {
	if (index === editableTabsValue.value) {
		// // FIXME  初始化container name
		data.containers[index] = deepClone(container) as V1Container;
	}
};

const props = defineProps({
	containers: Array<V1Container>,
});

watch(
	() => props.containers,
	() => {
		if (props.containers && props.containers.length != 0 && !isObjectValueEqual(data.containers, props.containers)) {
			data.loadFromParent = true;
			data.containers = deepClone(props.containers) as V1Container[];
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
	() => data.containers,
	() => {
		if (!data.loadFromParent) {
			emit('updateContainers', deepClone(data.containers));
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
