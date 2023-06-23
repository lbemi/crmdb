<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="45%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div class="box_body">
			<el-descriptions :title="data.configMap.metadata.name" :column="2" border>
				<el-descriptions-item label="命名空间" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{
					data.configMap.metadata.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center" :span="2">{{
					dateStrFormat(data.configMap.metadata.creationTimestamp)
				}}</el-descriptions-item>
				<el-descriptions-item label="标签" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in data.configMap.metadata.labels" :key="index">
						<el-tag class="label" type="success" size="small"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in data.configMap.metadata.annotations" :key="index">
						<span> {{ key }}:{{ item }} </span>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="数据" label-align="right" align="center" :span="2">
					<el-table :data="data.keyValues">
						<el-table-column prop="key" label="名称" />
						<el-table-column prop="value" label="值" />
					</el-table>
				</el-descriptions-item>
			</el-descriptions>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';
import { ConfigMap } from 'kubernetes-types/core/v1';
import { onMounted, reactive } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';

const data = reactive({
	visible: false,
	configMap: {
		metadata: {
			name: '',
			namespace: '',
		},
		data: {},
	} as ConfigMap,
	keyValues: [] as Array<{ key: string; value: string }>,
});

const props = defineProps({
	visible: Boolean,
	configMap: {
		type: Object as () => ConfigMap,
	},
	title: String,
});

const convertConfigMapTo = () => {
	let kvs = [] as Array<{ key: string; value: string }>;
	Object.keys(data.configMap.data).forEach((k) => {
		kvs.push({
			key: k,
			value: data.configMap.data![k],
		});
	});
	data.keyValues = kvs;
};

onMounted(() => {
	data.visible = props.visible;
	if (!isObjectValueEqual(props.configMap, {})) {
		data.configMap = props.configMap;
		convertConfigMapTo();
	}
});

const emit = defineEmits(['update:visible']);
const handleClose = () => {
	emit('update:visible', false);
};
</script>
<style scoped>
.box_body {
	margin-left: 20px;
	margin-top: 10px;
}
.footer {
	display: flex;
	margin-top: 50px;
	/*margin-left: 80px;*/
	justify-content: center;
}
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
