<template>
	<el-form-item :label="name">
		<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
			<el-form-item>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRow">新增</el-button>
				<el-text class="mx-1" type="info" tag="b">目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务</el-text>
			</el-form-item>
		</el-col>
		<div v-if="hosts">
			<el-form-item v-for="(host, i) in hosts" :key="i">
				<div style="display: flex; margin-bottom: 20px">
					<el-input v-model="props.hosts[i]" size="small" />
					<el-button v-if="i > 0" :icon="RemoveFilled" type="primary" size="small" text @click="onDelRow(i)" class="ml-2"></el-button>
				</div>
			</el-form-item>
		</div>
	</el-form-item>
</template>
<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';

const props = defineProps({
	hosts: {
		type: Array<string>,
		default: () => {
			return [];
		},
	},
	name: String,
});
const validateHost = (index: number, rule: any, value: any, callback: Function) => {
	if (props.hosts[index]) {
		callback(new Error('Host cannot be empty'));
	} else {
		callback();
	}
};

const onAddRow = () => {
	props.hosts.push('');
};
const onDelRow = (k: number) => {
	props.hosts.splice(k, 1);
};
</script>
<style scoped lang="scss"></style>
