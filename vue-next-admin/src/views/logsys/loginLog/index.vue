<template>
	<div class="system-menu-container layout-pd">
		<el-card shadow="hover">
			<div class="system-menu-search mb15">
				<el-input size="default" placeholder="请输入登录名模糊查询" style="max-width: 180px; margin-right: 10px" v-model="state.searchName" clearable>
				</el-input>
				<el-select v-model="state.searchStatus" size="default">
					<el-option v-for="item in status" :value="item.value" :label="item.label"> </el-option>
				</el-select>
				<el-button size="default" type="primary" class="ml10" @click="searchLog()">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
				<el-button v-auth="'logs:login:del'" size="default" type="danger" class="ml10" @click="onDelete()">
					<el-icon>
						<ele-Delete />
					</el-icon>
					批量删除
				</el-button>
				<el-button v-auth="'logs:login:del-all'" size="default" type="danger" class="ml10" @click="onDeleteAll()">
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
				<el-table-column label="ID" prop="id" />
				<el-table-column label="登录名" prop="username" />
				<el-table-column label="登录IP" prop="ipaddr" />
				<el-table-column
					label="登录状态"
					prop="status"
					:filters="[
						{ text: '成功', value: '1' },
						{ text: '失败', value: '-1' },
					]"
					:filter-method="filterTag"
					filter-placement="bottom-end"
				>
					<template #default="scope">
						<el-tag size="small" type="success" v-if="scope.row.status === '1'">成功</el-tag>
						<el-tag size="small" type="danger" v-else>失败</el-tag>
					</template>
				</el-table-column>
				<el-table-column label="浏览器" prop="browser" />
				<el-table-column label="登录平台" prop="os" />
				<el-table-column label="请求头" prop="remark" show-overflow-tooltip />
				<el-table-column prop="created_at" label="登录时间">
					<template #default="scope"> {{ dateStrFormat(scope.row.loginTime) }}</template>
				</el-table-column>
				<el-table-column label="操作" show-overflow-tooltip width="80">
					<template #default="scope">
						<!-- <el-button v-auth="'logs:login:del'" size="small" text type="primary" @click="onOpenEditApi('edit', scope.row)">详情</el-button> -->
						<el-button v-auth="'logs:login:del'" size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
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

<script setup lang="ts" name="loginLog">
import { ref, onMounted, reactive } from 'vue';
import { RouteRecordRaw } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import { dateStrFormat } from '@/utils/formatTime';
import { useLoinLogApi } from '@/api/logsys/loginLog';

type query = {
	page: number;
	limit: number;
	name?: string;
	status?: string;
};
// 定义变量内容
const loginLogApi = useLoinLogApi();
const menuDialogRef = ref();
const state = reactive({
	tableData: {
		data: [] as LoginLog[],
		loading: true,
		param: <query>{
			page: 1,
			limit: 10,
		},
		total: 0,
	},
	searchName: '',
	searchStatus: '0',
	select: [] as Array<LoginLog>,
});

const status = [
	{
		label: '所有状态',
		value: '0',
	},
	{
		label: '成功',
		value: '1',
	},
	{
		label: '失败',
		value: '-1',
	},
];

const filterTag = (value: string, row: LoginLog) => {
	return row.status === value;
};

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
			loginLogApi
				.deleteLoginLog({ ids: ids })
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
			await loginLogApi
				.deleteAllLoginLog()
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
const handleSelectionChange = (val: Array<LoginLog>) => {
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
	await loginLogApi
		.listLoginLog(state.tableData.param)
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
			await loginLogApi
				.deleteLoginLog({ ids: row.id })
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
