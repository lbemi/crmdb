<template>
	<div class="system-dept-dialog-container">
		<el-dialog class="my-dialog" :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="formRef" :model="state.hba" size="default" label-width="90px" :rules="formRules">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="规则名称" prop="rule_name">
							<el-input v-model="state.hba.rule_name" placeholder="请输入规则名称" clearable :disabled="state.isDetail"></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="选择账户" prop="account_id">
							<el-select v-model="state.hba.account_id" multiple placeholder="Select" style="width: 240px" :disabled="state.isDetail">
								<el-option v-for="item in state.accounts" :key="item.name" :label="item.name" :value="item.id" />
							</el-select>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="选择主机" prop="resource_id">
							<el-select v-model="state.hba.resource_id" multiple placeholder="Select" style="width: 240px" :disabled="state.isDetail">
								<el-option v-for="item in state.hosts" :key="item.ip" :label="item.ip" :value="item.id" />
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer" v-if="state.isDetail">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit(formRef)" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="hostDialog">
import { onMounted, reactive, ref } from 'vue';
import { Host } from '@/types/asset/hosts';
import { ElInput, ElMessage, FormInstance, FormRules } from 'element-plus';
import { useHostApi } from '@/api/asset/hosts';
import { useHBAApi } from '@/api/asset/hostBindAccount';
import { useAccountApi } from '@/api/asset/account';
import { deepClone } from '@/utils/other';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const hbaApi = useHBAApi();
const hostApi = useHostApi();
const accountApi = useAccountApi();
const formRef = ref<FormInstance>();
const state = reactive({
	isDetail: false,
	hba: {} as HostBingAccounts,
	hosts: [] as Host[], // 主机数据
	accounts: [] as Account[],
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, hba: HostBingAccounts) => {
	if (type === 'edit') {
		state.hba = deepClone(hba) as HostBingAccounts;
		state.dialog.title = '修改主机';
		state.dialog.submitTxt = '修 改';
	} else if (type === 'detail') {
		state.hba = deepClone(hba) as HostBingAccounts;
		state.dialog.title = '详情';
		state.isDetail = true;
	} else {
		state.dialog.title = '新增主机';
		state.dialog.submitTxt = '新 增';
	}
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	resetForm(formRef.value);
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate(async (valid) => {
		if (valid) {
			if (state.dialog.title === '新增主机') {
				await hbaApi
					.addHba(state.hba)
					.then((res: any) => {
						ElMessage.success(res.message);
						emit('refresh');
						closeDialog();
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
			} else {
				await hbaApi
					.updateHba(state.hba)
					.then((res: any) => {
						ElMessage.success(res.message);
						emit('refresh');
						closeDialog();
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
			}
		} else {
			console.log('error submit!');
			return false;
		}
	});
};

const resetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};

//添加标签
const formRules = reactive<FormRules>({
	resource_id: [{ required: true, message: '请选择主机', trigger: 'blur' }],
	rule_name: [{ required: true, message: '请输规则名称', trigger: 'blur' }],
	account_id: [{ required: true, message: '请选择帐户', trigger: 'change' }],
});

const getHosts = async () => {
	await hostApi
		.lisHost()
		.then((res: any) => {
			state.hosts = res.data.hosts;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
};
const getAccounts = async () => {
	await accountApi
		.listAccount()
		.then((res: any) => {
			state.accounts = res.data.data;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
};

onMounted(() => {
	getAccounts();
	getHosts();
});
// 暴露变量
defineExpose({
	openDialog,
});
</script>
<style></style>
