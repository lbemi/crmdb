<template>
	<el-drawer v-model="state.visible" @close="handleClose" size="44%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div class="box_body">
			<el-descriptions :title="state.persistentVolume.metadata!.name" :column="2" border>
				<el-descriptions-item label="创建时间" label-align="right" align="center" :span="1">{{
					dateStrFormat(state.persistentVolume.metadata!.creationTimestamp!)
				}}</el-descriptions-item>
				<el-descriptions-item label="状态" label-align="right" align="center" :span="1">
					{{ state.persistentVolume.status?.phase }}
				</el-descriptions-item>
				<el-descriptions-item label="标签" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.persistentVolume.metadata!.labels" :key="index">
						<el-tag class="label" type="success" size="small"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.persistentVolume.metadata!.annotations" :key="index">
						<span> {{ key }}:{{ item }} </span>
					</div>
				</el-descriptions-item>

				<el-descriptions-item label="访问模式" label-align="right" align="center" :span="1">
					<div v-for="(item, index) in state.persistentVolume.spec?.accessModes" :key="index">
						<span> {{ item }} </span>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="容量" label-align="right" align="center" :span="1">
					{{ state.persistentVolume.spec?.capacity?.storage }}
				</el-descriptions-item>
				<el-descriptions-item label="回收策略" label-align="right" align="center" :span="1">
					{{ state.persistentVolume.spec?.persistentVolumeReclaimPolicy }}
				</el-descriptions-item>
				<el-descriptions-item label="绑定的存储声明(PVC)" label-align="right" align="center" :span="1">
					{{ state.persistentVolume.spec?.claimRef?.namespace }}/{{ state.persistentVolume.spec?.claimRef?.name }}
				</el-descriptions-item>
				<el-descriptions-item label="存储类(StorageClass)" label-align="right" align="center" :span="1">
					{{ state.persistentVolume.spec?.storageClassName || '无' }}
				</el-descriptions-item>
			</el-descriptions>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';
import { PersistentVolume } from 'kubernetes-models/v1';
import { onMounted, reactive } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { dateStrFormat } from '@/utils/formatTime';

const state = reactive({
	visible: false,
	persistentVolume: {
		metadata: {
			name: '',
			namespace: '',
		},
	} as PersistentVolume,
	keyValues: [] as Array<{ key: string; value: string }>,
});

const props = defineProps({
	visible: Boolean,
	persistentVolume: {
		type: Object as () => PersistentVolume,
	},
	title: String,
});

onMounted(() => {
	state.visible = props.visible;
	if (props.persistentVolume && !isObjectValueEqual(props.persistentVolume, {})) {
		state.persistentVolume = props.persistentVolume;
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
