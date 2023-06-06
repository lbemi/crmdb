<template>
	<div class="system-menu-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="apiDialogFormRef" :model="state.ruleForm" size="default" label-width="80px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="菜单类型">
							<el-radio-group v-model.number="state.ruleForm.menuType">
								<el-radio-button label="3">API</el-radio-button>
								<el-radio-button label="2">按钮</el-radio-button>
							</el-radio-group>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="显示名称">
							<template #label>
								<el-tooltip effect="dark" content="菜单的描述信息仅做展示使用" placement="top-start">
									<span>显示名称</span>
								</el-tooltip>
							</template>
							<el-input v-model="state.ruleForm.memo" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="API分组">
							<el-input v-model="state.ruleForm.group" placeholder="分组" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="路由名称">
							<el-input v-model="state.ruleForm.name" placeholder="路由中的 name 值" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20" v-show="state.ruleForm.menuType === 2">
						<el-form-item label="权限标识">
							<el-input v-model="state.ruleForm.code" placeholder="请输入权限标识" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="路由路径">
							<el-input v-model="state.ruleForm.path" placeholder="路由中的 path 值" clearable></el-input>
						</el-form-item>
					</el-col>

					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="菜单排序">
							<el-input-number v-model="state.ruleForm.sequence" controls-position="right" placeholder="请输入排序" class="w100" />
						</el-form-item>
					</el-col>

					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item id="ss" v-model="state.ruleForm.menuType" label="请求方式:">
							<el-select v-model="state.ruleForm.method" clearable placeholder="请选择" value-key="value">
								<el-option v-for="item in state.methodOptions" :key="item.value" :label="item.label" :value="item.value" />
							</el-select>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="是否启用">
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
					<el-button type="primary" @click="onSubmit" size="small">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="systemMenuDialog">
import { reactive, ref } from 'vue';
import { useMenuApi } from '/@/api/system/menu';
import { ElMessage } from 'element-plus';

// import { setBackEndControlRefreshRoutes } from "/@/router/backEnd";

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
const menuApi = useMenuApi();

// 定义变量内容
const apiDialogFormRef = ref();
const state = reactive({
	// 参数请参考 `/src/router/route.ts` 中的 `dynamicRoutes` 路由菜单格式
	ruleForm: {
		id: 0,
		menuType: 3, // 菜单类型 1为菜单 2为按钮 3为API
		code: '',
		name: '', // 路由名称
		memo: '', //显示的名称
		sequence: 1, // 菜单排序
		path: '', // 路由路径
		redirect: '', // 路由重定向，有子集 children 时
		status: 1,
		group: '未分组',
		method: 'GET',
	},
	methodOptions: [
		{
			value: 'GET',
			label: 'GET',
		},
		{
			value: 'POST',
			label: 'POST',
		},
		{
			value: 'DELETE',
			label: 'DELETE',
		},
		{
			value: 'PUT',
			label: 'PUT',
		},
	],
	menuData: [], // 上级菜单数据
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
	params: {
		menuType: '1',
	},
});

// 打开弹窗
const openDialog = (type: string, row?: any) => {
	if (type === 'edit') {
		state.ruleForm = JSON.parse(JSON.stringify(row));

		state.dialog.title = '修改API';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增API';
		state.dialog.submitTxt = '新 增';
	}
	state.dialog.type = type;
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
	apiDialogFormRef.value.resetFields();
};

// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = async () => {
	if (state.dialog.type === 'add') {
		await menuApi
			.addMenu(state.ruleForm)
			.then(() => {
				ElMessage.success('添加成功');
				closeDialog(); // 关闭弹窗
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
		closeDialog(); // 关闭弹窗
	} else {
		await menuApi
			.updateMenu(state.ruleForm.id, state.ruleForm)
			.then(() => {
				ElMessage.success('修改成功');
				closeDialog(); // 关闭弹窗
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
		closeDialog(); // 关闭弹窗
	}
};

// 暴露变量
defineExpose({
	openDialog,
});
</script>
