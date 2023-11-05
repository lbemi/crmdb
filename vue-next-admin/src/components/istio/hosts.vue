<template>
	<el-form ref="hostRef" :model="state">
		<el-form-item :label="name">
			<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
				<el-form-item>
					<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRow">新增</el-button>
					<el-text class="mx-1" type="info" tag="b">目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务</el-text>
				</el-form-item>
			</el-col>
			<el-col>
				<div v-for="(host, i) in state.hosts" :key="i" class="mb20">
					<el-form-item :prop="`hosts[${i}].value`" :rules="formRules.value">
						<div style="display: flex">
							<el-input placeholder="请输入host" v-model="state.hosts[i].value" size="small" />
							<el-button v-if="i > 0" :icon="RemoveFilled" type="primary" size="small" text @click="onDelRow(i)" class="ml-2"></el-button>
						</div>
					</el-form-item>
				</div>
			</el-col>
		</el-form-item>
	</el-form>
</template>
<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { FormRules } from 'element-plus';
import { onMounted, ref, reactive } from 'vue';

interface host {
	value: string;
}
const hostRef = ref();
const state = reactive({
	hosts: [{ value: '' }] as Array<host>,
});

const formRules = reactive<FormRules>({
	value: [{ required: true, message: '请输入host', trigger: 'blur' }],
});
const onAddRow = () => {
	state.hosts.push({ value: '' });
};
const onDelRow = (k: number) => {
	state.hosts.splice(k, 1);
};

const props = defineProps({
	hosts: {
		type: Array<string>,
		default: () => {
			return [];
		},
	},
	name: String,
});

onMounted(() => {
	if (props.hosts.length > 0) {
		props.hosts?.forEach((item) => {
			state.hosts.push({
				value: item,
			});
		});
	}
});

const returnHosts = () => {
	let hosts = [] as Array<string>;
	state.hosts.forEach((item) => {
		hosts.push(item.value);
	});
	return hosts;
};
const validateHandler = async () => {
	let status = false;
	if (!hostRef.value) return false;
	await hostRef.value.validate((valid: boolean) => {
		status = valid;
	});
	return status;
};

defineExpose({
	returnHosts,
	validateHandler,
});
</script>
