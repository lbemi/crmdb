<template>
	<el-form-item label="HTTP">
		<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
			<el-form-item>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRow">新增</el-button>
				<el-text class="mx-1" type="info" tag="b">目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务</el-text>
			</el-form-item>
		</el-col>
		<div v-if="data.https">
			<el-form-item v-for="(http, i) in data.https" :key="i">
				<div style="display: flex; margin-bottom: 20px">
					<el-input v-model="http.route" size="small" />
					<el-button v-if="i > 0" :icon="RemoveFilled" type="primary" size="small" text @click="onDelRow(i)" class="ml-2"></el-button>
				</div>
			</el-form-item>
		</div>
	</el-form-item>
</template>
<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';

import { reactive, watch } from 'vue';

import { deepClone } from '@/utils/other';

const data = reactive({
	https: [] as Array<VirtualServiceHttp>,
});

const props = defineProps({
	https: Array<VirtualServiceHttp>,
});

const onAddRow = () => {
	data.https.push({} as VirtualServiceHttp);
};
const onDelRow = (index: number) => {
	data.https.splice(index, 1);
};

watch(
	() => props.https,
	() => {
		if (props.https && props.https.length > 0) {
			data.https = deepClone(props.https);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>

<style scoped lang="scss"></style>
