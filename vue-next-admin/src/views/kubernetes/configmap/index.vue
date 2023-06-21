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
					<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata.name" :value="item.metadata!.name!" />
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
						<el-button size="small" type="primary" text @click="configMapDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
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
						<el-button link type="primary" size="small" @click="configMapDetail(scope.row)">详情</el-button><el-divider direction="vertical" />
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
		<DrawDialog
			v-model:visible="data.draw.visible"
			:configMap="data.draw.configMap"
			:title="data.draw.title"
			@refresh="listConfigMap"
			v-if="data.draw.visible"
		/>
		<ConfigMapDetail v-model:visible="data.detail.visible" :configMap="data.detail.configMap" :title="data.detail.title" v-if="data.detail.visible" />
	</div>
</template>

<script setup lang="ts" name="k8sConfigMap">
import { ConfigMap } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, h, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete, List } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { useConfigMapApi } from '@/api/kubernetes/configMap';
import { deepClone } from '@/utils/other';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const DrawDialog = defineAsyncComponent(() => import('./component/draw.vue'));
const ConfigMapDetail = defineAsyncComponent(() => import('./component/detail.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const ConfigMapApi = useConfigMapApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	detail: {
		title: '',
		visible: false,
		configMap: {} as ConfigMap,
	},
	draw: {
		title: '',
		visible: false,
		configMap: {} as ConfigMap,
	},
	dialogVisible: false,
	codeData: {} as ConfigMap,
	loading: false,
	selectData: [] as ConfigMap[],
	configMaps: [] as ConfigMap[],
	tmpConfigMap: [] as ConfigMap[],
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
	listConfigMap();
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

	listConfigMap();
};
const handleChange = () => {
	listConfigMap();
};
const filterConfigMap = (configMaps: Array<ConfigMap>) => {
	const configMapList = [] as ConfigMap[];
	if (data.type === '1') {
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
const createConfigMap = () => {
	data.draw.title = '创建';
	data.draw.visible = true;
};
const deleteConfigMap = (configMap: ConfigMap) => {
	data.loading = true;
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${configMap.metadata!.name}`),
			h('span', null, ' 配置. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			ConfigMapApi.deleteConfigMap(configMap.metadata!.namespace!, configMap.metadata!.name!, { cloud: k8sStore.state.activeCluster })
				.then((res: any) => {
					listConfigMap();
					ElMessage.success(res.message);
				})
				.catch((e: any) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
	data.loading = false;
};
const configMapDetail = (configMap: ConfigMap) => {
	data.detail.title = '详情';
	data.detail.configMap = configMap;
	data.detail.visible = true;
};

const showYaml = (ConfigMap: ConfigMap) => {
	data.dialogVisible = true;
	delete ConfigMap.metadata?.managedFields;
	data.codeData = ConfigMap;
};
const updateConfigMapYaml = (code: any) => {
	console.log('更新ConfigMap', code);
};

const handleSelectionChange = () => {};
const updateConfigMap = (configMap: ConfigMap) => {
	data.draw.title = '编辑';
	data.draw.configMap = deepClone(configMap) as ConfigMap;
	data.draw.visible = true;
};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listConfigMap();
};
const listConfigMap = () => {
	data.loading = true;
	ConfigMapApi.listConfigMap(k8sStore.state.activeNamespace, data.query)
		.then((res: any) => {
			data.configMaps = res.data.data;
			data.tmpConfigMap = res.data.data;
			data.total = res.data.total;
		})
		.catch((e: any) => {
			ElMessage.error(e.message);
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
