<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-row :gutter="24">
					<el-col :span="18" style="display: flex">
						<el-text class="mx-1" :size="state.size">命名空间：</el-text>
						<el-select
							v-model="k8sStore.state.activeNamespace"
							style="max-width: 180px"
							class="m-2"
							placeholder="Select"
							:size="state.size"
							@change="handleChange"
							><el-option key="all" label="所有命名空间" value="all"></el-option>
							<el-option
								v-for="item in k8sStore.state.namespace"
								:key="item.metadata?.name"
								:label="item.metadata?.name"
								:value="item.metadata!.name!"
							/>
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
								<el-select v-model="state.type" placeholder="输入标签或者名称" style="width: 80px" :size="state.size">
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

						<el-button v-auth="'k8s:daemonset:add'" type="primary" :size="state.size" class="ml10" @click="createDaemonSet" :icon="Edit"
							>创建</el-button
						>
						<el-button
							v-auth="'k8s:daemonset:del'"
							type="danger"
							:size="state.size"
							class="ml10"
							:disabled="state.selectData.length == 0"
							@click="deleteDaemonSets(state.selectData)"
							:icon="Delete"
							>批量删除</el-button
						>
						<el-button type="success" :size="state.size" @click="refreshCurrentTagsView" style="margin-left: 10px">
							<el-icon>
								<ele-RefreshRight />
							</el-icon>
							刷新
						</el-button>
					</el-col>
					<el-col :span="6" style="display: flex"> </el-col>
				</el-row>
			</div>
			<el-table :data="state.daemonsets" style="width: 100%" @selection-change="handleSelectionChange" v-loading="state.loading" :size="state.size">
				<el-table-column type="selection" width="55" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column prop="metadata.name" label="名称">
					<template #default="scope">
						<el-button link type="primary" :size="state.size" @click="daemonsetDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
					</template>
				</el-table-column>

				<el-table-column prop="spec.replicas" label="Pods" width="150px" align="center">
					<template #header> <span>Pods</span><br /><span class="table-header">就绪/副本/失败</span> </template>

					<template #default="scope">
						<a style="color: green">{{ scope.row.status.numberReady }}</a
						>/ <a style="color: green">{{ scope.row.status.numberAvailable }}</a
						>/
						<a style="color: red">{{ scope.row.status.numberMisscheduled }}</a>
					</template>
				</el-table-column>
				<el-table-column label="镜像" show-overflow-tooltip>
					<template #default="scope">
						<div class="flex">
							<SvgIcon :name="'iconfont icon-jingxiangbanben'" class="mr5" style="color: #409eff" />
							<el-text :size="state.size" truncated type="" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{
								item.image.split('@')[0]
							}}</el-text>
						</div>
					</template>
				</el-table-column>

				<el-table-column label="标签">
					<template #default="scope">
						<el-tooltip placement="right" effect="light" v-if="scope.row.metadata.labels">
							<template #content>
								<div style="display: flex; flex-direction: column">
									<el-tag
										class="label"
										type=""
										v-for="(item, key, index) in scope.row.metadata.labels"
										:key="index"
										effect="plain"
										:size="state.size"
									>
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<el-tag type="" effect="plain" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" :size="state.size">
								<div>{{ key }}:{{ item }}</div>
							</el-tag>
						</el-tooltip>
					</template>
				</el-table-column>

				<el-table-column label="创建时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>
				<el-table-column fixed="right" label="操作" width="220px">
					<template #default="scope">
						<div style="display: flex; align-items: center">
							<el-button link type="primary" :size="state.size" @click="daemonsetDetail(scope.row)">详情</el-button>
							<el-button v-auth="'k8s:daemonset:edit'" link type="primary" :size="state.size" @click="handleUpdate(scope.row)">编辑</el-button>

							<el-button link type="primary" :size="state.size" @click="daemonsetDetail(scope.row)">监控</el-button>
							<el-divider direction="vertical" />
							<div>
								<el-dropdown :size="state.size">
									<span class="el-dropdown-link" style="font-size: 14px">
										更多<el-icon class="el-icon--right">
											<CaretBottom />
										</el-icon>
									</span>
									<template #dropdown>
										<el-dropdown-menu>
											<el-dropdown-item @click="showYaml(scope.row)">查看Yaml</el-dropdown-item>
											<el-dropdown-item @click="reDeploy(scope.row)">重新部署</el-dropdown-item>
											<el-dropdown-item @click="deleteDaemonSet(scope.row)">删除</el-dropdown-item>
										</el-dropdown-menu>
									</template>
								</el-dropdown>
							</div>
						</div>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<Pagination :total="state.total" @handlePageChange="handlePageChange" />
		</el-card>
		<YamlDialog
			v-model:dialogVisible="state.dialogVisible"
			:code-data="state.codeData"
			:resourceType="'daemonset'"
			@update="updateDaemonSet"
			v-if="state.dialogVisible"
		/>

		<CreateDialog
			v-model:dialogVisible="state.create.dialogVisible"
			:title="state.create.title"
			:daemonsetment="state.daemonset"
			@refresh="listDaemonSet()"
			v-if="state.create.dialogVisible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sDaemonSet">
import { reactive, onMounted, onBeforeUnmount, defineAsyncComponent, h } from 'vue';
import { Delete, Edit } from '@element-plus/icons-vue';
import { CaretBottom } from '@element-plus/icons-vue';
import { useDaemonsetApi } from '@/api/kubernetes/daemonset';
import { DaemonSet } from 'kubernetes-models/apps/v1';
import { PageInfo } from '@/types/kubernetes/common';
import { kubernetesInfo } from '@/stores/kubernetes';
import router from '@/router';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useWebsocketApi } from '@/api/kubernetes/websocket';
import YAML from 'js-yaml';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { deepClone } from '@/utils/other';
import { useThemeConfig } from '@/stores/themeConfig';

const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const CreateDialog = defineAsyncComponent(() => import('./component/dialog.vue'));
const theme = useThemeConfig();

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const state = reactive({
	size: theme.themeConfig.globalComponentSize,
	create: {
		dialogVisible: false,
		title: '',
	},
	daemonset: {} as DaemonSet,
	codeData: {} as DaemonSet,
	searchType: '1',
	search: '',
	dialogVisible: false,
	visible: false,
	scaleDaemonSet: <DaemonSet>{},
	query: <queryType>{
		cloud: '',
		page: 1,
		limit: 10,
	},
	namespace: '',
	loading: false,
	daemonsets: [] as DaemonSet[],
	selectData: [] as DaemonSet[],
	total: 0,
	type: '1',
	inputValue: '',
});
const route = useRoute();
const daemonsetApi = useDaemonsetApi();
const k8sStore = kubernetesInfo();
const socketApi = useWebsocketApi();

const ws = socketApi.createWebsocket('daemonset');
ws.onmessage = (e) => {
	if (e.data === 'ping') {
		return;
	} else {
		const object = JSON.parse(e.data);
		if (object.type === 'daemonset' && object.result.namespace === k8sStore.state.activeNamespace && object.cluster == k8sStore.state.activeCluster) {
			state.daemonsets = object.result.data;
		}
	}
};

const search = () => {
	state.loading = true;
	state.query.cloud = k8sStore.state.activeCluster;
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
	listDaemonSet();
	state.loading = false;
};

const reDeploy = (daemonset: DaemonSet) => {
	daemonsetApi
		.reDeployDaemonset(daemonset.metadata!.namespace!, daemonset.metadata!.name!, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('操作成功');
		})
		.catch((res: any) => {
			ElMessage.error(res.message);
		});
};

const handleChange = () => {
	listDaemonSet();
};
const updateDaemonSet = async (codeData: any) => {
	const updateData = YAML.load(codeData) as DaemonSet;
	delete updateData.status;
	delete updateData.metadata?.managedFields;
	await daemonsetApi
		.updateDaemonset(updateData, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('更新成功');
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	state.dialogVisible = false;
};

const deleteDaemonSets = (daeList: Array<DaemonSet>) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将批量删除多个 '),
			h('i', { style: 'color: teal' }, `DaemonSet副本`),
			h('span', null, ' . 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			daeList.forEach((dae: DaemonSet) => {
				if (dae.metadata) {
					daemonsetApi
						.deleteDaemonset(dae.metadata?.namespace!, dae.metadata?.name!, { cloud: k8sStore.state.activeCluster })
						.then(() => {})
						.catch(() => {
							ElMessage.error(`删除${dae.metadata?.name}失败`);
						});
				}
			});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};
const deleteDaemonSet = (dae: DaemonSet) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${dae.metadata?.name}`),
			h('span', null, ' daemonset. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			daemonsetApi.deleteDaemonset(dae.metadata?.namespace!, dae.metadata?.name!, { cloud: k8sStore.state.activeCluster }).then((res: any) => {
				if (res.code === 200) {
					ElMessage.success('删除成功');
				} else {
					ElMessage.error(`删除${dae.metadata?.name}失败`);
				}
			});
		})
		.catch(() => {
			ElMessage.error('取消');
		});
};

const showYaml = async (daemonset: DaemonSet) => {
	const dae = deepClone(daemonset) as DaemonSet;
	delete dae.metadata?.managedFields;
	state.codeData = dae;
	state.dialogVisible = true;
};

const handleSelectionChange = (value: any) => {
	state.selectData = value;
};

const listDaemonSet = async () => {
	state.namespace = k8sStore.state.activeNamespace;
	state.query.cloud = k8sStore.state.activeCluster;
	try {
		state.loading = true;
		await daemonsetApi.listDaemonset(k8sStore.state.activeNamespace, state.query).then((res) => {
			state.daemonsets = res.data.data;
			state.total = res.data.total;
		});
	} catch (e: any) {
		if (e.code != 5003) ElMessage.error(e.message);
	}
	state.loading = false;
};

mittBus.on('changeNamespace', () => {
	listDaemonSet();
});

const handlePageChange = (pageInfo: PageInfo) => {
	state.query.page = pageInfo.page;
	state.query.limit = pageInfo.limit;

	if (state.search != '') {
		search();
	} else {
		listDaemonSet();
	}
};
const daemonsetDetail = async (dae: DaemonSet) => {
	k8sStore.state.activeDaemonSet = dae;
	await router.push({
		name: 'k8sDaemonsetDetail',
		params: {
			name: dae.metadata?.name,
		},
	});
};

const createDaemonSet = () => {
	state.create.title = '创建daemonset';
	state.create.dialogVisible = true;
};

const handleUpdate = (daemonset: DaemonSet) => {
	k8sStore.state.creatDaemonSet.namespace = daemonset.metadata!.namespace!;
	const dae = deepClone(daemonset) as DaemonSet;
	delete dae.status;
	delete dae.metadata?.managedFields;
	state.daemonset = dae;
	state.create.title = '更新daemonset';
	state.create.dialogVisible = true;
};

onMounted(() => {
	listDaemonSet();
});

const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};

onBeforeUnmount(() => {
	ws.close();
});
</script>

<style scoped lang="scss">
.container {
	:deep(.el-card__body) {
		display: flex;
		flex-direction: column;
		flex: 1;
		overflow: auto;

		.el-table {
			flex: 1;
		}
	}
}

.el-input-number {
	width: 200px;
}

.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
