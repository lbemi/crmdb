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
import { V1Container, V1ContainerPort, V1EnvVar, V1SecurityContext } from '@kubernetes/client-node';
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
	containers: [
		{
			name: '',
			imagePullPolicy: 'IfNotPresent',
			securityContext: {
				// privileged: false,
			} as V1SecurityContext,
			livenessProbe: {},
			readinessProbe: {},
			startupProbe: {},
			env: [] as V1EnvVar[],
			ports: [] as V1ContainerPort,
			resources: {},
		} as V1Container,
	] as V1Container[],
	container: {
		name: '',
		imagePullPolicy: 'IfNotPresent',
		securityContext: {
			// privileged: false,
		} as V1SecurityContext,
		livenessProbe: {},
		readinessProbe: {},
		startupProbe: {},
		env: [] as V1EnvVar[],
		ports: [] as V1ContainerPort,
		resources: {},
	} as V1Container,
});

const getContainer = (index: number, container: V1Container) => {
	console.log('&&&&&&&&&&&&&&&&&container&&&&--------', container, data.containers[index], isObjectValueEqual(data.containers[index], container));
	if (index === editableTabsValue.value) {
		// // FIXME  初始化container name
		// if (!isObjectValueEqual(data.containers[index], container)) {
		console.log('2.&&&&&&&&&&&&&&&&&container&&&&&&&&&', container);
		data.containers[index] = container;
		// }
	}
};

const props = defineProps({
	containers: Array<V1Container>,
});

watch(
	() => props.containers,
	() => {
		if (props.containers && props.containers.length != 0 && !isObjectValueEqual(data.containers, props.containers)) {
			console.log('a.ZZZZZZZZZZZZZZZZZJHGDKJAGSKDBKJASD', props.containers);
			data.containers = deepClone(props.containers) as V1Container[];
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
const emit = defineEmits(['updateContainers']);

watch(
	() => [...data.containers],
	() => {
		console.log('3.containers jiesho;;;;;;', data.containers);
		emit('updateContainers', data.containers);
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
