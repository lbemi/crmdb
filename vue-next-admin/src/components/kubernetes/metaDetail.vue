<template>
	<div>
		<el-descriptions :column="1" direction="vertical">
			<el-descriptions-item label="标签">
				<el-card>
					<div v-if="!props.metaData?.labels">无数据</div>
					<div v-else v-for="(item, key, index) in props.metaData?.labels" :key="index" style="margin-bottom: 8px">
						<el-tag type="info" size="default"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-card>
			</el-descriptions-item>
			<el-descriptions-item label="注解">
				<el-card>
					<div v-if="!props.metaData?.annotations">无数据</div>
					<div v-else v-for="(item, key, index) in props.metaData?.annotations" :key="index" style="margin-bottom: 8px">
						<el-tag type="info" size="default"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-card>
			</el-descriptions-item>
		</el-descriptions>
	</div>
</template>

<script setup lang="ts">
import { ObjectMeta } from 'kubernetes-types/meta/v1';
type propsType = {
	metaData: ObjectMeta | undefined;
};
const props = defineProps<propsType>();
</script>

<style scoped>
/* .el-descriptions {
	margin-top: 10px;
} */
</style>
