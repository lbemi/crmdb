<template>
	<div class="form-adapt-container layout-pd">
		<el-card shadow="hover" header="创建Deployment">
			<div class="demo-collapse">
				<el-collapse v-model="activeName" @change="">
					<el-collapse-item title="基础信息" name="1">
						<template #title>
							基础信息<el-icon class="header-icon">
								<info-filled />
							</el-icon>
						</template>
						<el-form :model="state.meta" size="default" label-width="100px" class="mt20" label-position="left">
							<el-row :gutter="35">
								<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
									<el-form-item label="命名空间">
										<el-select v-model="state.meta.namespace" class="m-2" placeholder="Select">
											<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata!.name"
												:label="item.metadata!.name" :value="item.metadata!.name!" />
										</el-select>
									</el-form-item>
								</el-col>
								<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
									<el-space fill>
										<el-alert title="info alert" type="info" show-icon :closable="false" />
										<el-form-item label="应用名称" prop="name">
											<el-input v-model="state.meta.name" style="width: 220px" />
										</el-form-item>
										<el-alert type="info" show-icon :closable="false" title="基础信息"
											description="名称大小写不敏感" />
									</el-space>
								</el-col>

								<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
									<el-form-item label="副本数量">
										<el-input-number v-model="state.spec.replicas" :min="1" :max="100" />
									</el-form-item>
								</el-col>


								<el-col>
									<el-form-item label=" 标签">
										<el-button :icon="CirclePlusFilled" type="primary" size="small" text
											@click="onAddRow">新增</el-button>
									</el-form-item>
									<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
										<el-row v-for="(v, k) in state.form.list" :key="k">
											<el-col :span="5" class="mb20">
												<el-form-item label="键" :prop="`list[${k}].year`"
													:rules="[{ required: true, message: `年度不能为空`, trigger: 'blur' }]">
													<el-input v-model="state.form.list[k].year" style="width: 100%"
														placeholder="请输入key,不能重复">
													</el-input>
												</el-form-item>
											</el-col>
											<el-col :span="5" class="mb20" style="display: flex; align-items: center;">
												<el-form-item label="值" :prop="`list[${k}].month`"
													:rules="[{ required: true, message: `月度不能为空`, trigger: 'blur' }]">
													<el-input v-model="state.form.list[k].month" style="width: 100%"
														placeholder="请输入value">
													</el-input>
												</el-form-item>

											</el-col>
											<el-col :span="2">
												<el-form-item>
													<el-button :icon="RemoveFilled" type="primary" size="small" text
														@click="onDelRow(k)"></el-button>
												</el-form-item>
											</el-col>

										</el-row>
									</el-col>
								</el-col>

							</el-row>
						</el-form>
					</el-collapse-item>
					<el-collapse-item title="Feedback" name="2">
						<div>
							Operation feedback: enable the users to clearly perceive their
							operations by style updates and interactive effects;
						</div>
						<div>
							Visual feedback: reflect current state by updating or rearranging
							elements of the page.
						</div>
					</el-collapse-item>
					<el-collapse-item title="Efficiency" name="3">
						<div>
							Simplify the process: keep operating process simple and intuitive;
						</div>
						<div>
							Definite and clear: enunciate your intentions clearly so that the
							users can quickly understand and make decisions;
						</div>
						<div>
							Easy to identify: the interface should be straightforward, which helps
							the users to identify and frees them from memorizing and recalling.
						</div>
					</el-collapse-item>
					<el-collapse-item title="Controllability" name="4">
						<div>
							Decision making: giving advices about operations is acceptable, but do
							not make decisions for the users;
						</div>
						<div>
							Controlled consequences: users should be granted the freedom to
							operate, including canceling, aborting or terminating current
							operation.
						</div>
					</el-collapse-item>
				</el-collapse>
			</div>

		</el-card>

	</div>
</template>

<script setup lang="ts" name="pagesListAdapt">
import { defineAsyncComponent, reactive, ref } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { CirclePlusFilled, RemoveFilled, InfoFilled } from '@element-plus/icons-vue';


const Label = defineAsyncComponent(() => import('@/components/kubernetes/label.vue'));

const k8sStore = kubernetesInfo();

// 定义变量内容

const activeName = ref('1')
const state = reactive({
	meta: {
		name: '',
		namespace: '',

	},
	spec: {
		replicas: 1,
	},
	labelData: [{ key: 'app', value: '' }],
	form: {
		name: '',
		email: '',
		autograph: '',
		occupation: '',
		list: [
			{
				year: '',
				month: '',
				day: '',
			},
		],
		remarks: '',
	},
});

// 新增行
const onAddRow = () => {
	state.form.list.push({
		year: '',
		month: '',
		day: '',
	});
};
// 删除行
const onDelRow = (k: number) => {
	state.form.list.splice(k, 1);
};
</script>
