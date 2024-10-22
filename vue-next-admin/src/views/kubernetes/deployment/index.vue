<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-row :gutter="24">
					<el-col :span="21" style="display: flex">
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

						<el-button v-auth="'k8s:deployment:add'" type="primary" :size="state.size" class="ml10" @click="createDeployment" :icon="Edit"
							>创建</el-button
						>
						<el-button
							v-auth="'k8s:deployment:del'"
							type="danger"
							:size="state.size"
							class="ml10"
							:disabled="state.selectData.length == 0"
							@click="deleteDeployments(state.selectData)"
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
					<el-col :span="3" style="display: flex; align-items: center">
						自动刷新<el-switch v-model="state.autoRefresh" style="margin-left: 10px; margin-right: 10px" />
						<el-icon :class="{ 'is-loading': state.autoRefresh }"> <Refresh /> </el-icon>
					</el-col>
				</el-row>
			</div>
			<el-table :data="state.deployments" style="width: 100%" @selection-change="handleSelectionChange" v-loading="state.loading" :size="state.size">
				<el-table-column type="selection" width="55" />
				<el-table-column prop="metadata.namespace" label="命名空间" width="200px" v-if="k8sStore.state.activeNamespace === 'all'" />
				<el-table-column prop="metadata.name" label="名称">
					<template #default="scope">
						<el-button link type="primary" :size="state.size" @click="deployDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
						<div style="color: red">{{ depStatus(scope.row) }}</div>
					</template>
				</el-table-column>
				<el-table-column label="状态" width="130px">
					<template #default="scope">
						<div
							v-if="
								scope.row.status.replicas > scope.row.status.availableReplicas ||
								(scope.row.status.replicas > scope.row.spec.replicas && scope.row.status.availableReplicas)
							"
							style="display: flex; align-items: center"
						>
							<div style="display: inline-block; width: 12px; height: 12px; background: #e6a23c; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #e6a23c">Updating </span>
						</div>
						<div
							v-else-if="
								scope.row.spec.replicas == 0 || (scope.row.status.availableReplicas && scope.row.status.availableReplicas == scope.row.spec.replicas)
							"
							style="display: flex; align-items: center"
						>
							<div style="display: inline-block; width: 12px; height: 12px; background: #388c04; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #388c04">Running </span>
						</div>
						<div v-else style="display: flex; align-items: center">
							<div style="display: inline-block; width: 12px; height: 12px; background: #f56c6c; border-radius: 50%"></div>
							<span style="margin-left: 5px; font-size: 12px; color: #f56c6c">Waiting </span>
						</div>
					</template>
				</el-table-column>

				<el-table-column prop="spec.replicas" label="Pods" align="center">
					<template #header> <span>Pods</span><br /><span class="table-header">就绪/副本/失败</span> </template>
					<template #default="scope">
						<a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a
						>/ <a style="color: green">{{ scope.row.spec.replicas || '0' }}</a
						>/
						<a style="color: red">{{ scope.row.status.unavailableReplicas || '0' }}</a>
					</template>
				</el-table-column>
				<el-table-column label="镜像" show-overflow-tooltip>
					<template #default="scope">
						<div class="flex">
							<SvgIcon :name="'iconfont icon-jingxiangbanben'" class="mr5" style="color: #409eff" />
							<el-text :size="state.size" truncated type="info" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{
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

				<el-table-column label="创建时间">
					<template #default="scope">
						{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
					</template>
				</el-table-column>
				<el-table-column fixed="right" label="操作" width="270px">
					<template #default="scope">
						<div style="display: flex; align-items: center">
							<el-button link type="primary" :size="state.size" @click="deployDetail(scope.row)">详情</el-button>
							<el-button v-auth="'k8s:deployment:edit'" link type="primary" :size="state.size" @click="handleUpdate(scope.row)">编辑</el-button>
							<el-button v-auth="'k8s:deployment:scale'" link type="primary" :size="state.size" @click="openScaleDialog(scope.row)">伸缩</el-button>
							<el-button link type="primary" :size="state.size" @click="deployDetail(scope.row)">监控</el-button>
							<el-divider direction="vertical" />
							<el-dropdown>
								<span class="el-dropdown-link">
									更多
									<el-icon>
										<arrow-down />
									</el-icon>
								</span>
								<template #dropdown>
									<el-dropdown-menu>
										<el-dropdown-item @click="showYaml(scope.row)">查看Yaml</el-dropdown-item>
										<el-dropdown-item @click="reDeploy(scope.row)">重新部署</el-dropdown-item>
										<el-dropdown-item>编辑标签</el-dropdown-item>
										<el-dropdown-item>弹性伸缩</el-dropdown-item>
										<el-dropdown-item @click="rollBack(scope.row)">回滚</el-dropdown-item>
										<el-dropdown-item @click="deleteDeployment(scope.row)">删除</el-dropdown-item>
									</el-dropdown-menu>
								</template>
							</el-dropdown>
							<!--							<el-dropdown :size="state.size" style="width: 40px">-->
							<!--								<span class="el-dropdown-link">-->
							<!--									更多-->
							<!--									<el-icon class="el-icon&#45;&#45;right">-->
							<!--										<CaretBottom />-->
							<!--									</el-icon>-->
							<!--								</span>-->
							<!--								<template #dropdown>-->
							<!--									<el-dropdown-menu>-->
							<!--										<el-dropdown-item @click="showYaml(scope.row)">查看Yaml</el-dropdown-item>-->
							<!--										<el-dropdown-item @click="reDeploy(scope.row)">重新部署</el-dropdown-item>-->
							<!--										<el-dropdown-item>编辑标签</el-dropdown-item>-->
							<!--										<el-dropdown-item>弹性伸缩</el-dropdown-item>-->
							<!--										<el-dropdown-item @click="rollBack(scope.row)">回滚</el-dropdown-item>-->
							<!--										<el-dropdown-item @click="deleteDeployment(scope.row)">删除</el-dropdown-item>-->
							<!--									</el-dropdown-menu>-->
							<!--								</template>-->
							<!--							</el-dropdown>-->
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
			:resourceType="'deployment'"
			@update="updateDeployment"
			v-if="state.dialogVisible"
		/>
		<!--    伸缩-->
		<el-dialog v-model="state.visible" width="300px" @close="state.visible = false">
			<template #header>
				<span style="font-size: 14px"
					>{{ '修改: ' }} <a style="color: teal"> {{ state.scaleDeploy.metadata?.name }}</a> {{ ' 副本' }}</span
				>
			</template>
			<div style="text-align: center">
				<el-input-number v-if="state.scaleDeploy.spec" v-model="state.scaleDeploy.spec.replicas" :min="0" :max="100" :size="state.size" />
			</div>
			<template #footer>
				<span class="dialog-footer">
					<el-button text @click="state.visible = false" :size="state.size">取消</el-button>
					<el-button text type="primary" @click="scaleDeployment" :size="state.size"> 确定 </el-button>
				</span>
			</template>
		</el-dialog>
		<CreateDialog
			v-model:dialogVisible="state.create.dialogVisible"
			:title="state.create.title"
			:deployment="state.deployment"
			@refresh="listDeployment"
			v-if="state.create.dialogVisible"
		/>
	</div>
</template>

<script setup lang="ts" name="k8sDeployment">
import { reactive, onMounted, onBeforeUnmount, defineAsyncComponent, h } from 'vue';
import { Delete, Edit, Refresh } from '@element-plus/icons-vue';
import { ArrowDown } from '@element-plus/icons-vue';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { Deployment, DeploymentCondition } from 'kubernetes-models/apps/v1';
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
	autoRefresh: true,
	size: theme.themeConfig.globalComponentSize,
	create: {
		dialogVisible: false,
		title: '',
	},
	deployment: {} as Deployment,
	codeData: {} as Deployment,
	searchType: '1',
	search: '',
	dialogVisible: false,
	visible: false,
	scaleDeploy: <Deployment>{},
	query: <queryType>{
		cloud: '',
		page: 1,
		limit: 10,
	},
	namespace: '',
	loading: false,
	deployments: [] as Deployment[],
	selectData: [] as Deployment[],
	total: 0,
	type: '1',
	inputValue: '',
});
const route = useRoute();
const deploymentApi = useDeploymentApi();
const k8sStore = kubernetesInfo();
const socketApi = useWebsocketApi();

const ws = socketApi.createWebsocket('deployment');
ws.onmessage = (e) => {
	if (e.data === 'ping') {
		return;
	} else {
		const object = JSON.parse(e.data);
		if (
			object.type === 'deployment' &&
			object.result.namespace === k8sStore.state.activeNamespace &&
			object.cluster == k8sStore.state.activeCluster
		) {
			state.deployments = object.result.data;
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
	listDeployment();
	state.loading = false;
};

const depStatus = (deployment: Deployment) => {
	let runStatus = '';
	if (deployment.status) {
		deployment.status.conditions?.forEach((item: DeploymentCondition) => {
			if (item.type === 'Available' && item.status != 'True') {
				if (item.message) runStatus = item.message;
			}
		});
	}

	return runStatus;
};
const rollBack = (deployment: Deployment) => {
	const reversion = deployment.metadata?.annotations!['deployment.kubernetes.io/revision']!;
	parseInt(reversion, 10);
	deploymentApi
		.rollBackDeployment(deployment.metadata!.namespace!, deployment.metadata!.name!, parseInt(reversion, 10) - 1 + '', {
			cloud: k8sStore.state.activeCluster,
		})
		.then(() => {
			ElMessage.success('回滚成功');
		})
		.catch((res: any) => {
			ElMessage.error('回滚失败,' + res.message);
		});
};
const reDeploy = (deployment: Deployment) => {
	deploymentApi
		.reDeployDeployment(deployment.metadata!.namespace!, deployment.metadata!.name!, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('操作成功');
		})
		.catch((res: any) => {
			ElMessage.error(res.message);
		});
};

const handleChange = () => {
	listDeployment();
};
const updateDeployment = async (codeData: any) => {
	const updateData = YAML.load(codeData) as Deployment;
	delete updateData.status;
	delete updateData.metadata?.managedFields;
	await deploymentApi
		.updateDeployment(updateData, { cloud: k8sStore.state.activeCluster })
		.then(() => {
			ElMessage.success('更新成功');
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	state.dialogVisible = false;
};

const deleteDeployments = (depList: Array<Deployment>) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将批量删除多个 '),
			h('i', { style: 'color: teal' }, `Deployment副本`),
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
			depList.forEach((dep: Deployment) => {
				if (dep.metadata) {
					deploymentApi
						.deleteDeployment(dep.metadata?.namespace!, dep.metadata?.name!, { cloud: k8sStore.state.activeCluster })
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
const deleteDeployment = (dep: Deployment) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${dep.metadata?.name}`),
			h('span', null, ' deployment. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			deploymentApi.deleteDeployment(dep.metadata?.namespace!, dep.metadata?.name!, { cloud: k8sStore.state.activeCluster }).then((res: any) => {
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

const showYaml = async (deployment: Deployment) => {
	const dep = deepClone(deployment) as Deployment;
	delete dep.metadata?.managedFields;
	state.codeData = dep;
	state.dialogVisible = true;
};

const openScaleDialog = (dep: Deployment) => {
	state.scaleDeploy = deepClone(dep) as Deployment;
	state.visible = true;
};

const scaleDeployment = () => {
	deploymentApi
		.scaleDeployment(state.scaleDeploy.metadata?.namespace!, state.scaleDeploy.metadata?.name!, state.scaleDeploy.spec?.replicas!, {
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

const listDeployment = async () => {
	state.namespace = k8sStore.state.activeNamespace;
	state.query.cloud = k8sStore.state.activeCluster;
	try {
		state.loading = true;
		await deploymentApi.listDeployment(k8sStore.state.activeNamespace, state.query).then((res) => {
			state.deployments = res.data.data;
			state.total = res.data.total;
		});
	} catch (e: any) {
		if (e.code != 5003) ElMessage.error(e.message);
	}
	state.loading = false;
};

mittBus.on('changeNamespace', () => {
	listDeployment();
});

const handlePageChange = (pageInfo: PageInfo) => {
	state.query.page = pageInfo.page;
	state.query.limit = pageInfo.limit;

	if (state.search != '') {
		search();
	} else {
		listDeployment();
	}
};
const deployDetail = async (dep: Deployment) => {
	k8sStore.state.activeDeployment = dep;
	await router.push({
		name: 'k8sDeploymentDetail',
		params: {
			name: dep.metadata?.name,
		},
	});
};

const createDeployment = () => {
	router.push({
		name: 'deploymentCreate',
	});
};

const handleUpdate = (deployment: Deployment) => {
	k8sStore.state.creatDeployment.namespace = deployment.metadata!.namespace!;
	const dep = JSON.parse(JSON.stringify(deployment));
	delete dep.status;
	delete dep.metadata?.managedFields;
	state.deployment = dep;

	state.create.title = '更新deployment';
	state.create.dialogVisible = true;
};

onMounted(() => {
	listDeployment();
});

const refreshCurrentTagsView = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 0, ...route }));
};

onBeforeUnmount(() => {
	ws.close();
});
</script>

<style scoped lang="scss">
.example-showcase .el-dropdown + .el-dropdown {
	margin-left: 15px;
}
.example-showcase .el-dropdown-link {
	cursor: pointer;
	color: var(--el-color-primary);
	display: flex;
	align-items: center;
}

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
