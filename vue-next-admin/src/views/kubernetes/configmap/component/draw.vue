<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="48%">
		<template #header="{ close, titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ props.title }}</h4>
		</template>
		<div>
			<el-form :inline="true" class="demo-form-inline" v-model="data.configMap">
				<el-form-item label="命名空间：">
					<el-select v-model="data.configMap.metadata.namespace" style="max-width: 180px" size="small" class="m-2" placeholder="Select"
						><el-option key="all" label="所有命名空间" value="all"></el-option>
						<el-option
							v-for="item in k8sStore.state.namespace"
							:key="item.metadata?.name"
							:label="item.metadata.name"
							:value="item.metadata!.name!"
						/>
					</el-select>
				</el-form-item>
				<div>
					<el-form-item label="配置项名称:"><el-input size="small" v-model="data.configMap.metadata.name"></el-input> </el-form-item>
				</div>
				<el-form-item label="数据:">
					<div>
						<el-table :data="data.keyValues" style="width: 100%">
							<el-table-column label="名称" prop="name" width="180">
								<template #default="scope">
									<el-input v-model="scope.row.name" size="small" />
								</template>
							</el-table-column>
							<el-table-column label="值" prop="value" width="380">
								<template #default="scope">
									<el-input type="textarea" v-model="scope.row.value" size="small" />
								</template>
							</el-table-column>
							<el-table-column>
								<template #default="scope">
									<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.keyValues.splice(scope.$index, 1)"></el-button>
								</template>
							</el-table-column>
						</el-table>
					</div>
				</el-form-item>
				<div>
					<el-button size="small" @click="addKey()" style="width: 90%" type="primary" plain
						><el-icon><Plus /></el-icon>添加</el-button
					>
				</div>
			</el-form>
		</div>

		<template #footer>
			<div style="flex: auto">
				<el-button size="small">cancel</el-button>
				<el-button type="primary" size="small" @click="confirm">confirm</el-button>
			</div>
		</template>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';

import { ConfigMap } from 'kubernetes-types/core/v1';
import { computed, onMounted, reactive } from 'vue';
import { watch } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { RemoveFilled, Plus } from '@element-plus/icons-vue';

const k8sStore = kubernetesInfo();
const data = reactive({
	visible: false,
	configMap: {
		metadata: {
			name: '',
			namespace: '',
		},
		data: {},
	} as ConfigMap,
	keyValues: [] as Array<{ name: string; value: string }>,
});

const addKey = () => {
	data.keyValues.push({
		name: '',
		value: '',
	});
};
const emit = defineEmits(['update:visible']);

const props = defineProps({
	visible: Boolean,
	configMap: {
		type: Object as () => ConfigMap,
	},
	title: String,
});

const convertConfigMap = () => {
	data.keyValues.forEach((item) => {
		let obj = {};
		obj[item.name] = item.value;
		Object.assign(data.configMap.data, obj);
	});
};
const confirm = () => {
	convertConfigMap();
	console.log(data.configMap);
};
const handleClose = () => {
	emit('update:visible', false);
};

// onMounted(() => {
// 	data.visible = props.visible;
// 	if (props.configMap) {
// 		data.configMap = props.configMap;
// 	}
// });
watch(
	() => props,
	() => {
		console.log('----', props);
		data.visible = props.visible;
		if (props.configMap) {
			data.configMap = props.configMap;
		}
	},
	{
		deep: true,
		immediate: true,
	}
);
</script>
<style scoped>
.el-form {
	margin-left: 20px;
	margin-top: 10px;
}
</style>
