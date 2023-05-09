<template>
	<div class="system-role-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="500px">
			<el-form ref="roleAuthDialogFormRef" :model="state" size="default" label-width="90px">
				<el-row>
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-tabs v-model="state.activeName" tab-position="left">
							<el-tab-pane label="菜单权限" name="first">
								<el-checkbox v-model="state.menuExpand" @change="handleCheckedTreeExpand($event, 'menu')">展开/折叠 </el-checkbox>
								<el-checkbox v-model="state.menuNodeAll" @change="handleCheckedTreeNodeAll($event, 'menu')">全选/全不选 </el-checkbox>
								<el-tree
									class="tree-border"
									:data="state.menuData"
									show-checkbox
									ref="menuRef"
									node-key="id"
									empty-text="加载中，请稍后"
									:props="{
										label: 'memo',
										children: 'children',
									}"
									:default-checked-keys="state.menuCheckedKeys"
								></el-tree>
							</el-tab-pane>
							<el-tab-pane label="API权限" name="second">
								<el-checkbox v-model="state.apiExpand" @change="handleCheckedTreeExpand($event, 'api')">展开/折叠 </el-checkbox>
								<el-checkbox v-model="state.apiNodeAll" @change="handleCheckedTreeNodeAll($event, 'api')">全选/全不选 </el-checkbox>
								<el-tree
									class="tree-border"
									:data="state.apiData"
									show-checkbox
									ref="apiRef"
									node-key="id"
									empty-text="加载中，请稍后"
									:props="{
										children: 'children',
										label: 'memo',
									}"
									:default-checked-keys="state.apiCheckedKeys"
								></el-tree>
							</el-tab-pane>
						</el-tabs>
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

<script setup lang="ts" name="systemRoleAuthDialog">
import { reactive, ref } from 'vue';
import { useMenuApi } from '/@/api/system/menu';
import { unref } from 'vue-demi';
import { MenuType, RoleType } from '/@/types/views';
import { ElLoading, ElMessage } from 'element-plus';
import { useRoleApi } from '/@/api/system/role';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const menuApi = useMenuApi();
const roleApi = useRoleApi();
const menuRef = ref<HTMLElement | null>(null);
const apiRef = ref<HTMLElement | null>(null);
const roleAuthDialogFormRef = ref<HTMLElement | null>(null);
const state = reactive({
	loading: false,
	id: 0,
	apiExpand: false,
	apiNodeAll: false,
	menuExpand: false,
	menuNodeAll: false,
	apiAndMenuIDs: [],
	activeName: 'first',
	menuCheckedKeys: [],
	apiCheckedKeys: [],
	menuData: [],
	apiData: [],
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openAuthDialog = (row: RoleType) => {
	state.dialog.title = row.name + '角色-授权';
	state.dialog.submitTxt = '修 改';
	state.dialog.isShowDialog = true;
	state.id = row.id;
	state.apiCheckedKeys = [];
	state.menuCheckedKeys = [];
	state.activeName = 'first';

	listMenuAndButton();
	listApi();
	getRoleMenus();
	getRoleApis();
};

const listMenuAndButton = async () => {
	await menuApi.listMenu({ menuType: '1,2' }).then((res) => {
		state.menuData = res.data.menus;
	});
};

const getRoleMenus = async () => {
	roleApi.getRoleMenu(state.id, { menuType: '1,2' }).then((res) => {
		if (res.data) {
			res.data.forEach((item) => {
				state.menuCheckedKeys.push(item.id);
				if (item.children) {
					item.children.forEach((res) => {
						state.menuCheckedKeys.push(res.id);
					});
				}
			});
		}
	});
	// console.log("----",state.menuCheckedKeys)
};
const getRoleApis = async () => {
	roleApi.getRoleMenu(state.id, { menuType: '3' }).then((res) => {
		if (res.data) {
			res.data.forEach((item) => {
				state.apiCheckedKeys.push(item.id);
				if (item.children) {
					item.children.forEach((res) => {
						state.apiCheckedKeys.push(res.id);
					});
				}
			});
		}
	});
};
const listApi = async () => {
	await menuApi.listMenu({ menuType: '3', isTree: false }).then((res) => {
		state.apiData = buildApiTree(res.data.menus);
	});
};

const buildApiTree = (apis: MenuType[]) => {
	const apiObj: any = new Object();
	apis &&
		apis.map((item) => {
			if (Object.prototype.hasOwnProperty.call(apiObj, item.group)) {
				apiObj[item.group].push(item);
			} else {
				Object.assign(apiObj, { [item.group]: [item] });
			}
		});
	const apiTree = [];
	for (const key in apiObj) {
		const treeNode = {
			id: key,
			memo: key + '组', //描述
			children: apiObj[key],
		};
		apiTree.push(treeNode);
	}
	return apiTree;
};

// 树权限（全选/全不选）
const handleCheckedTreeNodeAll = (value: any, type: any) => {
	if (type == 'menu') {
		const formWrap = unref(menuRef) as any;
		formWrap.setCheckedNodes(value ? state.menuData : []);
	} else if (type == 'api') {
		const formWrap = unref(apiRef) as any;
		formWrap.setCheckedNodes(value ? state.apiData : []);
	}
};
// 树权限（展开/折叠）
const handleCheckedTreeExpand = (value: any, type: any) => {
	if (type === 'menu') {
		let treeList = state.menuData;
		for (let i = 0; i < treeList.length; i++) {
			const formWrap = unref(menuRef) as any;
			formWrap.store.nodesMap[treeList[i].id].expanded = value;
		}
	} else if (type === 'api') {
		let treeList = state.apiData;
		for (let i = 0; i < treeList.length; i++) {
			const formWrap = unref(apiRef) as any;
			formWrap.store.nodesMap[treeList[i].id].expanded = value;
		}
	}
};

// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = () => {
	const loading = ElLoading.service({
		lock: true,
		text: 'Loading',
		background: 'rgba(0, 0, 0, 0.7)',
	});
	getAllCheckIds();
	roleApi
		.setRoleAuth(state.id, { menuIDS: state.apiAndMenuIDs })
		.then(() => {
			ElMessage.success('授权成功');
			loading.close();
			closeDialog();
			emit('refresh');
		})
		.catch((res) => {
			ElMessage.error(res);
		});
};

// 所有菜单节点数据
const getMenuAllCheckedKeys = () => {
	const formWrap = unref(menuRef) as any;
	if (!formWrap) return;
	// 目前被选中的菜单节点
	let checkedKeys = formWrap.getCheckedKeys();
	// 半选中的菜单节点
	let halfCheckedKeys = formWrap.getHalfCheckedKeys();
	checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys);
	state.apiAndMenuIDs.push(...checkedKeys);
	return checkedKeys;
};
// 所有菜单节点数据
const getApiAllCheckedKeys = () => {
	const formWrap = unref(apiRef) as any;
	if (!formWrap) return;
	let checkArr: any[] = formWrap.getCheckedNodes(true);
	let apiIds: any[] = [];
	checkArr.forEach((res) => {
		apiIds.push(res.id);
	});
	state.apiAndMenuIDs.push(...apiIds);
	return apiIds;
};
const getAllCheckIds = () => {
	state.apiAndMenuIDs = [];
	getMenuAllCheckedKeys();
	getApiAllCheckedKeys();
	console.log('提交的权限列表', state.apiAndMenuIDs);
};

// 暴露变量
defineExpose({
	openAuthDialog,
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
