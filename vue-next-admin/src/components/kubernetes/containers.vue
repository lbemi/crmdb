<template>
	<div class="layout-pd">
		<div>
			<el-tabs v-model="editableTabsValue" type="card" editable class="demo-tabs" @edit="handleTabsEdit">
				<el-tab-pane v-for="(item,index) in editableTabs" :key="index" :label="item.title" :name="item.name" :closable='item.closeAble'>
					<container :ref="setItemRef"></container>
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup lang="ts">
import {defineAsyncComponent, reactive} from 'vue';
import {ref} from 'vue-demi';
import {V1Container} from '@kubernetes/client-node';
import type {TabPaneName} from 'element-plus';

const Container = defineAsyncComponent(() => import("./container.vue"))

let tabIndex = 1;
const editableTabsValue = ref('1');
const itemRefs = ref([])
//动态设置ref
const setItemRef =(el) =>{
	if(el) {
		itemRefs.value.push(el)
	}
}
const editableTabs = ref([
	{
		title: '容器' + tabIndex,
		name: '1',
		closeAble: false,
	}
]);

const handleTabsEdit = (targetName: TabPaneName | undefined, action: 'remove' | 'add') => {
	if (action === 'add') {
		const newTabName = `${++tabIndex}`;
		editableTabs.value.push({
			title: '容器' + newTabName,
			name: newTabName,
			closeAble: true
		});
		editableTabsValue.value = newTabName;
	} else if (action === 'remove') {
		const tabs = editableTabs;
		let activeName = editableTabsValue.value;
		if (activeName === targetName) {
			tabs.value.forEach((tab, index) => {
				if (tab.name === targetName) {
					const nextTab = tabs[index + 1] || tabs[index - 1];
					itemRefs.value.splice(index,1)
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
});

const getContainers = ()=>{
	console.log("********",itemRefs.value)
	data.containers=[]
	itemRefs.value.forEach((res,index) =>{
		data.containers[index] = res.getContainer()
	})
	return data.containers
}

defineExpose({
	getContainers
})
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
