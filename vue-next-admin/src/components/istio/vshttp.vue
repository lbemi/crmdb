<template>
	<el-form-item label="HTTP">
		<!--		<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">-->
		<el-form-item>
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRow">新增</el-button>
			<el-text class="mx-1" type="info" tag="b">目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务</el-text>
		</el-form-item>
		<!--		</el-col>-->
	</el-form-item>
	<el-card>
		<div v-for="(http, i) in data.https" :key="i">
			<el-form-item label="名称">
				<el-input v-model="http.name" size="small" />
			</el-form-item>
			<div v-for="(item, index) in http.match" :key="index">
				<el-card>
					<el-form-item label="匹配1"> </el-form-item>
					<el-form-item label="名称">
						<el-input label="匹配名称" v-model="item.name" size="small" />
					</el-form-item>
					<el-form-item label="匹配规则">
						<el-select v-model="data.matchUriType[index]" size="small">
							<el-option v-for="item in prefixType" :key="item.key" :label="item.key" :value="item.value" />
						</el-select>
						<el-input v-model="item.uri[data.matchUriType[index]]" size="small" />
					</el-form-item>
				</el-card>
			</div>
		</div>
	</el-card>
</template>
<script setup lang="ts">
import { CirclePlusFilled } from '@element-plus/icons-vue';

import { reactive, watch } from 'vue';

import { deepClone } from '@/utils/other';
import { VirtualServiceHttp } from '@/types/istio/http';

const data = reactive({
	matchUriType: [''],
	https: [] as Array<VirtualServiceHttp>,
});

const props = defineProps({
	https: Array<VirtualServiceHttp>,
});

const onAddRow = () => {
	data.https.push({
		name: '',
		route: [],
		match: [
			{
				name: '',
				uri: {},
				port: 0,
			},
		],
	});
};

const prefixType = [
	{ key: '精确匹配', value: 'exact' },
	{ key: '前缀匹配', value: 'prefix' },
	{ key: '正则匹配', value: 'regex' },
];
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
const returnHttps = () => {
	return data.https;
};

defineExpose({
	returnHttps,
});
</script>

<style scoped lang="scss"></style>
