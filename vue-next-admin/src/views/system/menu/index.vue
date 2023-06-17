<template>
	<div class="system-menu-container layout-pd">
		<el-card shadow="hover">
			<div class="system-menu-search mb15">
				<el-input size="default" placeholder="请输入菜单描述" style="max-width: 180px; margin-right: 10px" v-model="state.searchName" clearable>
				</el-input>
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
				<el-button size="default" type="success" class="ml10" @click="onOpenAddMenu('add')" v-auth="'sys:menu:add'">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增菜单
				</el-button>
				<el-button size="default" type="success" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="state.tableData.data"
				v-loading="state.tableData.loading"
				style="width: 100%"
				row-key="id"
				:tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
			>
				<el-table-column label="菜单描述" show-overflow-tooltip>
					<template #default="scope">
						<SvgIcon :name="scope.row.meta.icon" />
						<span class="ml10">{{ $t(scope.row.memo) }}</span>
					</template>
				</el-table-column>
				<el-table-column label="类型" show-overflow-tooltip width="90">
					<template #default="scope">
						<!-- <el-tag type="success" size="small">{{ scope.row.xx }}菜单</el-tag> -->
						<span v-if="scope.row.menuType == 1" style="font-size: 13px">
							<el-tag class="ml-2" type="success" effect="light">菜单</el-tag>
						</span>
						<span v-else style="font-size: 13px">
							<el-tag class="ml-2" type="danger" effect="light">按钮</el-tag>
						</span>
					</template>
				</el-table-column>
				<el-table-column prop="path" label="路由路径" show-overflow-tooltip></el-table-column>
				<el-table-column label="组件路径" show-overflow-tooltip>
					<template #default="scope">
						<span>{{ scope.row.component }}</span>
					</template>
				</el-table-column>
				<el-table-column label="权限标识" show-overflow-tooltip>
					<template #default="scope">
						<span>{{ scope.row.code }}</span>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="状态" width="80">
					<template #default="scope">
						<el-switch
							v-model="scope.row.status"
							v-auth="'sys:menu:status'"
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
						<el-button size="small" text type="primary" @click="onOpenEditMenu('edit', scope.row)" v-auth="'sys:menu:edit'">编辑</el-button>
						<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)" v-auth="'sys:menu:del'">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
		<MenuDialog ref="menuDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="systemMenu">
import { defineAsyncComponent, ref, onMounted, reactive } from 'vue';
import { RouteRecordRaw } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useMenuApi } from '@/api/system/menu';
import { dateStrFormat } from '@/utils/formatTime';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';

// import { setBackEndControlRefreshRoutes } from "@/router/backEnd";

// 引入组件
const MenuDialog = defineAsyncComponent(() => import('@/views/system/menu/dialog.vue'));

// 定义变量内容
type query = {
	menuType: string;
	memo?: string;
	status?: number;
};
const menuApi = useMenuApi();
const route = useRoute();
const menuDialogRef = ref();
const state = reactive({
	tableData: {
		data: [] as RouteRecordRaw[],
		loading: true,
	},
	params: <query>{
		menuType: '1', //获取菜单,1为目录,2为菜单
	},
	searchName: '',
	searchStatus: 0,
});

// 获取路由数据，真实请从接口获取
const getTableData = async () => {
	state.tableData.loading = true;
	if (state.searchName != '') {
		state.params.memo = state.searchName;
	} else {
		delete state.params.memo;
	}

	if (state.searchStatus != 0) {
		state.params.status = state.searchStatus;
	} else {
		delete state.params.status;
	}
	await menuApi
		.listMenu(state.params)
		.then((res) => {
			state.tableData.data = res.data.menus;
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	state.tableData.loading = false;
};
// 打开新增菜单弹窗
const onOpenAddMenu = (type: string) => {
	menuDialogRef.value.openDialog(type);
};
// 打开编辑菜单弹窗
const onOpenEditMenu = (type: string, row: RouteRecordRaw) => {
	menuDialogRef.value.openDialog(type, row);
};
// 删除菜单或按钮
const onTabelRowDel = (row: any) => {
	let text = row.menuType === 1 ? '菜单' : '按钮';
	ElMessageBox.confirm(`此操作将永久删除${text}：${row.name}, 是否继续?`, '警告', {
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
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
//修改状态
const changeStatus = async (obj: any) => {
	let text = obj.status === 1 ? '启用' : '停用';
	ElMessageBox({
		closeOnClickModal: false,
		closeOnPressEscape: false,
		title: '警告',
		message: '确认要 ' + text + ' "' + obj.name + '"菜单吗?',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
	})
		.then(async () => {
			await menuApi.updateMenuStatu(obj.id, obj.status).catch((e) => {
				ElMessage.error(e.message);
			});
		})
		.then(() => {
			ElMessage.success(text + '成功');
		})
		.catch(() => {
			obj.status = obj.status === 1 ? 2 : 1;
		});
};
// 页面加载时
onMounted(() => {
	getTableData();
});
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
</script>
