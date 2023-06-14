<template>
	<div class="system-menu-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px" @close="closeDialog" >
			<el-form ref="menuDialogFormRef" :model="state.ruleForm" size="default" label-width="80px" :rules="menuRules">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="上级菜单" prop="parentID">
							<el-cascader
								v-model="state.ruleForm.parentID"
								:options="state.menuData"
								class="w100"
								:props="{
									label: 'memo',
									value: 'id',
									checkStrictly: true,
									emitPath: false,
								}"
								clearable
								filterable
								placeholder="选择上级菜单"
								:show-all-levels="false"
							></el-cascader>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="菜单类型" prop="menuType">
							<el-radio-group v-model.number="state.ruleForm.menuType">
								<el-radio-button label="1">菜单</el-radio-button>
							</el-radio-group>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20" v-if="state.ruleForm.menuType == 1">
						<el-form-item label="菜单名称" prop="title">
							<el-input v-model="state.ruleForm.meta.title" placeholder="格式: message.router.xxx" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="显示名称" prop="memo">
							<template #label>
								<el-tooltip effect="dark" content="菜单的描述信息仅做展示使用" placement="top-start">
									<span>显示名称</span>
								</el-tooltip>
							</template>
							<el-input v-model="state.ruleForm.memo" clearable></el-input>
						</el-form-item>
					</el-col>
					<template v-if="state.ruleForm.menuType == 1">
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="路由名称" prop="name">
								<el-input v-model="state.ruleForm.name" placeholder="路由中的 name 值" clearable></el-input>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="路由路径" prop="path">
								<el-input v-model="state.ruleForm.path" placeholder="路由中的 path 值" clearable></el-input>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="重定向" prop="redirect" :disabled="state.ruleForm.parentID != 0 ? true : false">
								<el-input v-model="state.ruleForm.redirect" placeholder="请输入路由重定向" clearable></el-input>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="菜单图标" prop="icon">
								<IconSelector placeholder="请输入菜单图标" v-model="state.ruleForm.meta.icon" />
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="组件路径" prop="component">
								<el-input v-model="state.ruleForm.component" placeholder="组件路径" clearable></el-input>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="链接地址" prop="isLink">
								<el-input v-model="state.ruleForm.meta.isLink" placeholder="外链/内嵌时链接地址（http:xxx.com）" clearable :disabled="!state.isLink">
								</el-input>
							</el-form-item>
						</el-col>
					</template>
					<template v-if="state.ruleForm.menuType == 2">
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="权限标识" prop="code">
								<el-input v-model="state.ruleForm.code" placeholder="请输入权限标识" clearable></el-input>
							</el-form-item>
						</el-col>
					</template>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="菜单排序" prop="sequence">
							<el-input-number v-model="state.ruleForm.sequence" controls-position="right" placeholder="请输入排序" class="w100" />
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="路由路径" prop="path">
							<el-input v-model="state.ruleForm.path" placeholder="路由中的 path 值" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="k8s路由" prop="isK8s">
							<template #label>
								<el-tooltip effect="dark" content="k8s的子路由，只会在k8s的模式下才会显示" placement="top-start">
									<span>k8s路由</span>
								</el-tooltip>
							</template>
							<el-radio-group v-model="state.ruleForm.meta.isK8s">
								<el-radio :label="true">是</el-radio>
								<el-radio :label="false">否</el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>
					<template v-if="state.ruleForm.menuType == 1">
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="是否隐藏" prop="isHide">
								<el-radio-group v-model="state.ruleForm.meta.isHide">
									<el-radio :label="true">隐藏</el-radio>
									<el-radio :label="false">不隐藏</el-radio>
								</el-radio-group>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="页面缓存" prop="isKeepAlive">
								<el-radio-group v-model="state.ruleForm.meta.isKeepAlive">
									<el-radio :label="true">缓存</el-radio>
									<el-radio :label="false">不缓存</el-radio>
								</el-radio-group>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="是否固定" prop="isAffix">
								<el-radio-group v-model="state.ruleForm.meta.isAffix">
									<el-radio :label="true">固定</el-radio>
									<el-radio :label="false">不固定</el-radio>
								</el-radio-group>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="是否外链" prop="isLink">
								<el-radio-group v-model="state.isLink" :disabled="state.ruleForm.meta.isIframe">
									<el-radio :label="true">是</el-radio>
									<el-radio :label="false">否</el-radio>
								</el-radio-group>
							</el-form-item>
						</el-col>
						<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
							<el-form-item label="是否内嵌" prop="isIframe">
								<el-radio-group v-model="state.ruleForm.meta.isIframe" @change="onSelectIframeChange">
									<el-radio :label="true">是</el-radio>
									<el-radio :label="false">否</el-radio>
								</el-radio-group>
							</el-form-item>
						</el-col>
					</template>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="是否启用" prop="status">
							<el-radio-group v-model="state.ruleForm.status">
								<el-radio :label="1">启用</el-radio>
								<el-radio :label="2">不启用</el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item v-if="state.ruleForm.menuType === 2" id="ss" v-model="state.ruleForm.menuType" label="请求方式:" prop="menuType">
							<el-select v-model="state.ruleForm.method" clearable placeholder="请选择" value-key="value">
								<el-option v-for="item in state.methodOptions" :key="item.value" :label="item.label" :value="item.value" />
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="systemMenuDialog">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { useMenuApi } from '/@/api/system/menu';
import { ElMessage, FormInstance, FormRules } from 'element-plus';

// import { setBackEndControlRefreshRoutes } from "/@/router/backEnd";

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);
const menuApi = useMenuApi();
// 引入组件
const IconSelector = defineAsyncComponent(() => import('/@/components/iconSelector/index.vue'));

// 定义变量内容
const menuDialogFormRef = ref<FormInstance>();
const state = reactive({
	// 参数请参考 `/src/router/route.ts` 中的 `dynamicRoutes` 路由菜单格式
	ruleForm: {
		id: 0,
		parentID: 0, // 上级菜单
		menuType: 1, // 菜单类型 1为菜单 2为按钮
		name: '', // 路由名称
		memo: '', //显示的名称
		component: '', // 组件路径
		sequence: 1, // 菜单排序
		path: '', // 路由路径
		redirect: '', // 路由重定向，有子集 children 时
		status: 1,
		group: '',
		code: '', //权限标识
		method: 'GET',
		meta: {
			title: '', // 菜单名称
			icon: '', // 菜单图标
			isHide: false, // 是否隐藏
			isKeepAlive: true, // 是否缓存
			isAffix: false, // 是否固定
			isK8s: false, //是否是k8s路由
			isLink: '', // 外链/内嵌时链接地址（http:xxx.com），开启外链条件，`1、isLink: 链接地址不为空`
			isIframe: false, // 是否内嵌，开启条件，`1、isIframe:true 2、isLink：链接地址不为空`
		},
	},
	isLink: false, //是否是外链
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

// 获取菜单信息
const getMenuData = async () => {
	const res = await menuApi.listMenu(state.params);
	state.menuData = res.data.menus;
};

// 打开弹窗
const openDialog = (type: string, row?: any) => {
	if (type === 'edit') {
		state.ruleForm = JSON.parse(JSON.stringify(row));
		state.dialog.title = '修改菜单';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增菜单';
		state.dialog.submitTxt = '新 增';
		// 清空表单，此项需加表单验证才能使用
		// nextTick(() => {
		// 	menuDialogFormRef.value.resetFields();
		// });
	}
	state.dialog.type = type;
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
	menuDialogFormRef.value?.resetFields();
};
// 是否内嵌下拉改变
const onSelectIframeChange = () => {
	if (state.ruleForm.meta.isIframe) state.isLink = true;
	else state.isLink = false;
};

// TODO 添加验证规则
const menuRules = reactive<FormRules>({
	title:[{required:true,message:"请输入菜单名称",trigger: 'blur'}]
})

// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = async () => {
	if (state.dialog.type === 'add') {
		await menuApi
			.addMenu(state.ruleForm)
			.then((res) => {
				ElMessage.success(res.message);
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
			.then((res) => {
				ElMessage.success(res.message);
				closeDialog(); // 关闭弹窗
				emit('refresh');
			})
			.catch((e: any) => {
				ElMessage.error(e.message);
			});
		closeDialog(); // 关闭弹窗
	}
	// setBackEndControlRefreshRoutes() // 刷新菜单，未进行后端接口测试
};
// 页面加载时
onMounted(() => {
	getMenuData();
});

// 暴露变量
defineExpose({
	openDialog,
});
</script>
