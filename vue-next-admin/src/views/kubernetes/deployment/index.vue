<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<div class="mb15">
				<el-row :gutter="20">
					<el-col :span="18">
						命名空间:
						<el-select
							v-model="k8sStore.state.activeNamespace"
							style="max-width: 180px"
							class="m-2"
							placeholder="Select"
							size="small"
							@change="handleChange"
							><el-option key="all" label="所有命名空间" value="all"></el-option>
							<el-option
								v-for="item in k8sStore.state.namespace"
								:key="item.metadata?.name"
								:label="item.metadata?.name"
								:value="item.metadata!.name!"
							/>
						</el-select>

						<el-button type="primary" size="small" class="ml10" @click="createDeployment">创建</el-button>
						<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" @click="deleteDeployments(data.selectData)"
							>批量删除</el-button
						>
					</el-col>
					<el-col :span="6">
						<el-input size="small" placeholder="请输入集群名称" style="max-width: 180px" v-model="data.search"> </el-input>
						<el-button size="small" type="primary" class="ml10" @click="search">
							<el-icon>
								<ele-Search />
							</el-icon>
							查询
						</el-button>
					</el-col>
				</el-row>
			</div>

			<el-table
				:data="data.deployments"
				style="width: 100%"
				@selection-change="handleSelectionChange"
				v-loading="data.loading"
				max-height="100vh - 235px"
			>
				<el-table-column type="selection" width="55" />

				<el-table-column prop="metadata.name" label="名称" width="220px">
					<template #default="scope">
						<el-button link type="primary" @click="deployDetail(scope.row)"> {{ scope.row.metadata.name }}</el-button>
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
						<el-text truncated type="" v-for="(item, index) in scope.row.spec.template.spec.containers" :key="index">{{
							item.image.split('@')[0]
						}}</el-text>
					</template>
				</el-table-column>

				<el-table-column label="标签" width="180px" show-overflow-tooltip>
					<template #default="scope">
						<el-tooltip placement="right" effect="light">
							<template #content>
								<div style="display: flex; flex-direction: column">
									<el-tag class="label" type="" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" effect="plain" size="small">
										{{ key }}:{{ item }}
									</el-tag>
								</div>
							</template>
							<el-tag type="" effect="plain" v-for="(item, key, index) in scope.row.metadata.labels" :key="index" size="small">
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
							<el-button link type="primary" size="default" @click="deployDetail(scope.row)">详情</el-button>
							<el-button link type="primary" size="default" @click="deployDetail(scope.row)">编辑</el-button>
							<el-button link type="primary" size="default" @click="openScaleDialog(scope.row)">伸缩</el-button>
							<el-button link type="primary" size="default" @click="deployDetail(scope.row)">监控</el-button>
							<el-divider direction="vertical" />
							<el-dropdown>
								<span class="el-dropdown-link">
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
										<el-dropdown-item>日志</el-dropdown-item>
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
		<YamlDialog ref="yamlRef" :resourceType="'deployment'" :update-resource="updateDeployment" />
		<el-dialog v-model="data.dialogVisible" width="300px" @close="data.dialogVisible = false">
			<template #header>
				<span style="font-size: 16px">{{ '伸缩: ' + data.scaleDeploy.metadata?.name }}</span>
			</template>
			<div style="text-align: center">
				<el-input-number v-model="data.scaleDeploy.spec!.replicas" :min="0" :max="1000" size="default" w />
			</div>
			<template #footer>
				<span class="dialog-footer">
					<el-button text @click="data.dialogVisible = false" size="default">取消</el-button>
					<el-button text type="primary" @click="scaleDeployment" size="default"> 确定 </el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="k8sDeployment">
import { reactive, onMounted, onBeforeUnmount, defineAsyncComponent, ref, computed } from 'vue';
import { Check, Close } from '@element-plus/icons-vue';
import { CaretBottom } from '@element-plus/icons-vue';
import { useDeploymentApi } from '/@/api/kubernetes/deployment';
import { V1Deployment } from '@kubernetes/client-node';
import { PageInfo } from '/@/types/kubernetes/common';
import { kubernetesInfo } from '/@/stores/kubernetes';
import router from '/@/router';
import { ElMessage } from 'element-plus';
import { useWebsocketApi } from '/@/api/kubernetes/websocket';
import YAML from 'js-yaml';
import { dataTool } from 'echarts';

const YamlDialog = defineAsyncComponent(() => import('/@/components/yaml/index.vue'));
const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));

const yamlRef = ref();
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
	data.query.cloud = k8sStore.state.activeCluster;
	data.query.key = data.search;
	deploymentApi.searchDeployment(k8sStore.state.activeNamespace, data.query).then((res) => {
		if (res.code == 200) {
			data.deployments = res.data.data;
			data.total = res.data.total;
		}
	});
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
const updateDeployment = () => {
	const updateData = YAML.load(yamlRef.value.code) as V1Deployment;
	delete updateData.status;
	delete updateData.metadata?.managedFields;
	deploymentApi
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
	yamlRef.value.handleClose();
};

const deleteDeployments = (depList: Array<V1Deployment>) => {
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

	ElMessage.success('删除成功');
};
const deleteDeployment = (dep: V1Deployment) => {
	deploymentApi
		.deleteDeployment(dep.metadata?.namespace!, dep.metadata?.name!, { cloud: k8sStore.state.activeCluster })
		.then((res: any) => {
			if (res.code === 200) {
				ElMessage.success('删除成功');
			} else {
				ElMessage.error(`删除${dep.metadata?.name}失败`);
			}
		})
		.catch(() => {
			ElMessage.error(`删除${dep.metadata?.name}失败`);
		});
};

const showYaml = async (deployment: V1Deployment) => {
	delete deployment.metadata?.managedFields;
	yamlRef.value.openDialog(deployment);
};
const openScaleDialog = (dep: V1Deployment) => {
	data.scaleDeploy = dep;
	data.dialogVisible = true;
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
	data.dialogVisible = false;
};

const filterTableData = computed(() =>
	data.deployments.filter((item) => !data.search || item.metadata!.name!.toLowerCase().includes(data.search.toLowerCase()))
);
const data = reactive({
	search: '',
	dialogVisible: false,
	scaleDeploy: <V1Deployment>{},
	query: {
		cloud: '',
		page: 1,
		limit: 10,
		key: '',
		type: 0,
	},
	namespace: '',
	loading: false,
	deployments: [] as V1Deployment[],
	selectData: [] as V1Deployment[],
	total: 0,
});
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
		ElMessage.error(e.data.message);
	}
	data.loading = false;
};
const handleChange = () => {
	// data.namespace = activeNamespace
	// setActiveNamespace(data.namespace)
	//   setActiveNamespace(data.namespace)
	listDeployment();
};
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
