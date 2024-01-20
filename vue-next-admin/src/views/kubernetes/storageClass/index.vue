<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto layout-padding-view">
			<div class="mb15">
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
				<el-button type="primary" size="small" class="ml10" @click="createStorageClass" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" @click="deleteStorageClass" :icon="Delete"
					>批量删除</el-button
				>
				<el-button type="success" size="small" @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="data.StorageClasses"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="data.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button :size="theme.themeConfig.globalComponentSize" type="primary" text @click="storageClassDetail(scope.row)">
							{{ scope.row.metadata.name }}</el-button
						>
					</template>
				</el-table-column>
				<el-table-column label="是否默认">
					<template #default="scope">
						<el-tag v-if="scope.row.metadata.annotations['storageclass.kubernetes.io/is-default-class'] == 'true'" type="success" effect="plain"
							>是</el-tag
						>
						<el-tag v-else type="warning" effect="plain">否</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="provisioner" label="提供者" />
				<el-table-column prop="reclaimPolicy" label="回收策略" />
				<el-table-column prop="volumeBindingMode" label="绑定策略" />
				<el-table-column label="创建时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>

				<el-table-column fixed="right" label="操作" width="260px" flex>
					<template #default="scope">
						<el-button link type="primary" size="small" @click="storageClassDetail(scope.row)">详情</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="updateStorageClass(scope.row)">编辑</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button><el-divider direction="vertical" />
						<el-button :disabled="scope.row.metadata.name === 'kubernetes'" link type="danger" size="small" @click="deleteStorageClass(scope.row)"
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
			:resourceType="'StorageClass'"
			@update="updateStorageClassYaml"
			v-if="data.dialogVisible"
		/>
		<DrawDialog
			v-model:visible="data.draw.visible"
			:storageClass="data.draw.storageClass"
			:title="data.draw.title"
			@refresh="listStorageClass"
			v-if="data.draw.visible"
		/>
		<StorageClassDetail
			v-model:visible="data.detail.visible"
			:storageClass="data.detail.storageClass"
			:title="data.detail.title"
			v-if="data.detail.visible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sStorageClass">
import { defineAsyncComponent, h, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { useStorageClassApi } from '@/api/kubernetes/storageClass';
import { StorageClass } from 'kubernetes-types/storage/v1';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const DrawDialog = defineAsyncComponent(() => import('./component/draw.vue'));
const StorageClassDetail = defineAsyncComponent(() => import('./component/detail.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const StorageClassApi = useStorageClassApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	detail: {
		title: '',
		visible: false,
		storageClass: {} as StorageClass,
	},
	draw: {
		title: '',
		visible: false,
		storageClass: {} as StorageClass,
	},
	dialogVisible: false,
	codeData: {} as StorageClass,
	loading: false,
	selectData: [] as StorageClass[],
	StorageClasses: [] as StorageClass[],
	tmpStorageClass: [] as StorageClass[],
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
	listStorageClass();
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
	listStorageClass();
};
// const handleChange = () => {
// 	listStorageClass();
// };
// const filterStorageClass = (StorageClasses: Array<StorageClass>) => {
// 	const StorageClassList = [] as StorageClass[];
// 	if (data.query.type === '1') {
// 		StorageClasses.forEach((StorageClass: StorageClass) => {
// 			if (StorageClass.metadata?.name?.includes(data.query.key)) {
// 				StorageClassList.push(StorageClass);
// 			}
// 		});
// 	} else {
// 		StorageClasses.forEach((StorageClass: StorageClass) => {
// 			if (StorageClass.metadata?.labels) {
// 				for (let k in StorageClass.metadata.labels) {
// 					if (k.includes(data.query.key) || StorageClass.metadata.labels[k].includes(data.query.key)) {
// 						StorageClassList.push(StorageClass);
// 						break;
// 					}
// 				}
// 			}
// 		});
// 	}
// 	data.StorageClasses = StorageClassList;
// };
const createStorageClass = () => {
	data.draw.visible = true;
	data.draw.title = '创建StorageClass';
	data.draw.storageClass = {} as StorageClass;
};
const deleteStorageClass = (StorageClass: StorageClass) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${StorageClass.metadata!.name}`),
			h('span', null, ' StorageClass. 是否继续? '),
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
			StorageClassApi.deleteStorageClass(StorageClass.metadata!.name!, { cloud: k8sStore.state.activeCluster })
				.then((res: any) => {
					listStorageClass();
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
const showYaml = (StorageClass: StorageClass) => {
	data.dialogVisible = true;
	delete StorageClass.metadata?.managedFields;
	data.codeData = StorageClass;
	// yamlRef.value.openDialog(StorageClass);
};
const updateStorageClassYaml = (code: any) => {
	console.log('更新StorageClass', code);
};

const handleSelectionChange = () => {};
const updateStorageClass = (storageClass: StorageClass) => {
	data.draw.visible = true;
	data.draw.title = '修改';
	data.draw.storageClass = storageClass;
};
const storageClassDetail = (storageClass: StorageClass) => {
	data.detail.visible = true;
	data.detail.title = '详情';
	data.detail.storageClass = storageClass;
};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listStorageClass();
};
const listStorageClass = () => {
	data.loading = true;
	StorageClassApi.listStorageClass(data.query)
		.then((res: any) => {
			data.StorageClasses = res.data.data;
			data.tmpStorageClass = res.data.data;
			data.total = res.data.total;
		})
		.catch((e: any) => {
			if (e.code != 5003) ElMessage.error(e.message);
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
