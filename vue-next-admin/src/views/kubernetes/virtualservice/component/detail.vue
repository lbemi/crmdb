<template>
	<el-drawer v-model="data.visible" @close="handleClose" size="45%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div class="box_body">
			<el-descriptions :title="data.virtualService.metadata?.name" :column="2" border>
				<el-descriptions-item label="命名空间" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{
					data.virtualService.metadata?.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center" :span="2">{{
					dateStrFormat(data.virtualService.metadata?.creationTimestamp!)
				}}</el-descriptions-item>
				<el-descriptions-item label="标签" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in data.virtualService.metadata?.labels" :key="index">
						<el-tag class="label" type="success" size="small"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in data.virtualService.metadata?.annotations" :key="index">
						<span> {{ key }}:{{ item }} </span>
		w			</div>
				</el-descriptions-item>
			</el-descriptions>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';
import { VirtualService } from '@kubernetes-models/istio/networking.istio.io/v1beta1/VirtualService';
import { onMounted, reactive } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { dateStrFormat } from '@/utils/formatTime';

const data = reactive({
	visible: false,
	virtualService: {} as VirtualService,
	keyValues: [] as Array<{ key: string; value: string }>,
	binaryKeyValues: [] as Array<{ key: string; value: string }>,
});

const props = defineProps({
	visible: Boolean,
	virtualService: {
		type: Object as () => VirtualService,
	},
	title: String,
});

onMounted(() => {
	data.visible = props.visible;
	if (props.virtualService && !isObjectValueEqual(props.virtualService, data.virtualService)) {
		data.virtualService = props.virtualService;
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
