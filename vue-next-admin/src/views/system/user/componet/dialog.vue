<template>
	<div class="system-user-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px" @close="closeDialog(ruleFormRef)">
			<el-form ref="ruleFormRef" :model="state.ruleForm" size="default" label-width="90px" :rules="userFormRules">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="账户名称" prop="user_name">
							<el-input v-model="state.ruleForm.user_name" placeholder="请输入账户名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="邮箱" prop="email">
							<el-input v-model="state.ruleForm.email" placeholder="请输入" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" v-if="state.userId === 0" class="mb20">
						<el-form-item label="账户密码" prop="password">
							<el-input v-model="state.ruleForm.password" placeholder="请输入密码" type="password" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="用户状态" prop="status">
							<el-switch
								v-model="state.ruleForm.status"
								:active-value="1"
								:inactive-value="2"
								inline-prompt
								active-text="启"
								inactive-text="禁"
							></el-switch>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" v-if="state.userId === 0" class="mb20">
						<el-form-item label="确认密码" prop="confirmPassword">
							<el-input v-model="state.ruleForm.confirmPassword" placeholder="请再次输入密码" type="password" clearable></el-input>
						</el-form-item>
					</el-col>

					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="用户描述" prop="description" >
							<el-input v-model="state.ruleForm.description" type="textarea"  placeholder="请输入用户描述" maxlength="150"></el-input>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="closeDialog(ruleFormRef)" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit(ruleFormRef)" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="systemUserDialog">
import { reactive, ref } from 'vue';
import { FormInstance, FormRules, ElMessage } from 'element-plus';
import { useUserApi } from '/@/api/system/user';
import {RowUserType} from "/@/types/views";
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const userApi = useUserApi();

const ruleFormRef = ref<FormInstance>();

const state = reactive({
	ruleForm: {
		user_name: '',
		email: '',
		status: 1,
		password: '',
		description: '',
		confirmPassword: '', //二次密码确认
	},
	userId: 0,
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: RowUserType) => {
	if (type === 'edit') {
		state.userId = row.id;
		state.ruleForm.user_name = row.user_name;
		state.ruleForm.email = row.email;
		state.ruleForm.status = row.status;
		state.ruleForm.description = row.description;

		state.dialog.title = '修改用户';
		state.dialog.submitTxt = '修 改';
	} else {
		state.userId = 0;
		state.dialog.title = '新增用户';
		state.dialog.submitTxt = '新 增';
	}
	state.dialog.isShowDialog = true;
};

const onSubmit = async (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	await formEl.validate(async (valid: boolean) => {
		if (valid) {
			if (state.userId == 0) {
				await userApi.addUser(state.ruleForm).then(() => {
					ElMessage.success('添加成功');
					closeDialog(formEl);
					emit('refresh');
				});
			} else {
				await userApi.updateUser(state.userId, state.ruleForm).then(() => {
					ElMessage.success('修改成功');
					closeDialog(formEl);
					emit('refresh');
				});
			}
		} else {
			ElMessage.error('请正确填写');
		}
	});
};

// 关闭弹窗
const closeDialog = (formEl: FormInstance | undefined) => {

	if (!formEl) return;
	formEl.resetFields();
    state.dialog.isShowDialog = false;
	// userDialogFormRef.value.resetFields();
};

//自定义校验password规则
const validatePass = (rule: any, value: string, callback: any) => {
	if (value === '') {
		callback(new Error('请输入密码'));
	} else {
		if (state.ruleForm.confirmPassword !== '') {
			if (!ruleFormRef.value) return;
      if (ruleFormRef.value) {
        ruleFormRef.value.validateField('confirmPassword', () => null);
      }
		}
		callback();
	}
};

//自定义校验confirm password规则
const validateConfirmPass = (rule: any, value: string, callback: any) => {
	if (value === '') {
		callback(new Error('请再次输入密码'));
	} else if (value !== state.ruleForm.password) {
		callback(new Error('两次密码不匹配'));
	} else {
		callback();
	}
};

//自定义email邮箱规则
const validateEmail = (rule: any, value: string, callback: any) => {
	const regEmail =
		/^(([^<>()\\[\]\\.,;:\s@"]+(\.[^<>()\\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
	if (value) {
		if (regEmail.test(value)) {
			return callback();
		}
		callback(new Error('请输入正确的邮箱'));
	}
};

const userFormRules = reactive<FormRules>({
	user_name: [{ required: true, message: '请输入名字', trigger: 'blur' }],
	password: [
		{ required: true, validator: validatePass, trigger: 'blur' },
		{ min: 6, max: 12, message: '密码长度在6到12位之间', trigger: 'blur' },
	],
	confirmPassword: [{ required: true, validator: validateConfirmPass, trigger: 'blur' }],
	email: [{ validator: validateEmail, trigger: 'blur' }],
});

// 暴露变量
defineExpose({
	openDialog,
});
</script>
