<template>
	<div>
		<el-descriptions :column="1" direction="vertical">
			<el-descriptions-item label="标签">
				<el-card>
					<div v-for="(item, key, index) in props.metaData?.labels" :key="index" style="margin-bottom: 8px">
						<el-tag type="info" size="default"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-card>
			</el-descriptions-item>
			<el-descriptions-item label="注解">
				<el-card>
					<div v-for="(item, key, index) in props.metaData?.annotations" :key="index" style="margin-bottom: 8px">
						<el-tag type="info" size="default"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-card>
			</el-descriptions-item>
		</el-descriptions>
	</div>
</template>

<script setup lang="ts">
import { V1ObjectMeta } from '@kubernetes/client-node';

const props = defineProps({
	metaData: Object<V1ObjectMeta>,
});
</script>

<style scoped>
.el-descriptions {
	margin-top: 20px;
}
</style>
