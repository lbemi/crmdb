<template>
	<el-form :model="state" ref="ruleFormRef" class="demo-ruleForm" status-icon :rules="formRules">
		<el-form-item label="HTTP">
			<el-form-item>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRow">新增</el-button>
				<el-text class="mx-1" type="info" tag="b">目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务</el-text>
				<el-alert type="info" show-icon :closable="true">
					<p>目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务.</p>
				</el-alert>
			</el-form-item>
		</el-form-item>
		<div v-for="(http, k) in state.https" :key="k">
			<el-col :span="24" class="mb20">
				<el-card style="margin-left: 30px; margin-right: 30px">
					<el-row :gutter="20">
						<el-col :span="1" style="justify-items: center">
							<el-button :icon="RemoveFilled" type="primary" size="small" text @click="onDelRow(k)" class="ml-2"></el-button>
						</el-col>
						<el-col :span="23" class="mb20">
							<el-form-item
								label="http配置名称"
								:prop="`https[${k}].name`"
								:rules="[{ required: true, message: `http配置名称不能为空`, trigger: 'blur' }]"
							>
								<el-input v-model="state.https[k].name" style="max-width: 100px" />
							</el-form-item>
							<el-form-item label="匹配条件" style="display: flex">
								<div>
									<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddMatch(k)">新增</el-button>
								</div>
							</el-form-item>
						</el-col>
					</el-row>
					<div style="margin-left: 30px">
						<el-row :gutter="10" v-for="(item, index) in state.https[k].match" :key="index">
							<el-col :span="4" :offset="2" class="grid-content ep-bg-purple mb20">
								<el-form-item label="名称">
									<el-input placeholder="可以为空" v-model="item.name" size="small" style="width: 100px" />
								</el-form-item>
							</el-col>
							<el-col :span="6" class="grid-content ep-bg-purple mb20">
								<el-form-item label="匹配规则" :prop="`https[${k}].match[${index}].reg`" :rules="formRules.key">
									<el-select v-model="item.reg" size="small" style="width: 120px">
										<el-option v-for="item in prefixType" :key="item.key" :label="item.key" :value="item.value" />
									</el-select>
								</el-form-item>
							</el-col>
							<el-col :span="8" class="grid-content ep-bg-purple mb20">
								<el-form-item label="匹配路径" :prop="`https[${k}].match[${index}].value`" :rules="formRules.value">
									<el-input placeholder="匹配值" v-model="item.value" size="small" style="width: 120px" />
									<el-button :icon="RemoveFilled" type="primary" size="small" text @click="onDelMatch(k, index)" class="ml-2"></el-button>
								</el-form-item>
							</el-col>
							<el-col :span="4" />
						</el-row>
						<el-form-item> </el-form-item>
					</div>
				</el-card>
			</el-col>
		</div>
	</el-form>
	<el-button size="default" type="primary" @click="formCheck(ruleFormRef)">
		<SvgIcon name="iconfont icon-shuxing" />
		验证表单
	</el-button>
</template>
<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { reactive, ref } from 'vue';
import { VirtualServiceHttp } from '@/types/istio/http';
import { ElMessage, FormInstance, FormRules } from 'element-plus';

const ruleFormRef = ref<FormInstance>();
const state = reactive({
	check: false,
	https: [
		{
			name: '',
			match: [
				{
					name: '',
					reg: '',
					value: '',
				},
			],
		},
	],
});

const props = defineProps({
	https: Array<VirtualServiceHttp>,
});

const formRules = reactive<FormRules>({
	name: [
		{ required: true, message: 'Please input Activity name', trigger: 'blur' },
		{ min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
	],
	httpName: [
		{ required: true, message: '输入名称', trigger: 'blur' },
		// { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
	],
	key: [
		{
			required: true,
			message: '请选择匹配规则',
			trigger: 'change',
		},
	],
	value: [
		{
			required: true,
			message: '输入匹配路径',
			trigger: 'change',
		},
	],
});
const formCheck = (formEl: FormInstance | undefined) => {
	if (!formEl) return;

	formEl.validate((valid: boolean) => {
		if (valid) {
			ElMessage.success('验证成功');
			state.check = true;
		} else {
			state.check = false;
		}
	});
};
const onAddMatch = (index: number) => {
	state.https[index].match.push({
		name: '',
		reg: '',
		value: '',
	});
};
const onDelMatch = (index: number, i: number) => {
	state.https[index].match.splice(i, 1);
};

const onAddRow = () => {
	state.https.push({
		name: '',
		match: [
			{
				name: '',
				reg: '',
				value: '',
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
	state.https.splice(index, 1);
	state.https.splice(index, 1);
};

// watch(
// 	() => props.https,
// 	() => {
// 		if (props.https && props.https.length > 0) {
// 			// data.https = deepClone(props.https);
// 		}
// 	},
// 	{
// 		immediate: true,
// 		deep: true,
// 	}
// );

const returnHttps = () => {
	let res = [] as Array<VirtualServiceHttp>;
	state.https.forEach((item, index) => {
		let http = {} as VirtualServiceHttp;
		http.name = item.name;
		let match = [] as Array<Object>;
		state.https[index].match.forEach((v, i) => {
			match.push({
				name: state.https[index].match[i].name,
				uri: {
					[`${state.https[index].match[i].reg}`]: state.https[index].match[i].value,
				},
			});
		});
		http.match = match;
		res.push(http);
	});
	let check = state.check;
	console.log('--', state.check);
	return { res, check };
};

defineExpose({
	returnHttps,
	formCheck,
});
</script>

<style scoped lang="scss"></style>
