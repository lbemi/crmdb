<template>
	<div class="system-dept-container layout-padding">
		<el-row :gutter="24">
			<el-col :span="4" :xs="24">
				<el-card shadow="hover">
					<div class="head-container">
						<el-input v-model="state.groupName" placeholder="请输入分组名称" clearable prefix-icon="el-icon-search" style="margin-bottom: 20px" />
					</div>
					<div class="head-container">
						<el-tree
							:data="state.tableData.groups"
							node-key="id"
							:filter-node-method="filterNode"
							:props="state.defaultProps"
							:expand-on-click-node="false"
							ref="tree"
							default-expand-all
							@node-click="handleNodeClick"
						/>
					</div>
				</el-card>
			</el-col>
			<el-col :span="20" :xs="24">
				<el-card shadow="hover" class="layout-padding-auto">
					<div class="system-dept-search ml10">
						<el-input v-model="state.inputValue" placeholder="输入IP/标签/描述" clearable @change="search" style="width: 300px" size="default">
							<template #prepend>
								<el-select v-model="state.type" placeholder="输入IP/标签/描述" style="width: 80px" size="default">
									<el-option label="IP" value="0" size="default" />
									<el-option label="标签" value="1" size="default" />
									<el-option label="描述" value="2" size="default" />
								</el-select>
							</template>
							<template #append>
								<el-button size="default" @click="search">
									<el-icon>
										<ele-Search />
									</el-icon>
									查询
								</el-button>
							</template>
						</el-input>

						<el-button size="default" type="success" class="ml10" @click="onOpenAddHost('add')">
							<el-icon>
								<ele-FolderAdd />
							</el-icon>
							新增主机
						</el-button>
						<el-button type="warning" size="default" @click="refreshCurrentTagsView" style="margin-left: 10px">
							<el-icon>
								<ele-RefreshRight />
							</el-icon>
							刷新
						</el-button>
					</div>
					<el-table
						:data="state.tableData.hosts"
						v-loading="state.tableData.loading"
						style="width: 100%"
						row-key="id"
						default-expand-all
						:tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
					>
						<el-table-column prop="ip" label="主机IP" show-overflow-tooltip> </el-table-column>
						<el-table-column prop="remark" label="描述" show-overflow-tooltip> </el-table-column>

						<el-table-column prop="status" label="是否禁用" show-overflow-tooltip>
							<template #default="scope">
								<el-tag size="small" type="success" v-if="scope.row.status">启用</el-tag>
								<el-tag size="small" type="info" v-else>禁用</el-tag>
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
								<el-tag size="small" type="success" v-if="scope.row.enable_ssh">启用</el-tag>
								<el-tag size="small" type="info" v-else>禁用</el-tag>
							</template>
						</el-table-column>
						<el-table-column label="创建时间" show-overflow-tooltip>
							<template #default="scope">
								{{ dateStrFormat(scope.row.created_at) }}
							</template>
						</el-table-column>
						<el-table-column label="操作" show-overflow-tooltip width="150" fixed="right">
							<template #default="scope">
								<el-button size="small" text type="primary" @click="onOpenAddHost('add')">SSH</el-button>
								<el-button size="small" text type="primary" @click="onOpenEditHost('edit', scope.row)">修改</el-button>
								<el-button size="small" text type="primary" @click="onTabelRowDel(scope.row)">删除</el-button>
							</template>
						</el-table-column>
					</el-table>
					<!-- 分页区域 -->
					<Pagination :total="state.tableData.total" @handlePageChange="handlePageChange" />
				</el-card>
			</el-col>
		</el-row>

		<HostDialog ref="hostDialogRef" @refresh="getTableData()" />
	</div>
</template>

<script setup lang="ts" name="hosts">
import { defineAsyncComponent, ref, reactive, onMounted, h } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useHostApi } from '@/api/asset/hosts';
import { Host, HostState } from '@/types/asset/hosts';
import { useGroupApi } from '@/api/asset/group';
import { Group } from '@/types/asset/group';
import { PageInfo } from '@/types/kubernetes/common';
import { dateStrFormat } from '../../../utils/formatTime';
import { deepClone } from '@/utils/other';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';

// 引入组件
const HostDialog = defineAsyncComponent(() => import('./dialog.vue'));
const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));

const route = useRoute();
const hostApi = useHostApi();
const groupApi = useGroupApi();
const tree = ref('');
// 定义变量内容
const hostDialogRef = ref();
const state = reactive<HostState>({
	inputValue: '',
	type: '0',
	defaultProps: {
		children: 'children',
		label: 'name',
	},
	groupName: '',
	groupIds: '',
	tableData: {
		hosts: [],
		groups: [],
		total: 0,
		loading: false,
		param: {
			page: 1,
			limit: 10,
		},
	},
});
const search = () => {
	if (state.type == '0') {
		state.tableData.param.ip = state.inputValue;
		delete state.tableData.param.label;
		delete state.tableData.param.description;
	} else if (state.type == '1') {
		state.tableData.param.label = state.inputValue;
		delete state.tableData.param.ip;
		delete state.tableData.param.description;
	} else {
		state.tableData.param.description = state.inputValue;
		delete state.tableData.param.ip;
		delete state.tableData.param.label;
	}
	getTableData();
};
// 初始化表格数据
const getTableData = async () => {
	state.tableData.loading = true;
	state.tableData.hosts = [];
	await groupApi
		.lisGroup({ page: 0, limit: 0 })
		.then((res: any) => {
			state.tableData.groups = res.data.data;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	await hostApi
		.lisHost(state.tableData.param)
		.then((res: any) => {
			state.tableData.hosts = res.data.hosts;
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
// 节点单击事件
const handleNodeClick = (data: any) => {
	state.groupIds = '';
	getAllNode(data);
	state.tableData.param['groups'] = state.groupIds;
	getTableData();
};

// 筛选节点
const filterNode = (value: string, state: any) => {
	if (!value) return true;
	return state.groupName.includes(value);
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
const onOpenAddHost = (type: string) => {
	hostDialogRef.value.openDialog(type, state.tableData.groups);
};
// 打开编辑菜单弹窗
const onOpenEditHost = (type: string, host: Host) => {
	hostDialogRef.value.openDialog(type, state.tableData.groups, deepClone(host));
};
// 删除当前行
const onTabelRowDel = (row: Host) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [h('span', null, '此操作将删除 '), h('i', { style: 'color: teal' }, `${row.ip}`), h('span', null, ' 主机. 是否继续? ')]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			hostApi
				.deleteHost(row.id)
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
