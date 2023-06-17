<template>
	<div class="system-user-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-user-search mb15">
				<el-input size="default" placeholder="请输入用户名称" style="max-width: 180px; margin-right: 10px" v-model="state.searchName" clearable>
				</el-input>
				用户状态：
				<el-select v-model.number="state.searchStatus" size="default" style="width: 100px">
					<el-option v-for="item in status" :value="item.value" :label="item.label"> </el-option>
				</el-select>
				<el-button size="default" type="primary" class="ml10" @click="getTableData">
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
				<el-table-column label="操作" width="160">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="handleSetRole(scope.row)" v-auth="'sys:user:set'">分配角色</el-button>
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
		<UserSetRole
			v-model:visible="setRole.visible"
			:defaultCheckedRoles="setRole.defaultCheckedRoles"
			:title="setRole.title"
			v-model:userID="setRole.userID"
			v-model:roleList="setRole.roleList"
			v-if="setRole.visible"
		/>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useUserApi } from '@/api/system/user';
import { RoleType, RowUserType, SysUserState } from '@/types/views';
import { dateStrFormat } from '@/utils/formatTime';

// 引入组件
const UserDialog = defineAsyncComponent(() => import('./componet/dialog.vue'));
const UserSetRole = defineAsyncComponent(() => import('./componet/setRole.vue'));
type query = {
	page: number;
	limit: number;
	name?: string;
	status?: number;
};

// 定义变量内容
const userApi = useUserApi();
const userDialogRef = ref();
const state = reactive<SysUserState>({
	tableData: {
		data: [],
		total: 0,
		loading: false,
		param: <query>{
			page: 1,
			limit: 10,
		},
	},
	searchName: '',
	searchStatus: 0,
});

const setRole = reactive({
	visible: false,
	title: '分配角色',
	userID: 0,
	roleList: [],
	defaultCheckedRoles: [] as Array<number>,
});

const handleSetRole = async (user: RowUserType) => {
	setRole.title = `为【${user.user_name}】分配角色：`;
	setRole.userID = user.id;
	setRole.defaultCheckedRoles = [];

	await userApi.getUserRole(user.id).then((res) => {
		const roleList: Array<RoleType> = res.data;
		if (roleList !== null) {
			for (let i = 0; i < roleList.length; i++) {
				if (roleList[i].children !== null) {
					for (let j = 0; j < roleList[i].children.length; j++) {
						setRole.defaultCheckedRoles.push(roleList[i].children[j].id);
					}
				}
				setRole.defaultCheckedRoles.push(roleList[i].id);
			}
		}
	});
	setRole.visible = true;
};

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
const status = [
	{
		label: '所有状态',
		value: 0,
	},
	{
		label: '正常',
		value: 1,
	},
	{
		label: '禁用',
		value: 2,
	},
];
// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;
	if (state.searchName != '') {
		state.tableData.param.name = state.searchName;
	} else {
		delete state.tableData.param.name;
	}

	if (state.searchStatus != 0) {
		state.tableData.param.status = state.searchStatus;
	} else {
		delete state.tableData.param.status;
	}

	await userApi.listUser(state.tableData.param).then((res) => {
		state.tableData.data = res.data.users;
		state.tableData.total = res.data.total;
	});
	state.tableData.loading = false;
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
