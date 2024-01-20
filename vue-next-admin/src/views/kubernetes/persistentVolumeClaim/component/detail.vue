<template>
	<el-drawer v-model="state.visible" @close="handleClose" size="44%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div class="box_body">
			<el-descriptions :title="state.persistentVolumeClaim.metadata!.name" :column="2" border>
				<el-descriptions-item label="创建时间" label-align="right" align="center" :span="1">{{
					dateStrFormat(state.persistentVolumeClaim.metadata!.creationTimestamp!)
				}}</el-descriptions-item>
				<el-descriptions-item label="状态" label-align="right" align="center" :span="1">
					{{ state.persistentVolumeClaim.status?.phase }}
				</el-descriptions-item>
				<el-descriptions-item label="标签" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.persistentVolumeClaim.metadata!.labels" :key="index">
						<el-tag class="label" type="info" size="default"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.persistentVolumeClaim.metadata!.annotations" :key="index">
						<span> {{ key }}:{{ item }} </span>
					</div>
				</el-descriptions-item>

				<el-descriptions-item label="访问模式" label-align="right" align="center" :span="1">
					<div v-for="(item, index) in state.persistentVolumeClaim.spec?.accessModes" :key="index">
						<span> {{ item }} </span>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="容量" label-align="right" align="center" :span="1">
					{{ state.persistentVolumeClaim.status?.capacity?.storage }}
				</el-descriptions-item>

				<el-descriptions-item label="绑定的存储类(PV)" label-align="right" align="center" :span="1">
					{{ state.persistentVolumeClaim.spec?.volumeName }}
				</el-descriptions-item>
				<el-descriptions-item label="存储类(StorageClass)" label-align="right" align="center" :span="1">
					{{ state.persistentVolumeClaim.spec?.storageClassName || '无' }}
				</el-descriptions-item>
			</el-descriptions>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';
import { PersistentVolumeClaim } from 'kubernetes-types/core/v1';
import { onMounted, reactive } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { dateStrFormat } from '@/utils/formatTime';

const state = reactive({
	visible: false,
	persistentVolumeClaim: {
		metadata: {
			name: '',
			namespace: '',
		},
	} as PersistentVolumeClaim,
	keyValues: [] as Array<{ key: string; value: string }>,
});

const props = defineProps({
	visible: Boolean,
	persistentVolumeClaim: {
		type: Object as () => PersistentVolumeClaim,
	},
	title: String,
});

onMounted(() => {
	state.visible = props.visible;
	if (props.persistentVolumeClaim && !isObjectValueEqual(props.persistentVolumeClaim, {})) {
		state.persistentVolumeClaim = props.persistentVolumeClaim;
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

.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
