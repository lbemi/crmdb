<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto layout-padding-view">
			<div class="mb15">
				<el-text class="mx-1" :size="state.size">命名空间：</el-text>
				<el-select
					v-model="k8sStore.state.activeNamespace"
					style="max-width: 280px"
					class="m-2"
					placeholder="Select"
					:size="state.size"
					@change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata!.name!" />
				</el-select>
				<el-input
					v-model="state.inputValue"
					placeholder="输入标签或者名称"
					:size="state.size"
					clearable
					@change="search"
					style="width: 350px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="state.type" style="width: 80px" :size="state.size">
							<el-option label="标签" value="0" :size="state.size" />
							<el-option label="名称" value="1" :size="state.size" />
						</el-select>
					</template>
					<template #append>
						<el-button :size="state.size" @click="search">
							<el-icon>
								<ele-Search />
							</el-icon>
							查询
						</el-button>
					</template>
				</el-input>
				<el-button type="primary" :size="state.size" class="ml10" @click="createVirtualService" :icon="Edit">创建</el-button>
				<el-button type="danger" :size="state.size" class="ml10" :disabled="state.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" :size="state.size" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="state.pipelines"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="state.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button :size="state.size" type="primary" text @click="virtualServiceDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
					</template>
				</el-table-column>
				<el-table-column label="标签">
					<template #default="scope">
						<el-tooltip placement="right" effect="light" v-if="scope.row.metadata.labels">
							<template #content>
								<div style="display: flex; flex-direction: column">
									<el-tag
										class="label"
										effect="plain"
										type="info"
										v-for="(item, key, index) in scope.row.metadata.labels"
										:key="index"
										::size="state.size"
									>
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<el-icon><List /></el-icon>
						</el-tooltip>
						<a v-else>无</a>
					</template>
				</el-table-column>

				<el-table-column label="创建时间">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>

				<el-table-column fixed="right" label="操作" width="260px" flex>
					<template #default="scope">
						<el-button link type="primary" :size="state.size" @click="virtualServiceDetail(scope.row)">详情</el-button
						><el-divider direction="vertical" />
						<el-button link type="primary" :size="state.size" @click="updateVirtualService(scope.row)">编辑</el-button
						><el-divider direction="vertical" /> <el-button link type="primary" :size="state.size" @click="showYaml(scope.row)">查看YAML</el-button
						><el-divider direction="vertical" />
						<el-button :disabled="scope.row.metadata.name === 'kubernetes'" link type="danger" :size="state.size" @click="deletePipeline(scope.row)"
							>删除</el-button
						>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<pagination :total="state.total" @handlePageChange="handlePageChange"></pagination>
		</el-card>
		<YamlDialog
			v-model:dialogVisible="state.dialogVisible"
			:code-data="state.codeData"
			@update="updateVirtualServiceYaml"
			v-if="state.dialogVisible"
		/>
	</div>
</template>

<script setup lang="ts" name="pipeline">
import { VirtualService } from '@kubernetes-models/istio/networking.istio.io/v1beta1/VirtualService';
import { defineAsyncComponent, h, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete, List } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { deepClone } from '@/utils/other';
import { useTektonPipelinesApi } from '@/api/tekton/pipelines';

import { Pipeline } from '@/types/tekton/test/pipeline';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));

const pipelineApi = useTektonPipelinesApi();

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const route = useRoute();
const theme = useThemeConfig();

//定义数据
const state = reactive({
	size: theme.themeConfig.globalComponentSize,
	detail: {
		title: '',
		visible: false,
		virtualService: {} as VirtualService,
	},
	draw: {
		title: '',
		visible: false,
		virtualService: {} as VirtualService,
	},
	dialogVisible: false,
	codeData: {} as VirtualService,
	loading: false,
	selectData: [] as VirtualService[],
	pipelines: [] as Pipeline[],
	tmpVirtualService: [] as VirtualService[],
	total: 0,
	type: '1',
	inputValue: '',
	query: <queryType>{
		page: 1,
		limit: 10,
		cloud: k8sStore.state.activeCluster,
	},
});

onMounted(() => {
	listPipelines();
});

const search = () => {
	if (state.type == '1') {
		state.query.name = state.inputValue;
		delete state.query.label;
	} else if (state.type == '0') {
		state.query.label = state.inputValue;
		delete state.query.name;
	}
	if (state.inputValue === '') {
		delete state.query.label;
		delete state.query.name;
	}

	listPipelines();
};
const handleChange = () => {
	listPipelines();
};
const createVirtualService = () => {
	state.draw.title = '创建虚拟服务';
	state.draw.visible = true;
};
const deletePipeline = (pipeline: Pipeline) => {
	state.loading = true;
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${pipeline.metadata!.name}`),
			h('span', null, ' 任务. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			pipelineApi
				.deletePipeline(pipeline.metadata.namespace, pipeline.metadata.name, {
					cloud: k8sStore.state.activeCluster,
				})
				.then((res: any) => {
					listPipelines();
					ElMessage.success(res.message);
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
	state.loading = false;
};
const virtualServiceDetail = (virtualService: VirtualService) => {
	state.detail.title = '详情';
	state.detail.virtualService = virtualService;
	state.detail.visible = true;
};

const showYaml = (VirtualService: VirtualService) => {
	state.dialogVisible = true;
	delete VirtualService.metadata?.managedFields;
	state.codeData = VirtualService;
};
const updateVirtualServiceYaml = (code: any) => {
	console.log('更新Pipeline', code);
};

const handleSelectionChange = () => {};
const updateVirtualService = (virtualService: VirtualService) => {
	state.draw.title = '编辑';
	state.draw.virtualService = deepClone(virtualService) as VirtualService;
	state.draw.visible = true;
};
const handlePageChange = (page: PageInfo) => {
	state.query.page = page.page;
	state.query.limit = page.limit;
	listPipelines();
};
const listPipelines = () => {
	state.loading = true;
	pipelineApi
		.listPipeline(k8sStore.state.activeNamespace, state.query)
		.then((res: any) => {
			state.pipelines = res.data.data;
			state.total = res.data.total;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
		});
	state.loading = false;
};

const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};
</script>

<style scoped lang="scss">
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
