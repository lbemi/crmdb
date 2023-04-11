<template>
	<div class="layout-pd">
		<div>
			<!--			<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">-->
			<!--				<el-tab-pane v-for="(item,index) in editableTabs" :key="index" :label="item.title" :name="item.name" :closable='item.closeAble'>-->
			<!--					<container :ref="setItemRef" :container="data.containers[index]" @updateContainer='getContainer' />-->
			<!--				</el-tab-pane>-->
			<!--			</el-tabs>-->
			<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">
				<el-tab-pane v-for="(item, index) in data.containers" :key="index" :label="'容器' + (index+1)" :name="index+1" :closable="index == 0">
					<container :container="item" @updateContainer="getContainer" />
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, watch } from 'vue';
import { ref } from 'vue-demi';
import {V1Container, V1ContainerPort, V1EnvVar, V1SecurityContext} from '@kubernetes/client-node';
import type { TabPaneName } from 'element-plus';

const Container = defineAsyncComponent(() => import('./container.vue'));

let tabIndex = 0;
const editableTabsValue = ref('1');
const itemRefs = ref([]);
//动态设置ref
const setItemRef = (el) => {
	if (el) {
		itemRefs.value.push(el);
	}
};
const editableTabs = ref([
	{
		title: '容器' + tabIndex,
		name: '1',
		closeAble: false,
	},
]);

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
	if (action === 'add') {
		const newTabName = `${++tabIndex}`;
		// editableTabs.value.push({
		// 	title: '容器' + newTabName,
		// 	name: newTabName,
		// 	closeAble: true,
		// });
		data.containers.push(data.container)
		editableTabsValue.value = newTabName;
	} else if (action === 'remove') {
		const tabs = editableTabs;
		let activeName = editableTabsValue.value;
		if (activeName === targetName) {
			tabs.value.forEach((tab, index) => {
				if (tab.name === targetName) {
					const nextTab = tabs[index + 1] || tabs[index - 1];
					itemRefs.value.splice(index, 1);
					if (nextTab) {
						activeName = nextTab.name;
					}
				}
			});
		}

		editableTabsValue.value = activeName;
		editableTabs.value = tabs.value.filter((tab) => tab.name !== targetName);
	}
};

const data = reactive({
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
});

const getContainer = (container: V1Container) => {
	// console.log("********",itemRefs.value)
	// data.containers=[]
	// itemRefs.value.forEach((res,index) =>{
	// 	data.containers[index] = res.getContainer()
	// })
	// return data.containers
	const index = parseInt(editableTabsValue.value);
	data.containers[index] = container;
};

const props = defineProps({
	containers: Array<V1Container>,
});

watch(
	props.containers,
	() => {
		console.log("传递过来的containers：",props.containers)
		if (props.containers) {
			data.containers = props.containers;
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
const emit = defineEmits(['updateContainers'])

watch(()=>data.containers ,()=> {
	emit('updateContainers',data.containers)
});
// defineExpose({
// 	getContainers,
// });
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
