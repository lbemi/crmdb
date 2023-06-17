/** * Created by lei on 2022/12/15 */
<template>
	<div>
		<el-dialog v-model="visible" :title="title" width="30%" :before-close="handleClose">
			<Labels :labelData="labels" @on-click="getLabels" v-if="visible" />
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="handleClose()">取消</el-button>
					<el-button type="primary" @click="handleConfirm()"> 更新 </el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import Labels from '@/components/label/index.vue';

import { ElMessage } from 'element-plus';
import { ref, watch, reactive, toRefs } from 'vue';
import { useNodeApi } from '@/api/kubernetes/node';
import { Node } from '@/types/kubernetes/cluster';

interface label {
	key: string;
	value: string;
}
const props = defineProps<{
	visible: boolean;
	title: string;
	data: Node;
	cloud: string;
}>();
const { visible } = toRefs(props);
const patchData = reactive({
	name: '',
	labels: {},
});
// const nodeData = ref({} as Node)

const labels = ref<label[]>([]);
const nodeApi = useNodeApi();
const emits = defineEmits(['update:visible', 'valueChange']);
const handleClose = () => {
	emits('update:visible', false);
};

const handleConfirm = async () => {
	// console.log('********', nodeData.value.metadata.labels, props.cloud)
	// await nodeApi.patch.request({ cloud: props.cloud }, patchData).then((res) => {
	//   ElMessage.success(res.data.message)
	// })
	await nodeApi
		.updateLabel({ cloud: props.cloud }, patchData)
		.then((res) => {
			ElMessage.success(res.message);
		})
		.catch((res) => {
			ElMessage.error(res.message);
		});
	emits('valueChange');
	handleClose();
};

const getLabels = (labels: { [index: string]: string }) => {
	patchData.labels = labels;
};

watch(
	() => props.data,
	() => {
		// nodeData.value = props.data
		labels.value = [];
		patchData.name = props.data.metadata?.name!;
		if (props.data) {
			for (let key in props.data.metadata?.labels) {
				const l: label = {
					key: key,
					value: props.data.metadata?.labels![key],
				};
				labels.value.push(l);
			}
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>
<style scoped lang="less"></style>
