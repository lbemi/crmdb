<template>
	<el-form :model="state" ref="ruleFormRef" style="margin-left: 35px" status-icon :rules="formRules">
		<el-form-item label="HTTP">
			<el-form-item>
				<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRow">新增</el-button>
				<el-text class="mx-1" type="info" tag="b">目标主机,通配符前缀的服务名或IP。适用于 HTTP 和 TCP服务</el-text>
			</el-form-item>
		</el-form-item>
		<div v-for="(http, k) in state.https" :key="k">
			<el-col :span="24" class="mb20">
				<el-card style="margin-left: 30px; margin-right: 30px" class="box-card">
					<el-row :gutter="20" align="middle">
						<el-col :span="1" style="display: flex">
							<el-button :icon="RemoveFilled" type="primary" size="small" text @click="onDelRow(k)" class="ml-2" />
						</el-col>
						<el-col :span="23" class="mb20">
							<el-form-item
								label="http配置名称"
								:prop="`https[${k}].name`"
								:rules="[{ required: true, message: `http配置名称不能为空`, trigger: 'blur' }]"
							>
								<el-input v-model="state.https[k].name" style="max-width: 100px" />
							</el-form-item>
							<el-col class="mb20">
								<el-card>
									<el-form-item label="匹配条件" style="display: flex">
										<div>
											<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddMatch(k)">新增</el-button>
										</div>
									</el-form-item>
									<div style="margin-left: 30px">
										<el-row :gutter="10" v-for="(item, index) in state.https[k].match" :key="index">
											<el-col :span="5" :offset="1" class="grid-content ep-bg-purple mb20">
												<el-form-item label="名称">
													<el-input placeholder="可以为空" v-model="item.name" size="small" style="width: 100px" />
												</el-form-item>
											</el-col>
											<el-col :span="7" class="grid-content ep-bg-purple mb20">
												<el-form-item label="匹配规则" :prop="`https[${k}].match[${index}].reg`" :rules="formRules.key">
													<el-select v-model="item.reg" size="small" style="width: 120px">
														<el-option v-for="item in prefixType" :key="item.key" :label="item.key" :value="item.value" />
													</el-select>
												</el-form-item>
											</el-col>
											<el-col :span="10" class="grid-content ep-bg-purple mb20">
												<el-form-item label="匹配路径" :prop="`https[${k}].match[${index}].value`" :rules="formRules.value">
													<el-input placeholder="匹配值" v-model="item.value" size="small" style="width: 120px" />
													<el-button :icon="RemoveFilled" type="primary" size="small" text @click="onDelMatch(k, index)" class="ml-2"></el-button>
												</el-form-item>
											</el-col>
											<el-col :span="1" />
										</el-row>
										<el-form-item> </el-form-item>
									</div>
								</el-card>
							</el-col>
							<el-col class="mb20">
								<el-card>
									<el-form-item label="路由配置" style="display: flex">
										<div>
											<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddRoute(k)">新增</el-button>
										</div>
									</el-form-item>
									<div>
										<el-row :gutter="10" v-for="(item, j) in state.https[k].route" :key="j">
											<el-card>
												<el-col :span="24" class="grid-content ep-bg-purple mb20">
													<el-form-item label="目标">
														<el-button :icon="CirclePlusFilled" type="primary" size="small" text @click="onAddDest(k, j)">新增</el-button>
														<el-button :icon="RemoveFilled" type="danger" size="small" text @click="onDelroute(k, j)" class="ml-2"
															>删除此目标</el-button
														>
													</el-form-item>
													<el-card v-for="(dest, i) in state.https[k].route[j].destination" :key="i" class="mb10">
														<el-row :gutter="10">
															<el-col :span="8" :offset="1" class="grid-content ep-bg-purple mb20">
																<el-form-item label="服务名" :prop="`https[${k}].route[${j}].destination[${i}].host`" :rules="formRules.host">
																	<el-input placeholder="服务名" v-model="dest.host" size="small" />
																</el-form-item>
															</el-col>
															<el-col :span="4" class="grid-content ep-bg-purple mb20">
																<el-form-item label="子集">
																	<el-input placeholder="子集" v-model="dest.subset" size="small" />
																</el-form-item>
															</el-col>
															<el-col :span="3" class="grid-content ep-bg-purple mb20">
																<el-form-item label="端口">
																	<el-input placeholder="0-100" v-model.number="dest.port.number" size="small" />
																</el-form-item>
															</el-col>
															<el-col :span="7" class="grid-content ep-bg-purple mb20">
																<el-form-item label="权重">
																	<el-input placeholder="0-100" v-model.number="dest.weight" size="small" style="width: 55px" />
																	<el-button
																		:icon="RemoveFilled"
																		type="primary"
																		size="small"
																		text
																		@click="onDelDest(k, j, i)"
																		class="ml-2"
																	></el-button>
																</el-form-item>
															</el-col>
															<el-col :span="24" class="grid-content ep-bg-purple mb20">
																<el-form-item label="header操作">
																	<el-button
																		:icon="CirclePlusFilled"
																		type="primary"
																		size="small"
																		text
																		@click="onAddHeaderOperator(k, j, i)"
																		:disabled="state.https[k].route[j].destination[i].headers.length === 2"
																		>新增</el-button
																	>
																	<el-button :icon="RemoveFilled" type="danger" size="small" text @click="onDelHeaderOperator(k, j, i)" class="ml-2"
																		>删除此header</el-button
																	>
																</el-form-item>
																<div v-if="state.https[k].route[j].destination[i].headers.length > 0">
																	<el-row :gutter="20" v-for="(header, p) in state.https[k].route[j].destination[i].headers" :key="p">
																		<el-col :span="5" :offset="1" class="grid-content ep-bg-purple mb20">
																			<el-form-item
																				label="类型"
																				:prop="`https[${k}].route[${j}].destination[${i}].headers[${p}].type`"
																				:rules="formRules.headerType"
																			>
																				<el-select v-model="header.type" size="small">
																					<el-option v-for="item in headerType" :key="item.key" :label="item.key" :value="item.value" />
																				</el-select>
																			</el-form-item>
																		</el-col>
																		<el-col :span="5" class="grid-content ep-bg-purple mb20">
																			<el-form-item
																				label="操作"
																				:prop="`https[${k}].route[${j}].destination[${i}].headers[${p}].option`"
																				:rules="formRules.headerOperType"
																			>
																				<el-select v-model="header.option" size="small">
																					<el-option v-for="item in headerOperType" :key="item.key" :label="item.key" :value="item.value" />
																				</el-select>
																			</el-form-item>
																		</el-col>
																		<el-col :span="5" class="grid-content ep-bg-purple mb20">
																			<el-form-item
																				label="key"
																				:prop="`https[${k}].route[${j}].destination[${i}].headers[${p}].target.key`"
																				:rules="formRules.headerKey"
																			>
																				<el-input placeholder="输入key" v-model="header.target.key" size="small" />
																			</el-form-item>
																		</el-col>
																		<el-col :span="5" class="grid-content ep-bg-purple mb20" v-if="header.option != 'remove'">
																			<el-form-item
																				label="value"
																				:prop="`https[${k}].route[${j}].destination[${i}].headers[${p}].target.value`"
																				:rules="formRules.headerValue"
																			>
																				<el-input placeholder="输入value" v-model="header.target.value" size="small" />
																			</el-form-item>
																		</el-col>
																		<el-col :span="3" class="grid-content ep-bg-purple mb20">
																			<el-button
																				:icon="RemoveFilled"
																				type="primary"
																				size="small"
																				text
																				@click="onDelHeader(k, j, i, p)"
																				class="ml-2"
																			></el-button>
																		</el-col>
																	</el-row>
																</div>
															</el-col>
														</el-row>
													</el-card>
												</el-col>
											</el-card>
										</el-row>
									</div>
								</el-card>
							</el-col>
						</el-col>
					</el-row>
				</el-card>
			</el-col>
		</div>
	</el-form>
</template>
<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { reactive, ref } from 'vue';
import { VirtualServiceHttp } from '@/types/istio/http';
import { FormInstance, FormRules } from 'element-plus';

interface Header {
	type: string;
	option: string;
	target: {
		key: string;
		value: string;
	};
}
const ruleFormRef = ref<FormInstance>();
const state = reactive({
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
			route: [
				{
					destination: [
						{
							weight: 100,
							host: '',
							subset: '',
							port: {
								number: 0,
							},
							headers: [] as Array<Header>,
						},
					],
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

	httpName: [{ required: true, message: '输入名称', trigger: 'blur' || 'change' }],
	headerType: [{ required: true, message: '请选择', trigger: 'blur' || 'change' }],
	headerOperType: [{ required: true, message: '请选择操作类型', trigger: 'blur' || 'change' }],
	headerKey: [
		{
			required: true,
			message: '请输入key',
			trigger: 'blur',
		},
	],
	headerValue: [
		{
			required: true,
			message: '请输入value',
			trigger: 'blur',
		},
	],

	key: [
		{
			required: true,
			message: '请选择匹配规则',
			trigger: 'blur',
		},
	],
	value: [
		{
			required: true,
			message: '输入匹配路径',
			trigger: 'blur',
		},
	],
	host: [
		{
			required: true,
			message: '输入服务名',
			trigger: 'blur',
		},
	],
	subset: [
		{
			required: true,
			message: '输入子集',
			trigger: 'blur',
		},
	],
});
const validateHandler = async () => {
	let status = false;
	if (!ruleFormRef.value) return false;
	await ruleFormRef.value.validate((valid: boolean) => {
		status = valid;
	});
	return status;
};
const onAddMatch = (index: number) => {
	state.https[index].match.push({
		name: '',
		reg: '',
		value: '',
	});
};
const onAddRoute = (k: number) => {
	state.https[k].route.push({
		destination: [
			{
				weight: 100,
				host: '',
				subset: '',
				port: {
					number: 0,
				},
				headers: [],
			},
		],
	});
};
const onDelroute = (k: number, j: number) => {
	state.https[k].route.splice(j, 1);
};
const onAddDest = (k: number, j: number) => {
	state.https[k].route[j].destination.push({
		weight: 100,
		host: '',
		subset: '',
		port: {
			number: 0,
		},
		headers: [],
	});
};
const onDelDest = (k: number, j: number, i: number) => {
	state.https[k].route[j].destination.splice(i, 1);
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
		route: [
			{
				destination: [
					{
						host: '',
						subset: '',
						weight: 0,
						port: {
							number: 0,
						},
						headers: [],
						// headers2: {
						// 	request: {
						// 		action: '',
						// 		key: '',
						// 		value: '',
						// 		keys: [''],
						// 	},
						// 	response: {
						// 		action: '',
						// 		key: '',
						// 		value: '',
						// 		keys: [''],
						// 	},
						// },
					},
				],
			},
		],
	});
};

const onAddHeaderOperator = (k: number, j: number, i: number) => {
	state.https[k].route[j].destination[i].headers.push({
		type: '',
		option: '',
		target: {
			key: '',
			value: '',
		},
	});
};
const onDelHeaderOperator = (k: number, j: number, i: number) => {
	state.https[k].route[j].destination[i].headers.splice(j, 1);
};
const onDelHeader = (k: number, j: number, i: number, p: number) => {
	state.https[k].route[j].destination[i].headers.splice(p, 1);
};

const prefixType = [
	{ key: '精确匹配', value: 'exact' },
	{ key: '前缀匹配', value: 'prefix' },
	{ key: '正则匹配', value: 'regex' },
];
const headerType = [
	{
		key: '请求',
		value: 'request',
	},
	{
		key: '响应',
		value: 'response',
	},
];
const headerOperType = [
	{
		key: '设置',
		value: 'set',
	},
	{
		key: '新增',
		value: 'add',
	},
	{
		key: '删除',
		value: 'remove',
	},
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
		item.match.forEach((v, i) => {
			match.push({
				name: state.https[index].match[i].name,
				uri: {
					[`${state.https[index].match[i].reg}`]: state.https[index].match[i].value,
				},
			});
		});

		let route = {};
		item.route.forEach((v) => {
			let dest = [] as Array<Object>;
			v.destination.forEach((d) => {
				let headers = {};
				if (d.headers.length > 0 && d.headers.length < 3) {
					let request = {};
					let response = {};
					d.headers.forEach((h) => {
						if (h.type === 'request') {
							if (h.option === 'remove') {
								request = {
									[`${h.type}`]: {
										[`${h.option}`]: [h.target.key],
									},
								};
							} else {
								request = {
									[`${h.type}`]: {
										[`${h.option}`]: {
											[`${h.target.key}`]: h.target.value,
										},
									},
								};
							}
						} else {
							if (h.option === 'remove') {
								response = {
									[`${h.type}`]: {
										[`${h.option}`]: [h.target.key],
									},
								};
							} else {
								response = {
									[`${h.type}`]: {
										[`${h.option}`]: {
											[`${h.target.key}`]: h.target.value,
										},
									},
								};
							}
						}
					});
					headers = {
						...request,
						...response,
					};
					dest.push({
						destination: {
							host: d.host,
							subset: d.subset,
							port: {
								number: d.port.number,
							},
						},
						weight: d.weight,
						headers: headers,
					});
				} else {
					dest.push({
						destination: {
							host: d.host,
							subset: d.subset,
							port: {
								number: d.port.number,
							},
						},
						weight: d.weight,
					});
				}
			});
			route = dest;
		});
		http.match = match;
		http.route = route;
		res.push(http);
	});
	return res;
};

defineExpose({
	returnHttps,
	validateHandler,
});
</script>

<style scoped lang="scss"></style>
