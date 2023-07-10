<template>
	<div class="system-dept-container layout-padding">
		<el-row :gutter="24">
			<el-col :span="24" :xs="24">
				<el-card shadow="hover" class="layout-padding-auto">
					<div class="system-dept-search mb15">
						<el-input v-model="state.inputValue" placeholder="输入标签或者名称" clearable @change="search" style="width: 330px; margin-left: 10px">
							<template #prepend>
								<el-select v-model="state.type" style="width: 100px">
									<el-option label="名称" value="0" />
									<el-option label="登录名" value="1" />
								</el-select>
							</template>
							<template #append>
								<el-button size="small" @click="search">
									<el-icon>
										<ele-Search />
									</el-icon>
									查询
								</el-button>
							</template>
						</el-input>
						<el-button size="default" type="primary" class="ml10">
							<el-icon>
								<ele-Search />
							</el-icon>
							查询
						</el-button>
						<el-button size="default" type="success" class="ml10" @click="onOpenAddAccount('add')">
							<el-icon>
								<ele-FolderAdd />
							</el-icon>
							新增帐号
						</el-button>
						<el-button type="warning" size="default" @click="refreshCurrentTagsView" style="margin-left: 10px">
							<el-icon>
								<ele-RefreshRight />
							</el-icon>
							刷新
						</el-button>
					</div>
					<el-table
						:data="state.tableData.accounts"
						v-loading="state.tableData.loading"
						style="width: 100%"
						row-key="id"
						default-expand-all
						:tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
					>
						<el-table-column prop="name" label="名称" show-overflow-tooltip> </el-table-column>
						<el-table-column prop="user_name" label="登录名" show-overflow-tooltip> </el-table-column>
						<el-table-column prop="auth_method" label="类型" show-overflow-tooltip>
							<template #default="scope">
								<el-tag type="success" v-if="scope.row.auth_method == '01'">默认</el-tag>
								<el-tag v-else>密钥</el-tag>
							</template>
						</el-table-column>
						<el-table-column prop="status" label="是否禁用" show-overflow-tooltip>
							<template #default="scope">
								<el-tag type="success" v-if="scope.row.status">启用</el-tag>
								<el-tag type="info" v-else>禁用</el-tag>
							</template>
						</el-table-column>
						<el-table-column prop="password" label="密码" show-overflow-tooltip align="center">
							<template #default="scope">
								<el-icon color="#f56c6c"><View /></el-icon>
								<a @click="copyText(scope.row.password)">{{ scope.row.password }}</a>
							</template>
						</el-table-column>

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
								<el-button size="small" text type="primary" @click="onOpenEditAccount('edit', scope.row)">修改</el-button>
								<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
							</template>
						</el-table-column>
					</el-table>
					<!-- 分页区域 -->
					<Pagination :total="state.tableData.total" @handlePageChange="handlePageChange" />
				</el-card>
			</el-col>
		</el-row>

		<AccountDialog ref="deptDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="accounts">
import { defineAsyncComponent, ref, reactive, onMounted, h } from 'vue';
import { View } from '@element-plus/icons-vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { PageInfo } from '@/types/kubernetes/common';
import { dateStrFormat } from '@/utils/formatTime';
import { useAccountApi } from '@/api/asset/account';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { deepClone } from '@/utils/other';

// 引入组件
const AccountDialog = defineAsyncComponent(() => import('./dialog.vue'));
const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
import commonFunction from '@/utils/commonFunction';
import { HostParams } from '@/types/asset/hosts';

// 定义变量内容
const { copyText } = commonFunction();
const route = useRoute();
const accountApi = useAccountApi();
const deptDialogRef = ref();
const state = reactive({
	inputValue: '',
	type: '0',
	show: true,
	defaultProps: {
		children: 'children',
		label: 'name',
	},
	groupName: '',
	groupIds: '',
	tableData: {
		accounts: [] as Account[],
		total: 0,
		loading: false,
		param: {
			page: 1,
			limit: 10,
		} as HostParams,
	},
});

const search = () => {
	if (state.type == '0') {
		state.tableData.param.name = state.inputValue;
		delete state.tableData.param.userName;
	} else if (state.type == '1') {
		state.tableData.param.userName = state.inputValue;
		delete state.tableData.param.name;
	}
	getTableData();
};
// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;
	state.tableData.accounts = [];
	await accountApi
		.listAccount(state.tableData.param)
		.then((res: any) => {
			state.tableData.accounts = res.data.data;
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

// 打开新增菜单弹窗
const onOpenAddAccount = (type: string) => {
	deptDialogRef.value.openDialog(type);
};
// 打开编辑菜单弹窗
const onOpenEditAccount = (type: string, row: Account) => {
	deptDialogRef.value.openDialog(type, deepClone(row));
};
// 删除当前行
const onTabelRowDel = (row: Account) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [h('span', null, '此操作将删除 '), h('i', { style: 'color: teal' }, `${row.name}`), h('span', null, ' 主机. 是否继续? ')]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			accountApi
				.deleteAccount(row.id)
				.then(() => {
					ElMessage.success('删除成功');
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
			getTableData();
		})
		.catch(() => {
			ElMessage.info('已取消删除');
		});
};

const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
// 页面加载时
onMounted(() => {
	getTableData();
});
</script>
