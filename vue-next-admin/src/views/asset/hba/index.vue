<template>
	<div class="system-dept-container layout-padding">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="system-dept-search mb15">
				<el-input size="default" placeholder="请输入主机ip" style="max-width: 180px"> </el-input>
				<el-button size="default" type="primary" class="ml10">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button size="default" type="success" class="ml10" @click="onOpenAddHba('add')">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增规则
				</el-button>
				<el-button type="warning" size="default" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="state.tableData.hbas"
				v-loading="state.tableData.loading"
				style="width: 100%"
				row-key="id"
				default-expand-all
				:tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
			>
				<el-table-column prop="id" label="ID" show-overflow-tooltip> </el-table-column>
				<el-table-column prop="rule_name" label="规则名称" show-overflow-tooltip> </el-table-column>

				<!-- <el-table-column prop="account_id" label="账户ID" show-overflow-tooltip></el-table-column>
				<el-table-column prop="resource_id" label="主机ID" show-overflow-tooltip></el-table-column> -->

				<el-table-column label="创建时间" show-overflow-tooltip>
					<template #default="scope">
						{{ dateStrFormat(scope.row.created_at) }}
					</template>
				</el-table-column>
				<el-table-column label="更新时间" show-overflow-tooltip>
					<template #default="scope">
						{{ dateStrFormat(scope.row.updated_at) }}
					</template>
				</el-table-column>
				<el-table-column label="操作" show-overflow-tooltip fixed="right">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onOpenEditDept('detail', scope.row)">详情</el-button>
						<el-button size="small" text type="primary" @click="onOpenEditDept('edit', scope.row)">修改</el-button>
						<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<Pagination :total="state.tableData.total" @handlePageChange="handlePageChange" />
		</el-card>
		<HbaDialog ref="hbaDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="hba">
import { defineAsyncComponent, ref, reactive, onMounted } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { DeptTreeType } from '@/types/views';
import { Group } from '@/types/asset/group';
import { PageInfo } from '@/types/kubernetes/common';
import { dateStrFormat } from '../../../utils/formatTime';
import { useHBAApi } from '@/api/asset/hostBindAccount';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';

// 引入组件
const HbaDialog = defineAsyncComponent(() => import('./dialog.vue'));
const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));

// 定义变量内容
const hbaApi = useHBAApi();
const route = useRoute();
const hbaDialogRef = ref();
const state = reactive({
	defaultProps: {
		children: 'children',
		label: 'name',
	},
	groupName: '',
	groupIds: '',
	tableData: {
		hbas: [],
		hosts: [],
		accounts: [],
		total: 0,
		loading: false,
		param: {
			page: 1,
			limit: 10,
		} as ParamType,
	},
});

// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;
	state.tableData.hbas = [];
	await hbaApi
		.lisHba(state.tableData.param)
		.then((res: any) => {
			state.tableData.hbas = res.data.data;
			state.tableData.total = res.data.total;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	state.tableData.loading = false;
};

// 分页点击事件
const handlePageChange = (pageInfo: PageInfo) => {
	state.tableData.param.page = pageInfo.page;
	state.tableData.param.limit = pageInfo.limit;
	getTableData();
};

const getAllNode = (group: Group) => {
	if (group.children) {
		group.children.forEach((g) => {
			getAllNode(g);
		});
	}
	if (state.groupIds.length > 0) {
		state.groupIds = state.groupIds + ',' + group.id;
	} else {
		state.groupIds = state.groupIds + group.id;
	}
};

// 打开新增菜单弹窗
const onOpenAddHba = (type: string) => {
	hbaDialogRef.value.openDialog(type, state.tableData.hosts, state.tableData.accounts);
	// deptDialogRef.value.openDialog(type, state.tableData.groups);
};

// 打开编辑菜单弹窗
const onOpenEditDept = (type: string, row: HostBingAccounts) => {
	hbaDialogRef.value.openDialog(type, row);
};

// 删除当前行
const onTabelRowDel = (row: DeptTreeType) => {
	ElMessageBox.confirm(`此操作将永久删除部门：${row.deptName}, 是否继续?`, '提示', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			getTableData();
			ElMessage.success('删除成功');
		})
		.catch(() => {});
};
const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
// 页面加载时
onMounted(() => {
	getTableData();
});
</script>
