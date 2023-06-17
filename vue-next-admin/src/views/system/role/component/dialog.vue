<template>
	<div class="system-role-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px" @close="onCancel">
			<el-form ref="roleDialogFormRef" :model="state.ruleForm" size="default" label-width="90px" :rules="roleFormRules">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="角色名称"  prop="name">
							<el-input v-model="state.ruleForm.name"  placeholder="请输入角色名称" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="角色描述" prop="memo">
							<template #label>
								<el-tooltip effect="dark" content="角色的描述信息" placement="top-start">
									<span>角色描述</span>
								</el-tooltip>
							</template>
							<el-input v-model="state.ruleForm.memo" placeholder="请输入角色描述" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="排序" prop="sequence">
							<el-input-number
								v-model="state.ruleForm.sequence"
								:min="0"
								:max="999"
								controls-position="right"
								placeholder="请输入排序"
								class="w100"
							/>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="是否启用" prop="status">
							<el-radio-group v-model="state.ruleForm.status">
								<el-radio :label="1">启用</el-radio>
								<el-radio :label="2">不启用</el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>

				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="small">取 消</el-button>
					<el-button type="primary" @click="onSubmit(roleDialogFormRef)" size="small">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="systemRoleDialog">
import { reactive, ref } from 'vue';
import {  RoleType } from '@/types/views';
import {ElMessage, FormInstance, FormRules} from 'element-plus';
import {useRoleApi} from "@/api/system/role";

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const roleApi = useRoleApi();
const roleDialogFormRef = ref();

const state = reactive({
	ruleForm: {
		name: '', // 角色名称
		memo: '', // 角色描述
		sequence: 1, // 排序
		status: 1, // 角色状态
	},
  roleID: 0,
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: RoleType) => {
	if (type === 'edit') {
		state.ruleForm =JSON.parse(JSON.stringify(row));
    state.roleID = row.id;
		state.dialog.title = '修改角色';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增角色';
		state.dialog.submitTxt = '新 增';

	}
  state.dialog.type = type
	state.dialog.isShowDialog = true;
};

// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
  roleDialogFormRef.value.resetFields();
};

// 取消
const onCancel = () => {
	closeDialog();
};

// 提交
const onSubmit =async (formEl: FormInstance|undefined) => {
  if(!formEl) return;
  await formEl.validate( async (valid: boolean)=>{
    if (valid) {
      if (state.dialog.type==="add") {
        await  roleApi.addRole(state.ruleForm).then(()=>{
          ElMessage.success("添加成功")
        }).catch((res) => {
          ElMessage.error(res.message)
        })
      } else {
        await roleApi.updateRole(state.roleID, state.ruleForm).then(()=>{
          ElMessage.success("修改成功")
        }).catch((res) => {
          ElMessage.error(res.message)
        })
      }
      closeDialog();
      emit('refresh');
    }
  })

};

const roleFormRules = reactive<FormRules>({
  name: [{ required: true, message: '请输入角色', trigger: 'blur' }],
});
// 暴露变量
defineExpose({
	openDialog,
});
</script>

<style scoped lang="scss">
.system-role-dialog-container {
	.menu-data-tree {
		width: 100%;
		border: 1px solid var(--el-border-color);
		border-radius: var(--el-input-border-radius, var(--el-border-radius-base));
		padding: 5px;
	}
}
</style>
