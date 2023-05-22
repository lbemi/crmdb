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
					v-model="data.query.key"
					placeholder="输入标签或者名称"
					size="small"
					clearable
					@change="search"
					style="max-width: 300px; margin-left: 10px"
				>
					<template #prepend>
						<el-select v-model="data.query.type" placeholder="输入标签或者名称" style="max-width: 120px" size="small">
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
				<el-button type="primary" size="small" class="ml10" @click="createConfigMap" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="data.configMaps"
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

				<el-table-column fixed="right" label="操作" width="260px" flex>
					<template #default="scope">
						<el-button link type="primary" size="small" @click="updateConfigMap(scope.row)">详情</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="updateConfigMap(scope.row)">编辑</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button><el-divider direction="vertical" />
						<el-button :disabled="scope.row.metadata.name === 'kubernetes'" link type="danger" size="small" @click="deleteConfigMap(scope.row)"
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
			:resourceType="'configMap'"
			@update="updateConfigMapYaml"
			v-if="data.dialogVisible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sSecret">
import { ConfigMap } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, h, onMounted, reactive, ref } from 'vue';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { ResponseType } from '/@/types/response';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '/@/utils/formatTime';
import { PageInfo } from '/@/types/kubernetes/common';
import { Edit, Delete, List } from '@element-plus/icons-vue';
import { useThemeConfig } from '/@/stores/themeConfig';
import { useConfigMapApi } from '/@/api/kubernetes/configMap';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));

const k8sStore = kubernetesInfo();
const ConfigMapApi = useConfigMapApi();
const configMapApi = useConfigMapApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	dialogVisible: false,
	codeData: {} as ConfigMap,
	loading: false,
	selectData: [] as ConfigMap[],
	configMaps: [] as ConfigMap[],
	tmpConfigMap: [] as ConfigMap[],
	total: 0,
	query: {
		page: 1,
		limit: 10,
		key: '',
		type: '1',
		cloud: k8sStore.state.activeCluster,
	},
});
onMounted(() => {
	listConfigMap();
});

const search = () => {
	filterConfigMap(data.tmpConfigMap);
};
const handleChange = () => {
	listConfigMap();
};
const filterConfigMap = (configMaps: Array<ConfigMap>) => {
	const configMapList = [] as ConfigMap[];
	if (data.query.type === '1') {
		configMaps.forEach((configMap: ConfigMap) => {
			if (configMap.metadata?.name?.includes(data.query.key)) {
				configMapList.push(configMap);
			}
		});
	} else {
		configMaps.forEach((configMap: ConfigMap) => {
			if (configMap.metadata?.labels) {
				for (let k in configMap.metadata.labels) {
					if (k.includes(data.query.key) || configMap.metadata.labels[k].includes(data.query.key)) {
						configMapList.push(configMap);
						break;
					}
				}
			}
		});
	}
	data.configMaps = configMapList;
};
const createConfigMap = () => {};
const deleteConfigMap = (configMap: ConfigMap) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${configMap.metadata!.name}`),
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
			ConfigMapApi.deleteConfigMap(configMap.metadata!.namespace!, configMap.metadata!.name!, { cloud: k8sStore.state.activeCluster })
				.then((res) => {
					listConfigMap();
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
const showYaml = (ConfigMap: ConfigMap) => {
	data.dialogVisible = true;
	delete ConfigMap.metadata?.managedFields;
	data.codeData = ConfigMap;
	// yamlRef.value.openDialog(ConfigMap);
};
const updateConfigMapYaml = (code: any) => {
	console.log('更新ConfigMap', code);
};

const handleSelectionChange = () => {};
const updateConfigMap = (ervice: ConfigMap) => {};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listConfigMap();
};
const listConfigMap = () => {
	ConfigMapApi.listConfigMap(k8sStore.state.activeNamespace, data.query)
		.then((res: ResponseType) => {
			if (res.code === 200) {
				data.configMaps = res.data.data;
				data.tmpConfigMap = res.data.data;
				data.total = res.data.total;
			}
		})
		.catch((e) => {
			ElMessage.error(e);
		});
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
