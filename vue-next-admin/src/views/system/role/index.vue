<template>
	<div class="system-role-container layout-padding">
		<div class="system-role-padding layout-padding-auto layout-padding-view">
			<div class="system-user-search mb15">
				<el-input v-model="state.tableData.param.search" size="default" placeholder="请输入角色名称" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button size="default" type="success" class="ml10" @click="onOpenAddRole('add')" v-auth="'sys:role:add'">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增角色
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
				<el-table-column prop="id" label="角色ID" width="130" />
				<el-table-column prop="name" label="角色名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="sequence" label="排序" show-overflow-tooltip></el-table-column>
				<el-table-column prop="status" label="状态" width="80">
					<template #default="scope">
						<el-switch
							v-model="scope.row.status"
							v-auth="'sys:role:status'"
							class="ml-2"
							style="--el-switch-on-color: #409eff; --el-switch-off-color: #ff4949"
							:active-value="1"
							:inactive-value="2"
							size="small"
							inline-prompt
							active-text="启用"
							inactive-text="禁用"
							width="45px"
							@change="changeStatus(scope.row)"
						/>
					</template>
				</el-table-column>
				<el-table-column prop="memo" label="角色描述" show-overflow-tooltip></el-table-column>
				<el-table-column prop="created_at" label="创建时间" show-overflow-tooltip>
					<template #default="scope"> {{ dateStrFormat(scope.row.created_at) }}</template>
				</el-table-column>
				<el-table-column label="操作" width="140">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onOpenEditRole('edit', scope.row)" v-auth="'sys:role:edit'">修改</el-button>
						<el-button size="small" text type="primary" @click="onOpenAuthRole(scope.row)" v-auth="'sys:role:set'">授权</el-button>
						<el-button size="small" text type="primary" @click="onRowDel(scope.row)" v-auth="'sys:role:del'">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
			<el-pagination
				@size-change="onHandleSizeChange"
				@current-change="onHandleCurrentChange"
				class="mt15"
				:pager-count="5"
				:page-sizes="[10, 20, 30]"
				v-model:current-page="state.tableData.param.page"
				background
				v-model:page-size="state.tableData.param.limit"
				layout="total, sizes, prev, pager, next, jumper"
				:total="state.tableData.total"
			>
			</el-pagination>
		</div>
		<RoleDialog ref="roleDialogRef" @refresh="getTableData()" />
		<RoleAuthDialog ref="roleAuthDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="systemRole">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useRoleApi } from '/@/api/system/role';
import { SysRoleState, RoleType } from '/@/types/views';
import { dateStrFormat } from '/@/utils/formatTime';

// 引入组件
const RoleDialog = defineAsyncComponent(() => import('/@/views/system/role/component/dialog.vue'));
const RoleAuthDialog = defineAsyncComponent(() => import('/@/views/system/role/component/authDialog.vue'));
// 定义变量内容
const roleApi = useRoleApi();
const roleDialogRef = ref();
const roleAuthDialogRef = ref();
const state = reactive<SysRoleState>({
	tableData: {
		data: [],
		total: 0,
		loading: false,
		param: {
			search: '',
			page: 1,
			limit: 10,
		},
	},
});
// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;
	await roleApi.listRole(state.tableData.param).then((res: any) => {
		state.tableData.data = res.data.data;
		state.tableData.total = res.data.total;
	});
	setTimeout(() => {
		state.tableData.loading = false;
	}, 500);
};
// 打开新增角色弹窗
const onOpenAddRole = (type: string) => {
	roleDialogRef.value.openDialog(type);
};
// 打开修改角色弹窗
const onOpenEditRole = (type: string, row: Object) => {
	roleDialogRef.value.openDialog(type, row);
};
const onOpenAuthRole = (row: object) => {
	roleAuthDialogRef.value.openAuthDialog(row);
};
//修改状态
const changeStatus = async (obj: any) => {
	let text = obj.status === 1 ? '启用' : '停用';
	ElMessageBox({
		closeOnClickModal: false,
		closeOnPressEscape: false,
		title: '警告',
		message: '确认要 ' + text + ' "' + obj.name + '"角色吗?',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
	}).then(async () => {
		await roleApi
			.updateRoleStatus(obj.id, obj.status)
			.then(() => {
				ElMessage.success(text + '成功');
			})
			.catch((e) => {
				obj.status = obj.status === 1 ? 2 : 1;
				ElMessage.error(e.message);
			});
	});
};
// 删除角色
const onRowDel = (row: RoleType) => {
	ElMessageBox.confirm(`此操作将永久删除角色名称：“${row.name}”，是否继续?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(async () => {
			await roleApi.deleteRole(row.id);
			getTableData();
			ElMessage.success('删除成功');
		})
		.catch(() => {});
};
// 分页改变
const onHandleSizeChange = (val: number) => {
	state.tableData.param.limit = val;
	getTableData();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
	state.tableData.param.page = val;
	getTableData();
};
// 页面加载时
onMounted(() => {
	getTableData();
});
</script>

<style scoped lang="scss">
.system-role-container {
	.system-role-padding {
		padding: 15px;
		.el-table {
			flex: 1;
		}
	}
}
</style>
