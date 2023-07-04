<template>
	<div class="system-dept-container layout-padding">
		<el-row gutter="24">
			<el-col :span="4" :xs="24">
				<el-card shadow="always">
					<div class="head-container">
						<el-input v-model="state.deptName" placeholder="请输入部门名称" clearable prefix-icon="el-icon-search" style="margin-bottom: 20px" />
					</div>
					<div class="head-container">
						<el-tree
							:data="state.deptOptions"
							:props="state.defaultProps"
							node-key="deptId"
							:expand-on-click-node="false"
							:filter-node-method="filterNode"
							ref="tree"
							default-expand-all
							@node-click="handleNodeClick"
						/>
					</div>
				</el-card>
			</el-col>
			<el-col :span="20" :xs="24">
				<el-card shadow="hover" class="layout-padding-auto">
					<div class="system-dept-search mb15">
						<el-input size="default" placeholder="请输入部门名称" style="max-width: 180px"> </el-input>
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
							新增部门
						</el-button>
					</div>
					<el-table
						:data="state.tableData.data"
						v-loading="state.tableData.loading"
						style="width: 100%"
						row-key="id"
						default-expand-all
						:tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
					>
						<el-table-column prop="ip" label="主机IP" show-overflow-tooltip> </el-table-column>
						<el-table-column prop="remark" label="描述" show-overflow-tooltip> </el-table-column>
						<el-table-column label="排序" show-overflow-tooltip width="80">
							<template #default="scope">
								{{ scope.$index }}
							</template>
						</el-table-column>
						<el-table-column prop="status" label="是否禁用" show-overflow-tooltip>
							<template #default="scope">
								<el-tag type="success" v-if="scope.row.status">启用</el-tag>
								<el-tag type="info" v-else>禁用</el-tag>
							</template>
						</el-table-column>
						<el-table-column prop="port" label="端口" show-overflow-tooltip></el-table-column>
						<el-table-column label="标签" show-overflow-tooltip>
							<template #default="scope">
								<div v-for="(tag, index) in scope.row.labels" :key="index" class="tag">
									<el-tag effect="plain" size="small">{{ tag }}</el-tag>
								</div>
							</template>
						</el-table-column>
						<el-table-column label="SSH" show-overflow-tooltip>
							<template #default="scope">
								<el-tag type="success" v-if="scope.row.enable_ssh">启用</el-tag>
								<el-tag type="info" v-else>禁用</el-tag>
							</template>
						</el-table-column>
						<el-table-column prop="created_at" label="创建时间" show-overflow-tooltip></el-table-column>
						<el-table-column label="操作" show-overflow-tooltip width="140">
							<template #default="scope">
								<el-button size="small" text type="primary" @click="onOpenAddDept('add')">新增</el-button>
								<el-button size="small" text type="primary" @click="onOpenEditDept('edit', scope.row)">修改</el-button>
								<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-card>
			</el-col>
		</el-row>

		<DeptDialog ref="deptDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="hosts">
import { defineAsyncComponent, ref, reactive, onMounted } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { DeptTreeType } from '@/types/views';
import { useHostApi } from '@/api/asset/hosts';
import { HostState } from '@/types/asset/hosts';

// 引入组件
const DeptDialog = defineAsyncComponent(() => import('@/views/system/dept/dialog.vue'));

const hostApi = useHostApi();
// 定义变量内容
const deptDialogRef = ref();
const state = reactive<HostState>({
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

// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;
	state.tableData.data = [];
	await hostApi
		.lisHost(state.tableData.param)
		.then((res: any) => {
			state.tableData.data = res.data.hosts;
			state.tableData.total = res.data.total;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	state.tableData.loading = false;
};
// 打开新增菜单弹窗
const onOpenAddDept = (type: string) => {
	deptDialogRef.value.openDialog(type);
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
});
</script>
