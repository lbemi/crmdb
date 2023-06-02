<template>
	<div class="system-menu-container layout-pd">
		<el-card shadow="hover">
			<div class="system-menu-search mb15">
				<el-input size="default" placeholder="请输入API名称" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button size="default" type="success" class="ml10" @click="onOpenAddApi('add')">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增API
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%" row-key="path">
				<el-table-column label="ID" show-overflow-tooltip>
					<template #default="scope">
						<span class="ml10">{{ $t(scope.row.id) }}</span>
					</template>
				</el-table-column>
				<el-table-column label="分组" show-overflow-tooltip width="100">
					<template #default="scope">
						<el-tag class="ml-2" effect="plain">{{ scope.row.group }}</el-tag>
					</template>
				</el-table-column>
				<el-table-column label="描述" show-overflow-tooltip>
					<template #default="scope">
						<span class="ml10">{{ $t(scope.row.memo) }}</span>
					</template>
				</el-table-column>
				<el-table-column prop="path" label="请求路径" show-overflow-tooltip></el-table-column>

				<el-table-column label="请求方法" show-overflow-tooltip>
					<template #default="scope">
						<span v-if="scope.row.method === 'GET'" style="font-size: 13px" text="bold">
							<el-tag class="ml-2">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'POST'" style="font-size: 13px">
							<el-tag class="ml-2" type="success">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'DELETE'" style="font-size: 13px">
							<el-tag class="ml-2" type="danger">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'PUT'" style="font-size: 13px">
							<el-tag class="ml-2" type="warning">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else style="font-size: 13px">
							<el-tag class="ml-2">{{ scope.row.method }}</el-tag>
						</span>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="状态" width="80">
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
							@change="changeStatus(scope.row)"
						/>
					</template>
				</el-table-column>
				<el-table-column label="排序" show-overflow-tooltip width="80">
					<template #default="scope">
						{{ scope.row.sequence }}
					</template>
				</el-table-column>
				<el-table-column prop="created_at" label="创建时间" show-overflow-tooltip>
					<template #default="scope"> {{ dateStrFormat(scope.row.created_at) }}</template>
				</el-table-column>
				<el-table-column label="操作" show-overflow-tooltip width="140">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onOpenAddApi('add')">新增</el-button>
						<el-button size="small" text type="primary" @click="onOpenEditApi('edit', scope.row)">修改</el-button>
						<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
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
		<MenuDialog ref="menuDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="systemApi">
import { defineAsyncComponent, ref, onMounted, reactive } from 'vue';
import { RouteRecordRaw } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useMenuApi } from '/@/api/system/menu';

// import { setBackEndControlRefreshRoutes } from "/@/router/backEnd";

// 引入组件
const MenuDialog = defineAsyncComponent(() => import('./dialog.vue'));

// 定义变量内容
const menuApi = useMenuApi();
const menuDialogRef = ref();
const state = reactive({
	tableData: {
		data: [] as RouteRecordRaw[],
		loading: true,
		param: {
			page: 1,
			limit: 10,
			menuType: 3,
			isTree: false,
		},
		total: 0,
	},
});

// 获取路由数据，真实请从接口获取
const getTableData = async () => {
	state.tableData.loading = true;
	await menuApi
		.listMenu(state.tableData.param)
		.then((res) => {
			state.tableData.data = res.data.menus;
			state.tableData.total = res.data.total;
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	// state.tableData.data = routesList.value;
	setTimeout(() => {
		state.tableData.loading = false;
	}, 500);
};
// 打开新增菜单弹窗
const onOpenAddApi = (type: string) => {
	menuDialogRef.value.openDialog(type);
};
// 打开编辑菜单弹窗
const onOpenEditApi = (type: string, row: RouteRecordRaw) => {
	menuDialogRef.value.openDialog(type, row);
};
// 删除菜单或按钮
const onTabelRowDel = (row: any) => {
	ElMessageBox.confirm(`此操作将永久删除API：${row.name}, 是否继续?`, '警告', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(async () => {
			await menuApi.deleteMenu(row.id);
			ElMessage.success('删除成功');
			getTableData();
			//await setBackEndControlRefreshRoutes() // 刷新菜单，未进行后端接口测试
		})
		.catch(() => {});
};
const changeStatus = async (obj: any) => {
	let text = obj.status === 1 ? '启用' : '停用';
	ElMessageBox({
		closeOnClickModal: false,
		closeOnPressEscape: false,
		title: '警告',
		message: '确认要 ' + text + ' "' + obj.name + '" API吗?',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
	})
		.then(async () => {
			return await menuApi.updateMenuStatu(obj.id, obj.status);
		})
		.then(() => {
			ElMessage.success(text + '成功');
		})
		.catch(() => {
			obj.status = obj.status === 1 ? 2 : 1;
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
