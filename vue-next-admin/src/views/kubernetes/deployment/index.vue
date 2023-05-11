<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-row :gutter="24">
					<el-col :span="18" style="display: flex; justify-content;: center">
						<el-text class="mx-1" :size="theme.themeConfig.globalComponentSize">命名空间：</el-text>
						<el-select
							v-model="k8sStore.state.activeNamespace"
							style="max-width: 180px"
							class="m-2"
							placeholder="Select"
							:size="theme.themeConfig.globalComponentSize"
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
							v-model="data.search"
							placeholder="输入标签或者名称"
							:size="theme.themeConfig.globalComponentSize"
							clearable
							@change="search"
							style="width: 250px; margin-left: 10px"
						>
							<template #prepend>
								<el-select v-model="data.searchType" placeholder="输入标签或者名称" style="width: 80px" :size="theme.themeConfig.globalComponentSize">
									<el-option label="标签" value="0" :size="theme.themeConfig.globalComponentSize" />
									<el-option label="名称" value="1" :size="theme.themeConfig.globalComponentSize" />
								</el-select>
							</template>
							<template #append>
								<el-button :size="theme.themeConfig.globalComponentSize" @click="search">
									<el-icon>
										<ele-Search />
									</el-icon>
									查询
								</el-button>
							</template>
						</el-input>

						<el-button type="primary" :size="theme.themeConfig.globalComponentSize" class="ml10" @click="createDeployment" :icon="Edit"
							>创建</el-button
						>
						<el-button
							type="danger"
							:size="theme.themeConfig.globalComponentSize"
							class="ml10"
							:disabled="data.selectData.length == 0"
							@click="deleteDeployments(data.selectData)"
							:icon="Delete"
							>批量删除</el-button
						>
						<el-button type="success" :size="theme.themeConfig.globalComponentSize" @click="refreshCurrentTagsView" style="margin-left: 10px">
							<el-icon>
								<ele-RefreshRight />
							</el-icon>
							刷新
						</el-button>
					</el-col>
					<el-col :span="6" style="display: flex"> </el-col>
				</el-row>
			</div>
			<el-table
				:data="data.deployments"
				style="width: 100%"
				@selection-change="handleSelectionChange"
				v-loading="data.loading"
				:size="theme.themeConfig.globalComponentSize"
			>
				<el-table-column type="selection" width="55" />

				<el-table-column prop="metadata.name" label="名称">
					<template #default="scope">
						<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="deployDetail(scope.row)">
							{{ scope.row.metadata.name }}</el-button
						>
					</template>
				</el-table-column>
				<el-table-column label="状态" width="90px">
					<template #default="scope">
						<el-button
							v-if="scope.row.status.availableReplicas && scope.row.status.availableReplicas == scope.row.status.replicas"
							type="success"
							:icon="Check"
							size="small"
							circle
						/>
						<el-button v-else type="danger" :icon="Close" size="small" circle />
					</template>
				</el-table-column>
				<el-table-column prop="spec.replicas" label="Pods" width="150px" align="center">
					<template #header> <span>Pods</span><br /><span style="font-size: 10px; font-weight: 50">就绪/副本/失败</span> </template>

					<template #default="scope">
						<a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a
						>/ <a style="color: green">{{ scope.row.status.replicas || '0' }}</a
						>/
						<a style="color: red">{{ scope.row.status.unavailableReplicas || '0' }}</a>
					</template>
				</el-table-column>
				<el-table-column label="镜像" show-overflow-tooltip>
					<template #default="scope">
						<el-text
							:size="theme.themeConfig.globalComponentSize"
							truncated
							type=""
							v-for="(item, index) in scope.row.spec.template.spec.containers"
							:key="index"
							>{{ item.image.split('@')[0] }}</el-text
						>
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
										:size="theme.themeConfig.globalComponentSize"
									>
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<el-tag
								type=""
								effect="plain"
								v-for="(item, key, index) in scope.row.metadata.labels"
								:key="index"
								:size="theme.themeConfig.globalComponentSize"
							>
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
				<el-table-column fixed="right" label="操作">
					<template #default="scope">
						<div style="display: flex; align-items: center">
							<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="deployDetail(scope.row)">详情</el-button>
							<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="deployDetail(scope.row)">编辑</el-button>
							<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="openScaleDialog(scope.row)">伸缩</el-button>
							<el-button link type="primary" :size="theme.themeConfig.globalComponentSize" @click="deployDetail(scope.row)">监控</el-button>
							<el-divider direction="vertical" />
							<el-dropdown :size="theme.themeConfig.globalComponentSize">
								<span class="el-dropdown-link" style="font-size: 12px">
									更多<el-icon class="el-icon--right"><CaretBottom /></el-icon>
								</span>
								<template #dropdown>
									<el-dropdown-menu>
										<el-dropdown-item @click="showYaml(scope.row)">查看Yaml</el-dropdown-item>
										<el-dropdown-item @click="reDeploy(scope.row)">重新部署</el-dropdown-item>
										<el-dropdown-item>编辑标签</el-dropdown-item>
										<el-dropdown-item>节点亲和性</el-dropdown-item>
										<el-dropdown-item>弹性伸缩</el-dropdown-item>
										<el-dropdown-item>调度容忍度</el-dropdown-item>
										<el-dropdown-item>升级策略</el-dropdown-item>
										<el-dropdown-item>复制创建</el-dropdown-item>
										<el-dropdown-item @click="rollBack(scope.row)">回滚</el-dropdown-item>
										<el-dropdown-item @click="deleteDeployment(scope.row)">删除</el-dropdown-item>
									</el-dropdown-menu>
								</template>
							</el-dropdown>
						</div>
					</template>
				</el-table-column>
			</el-table>
			<!-- 分页区域 -->
			<Pagination :total="data.total" @handlePageChange="handlePageChange" />
		</el-card>
		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:code-data="data.codeData"
			:resourceType="'deployment'"
			@update="updateDeployment"
			v-if="data.dialogVisible"
		/>
		<!-- <YamlDialog ref="yamlRef" :resourceType="'deployment'" :update-resource="updateDeployment" /> -->
		<el-dialog v-model="data.visible" width="300px" @close="data.visible = false">
			<template #header>
				<span style="font-size: 16px">{{ '伸缩: ' + data.scaleDeploy.metadata?.name }}</span>
			</template>
			<div style="text-align: center">
				<el-input-number v-model="data.scaleDeploy.spec!.replicas" :min="0" :max="1000" size="default" w />
			</div>
			<template #footer>
				<span class="dialog-footer">
					<el-button text @click="data.visible = false" size="default">取消</el-button>
					<el-button text type="primary" @click="scaleDeployment" size="default"> 确定 </el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="k8sDeployment">
import { reactive, onMounted, onBeforeUnmount, defineAsyncComponent, ref, computed, onUnmounted, h } from 'vue';
import { Check, Close, Delete, Edit, Search } from '@element-plus/icons-vue';
import { CaretBottom } from '@element-plus/icons-vue';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1Deployment } from '@kubernetes/client-node';
import { PageInfo } from '/@/types/kubernetes/common';
import { kubernetesInfo } from '/@/stores/kubernetes';
import router from '/@/router';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';
import YAML from 'js-yaml';
import mittBus from '/@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '/@/utils/formatTime';
import { deepClone, globalComponentSize } from '/@/utils/other';
import { useThemeConfig } from '/@/stores/themeConfig';
const theme = useThemeConfig();

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));
const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));

const data = reactive({
	codeData: {} as V1Deployment,
	searchType: '1',
	search: '',
	dialogVisible: false,
	visible: false,
	scaleDeploy: <V1Deployment>{},
	query: {
		cloud: '',
		page: 1,
		limit: 10,
		key: '',
		type: '0',
	},
	namespace: '',
	loading: false,
	deployments: [] as V1Deployment[],
	selectData: [] as V1Deployment[],
	total: 0,
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
			data.deployments = object.result.data;
		}
	}
};

const search = () => {
	data.loading = true;
	data.query.cloud = k8sStore.state.activeCluster;
	data.query.key = data.search;
	data.query.type = data.searchType;

	if (data.query.key != '') {
		deploymentApi.searchDeployment(k8sStore.state.activeNamespace, data.query).then((res) => {
			if (res.code == 200) {
				data.deployments = res.data.data;
				data.total = res.data.total;
			}
		});
	} else {
		listDeployment();
	}
	data.loading = false;
};

const rollBack = (deployment: V1Deployment) => {
	const reversion = deployment.metadata?.annotations!['deployment.kubernetes.io/revision']!;
	parseInt(reversion, 10);
	deploymentApi
		.rollBackDeployment(deployment.metadata!.namespace!, deployment.metadata!.name!, parseInt(reversion, 10) - 1 + '', {
			cloud: k8sStore.state.activeCluster,
		})
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('回滚成功');
			} else {
				ElMessage.error('回滚失败,' + res.message);
			}
		})
		.catch((res) => {
			ElMessage.error('回滚失败,' + res.message);
		});
};
const reDeploy = (deployment: V1Deployment) => {
	deploymentApi
		.reDeployDeployment(deployment.metadata!.namespace!, deployment.metadata!.name!, { cloud: k8sStore.state.activeCluster })
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('操作成功');
			} else {
				ElMessage.error(res.message);
			}
		})
		.catch((res: any) => {
			ElMessage.error(res.message);
		});
};

const handleChange = () => {
	listDeployment();
};
const updateDeployment = async (codeData: any) => {
	const updateData = YAML.load(codeData) as V1Deployment;
	delete updateData.status;
	delete updateData.metadata?.managedFields;
	await deploymentApi
		.updateDeployment(updateData, { cloud: k8sStore.state.activeCluster })
		.then((res) => {
			if (res.code == 200) {
				ElMessage.success('更新成功');
			} else {
				ElMessage.error(res.message);
			}
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
	data.dialogVisible = false;
};

const deleteDeployments = (depList: Array<V1Deployment>) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将批量删除多个 '),
			h('i', { style: 'color: teal' }, `Deplpoment副本`),
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
			depList.forEach((dep: V1Deployment) => {
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
const deleteDeployment = (dep: V1Deployment) => {
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

const showYaml = async (deployment: V1Deployment) => {
	const dep = deepClone(deployment);
	delete dep.metadata?.managedFields;
	data.codeData = dep;
	data.dialogVisible = true;
};
const openScaleDialog = (dep: V1Deployment) => {
	data.scaleDeploy = deepClone(dep);
	data.visible = true;
};

const scaleDeployment = () => {
	deploymentApi
		.scaleDeployment(data.scaleDeploy.metadata?.namespace!, data.scaleDeploy.metadata?.name!, data.scaleDeploy.spec?.replicas!, {
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
	data.visible = false;
};

const filterTableData = computed(() =>
	data.deployments.filter((item) => !data.search || item.metadata!.name!.toLowerCase().includes(data.search.toLowerCase()))
);

const handleSelectionChange = (value: any) => {
	data.selectData = value;
};

const listDeployment = async () => {
	data.namespace = k8sStore.state.activeNamespace;
	data.query.cloud = k8sStore.state.activeCluster;
	try {
		data.loading = true;
		await deploymentApi.listDeployment(k8sStore.state.activeNamespace, data.query).then((res) => {
			data.deployments = res.data.data;
			data.total = res.data.total;
		});
	} catch (e: any) {
		ElMessage.error(e);
	}
	data.loading = false;
};

mittBus.on('changeNamespace', () => {
	listDeployment();
});

const handlePageChange = (pageInfo: PageInfo) => {
	data.query.page = pageInfo.page;
	data.query.limit = pageInfo.limit;

	if (data.search != '') {
		search();
	} else {
		listDeployment();
	}
};
const deployDetail = async (dep: V1Deployment) => {
	k8sStore.state.activeDeployment = dep;
	router.push({
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
onMounted(() => {
	listDeployment();
});
onUnmounted(() => {
	mittBus.off('changeNamespace', () => {});
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
