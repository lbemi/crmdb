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

						<el-button v-auth="'k8s:cronjob:add'" type="primary" :size="state.size" class="ml10" @click="createCronJob" :icon="Edit">创建</el-button>
						<el-button
							v-auth="'k8s:cronjob:del'"
							type="danger"
							:size="state.size"
							class="ml10"
							:disabled="state.selectData.length == 0"
							@click="deleteCronJobs(state.selectData)"
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
			<el-table :data="state.cronjobs" style="width: 100%" @selection-change="handleSelectionChange" v-loading="state.loading" :size="state.size">
				<el-table-column type="selection" width="55" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column prop="metadata.name" label="名称">
					<template #default="scope">
						<el-button link type="primary" :size="state.size" @click="deployDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
						<div style="color: red">{{ depStatus(scope.row) }}</div>
					</template>
				</el-table-column>
				<el-table-column label="状态" prop="spec.concurrencyPolicy" width="130px"> </el-table-column>
				<el-table-column label="调度规则" prop="spec.schedule"> </el-table-column>
				<el-table-column label="是否挂起" prop="spec.suspend"> </el-table-column>
				<!--				<el-table-column prop="spec.replicas" label="Pods" width="150px" align="center">-->
				<!--					<template #header> <span>Pods</span><br /><span style="font-size: 10px; font-weight: 50">就绪/副本/失败</span> </template>-->

				<!--					<template #default="scope">-->
				<!--						<a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a-->
				<!--						>/ <a style="color: green">{{ scope.row.status.replicas || '0' }}</a-->
				<!--						>/-->
				<!--						<a style="color: red">{{ scope.row.status.unavailableReplicas || '0' }}</a>-->
				<!--					</template>-->
				<!--				</el-table-column>-->
				<!--				<el-table-column label="镜像" show-overflow-tooltip>-->
				<!--					<template #default="scope">-->
				<!--						<el-text :size="state.size" truncated type="" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{-->
				<!--							item.image.split('@')[0]-->
				<!--						}}</el-text>-->
				<!--					</template>-->
				<!--				</el-table-column>-->

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
				<el-table-column label="上次执行时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.status.lastScheduleTime) }}
					</template>
				</el-table-column>
				<el-table-column label="上次执行成功时间" width="170px">
					<template #default="scope">
						{{ dateStrFormat(scope.row.status.lastSuccessfulTime) }}
					</template>
				</el-table-column>
				<el-table-column fixed="right" label="操作" width="165px">
					<template #default="scope">
						<div style="display: flex; align-items: center">
							<el-button link type="primary" :size="state.size" @click="deployDetail(scope.row)">详情</el-button>
							<el-button v-auth="'k8s:cronjob:edit'" link type="primary" :size="state.size" @click="handleUpdate(scope.row)">编辑</el-button>
							<el-divider direction="vertical" />
							<div class="flex-auto">
								<el-dropdown :size="state.size">
									<span class="el-dropdown-link" style="font-size: 12px">
										更多<el-icon class="el-icon--right">
											<CaretBottom />
										</el-icon>
									</span>
									<template #dropdown>
										<el-dropdown-menu>
											<el-dropdown-item @click="showYaml(scope.row)">查看Yaml</el-dropdown-item>
											<el-dropdown-item @click="reDeploy(scope.row)">重新部署</el-dropdown-item>
											<el-dropdown-item>编辑标签</el-dropdown-item>
											<el-dropdown-item>弹性伸缩</el-dropdown-item>
											<el-dropdown-item @click="rollBack(scope.row)">回滚</el-dropdown-item>
											<el-dropdown-item @click="deleteCronJob(scope.row)">删除</el-dropdown-item>
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
			:resourceType="'cronjob'"
			@update="updateCronJob"
			v-if="state.dialogVisible"
		/>

		<CreateDialog
			v-model:dialogVisible="state.create.dialogVisible"
			:title="state.create.title"
			:cronjob="state.cronjob"
			@refresh="listCronJob()"
			v-if="state.create.dialogVisible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sCronJob">
import { reactive, onMounted, onBeforeUnmount, defineAsyncComponent, h } from 'vue';
import { Delete, Edit } from '@element-plus/icons-vue';
import { CaretBottom } from '@element-plus/icons-vue';
import { useCronJobApi } from '@/api/kubernetes/cronjob';
import { CronJob } from 'kubernetes-models/batch/v1';
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
	cronjob: {} as CronJob,
	codeData: {} as CronJob,
	searchType: '1',
	search: '',
	dialogVisible: false,
	visible: false,
	scaleDeploy: <CronJob>{},
	query: <queryType>{
		cloud: '',
		page: 1,
		limit: 10,
	},
	namespace: '',
	loading: false,
	cronjobs: [] as CronJob[],
	selectData: [] as CronJob[],
	total: 0,
	type: '1',
	inputValue: '',
});
const route = useRoute();
const cronjobApi = useCronJobApi();
const k8sStore = kubernetesInfo();
const socketApi = useWebsocketApi();

const ws = socketApi.createWebsocket('cronjob');
ws.onmessage = (e) => {
	if (e.data === 'ping') {
		return;
	} else {
		const object = JSON.parse(e.data);
		if (object.type === 'cronjob' && object.result.namespace === k8sStore.state.activeNamespace && object.cluster == k8sStore.state.activeCluster) {
			state.cronjobs = object.result.data;
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
	listCronJob();
	state.loading = false;
};

const depStatus = (cronjob: CronJob) => {
	let runStatus = '';
	if (cronjob.status) {
		cronjob.status.conditions?.forEach((item: CronJobCondition) => {
			if (item.type === 'Available' && item.status != 'True') {
				if (item.message) runStatus = item.message;
			}
		});
	}

	return runStatus;
};
const rollBack = (cronjob: CronJob) => {
	const reversion = cronjob.metadata?.annotations!['cronjob.kubernetes.io/revision']!;
	parseInt(reversion, 10);
	cronjobApi
		.rollBackCronJob(cronjob.metadata!.namespace!, cronjob.metadata!.name!, parseInt(reversion, 10) - 1 + '', {
			cloud: k8sStore.state.activeCluster,
		})
		.then(() => {
			ElMessage.success('回滚成功');
		})
		.catch((res: any) => {
			ElMessage.error('回滚失败,' + res.message);
		});
};
const reDeploy = (cronjob: CronJob) => {
	cronjobApi
		.reDeployCronJob(cronjob.metadata!.namespace!, cronjob.metadata!.name!, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('操作成功');
		})
		.catch((res: any) => {
			ElMessage.error(res.message);
		});
};

const handleChange = () => {
	listCronJob();
};
const updateCronJob = async (codeData: any) => {
	const updateData = YAML.load(codeData) as CronJob;
	delete updateData.status;
	delete updateData.metadata?.managedFields;
	await cronjobApi
		.updateCronJob(updateData, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('更新成功');
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	state.dialogVisible = false;
};

const deleteCronJobs = (depList: Array<CronJob>) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将批量删除多个 '),
			h('i', { style: 'color: teal' }, `CronJob副本`),
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
			depList.forEach((dep: CronJob) => {
				if (dep.metadata) {
					cronjobApi
						.deleteCronJob(dep.metadata?.namespace!, dep.metadata?.name!, { cloud: k8sStore.state.activeCluster })
						.then(() => {})
						.catch(() => {
							ElMessage.error(`删除${dep.metadata?.name}失败`);
						});
				}
			});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};
const deleteCronJob = (dep: CronJob) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${dep.metadata?.name}`),
			h('span', null, ' cronjob. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			cronjobApi.deleteCronJob(dep.metadata?.namespace!, dep.metadata?.name!, { cloud: k8sStore.state.activeCluster }).then((res: any) => {
				if (res.code === 200) {
					ElMessage.success('删除成功');
				} else {
					ElMessage.error(`删除${dep.metadata?.name}失败`);
				}
			});
		})
		.catch(() => {
			ElMessage.error('取消');
		});
};

const showYaml = async (cronjob: CronJob) => {
	const dep = deepClone(cronjob) as CronJob;
	delete dep.metadata?.managedFields;
	state.codeData = dep;
	state.dialogVisible = true;
};

const openScaleDialog = (dep: CronJob) => {
	state.scaleDeploy = deepClone(dep) as CronJob;
	state.visible = true;
};

const scaleCronJob = () => {
	cronjobApi
		.scaleCronJob(state.scaleDeploy.metadata?.namespace!, state.scaleDeploy.metadata?.name!, state.scaleDeploy.spec?.replicas!, {
			cloud: k8sStore.state.activeCluster,
		})
		.then((res: any) => {
			if (res.code === 200) {
				ElMessage.success('操作成功');
			} else {
				ElMessage.error('伸缩失败');
			}
		})
		.catch(() => {
			ElMessage.error('伸缩失败');
		});
	state.visible = false;
};

const handleSelectionChange = (value: any) => {
	state.selectData = value;
};

const listCronJob = async () => {
	state.namespace = k8sStore.state.activeNamespace;
	state.query.cloud = k8sStore.state.activeCluster;
	try {
		state.loading = true;
		await cronjobApi.listCronJob(k8sStore.state.activeNamespace, state.query).then((res) => {
			state.cronjobs = res.data.data;
			state.total = res.data.total;
		});
	} catch (e: any) {
		if (e.code != 5003) ElMessage.error(e.message);
	}
	state.loading = false;
};

mittBus.on('changeNamespace', () => {
	listCronJob();
});

const handlePageChange = (pageInfo: PageInfo) => {
	state.query.page = pageInfo.page;
	state.query.limit = pageInfo.limit;

	if (state.search != '') {
		search();
	} else {
		listCronJob();
	}
};
const deployDetail = async (dep: CronJob) => {
	k8sStore.state.activeCronJob = dep;
	await router.push({
		name: 'k8sCronJobDetail',
		params: {
			name: dep.metadata?.name,
		},
	});
};

const createCronJob = () => {
	state.create.title = '创建cronjob';
	state.create.dialogVisible = true;
};

const handleUpdate = (cronjob: CronJob) => {
	k8sStore.state.createCronJob.namespace = cronjob.metadata!.namespace!;
	const dep = deepClone(cronjob) as CronJob;
	delete dep.status;
	delete dep.metadata?.managedFields;
	state.cronjob = dep;
	state.create.title = '更新cronjob';
	state.create.dialogVisible = true;
};

onMounted(() => {
	listCronJob();
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
