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
					<el-option
						v-for="item in k8sStore.state.namespace"
						:key="item.metadata?.name"
						:label="item.metadata?.name!"
						:value="item.metadata!.name!"
					/>
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
				<el-button type="primary" size="small" plain class="ml10" @click="createPersistentVolumeClaim" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" plain class="ml10" :disabled="data.selectData.length == 0" :icon="Delete">批量删除</el-button>
				<el-button type="success" size="small" plain @click="refreshCurrentTagsView" style="margin-left: 10px">
					<el-icon>
						<ele-RefreshRight />
					</el-icon>
					刷新
				</el-button>
			</div>
			<el-table
				:data="data.PersistentVolumeClaims"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="data.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button :size="theme.themeConfig.globalComponentSize" text @click="persistentVolumeClaimDetail(scope.row)">
							{{ scope.row.metadata.name }}</el-button
						>
						<div>
							<el-text type="info">{{ scope.row.spec.storageClassName }}</el-text>
						</div>
					</template>
				</el-table-column>
				<el-table-column prop="status.phase" label="状态" />
				<el-table-column prop="status.capacity.storage" label="容量" />
				<el-table-column prop="spec.accessModes" label="访问模式" />
				<el-table-column label="关联的存储卷">
					<template #default="scope">
						<el-link :underline="false">{{ scope.row.spec.volumeName }}</el-link>
					</template>
				</el-table-column>
				<el-table-column label="创建时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>
				<el-table-column fixed="right" label="操作" width="260px" flex>
					<template #default="scope">
						<el-button link type="primary" size="small" @click="persistentVolumeClaimDetail(scope.row)">详情</el-button
						><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="updatePersistentVolumeClaim(scope.row)">编辑</el-button
						><el-divider direction="vertical" /> <el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button
						><el-divider direction="vertical" />
						<el-button
							:disabled="scope.row.metadata.name === 'kubernetes'"
							link
							type="danger"
							size="small"
							@click="deletePersistentVolumeClaim(scope.row)"
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
			:resourceType="'PersistentVolumeClaim'"
			@update="updatePersistentVolumeClaimYaml"
			v-if="data.dialogVisible"
		/>
		<DrawDialog
			:visible="data.draw.visible"
			:secret="data.draw.persistentVolumeClaim"
			:title="data.draw.title"
			@refresh="listPersistentVolumeClaim"
			v-if="data.draw.visible"
		/>
		<PersistentVolumeClaimDetail
			v-model:visible="data.detail.visible"
			:persistentVolumeClaim="data.detail.persistentVolumeClaim"
			:title="data.detail.title"
			v-if="data.detail.visible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sPersistentVolumeClaim">
import { PersistentVolumeClaim } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, h, onMounted, reactive } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { usePersistentVolumeClaimApi } from '@/api/kubernetes/persitentVolumeClaim';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const DrawDialog = defineAsyncComponent(() => import('./component/draw.vue'));
const PersistentVolumeClaimDetail = defineAsyncComponent(() => import('./component/detail.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const PersistentVolumeClaimApi = usePersistentVolumeClaimApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	detail: {
		title: '',
		visible: false,
		persistentVolumeClaim: {} as PersistentVolumeClaim,
	},
	draw: {
		title: '',
		visible: false,
		persistentVolumeClaim: {} as PersistentVolumeClaim,
	},
	dialogVisible: false,
	codeData: {} as PersistentVolumeClaim,
	loading: false,
	selectData: [] as PersistentVolumeClaim[],
	PersistentVolumeClaims: [] as PersistentVolumeClaim[],
	tmpPersistentVolumeClaim: [] as PersistentVolumeClaim[],
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
	listPersistentVolumeClaim();
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
	listPersistentVolumeClaim();
};
const handleChange = () => {
	listPersistentVolumeClaim();
};
// const filterPersistentVolumeClaim = (PersistentVolumeClaims: Array<PersistentVolumeClaim>) => {
// 	const PersistentVolumeClaimList = [] as PersistentVolumeClaim[];
// 	if (data.query.type === '1') {
// 		PersistentVolumeClaims.forEach((PersistentVolumeClaim: PersistentVolumeClaim) => {
// 			if (PersistentVolumeClaim.metadata?.name?.includes(data.query.key)) {
// 				PersistentVolumeClaimList.push(PersistentVolumeClaim);
// 			}
// 		});
// 	} else {
// 		PersistentVolumeClaims.forEach((PersistentVolumeClaim: PersistentVolumeClaim) => {
// 			if (PersistentVolumeClaim.metadata?.labels) {
// 				for (let k in PersistentVolumeClaim.metadata.labels) {
// 					if (k.includes(data.query.key) || PersistentVolumeClaim.metadata.labels[k].includes(data.query.key)) {
// 						PersistentVolumeClaimList.push(PersistentVolumeClaim);
// 						break;
// 					}
// 				}
// 			}
// 		});
// 	}
// 	data.PersistentVolumeClaims = PersistentVolumeClaimList;
// };
const createPersistentVolumeClaim = () => {
	data.draw.visible = true;
	data.draw.title = '创建PersistentVolumeClaim';
	data.draw.persistentVolumeClaim = {};
};
const deletePersistentVolumeClaim = (PersistentVolumeClaim: PersistentVolumeClaim) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${PersistentVolumeClaim.metadata!.name}`),
			h('span', null, ' PersistentVolumeClaim. 是否继续? '),
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
			PersistentVolumeClaimApi.deletePersistentVolumeClaim(PersistentVolumeClaim.metadata!.namespace!, PersistentVolumeClaim.metadata!.name!, {
				cloud: k8sStore.state.activeCluster,
			})
				.then((res: any) => {
					listPersistentVolumeClaim();
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
const showYaml = (PersistentVolumeClaim: PersistentVolumeClaim) => {
	data.dialogVisible = true;
	delete PersistentVolumeClaim.metadata?.managedFields;
	data.codeData = PersistentVolumeClaim;
	// yamlRef.value.openDialog(PersistentVolumeClaim);
};
const updatePersistentVolumeClaimYaml = (code: any) => {
	console.log('更新PersistentVolumeClaim', code);
};

const handleSelectionChange = () => {};
const updatePersistentVolumeClaim = (persistentVolumeClaim: PersistentVolumeClaim) => {
	data.draw.visible = true;
	data.draw.title = '修改';
	data.draw.persistentVolumeClaim = persistentVolumeClaim;
};
const persistentVolumeClaimDetail = (persistentVolumeClaim: PersistentVolumeClaim) => {
	data.detail.visible = true;
	data.detail.title = '详情';
	data.detail.persistentVolumeClaim = persistentVolumeClaim;
};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listPersistentVolumeClaim();
};
const listPersistentVolumeClaim = () => {
	data.loading = true;
	PersistentVolumeClaimApi.listPersistentVolumeClaim(k8sStore.state.activeNamespace, data.query)
		.then((res: any) => {
			data.PersistentVolumeClaims = res.data.data;
			data.tmpPersistentVolumeClaim = res.data.data;
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
