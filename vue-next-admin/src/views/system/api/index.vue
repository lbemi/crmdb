<template>
	<div class="system-menu-container layout-pd">
		<el-card shadow="hover">
			<div class="system-menu-search mb15">
				<el-input size="default" placeholder="请输入API名称" v-model="state.searchName" clearable style="max-width: 180px; margin-right: 10px">
				</el-input>
				分组：
				<el-input size="default" placeholder="请输入API分组" v-model="state.searchGroup" clearable style="max-width: 180px; margin-right: 10px">
				</el-input>
				<!-- 分组2：
				<el-select v-model.number="state.searchGroup" size="default" style="width: 100px">
					<el-option v-for="item in state.groupList" :value="item" :label="item"> </el-option>
				</el-select> -->
				类型：
				<el-select v-model.number="state.searchType" size="default" style="width: 100px">
					<el-option v-for="item in type" :value="item.value" :label="item.label"> </el-option>
				</el-select>
				状态：
				<el-select v-model.number="state.searchStatus" size="default" style="width: 100px">
					<el-option v-for="item in status" :value="item.value" :label="item.label"> </el-option>
				</el-select>
				<el-button size="default" type="primary" class="ml10" @click="getTableData">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button v-auth="'sys:menu:add'" size="default" type="success" class="ml10" @click="onOpenAddApi('add')">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增API
				</el-button>
				<el-button type="success" size="default" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%" row-key="id">
				<el-table-column label="ID" show-overflow-tooltip width="150">
					<template #default="scope">
						<span class="ml10">{{ $t(scope.row.id) }}</span>
					</template>
				</el-table-column>
				<el-table-column
					label="分组"
					show-overflow-tooltip
					width="120"
					:filters="[
						{ text: 'Home', value: 'Home' },
						{ text: 'Office', value: 'Office' },
					]"
					:filter-method="filterGroup"
					filter-placement="bottom-end"
				>
					<template #default="scope">
						<el-tag size="small" class="ml-2" effect="plain">{{ scope.row.group }}</el-tag>
					</template>
				</el-table-column>
				<el-table-column label="描述" show-overflow-tooltip>
					<template #default="scope">
						<span class="ml10">{{ $t(scope.row.memo) }}</span>
					</template>
				</el-table-column>
				<el-table-column label="类型" show-overflow-tooltip width="90">
					<template #default="scope">
						<!-- <el-tag type="success" size="small">{{ scope.row.xx }}菜单</el-tag> -->
						<span v-if="scope.row.menuType == 3" style="font-size: 13px">
							<el-tag size="small" class="ml-2" type="success" effect="light">API</el-tag>
						</span>
						<span v-else style="font-size: 13px">
							<el-tag size="small" class="ml-2" type="danger" effect="light">按钮</el-tag>
						</span>
					</template>
				</el-table-column>
				<el-table-column prop="path" label="请求路径" show-overflow-tooltip></el-table-column>

				<el-table-column label="请求方法" show-overflow-tooltip>
					<template #default="scope">
						<span v-if="scope.row.method === 'GET'" style="font-size: 13px" text="bold">
							<el-tag size="small" class="ml-2">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'POST'" style="font-size: 13px">
							<el-tag size="small" class="ml-2" type="success">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'DELETE'" style="font-size: 13px">
							<el-tag size="small" class="ml-2" type="danger">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'PUT'" style="font-size: 13px">
							<el-tag size="small" class="ml-2" type="warning">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else-if="scope.row.method === 'PATCH'" style="font-size: 13px">
							<el-tag size="small" class="ml-2">{{ scope.row.method }}</el-tag>
						</span>
						<span v-else style="font-size: 13px">
							<el-tag size="small" class="ml-2">{{ scope.row.method }}</el-tag>
						</span>
					</template>
				</el-table-column>
				<el-table-column prop="code" label="权限标识" show-overflow-tooltip></el-table-column>
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
				<el-table-column prop="created_at" label="创建时间" sortable show-overflow-tooltip>
					<template #default="scope"> {{ dateStrFormat(scope.row.created_at) }}</template>
				</el-table-column>
				<el-table-column label="操作" show-overflow-tooltip width="140">
					<template #default="scope">
						<el-button v-auth="'sys:menu:add'" size="small" text type="primary" @click="onOpenAddApi('add')">新增</el-button>
						<el-button v-auth="'sys:menu:edit'" size="small" text type="primary" @click="onOpenEditApi('edit', scope.row)">修改</el-button>
						<el-button v-auth="'sys:menu:del'" size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
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
import { RouteRecordRaw, useRoute } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useMenuApi } from '/@/api/system/menu';
import { dateStrFormat } from '/@/utils/formatTime';
import mittBus from '/@/utils/mitt';

// import { setBackEndControlRefreshRoutes } from "/@/router/backEnd";

// 引入组件
const MenuDialog = defineAsyncComponent(() => import('./dialog.vue'));

// 定义变量内容
type query = {
	page: number;
	limit: number;
	menuType: string;
	isTree: boolean;
	memo?: string;
	status?: number;
	group?: string;
};
const route = useRoute();
const menuApi = useMenuApi();
const menuDialogRef = ref();
const state = reactive({
	tableData: {
		data: [] as RouteRecordRaw[],
		loading: true,
		param: <query>{
			page: 1,
			limit: 10,
			menuType: '2,3',
			isTree: false,
		},
		total: 0,
	},
	searchName: '',
	searchStatus: 0,
	searchType: 0,
	searchGroup: '',
	groupList: [''],
});
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
// 获取路由数据，真实请从接口获取
const getTableData = async () => {
	state.tableData.loading = true;
	if (state.searchName != '') {
		state.tableData.param.memo = state.searchName;
	} else {
		delete state.tableData.param.memo;
	}
	if (state.searchGroup != '') {
		state.tableData.param.group = state.searchGroup;
	} else {
		delete state.tableData.param.group;
	}

	if (state.searchStatus != 0) {
		state.tableData.param.status = state.searchStatus;
	} else {
		delete state.tableData.param.status;
	}

	if (state.searchType != 0) {
		state.tableData.param.menuType = state.searchType + '';
	} else {
		state.tableData.param.menuType = '2,3';
	}

	await menuApi
		.listMenu(state.tableData.param)
		.then((res) => {
			state.tableData.data = res.data.menus;
			state.tableData.total = res.data.total;
			filterGroup();
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	state.tableData.loading = false;
};
// 打开新增菜单弹窗
const onOpenAddApi = (type: string) => {
	menuDialogRef.value.openDialog(type);
};
// 打开编辑菜单弹窗
const onOpenEditApi = (type: string, row: RouteRecordRaw) => {
	menuDialogRef.value.openDialog(type, row);
};
// TODO
const filterGroup = () => {
	const groupList = [''];
	state.tableData.data.forEach((item: any) => {
		if (groupList.indexOf(item.group) < 0) {
			groupList.push(item.group);
		}
	});
	state.groupList = groupList;
};
// 删除菜单或按钮
const onTabelRowDel = (row: any) => {
	ElMessageBox.confirm(`此操作将永久删除API：${row.name}, 是否继续?`, '警告', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
		buttonSize: 'small',
	})
		.then(async () => {
			await menuApi
				.deleteMenu(row.id)
				.then((res) => {
					ElMessage.success('删除成功');
					getTableData();
				})
				.catch((e) => {
					ElMessage.error(e.message);
				});
			//await setBackEndControlRefreshRoutes() // 刷新菜单，未进行后端接口测试
		})
		.catch(() => {
			ElMessage.info('取消');
		});
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
		buttonSize: 'small',
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
const type = [
	{
		label: '所有类型',
		value: 0,
	},
	{
		label: '按钮',
		value: 2,
	},
	{
		label: 'API',
		value: 3,
	},
];
// 页面加载时
onMounted(() => {
	getTableData();
});
</script>
