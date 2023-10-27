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
					<el-table class="custom-table" v-if="data.configMap.data" :data="data.keyValues">
						<el-table-column prop="key" label="名称" />
						<el-table-column prop="value" label="值">
							<template #default="scope">
								<div class="custom-cell" :class="{ 'max-height': scope.row.value.length > 200 }">
									{{ scope.row.value }}
								</div>
							</template>
						</el-table-column>
					</el-table>
					<el-table class="custom-table" v-if="data.configMap.binaryData" :data="data.binaryKeyValues">
						<el-table-column prop="key" label="名称" />
						<el-table-column prop="value" label="值">
							<template #default="scope">
								<div class="custom-cell" :class="{ 'max-height': scope.row.value.length > 200 }">
									{{ scope.row.value }}
								</div>
							</template>
						</el-table-column>
					</el-table>
				</el-descriptions-item>
			</el-descriptions>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';
import { ConfigMap } from '@/types/kubernetes-types/core/v1';
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
		binaryData: {},
	} as ConfigMap,
	keyValues: [] as Array<{ key: string; value: string }>,
	binaryKeyValues: [] as Array<{ key: string; value: string }>,
});

const props = defineProps({
	visible: Boolean,
	configMap: {
		type: Object as () => ConfigMap,
	},
	title: String,
});

const convertConfigMapTo = (obj: { [name: string]: string }) => {
	let kvs = [] as Array<{ key: string; value: string }>;
	Object.keys(obj).forEach((k) => {
		kvs.push({
			key: k,
			value: obj[k],
		});
	});
	return kvs;
};

onMounted(() => {
	data.visible = props.visible;
	if (!isObjectValueEqual(props.configMap, {})) {
		data.configMap = props.configMap;
		if (data.configMap.data) {
			data.keyValues = convertConfigMapTo(data.configMap.data);
		}
		if (data.configMap.binaryData) {
			data.binaryKeyValues = convertConfigMapTo(data.configMap.binaryData);
		}
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
.custom-cell {
	overflow: auto;
}

.custom-cell.max-height {
	max-height: 200px; /* 设置最大高度为 200px，可以根据需要进行调整 */
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
