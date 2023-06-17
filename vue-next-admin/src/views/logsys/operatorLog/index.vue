<template>
	<div class="system-menu-container layout-pd">
		<el-card shadow="hover">
			<div class="system-menu-search mb15">
				操作人：
				<el-input size="default" placeholder="请输入操作人模糊查询" style="max-width: 180px; margin-right: 10px" v-model="state.searchName" clearable>
				</el-input>
				模块名称：
				<el-input
					size="default"
					placeholder="请输入模块名模糊查询"
					style="max-width: 180px; margin-right: 10px"
					v-model="state.searchTitle"
					clearable
				>
				</el-input>
				操作类型：
				<el-select v-model="state.searchType" size="default" style="max-width: 180px; width: 100px; margin-right: 10px">
					<el-option v-for="item in type" :value="item.value" :label="item.label"> </el-option>
				</el-select>
				操作状态：
				<el-select v-model="state.searchStatus" size="default" style="width: 100px">
					<el-option v-for="item in status" :value="item.value" :label="item.label"> </el-option>
				</el-select>
				<el-button size="default" type="primary" class="ml10" @click="searchLog()">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button v-auth="'logs:operator:del'" size="default" type="danger" class="ml10" @click="onDelete()">
					<el-icon>
						<ele-Delete />
					</el-icon>
					批量删除
				</el-button>
				<el-button v-auth="'logs:operator:del-all'" size="default" type="danger" class="ml10" @click="onDeleteAll()">
					<el-icon>
						<ele-Delete />
					</el-icon>
					全部清除
				</el-button>
			</div>

			<el-table
				:data="state.tableData.data"
				v-loading="state.tableData.loading"
				style="width: 100%"
				row-key="path"
				@selection-change="handleSelectionChange"
			>
				<el-table-column type="selection" width="55" />
				<el-table-column label="ID" prop="id" width="80" />
				<el-table-column label="模块名称" prop="title" width="120" />
				<el-table-column label="操作类型" prop="businessType" width="120">
					<template #default="scope">
						<div v-if="scope.row.businessType === '03'"><el-tag size="small" type="danger">删除</el-tag></div>
						<div v-else-if="scope.row.businessType === '02'"><el-tag size="small" type="warning">修改</el-tag></div>
						<div v-else><el-tag size="small">添加</el-tag></div>
					</template>
				</el-table-column>
				<!-- <el-table-column label="Method" width="120">
					<template #default="scope">
						<el-tag effect="plain" size="small" :type="scope.row.method === 'PUT' ? 'warning' : 'danger'">{{ scope.row.method }}</el-tag>
					</template>
				</el-table-column> -->
				<el-table-column label="IP" prop="ip" width="180" />
				<el-table-column label="请求地址" prop="url" />
				<el-table-column label="操作状态" width="120">
					<template #default="scope">
						<el-tooltip :content="'错误信息：' + scope.row.errMsg" placement="right" effect="light">
							<el-tag size="small" :type="scope.row.status === 200 ? 'success' : 'danger'">{{ scope.row.status }}</el-tag>
						</el-tooltip>
					</template>
				</el-table-column>
				<el-table-column label="操作人" prop="name" width="120" />
				<el-table-column prop="updated_at" label="操作时间" width="170">
					<template #default="scope"> {{ dateStrFormat(scope.row.updated_at) }}</template>
				</el-table-column>
				<el-table-column label="操作" show-overflow-tooltip width="120">
					<template #default="scope">
						<!-- <el-button v-auth="'logs:operator:del'" size="small" text type="primary" @click="onOpenEditApi('edit', scope.row)">详情</el-button> -->
						<el-button v-auth="'logs:operator:del'" size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
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
	</div>
</template>

<script setup lang="ts" name="operatorLog">
import { ref, onMounted, reactive } from 'vue';
import { RouteRecordRaw } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import { dateStrFormat } from '@/utils/formatTime';
import { useOperatorLogApi } from '@/api/logsys/opratorLog';

type query = {
	page: number;
	limit: number;
	name?: string;
	status?: string;
	title?: string;
	type?: string;
};
// 定义变量内容
const opeatorLogApi = useOperatorLogApi();
const menuDialogRef = ref();
const state = reactive({
	tableData: {
		data: [] as OperatorLog[],
		loading: true,
		param: <query>{
			page: 1,
			limit: 10,
		},
		total: 0,
	},
	searchName: '',
	searchStatus: '0',
	searchType: '0',
	searchTitle: '',
	select: [] as Array<OperatorLog>,
});
const status = [
	{
		label: '所有状态',
		value: '0',
	},
	{
		label: '正常',
		value: 'normal',
	},
	{
		label: '非正常',
		value: 'failed',
	},
];
const type = [
	{
		label: '所有类型',
		value: '0',
	},
	{
		label: '删除',
		value: '03',
	},
	{
		label: '新增',
		value: '01',
	},
	{
		label: '修改',
		value: '02',
	},
];

const onDelete = () => {
	let ids = '';
	ElMessageBox.confirm(`批量删除日志, 是否继续?`, '警告', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			let count = 0;
			state.select.forEach((item) => {
				count++;
				if (count == state.select.length) {
					ids += item.id + '';
				} else {
					ids += item.id + ',';
				}
			});
			opeatorLogApi
				.deleteOperatorLog({ ids: ids })
				.then((res) => {
					ElMessage.success(res.message);
					getTableData();
				})
				.catch((e) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};

const onDeleteAll = () => {
	ElMessageBox.confirm(`此操作将删除所有日志，请谨慎操作, 是否继续?`, '警告', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(async () => {
			await opeatorLogApi
				.deleteAllOperatorLog()
				.then((res) => {
					ElMessage.success(res.message);
					getTableData();
				})
				.catch((e) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};
const handleSelectionChange = (val: Array<OperatorLog>) => {
	state.select = val;
};
// 获取路由数据，真实请从接口获取
const getTableData = async () => {
	state.tableData.loading = true;
	if (state.searchName != '') {
		state.tableData.param.name = state.searchName;
	} else {
		delete state.tableData.param.name;
	}

	if (state.searchStatus != '0') {
		state.tableData.param.status = state.searchStatus;
	} else {
		delete state.tableData.param.status;
	}

	if (state.searchTitle != '') {
		state.tableData.param.title = state.searchTitle;
	} else {
		delete state.tableData.param.title;
	}

	if (state.searchType != '0') {
		state.tableData.param.type = state.searchType;
	} else {
		delete state.tableData.param.type;
	}

	await opeatorLogApi
		.listOperatorLog(state.tableData.param)
		.then((res) => {
			state.tableData.data = res.data.data;
			state.tableData.total = res.data.total;
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	// state.tableData.data = routesList.value;
	state.tableData.loading = false;
};

const searchLog = async () => {
	getTableData();
};

// 打开编辑菜单弹窗
const onOpenEditApi = (type: string, row: RouteRecordRaw) => {
	menuDialogRef.value.openDialog(type, row);
};

// 删除菜单或按钮
const onTabelRowDel = (row: any) => {
	ElMessageBox.confirm(`删除日志, 是否继续?`, '警告', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(async () => {
			await opeatorLogApi
				.deleteOperatorLog({ ids: row.id })
				.then((res) => {
					ElMessage.success(res.message);
					getTableData();
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
