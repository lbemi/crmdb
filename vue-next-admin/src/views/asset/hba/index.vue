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
				<el-button size="default" type="success" class="ml10" @click="onOpenAddDept('add')">
					<el-icon>
						<ele-FolderAdd />
					</el-icon>
					新增规则
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
						<el-button size="small" text type="primary" @click="onOpenAddDept('add')">详情</el-button>
						<el-button size="small" text type="primary" @click="onOpenEditDept('edit', scope.row)">修改</el-button>
						<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<Pagination :total="state.tableData.total" @handlePageChange="handlePageChange" />
		</el-card>
		<DeptDialog ref="deptDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="hba">
import { defineAsyncComponent, ref, reactive, onMounted } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { DeptTreeType } from '@/types/views';
import { useHostApi } from '@/api/asset/hosts';
import { HostState } from '@/types/asset/hosts';
import { useGroupApi } from '@/api/asset/group';
import { Group } from '@/types/asset/group';
import { PageInfo } from '@/types/kubernetes/common';
import { dateStrFormat } from '../../../utils/formatTime';
import { useHBAApi } from '@/api/asset/hostBindAccount';
import { useAccountApi } from '@/api/asset/account';

// 引入组件
const DeptDialog = defineAsyncComponent(() => import('./dialog.vue'));
const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));

const hostApi = useHostApi();
const accountApi = useAccountApi();
const hbaApi = useHBAApi();
const tree = ref('');
// 定义变量内容
const deptDialogRef = ref();
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

const getHosts = async () => {
	state.tableData.loading = true;
	await hostApi
		.lisHost()
		.then((res: any) => {
			state.tableData.hosts = res.data.data;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	state.tableData.loading = false;
};
const getAccounts = async () => {
	state.tableData.loading = true;
	await accountApi
		.listAccount()
		.then((res: any) => {
			state.tableData.accounts = res.data.data;
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
const onOpenAddDept = (type: string) => {
	// deptDialogRef.value.openDialog(type, state.tableData.groups);
};
// 打开编辑菜单弹窗
const onOpenEditDept = (type: string, row: DeptTreeType) => {
	deptDialogRef.value.openDialog(type, row);
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
// 页面加载时
onMounted(() => {
	getTableData();
	getAccounts();
	getHosts();
});
</script>
