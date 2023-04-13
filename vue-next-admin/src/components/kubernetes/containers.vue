
<template>
	<div class="layout-pd">
		<div>
			<!--			<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">-->
			<!--				<el-tab-pane v-for="(item,index) in editableTabs" :key="index" :label="item.title" :name="item.name" :closable='item.closeAble'>-->
			<!--					<container :ref="setItemRef" :container="data.containers[index]" @updateContainer='getContainer' />-->
			<!--				</el-tab-pane>-->
			<!--			</el-tabs>-->
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

const Container = defineAsyncComponent(() => import('./container.vue'));

const editableTabsValue = ref(0);
// const itemRefs = ref([]);
// // //动态设置ref
// const setItemRef = (el) => {
// 	if (el) {
// 		itemRefs.value.push(el);
// 	}
// };
// const editableTabs = ref([
// 	{
// 		title: '容器' + tabIndex,
// 		name: '1',
// 		closeAble: false,
// 	},
// ]);

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
	if (action === 'add') {
		const newTabName = data.containers.length;
		// editableTabs.value.push({
		// 	title: '容器' + newTabName,
		// 	name: newTabName,
		// 	closeAble: true,
		// });
		data.containers.push(data.container);
		editableTabsValue.value = newTabName ;
		console.log('当前活动页面：', editableTabsValue.value);
	} else if (action === 'remove') {
		console.log('我要删除第：', targetName);
		if (targetName === 0) {
			return;
		}
		const tabs = data.containers;
		let activeName = editableTabsValue.value;
		if (activeName === targetName) {
			// tabs.forEach((tab, index) => {
			// 	if ((index + 1) === targetName) {
			//
			// 		// const nextTab = tabs[index + 1] || tabs[index - 1];
			// 		let nextTab = 0;
			// 		if (tabs[index + 2]) {
			// 			nextTab = index + 1 ;
			// 		} else if (tabs[index - 1]) {
			// 			nextTab = index - 1 ;
			// 		}
			// 		// itemRefs.value.splice(index, 1);
			// 		if (nextTab) {
			// 			activeName = nextTab;
			// 		}
			// 	}

			// });
			activeName = activeName - 1;
		}
		tabs.forEach((tab, index) => {
			console.log('我要删除：', index, targetName, activeName);
			if (index === targetName) {
				console.log('我在删除：', index, targetName);
				tabs.splice(index, 1);
			}
		});
		data.containers = tabs;
		editableTabsValue.value = activeName;
		console.log(editableTabsValue.value);
		// data.containers = tabs.filter((tab,index) => (index+1) !== targetName);
		// data.containers = tabs.splice(activeName,1);
	}
};

const data = reactive({
	containers: [
		{
      name: '',
      imagePullPolicy: 'ifNotPresent',
			securityContext: {
				privileged: false,
			} as V1SecurityContext,
			livenessProbe: {},
			readinessProbe: {},
			startupProbe: {},
			env: [] as V1EnvVar[],
			ports: [] as V1ContainerPort,
      resources: {}
		} as V1Container,
	] as V1Container[],
	container: {
    name: '',
    imagePullPolicy: 'ifNotPresent',
		securityContext: {
			privileged: false,
		} as V1SecurityContext,
		livenessProbe: {},
		readinessProbe: {},
		startupProbe: {},
		env: [] as V1EnvVar[],
		ports: [] as V1ContainerPort,
    resources: {}
	} as V1Container,
});

const getContainer = (index:number,container: V1Container) => {
	// console.log("********",itemRefs.value)
	// data.containers=[]
	// itemRefs.value.forEach((res,index) =>{
	// 	data.containers[index] = res.getContainer()
	// })
	// return data.containers
	// const index = parseInt(editableTabsValue.value);
	console.log('接受到容器container的数据：', container,"当前活动页面：",editableTabsValue.value,index);
	if(index === editableTabsValue.value) {
		data.containers[index] = container;
		console.log('接受到容器的数据：', data.containers);
	}

};

const props = defineProps({
	containers: Array<V1Container>,
});

watch(
		()=>props.containers,
	() => {
		console.log('传递过来的containers：', props.containers);
		if (props.containers && props.containers.length != 0) {
			console.log('传递过来的containers：', props.containers);
			data.containers = props.containers;
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
		console.log('出发更新depliyment Containers，', data.containers);
		emit('updateContainers', data.containers);
	},
	{
		immediate: true,
		deep: true,
	}
);
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
