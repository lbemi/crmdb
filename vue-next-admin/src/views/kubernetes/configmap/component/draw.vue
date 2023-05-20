<template>
	<el-drawer v-model="data.visible" :show-close="false" @close="handleClose">
		<!-- <template #header="{ close, titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ props.title }}</h4>
			<el-button type="danger" @click="handleClose">
				<el-icon class="el-icon--left"><CircleCloseFilled /></el-icon>
				Close
			</el-button>
		</template> -->
		asdas
		<!-- <el-descriptions class="margin-top" title="With border" :column="3" border>
			<template #extra>
				<el-button type="primary">Operation</el-button>
			</template>
			<el-descriptions-item>
				<template #label>
					<div class="cell-item">
						<el-icon :style="iconStyle">
							<user />
						</el-icon>
						Username
					</div>
				</template>
				kooriookami
			</el-descriptions-item>
			<el-descriptions-item>
				<template #label>
					<div class="cell-item">
						<el-icon :style="iconStyle">
							<iphone />
						</el-icon>
						Telephone
					</div>
				</template>
				18100000000
			</el-descriptions-item>
			<el-descriptions-item>
				<template #label>
					<div class="cell-item">
						<el-icon :style="iconStyle">
							<location />
						</el-icon>
						Place
					</div>
				</template>
				Suzhou
			</el-descriptions-item>
			<el-descriptions-item>
				<template #label>
					<div class="cell-item">
						<el-icon :style="iconStyle">
							<tickets />
						</el-icon>
						Remarks
					</div>
				</template>
				<el-tag size="small">School</el-tag>
			</el-descriptions-item>
			<el-descriptions-item>
				<template #label>
					<div class="cell-item">
						<el-icon :style="iconStyle">
							<office-building />
						</el-icon>
						Address
					</div>
				</template>
				No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province
			</el-descriptions-item>
		</el-descriptions> -->
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElButton, ElDrawer } from 'element-plus';
import { CircleCloseFilled } from '@element-plus/icons-vue';
import { V1ConfigMap } from '@kubernetes/client-node';
import { computed, reactive } from 'vue';
import { watch } from 'vue';

const data = reactive({
	visible: false,
	configMap: {} as V1ConfigMap,
});

const emit = defineEmits(['update:visible']);

const props = defineProps({
	visible: Boolean,
	configMap: V1ConfigMap,
	title: String,
});

const iconStyle = computed(() => {
	const marginMap = {
		large: '8px',
		default: '6px',
		small: '4px',
	};
	return {
		marginRight: marginMap['default'] || marginMap.default,
	};
});

const handleClose = () => {
	emit('update:visible', false);
};

watch(
	() => props,
	() => {
		data.visible = props.visible;
		if (props.configMap) {
			data.configMap = props.configMap;
		}
	},
	{
		immediate: true,
	}
);
</script>
<style scoped>
.el-descriptions {
	margin-top: 20px;
}
.cell-item {
	display: flex;
	align-items: center;
}
.margin-top {
	margin-top: 20px;
}
</style>
