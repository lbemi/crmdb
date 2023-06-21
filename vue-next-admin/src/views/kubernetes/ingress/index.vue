<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto layout-padding-view">
			<div class="mb15">
				<el-text class="mx-1" :size="theme.themeConfig.globalComponentSize">命名空间：</el-text>
				<el-select
					v-model="k8sStore.state.activeNamespace"
					style="max-width: 280px"
					class="m-2"
					placeholder="Select"
					size="small"
					@change="handleChange"
					><el-option key="all" label="所有命名空间" value="all"></el-option>
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata!.name!" />
				</el-select>
				<el-input
					v-model="data.inputValue"
					placeholder="输入标签或者名称"
					size="small"
					clearable
					@change="search"
					style="max-width: 300px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="data.type" placeholder="输入标签或者名称" style="max-width: 120px" size="small">
							<el-option label="标签" value="0" size="small" />
							<el-option label="名称" value="1" size="small" />
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
				<el-button type="primary" size="small" class="ml10" @click="createIngress" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="data.services"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="data.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button :size="theme.themeConfig.globalComponentSize" type="primary" text> {{ scope.row.metadata.name }}</el-button>
					</template>
				</el-table-column>
				<el-table-column prop="metadata.namespace" label="命名空间" />
				<el-table-column prop="spec.type" label="类型" />
				<el-table-column label="外部访问IP">
					<template #default="scope">
						<el-link
							target="_blank"
							type="primary"
							:size="theme.themeConfig.globalComponentSize"
							v-if="scope.row.status.loadBalancer.ingress"
							v-for="item in scope.row.status.loadBalancer.ingress"
							:href="'http://' + item.ip"
						>
							{{ item.ip }}</el-link
						>
					</template>
				</el-table-column>
				<el-table-column label="规则" style="display: flex" min-width="200px">
					<template #default="scope">
						<div v-for="item in scope.row.spec.rules">
							<div v-for="path in item.http.paths">
								<el-link :size="theme.themeConfig.globalComponentSize" target="_blank" type="primary" :href="'http://' + item.host + path.path"
									>{{ item.host }}{{ path.path }}</el-link
								>
								<a> -> {{ path.backend.service.name }}:{{ path.backend.service.port.number }}</a>
							</div>
						</div>
					</template>
				</el-table-column>

				<el-table-column label="标签" width="70px">
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
										:size="theme.themeConfig.globalComponentSize"
									>
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<el-icon><List /></el-icon>
						</el-tooltip>
					</template>
				</el-table-column>

				<el-table-column label="创建时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>

				<el-table-column fixed="right" label="操作" width="160px">
					<template #default="scope">
						<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="updateIngress(scope.row)">详情</el-button
						><el-divider direction="vertical" />
						<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="updateIngress(scope.row)">编辑</el-button
						><el-divider direction="vertical" />
						<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="showYaml(scope.row)">查看YAML</el-button
						><el-divider direction="vertical" />
						<el-button
							:disabled="scope.row.metadata.name === 'kubernetes'"
							link
							type="danger"
							:size="theme.themeConfig.globalComponentSize"
							@click="deleteIngress(scope.row)"
							>删除</el-button
						>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<pagination :total="data.total" @handlePageChange="handlePageChange"></pagination>
		</el-card>
		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:code-data="data.codeData"
			:resourceType="'ingress'"
			@update="updateIngressYaml"
			v-if="data.dialogVisible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sIngress">
import { Ingress } from 'kubernetes-types/networking/v1';
import { defineAsyncComponent, h, onMounted, reactive, ref } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ResponseType } from '@/types/response';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete, List } from '@element-plus/icons-vue';
import { useIngressApi } from '@/api/kubernetes/ingress';
import { useThemeConfig } from '@/stores/themeConfig';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const ingressApi = useIngressApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	dialogVisible: false,
	codeData: {} as Ingress,
	loading: false,
	selectData: [] as Ingress[],
	services: [] as Ingress[],
	tmpIngress: [] as Ingress[],
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
	listIngress();
});

const search = () => {
	if (data.type == '1') {
		data.query.name = data.inputValue;
		delete data.query.label;
	} else if (data.type == '0') {
		data.query.label = data.inputValue;
		delete data.query.name;
	}
	if (data.inputValue === '') {
		delete data.query.label;
		delete data.query.name;
	}
	listIngress();
};
const handleChange = () => {
	listIngress();
};

const createIngress = () => {};
const deleteIngress = (service: Ingress) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${service.metadata!.name}`),
			h('span', null, ' 服务. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			data.loading = true;
			ingressApi
				.deleteIngress({ cloud: k8sStore.state.activeCluster }, service.metadata!.name!, service.metadata!.namespace!)
				.then((res: any) => {
					listIngress();
					ElMessage.success(res.message);
				})
				.catch((e) => {
					ElMessage.error(e);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
	data.loading = false;
};
const showYaml = (ingress: Ingress) => {
	data.dialogVisible = true;
	delete ingress.metadata?.managedFields;
	data.codeData = ingress;
	// yamlRef.value.openDialog(ingress);
};
const updateIngressYaml = (code: any) => {};

const handleSelectionChange = () => {};
const updateIngress = (ervice: Ingress) => {};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listIngress();
};
const listIngress = () => {
	data.loading = true;
	ingressApi
		.listIngress(k8sStore.state.activeNamespace, data.query)
		.then((res: ResponseType) => {
			if (res.code === 200) {
				data.services = res.data.data;
				data.tmpIngress = res.data.data;
				data.total = res.data.total;
			}
		})
		.catch((e) => {
			ElMessage.error(e);
		});

	data.loading = false;
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
