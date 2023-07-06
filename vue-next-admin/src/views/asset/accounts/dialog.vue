<template>
	<div class="system-dept-dialog-container">
		<el-dialog class="my-dialog" :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="formRef" :model="state.account" size="default" label-width="90px" :rules="formRules">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户" prop="name">
							<el-input v-model="state.account.name" placeholder="名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="登录名" prop="user_name">
							<el-input v-model="state.account.user_name" placeholder="请输入登录名" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="密码" prop="password">
							<el-input type="password" v-model.number="state.account.password" placeholder="请输密码" clearable></el-input>
						</el-form-item>
					</el-col>

					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户状态" prop="status">
							<el-radio-group v-model="state.account.status">
								<el-radio :label="1">启用</el-radio>
								<el-radio :label="2">不启用</el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel()" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit(formRef)" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="accountDialog">
import { reactive, ref } from 'vue';
import { ElInput, ElMessage, FormInstance, FormRules } from 'element-plus';
import { useAccountApi } from '@/api/asset/account';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const accountApi = useAccountApi();
const formRef = ref<FormInstance>();
const state = reactive({
	account: {
		name: '',
		status: 1,
		auth_method: '01',
	} as Account,
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: Account) => {
	if (type === 'edit') {
		state.account = row;
		state.dialog.title = '修改账户';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增账户';
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
			console.log(state.account);
			if (state.dialog.title === '新增账户') {
				await accountApi
					.addAccount(state.account)
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

const validateIP = (rule: any, value: any, callback: any) => {
	if (value === '' || typeof value === 'undefined' || value == null) {
		callback(new Error('请输入正确的IP地址'));
	} else {
		const reg =
			/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
		if (!reg.test(value) && value !== '') {
			callback(new Error('请输入正确的IP地址'));
		} else {
			callback();
		}
	}
};

const formRules = reactive<FormRules>({
	ip: [{ required: true, validator: validateIP, trigger: 'blur' }],
	port: [{ required: true, message: '请输入端口', trigger: 'blur' }],
	group_id: [{ required: true, message: '请选择分组', trigger: 'change' }],
});

// 暴露变量
defineExpose({
	openDialog,
});
</script>
<style></style>
