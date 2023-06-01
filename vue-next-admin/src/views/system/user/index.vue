<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
				<el-input size="default" placeholder="请输入用户名称" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button size="default" type="success" class="ml10" @click="onOpenAddUser('add')" v-auth="'sys:user:add'">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增用户
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
				<el-table-column prop="id" label="用户ID" show-overflow-tooltip></el-table-column>
				<el-table-column prop="user_name" label="用户名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="email" label="邮箱" show-overflow-tooltip></el-table-column>
				<el-table-column prop="status" label="用户状态" show-overflow-tooltip>
					<template #default="scope">
						<el-switch
							v-model="scope.row.status"
							class="ml-2"
							style="--el-switch-on-color: #409eff; --el-switch-off-color: #ff4949"
							:active-value="1"
							:inactive-value="2"
							size="small"
							inline-prompt
							active-text="启用"
							inactive-text="禁用"
							width="45px"
							@click="changeStatus(scope.row)"
							v-auth="'sys:user:status'"
						/>
					</template>
				</el-table-column>
				<el-table-column prop="description" label="用户描述" show-overflow-tooltip></el-table-column>
				<el-table-column prop="created_at" label="创建时间" show-overflow-tooltip>
					<template #default="scope"> {{ dateStrFormat(scope.row.created_at) }}</template>
				</el-table-column>
				<el-table-column label="操作" width="100">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onOpenEditUser('edit', scope.row)" v-auth="'sys:user:edit'">修改</el-button>
						<el-button size="small" text type="primary" @click="onRowDel(scope.row)" v-auth="'sys:user:del'">删除</el-button>
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
		</el-card>
		<UserDialog ref="userDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useUserApi } from '/@/api/system/user';
import { RowUserType, SysUserState } from '/@/types/views';
import { dateStrFormat } from '/@/utils/formatTime';

// 引入组件
const UserDialog = defineAsyncComponent(() => import('./componet/dialog.vue'));

// 定义变量内容
const userApi = useUserApi();
const userDialogRef = ref();
const state = reactive<SysUserState>({
	tableData: {
		data: [],
		total: 0,
		loading: false,
		param: {
			page: 1,
			limit: 10,
		},
	},
});

const changeStatus = async (obj: any) => {
	await userApi
		.updateStatus(obj.id, obj.status)
		.then((res) => {
			getTableData();
			ElMessage.success(res.message);
		})
		.catch(() => {
			obj.status = 1;
		});
};

// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;

	await userApi.listUser(state.tableData.param).then((res) => {
		state.tableData.data = res.data.users;
		state.tableData.total = res.data.total;
	});

	setTimeout(() => {
		state.tableData.loading = false;
	}, 500);
};
// 打开新增用户弹窗
const onOpenAddUser = (type: string) => {
	userDialogRef.value.openDialog(type);
};
// 打开修改用户弹窗
const onOpenEditUser = (type: string, row: RowUserType) => {
	userDialogRef.value.openDialog(type, JSON.parse(JSON.stringify(row)));
};
// 删除用户
const onRowDel = (row: RowUserType) => {
	ElMessageBox.confirm(`此操作将永久删除账户名称：“${row.user_name}”，是否继续?`, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			userApi
				.deleteUser(row.id)
				.then((res) => {
					if (res.code === 200) {
						ElMessage.success('删除成功');
						getTableData();
					} else {
						ElMessage.error(res.message);
					}
				})
				.catch((e) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
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
.system-user-container {
	:deep(.el-card__body) {
		display: flex;
		flex-direction: column;
		flex: 1;
		overflow: auto;
		.el-table {
			flex: 1;
		}
	}
}
</style>
