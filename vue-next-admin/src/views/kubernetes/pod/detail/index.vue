<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<el-row :gutter="20">
				<el-col :span="16">
					<el-button type="info" :icon="ArrowLeft" text @click="backRoute">返回</el-button>
					<span style="font-weight: 35">{{ podStore.state.podDetail?.metadata?.name }}</span></el-col
				>
				<el-col :span="8"
					><el-button type="primary" size="small" :icon="Edit" @click="showYaml">编辑</el-button>
					<el-button type="primary" size="small" :icon="View" @click="showYaml">查看YAML</el-button>
					<el-button type="danger" size="small" :icon="Delete" @click="deletePod(podStore.state.podDetail)">删除</el-button>
					<el-button type="success" size="small" @click="podStore.refreshPodDetail" style="margin-left: 10px">
						<el-icon>
							<ele-RefreshRight />
						</el-icon>
						刷新
					</el-button>
				</el-col>
			</el-row>

			<el-descriptions :column="4" border class="desc-body">
				<el-descriptions-item label="名称" :span="2" label-align="right" label-class-name="my-label" class-name="my-content" width="150px">
					{{ podStore.state.podDetail?.metadata?.name }}
				</el-descriptions-item>
				<el-descriptions-item label="命名空间" label-align="right">{{ podStore.state.podDetail?.metadata?.namespace }}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right">{{
					dateStrFormat(podStore.state.podDetail.metadata!.creationTimestamp!.toString())
				}}</el-descriptions-item>
				<el-descriptions-item label="镜像" :span="2" label-align="right">
					<div v-for="(item, index) in podStore.state.podDetail?.spec?.containers" :key="index">
						<el-tag size="small" round effect="plain">{{ item.image?.split('@')[0] }}</el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="所在节点及IP" label-align="right">
					<div>{{ podStore.state.podDetail?.spec?.nodeName }}</div>
					<div>{{ podStore.state.podDetail?.status?.hostIP }}</div>
				</el-descriptions-item>
				<el-descriptions-item label="Pod IP" label-align="right">
					{{ podStore.state.podDetail?.status?.podIP }}
				</el-descriptions-item>
				<el-descriptions-item label="状态" :span="2" label-align="right" align="center">
					<a v-html="podStatus(podStore.state.podDetail?.status!)" />
					<el-link type="primary" :underline="false" @click="data.iShow = !data.iShow" style="font-size: 12px; margin-left: 5px"
						>展开现状详情<el-icon> <CaretBottom /> </el-icon
					></el-link>
				</el-descriptions-item>
				<el-descriptions-item label="重启次数" label-align="right">
					<template #default>
						<div v-if="podStore.state.podDetail?.status?.containerStatuses">{{ podRestart(podStore.state.podDetail?.status) }}</div>
					</template>
				</el-descriptions-item>
				<el-descriptions-item label="QoS类别" label-align="right">
					<template #default>
						{{ podStore.state.podDetail?.status?.qosClass }}
					</template>
				</el-descriptions-item>
			</el-descriptions>

			<div v-show="data.iShow">
				<el-divider />
				<el-table :data="podStore.state.podDetail.status?.conditions" stripe style="width: 100%">
					<el-table-column prop="type" label="类型" />
					<el-table-column prop="status" label="状态">
						<template #default="scope">
							<el-tag v-if="scope.row.status === 'True'" type="success"> 正常</el-tag>
							<el-tag type="danger" v-else> 异常</el-tag>
						</template>
					</el-table-column>
					<el-table-column prop="lastTransitionTime" label="更新时间">
						<template #default="scope">
							{{ dateStrFormat(scope.row.lastTransitionTime) }}
						</template>
					</el-table-column>
					<el-table-column prop="reason" label="内容" />
					<el-table-column prop="message" label="消息" />
				</el-table>
			</div>

			<!-- <el-divider /> -->
			<el-tabs v-model="data.activeName" class="demo-tabs" @tab-click="handleClick">
				<el-tab-pane label="容器" name="first">
					<el-space wrap v-if="data.containers">
						<el-card class="box-card" v-for="(container, index) in data.containers" :key="index">
							<template #header>
								<div style="display: flex; justify-content: space-between">
									<h3><SvgIcon name="iconfont icon-container-" class="svg" />{{ container.name }}</h3>
									<div>
										<el-button link type="primary" size="small" @click="jumpPodExec()">终端</el-button>
										<el-divider direction="vertical" />
										<el-button link type="primary" size="small" @click="jumpPodLog()">日志</el-button>
									</div>
								</div>
								<div>
									<div class="image">
										<el-text type="info" size="small" show-overflow-tooltip>{{ '镜像: ' + container.image }}</el-text>
									</div>
								</div>
							</template>
							<div v-if="data.containerStatuses">
								<div v-for="c in data.containerStatuses" :key="c.name">
									<div v-if="c.name == container.name">
										<h4>状态</h4>
										<div style="display: flex; justify-content: space-between; margin: 10px 0">
											<!--											<p v-if="!c.state?.running" v-html="podStatus(podStore.state.podDetail.status!)"></p>-->
											<el-text size="default" :type="c.ready || c.ready || c.state?.terminated?.reason === 'Completed' ? 'success' : 'danger'">
												<div v-if="c.state">
													<div v-if="c.state.waiting">{{ !c.ready ? c.state?.waiting?.reason : 'Running' }}</div>
													<div v-if="c.state?.terminated">{{ !c.ready ? c.state?.terminated?.reason : 'Running' }}</div>
													<div v-if="c.state?.running">{{ 'Running' }}</div>
												</div>

												<div v-if="c.lastState">
													<!--												<div v-if="item">{{ c.ready ? 'Running' : item!.reason }}</div>-->
													<div v-if="c.lastState.waiting">{{ !c.ready ? c.lastState?.waiting?.reason : 'Running' }}</div>
													<div v-if="c.lastState.terminated">{{ !c.ready ? c.lastState?.terminated?.reason : 'Running' }}</div>
													<div v-if="c.lastState.running">{{ 'Running' }}</div>
												</div>
											</el-text>
											<el-text size="default" v-if="data.containerStatuses" :type="c.started ? 'success' : 'danger'">
												{{ c.started ? 'Started' : 'NotStarted' }}
											</el-text>
											<el-text size="default" v-if="data.containerStatuses" :type="c.ready ? 'success' : 'danger'">
												{{ c.ready ? 'Ready' : 'NotReady' }}
											</el-text>
										</div>
										<div style="display: flex">
											<div>
												<h4>重启次数</h4>
												<div style="margin: 10px 0">
													<el-text size="large">
														<div v-if="data.containerStatuses">{{ c.restartCount }}</div>
														<div v-else>0</div>
													</el-text>
												</div>
											</div>
											<div style="margin-left: 120px">
												<h4>上次重启时间</h4>
												<div style="margin: 10px 0">
													<el-text size="large" v-if="!isObjectValueEqual(c.lastState, {})">
														<div v-if="c.lastState?.terminated">
															{{ dateStrFormat(c.lastState!.terminated!.startedAt!) }}
														</div>
														<!-- {{ c.lastState?.terminated?.startedAt || '-' }} -->
													</el-text>
													<el-text size="large" v-else> - </el-text>
												</div>
											</div>
										</div>

										<h4>错误原因</h4>
										<div style="margin: 10px 0">
											<div style="color: red" v-if="!isObjectValueEqual(c.lastState, {}) && c.lastState?.terminated?.message">
												<div style="color: red">
													<p>
														{{ c.lastState?.terminated?.message || '-' }}
													</p>
												</div>
											</div>
											<div v-else-if="c.state">
												<div style="color: red">
													{{ c.state?.waiting?.message || '-' }}
												</div>
											</div>
											<div v-else>-</div>
										</div>
										<div style="display: flex; margin-top: 5px">
											<!-- <div>CPU资源</div> -->
											<div style="">
												<h4>CPU资源</h4>
												<div style="margin: 10px 15px">
													<div>
														<el-text size="small" v-if="container.resources"> CPU需求: {{ container.resources?.requests?.cpu || '-' }} </el-text>
													</div>
													<div>
														<el-text size="small" v-if="container.resources"> CPU限制: {{ container.resources?.limits?.cpu || '-' }} </el-text>
													</div>
												</div>
											</div>
											<div style="margin-left: 85px">
												<h4>Memory资源</h4>
												<div style="margin: 10px 15px">
													<div>
														<el-text size="small" v-if="container.resources">
															Memory需求: {{ container.resources?.requests?.memory || '-' }}
														</el-text>
													</div>
													<div>
														<el-text size="small" v-if="container.resources"> Memory限制: {{ container.resources?.limits?.memory || '-' }} </el-text>
													</div>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
							<!--							<div v-else>-->
							<!--								<div v-for="c in podStore.state.podDetail.status?.conditions" :key="c.type">-->
							<!--									<h4>状态</h4>-->
							<!--									<div style="display: flex; justify-content: space-between; margin: 10px 0">-->
							<!--										<el-text size="default" type="danger">-->
							<!--											<div>{{ c.reason }}</div>-->
							<!--										</el-text>-->
							<!--										<el-text type="danger">-->
							<!--											{{ 'NotStarted' }}-->
							<!--										</el-text>-->
							<!--										<el-text type="danger">-->
							<!--											{{ 'NotReady' }}-->
							<!--										</el-text>-->
							<!--									</div>-->
							<!--									<div style="display: flex">-->
							<!--										<div>-->
							<!--											<h4>重启次数</h4>-->
							<!--											<div style="margin: 10px 0">-->
							<!--												<el-text size="large">-->
							<!--													<div v-if="data.containerStatuses">{{ c.restartCount || '0' }}</div>-->
							<!--													<div v-else>0</div>-->
							<!--												</el-text>-->
							<!--											</div>-->
							<!--										</div>-->
							<!--										<div style="margin-left: 120px">-->
							<!--											<h4>上次重启时间</h4>-->
							<!--											<div style="margin: 10px 0">-->
							<!--												<el-text size="large"> - </el-text>-->
							<!--											</div>-->
							<!--										</div>-->
							<!--									</div>-->
							<!--									<div>-->
							<!--										<h4>错误原因</h4>-->
							<!--										<div style="margin: 10px 0">-->
							<!--											<div style="color: red">-->
							<!--												{{ c.message }}-->
							<!--											</div>-->
							<!--										</div>-->
							<!--										<div style="display: flex; margin-top: 5px">-->
							<!--											&lt;!&ndash; <div>CPU资源</div> &ndash;&gt;-->
							<!--											<div style="">-->
							<!--												<h4>CPU资源</h4>-->
							<!--												<div style="margin: 10px 15px">-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources"> CPU需求: {{ container.resources?.requests?.cpu || '-' }} </el-text>-->
							<!--													</div>-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources"> CPU限制: {{ container.resources?.limits?.cpu || '-' }} </el-text>-->
							<!--													</div>-->
							<!--												</div>-->
							<!--											</div>-->
							<!--											<div style="margin-left: 85px">-->
							<!--												<h4>Memory资源</h4>-->
							<!--												<div style="margin: 10px 15px">-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources">-->
							<!--															Memory需求: {{ container.resources?.requests?.memory || '-' }}-->
							<!--														</el-text>-->
							<!--													</div>-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources"> Memory限制: {{ container.resources?.limits?.memory || '-' }} </el-text>-->
							<!--													</div>-->
							<!--												</div>-->
							<!--											</div>-->
							<!--										</div>-->
							<!--									</div>-->
							<!--								</div>-->
							<!--							</div>-->
						</el-card>
						<!-- init容器 -->
						<el-card class="box-card" v-for="(container, index) in data.intiContainers" :key="index">
							<template #header>
								<div style="display: flex; justify-content: space-between">
									<div>
										<h3>
											<SvgIcon name="iconfont icon-container-" class="svg" />{{ container.name
											}}<el-text style="margin-left: 5px" size="small" class="mx-1" type="danger">初始化容器</el-text>
										</h3>
									</div>
									<div>
										<el-button link type="primary" size="small" @click="jumpPodLog()">日志</el-button>
									</div>
								</div>
								<div class="image">
									<el-text type="info" size="small" show-overflow-tooltip>{{ '镜像: ' + container.image }}</el-text>
								</div>
							</template>
							<div v-if="data.initContainerStatus">
								<div v-for="c in data.initContainerStatus" :key="c.name">
									<div v-if="c.name === container.name">
										<h4>状态</h4>
										<div style="display: flex; justify-content: space-between; margin: 10px 0">
											<!--										<p v-html="podStatus(podStore.state.podDetail.status!)" v-if="!data.initContainerStatus"></p>-->
											<el-text size="default" :type="c.ready || c.state?.terminated?.reason === 'Completed' ? 'success' : 'danger'">
												<div v-if="c.state">
													<div v-if="c.state.waiting">{{ !c.ready ? c.state?.waiting?.reason : 'Running' }}</div>
													<div v-if="c.ready && c.state?.terminated">{{ c.state?.terminated?.reason }}</div>
													<div v-if="c.state?.running">{{ 'Running' }}</div>
												</div>

												<div v-if="c.lastState">
													<!--												<div v-if="item">{{ c.ready ? 'Running' : item!.reason }}</div>-->
													<div v-if="c.lastState.waiting">{{ !c.ready ? c.lastState?.waiting?.reason : 'Running' }}</div>
													<div v-if="c.lastState.terminated">{{ !c.ready ? c.lastState?.terminated?.reason : 'Running' }}</div>
													<div v-if="c.lastState.running">{{ 'Running' }}</div>
												</div>
											</el-text>
											<el-text size="default" v-if="data.containerStatuses" :type="c.started ? 'success' : 'danger'">
												{{ c.started ? 'Started' : 'NotStarted' }}
											</el-text>
											<el-text size="default" v-if="data.containerStatuses" :type="c.ready ? 'success' : 'warning'">
												{{ c.ready ? 'Ready' : 'NotReady' }}
											</el-text>
										</div>
										<div style="display: flex">
											<div>
												<h4>重启次数</h4>
												<div style="margin: 10px 0">
													<el-text size="large">
														<div v-if="data.initContainerStatus">{{ c.restartCount }}</div>
														<div v-else>0</div>
													</el-text>
												</div>
											</div>
											<div style="margin-left: 120px">
												<h4>上次重启时间</h4>
												<div style="margin: 10px 0">
													<el-text size="large" v-if="!isObjectValueEqual(c.lastState, {})">
														<div v-if="c.lastState?.terminated">
															{{ dateStrFormat(c.lastState!.terminated!.startedAt!) }}
														</div>
														<!-- {{ c.lastState?.terminated?.startedAt || '-' }} -->
													</el-text>
													<el-text size="large" v-else> - </el-text>
												</div>
											</div>
										</div>

										<h4>异常原因</h4>
										<div style="margin: 10px 0">
											<div style="color: #ff0000" v-if="!isObjectValueEqual(c.lastState, {}) && c.lastState?.terminated?.message">
												<div style="color: red">
													<p>
														{{ c.lastState?.terminated?.message || '-' }}
													</p>
												</div>
											</div>
											<div v-else-if="c.state">
												<div style="color: red">
													{{ c.state?.waiting?.message || '-' }}
												</div>
											</div>
											<div v-else>-</div>
										</div>
										<div style="display: flex; margin-top: 5px">
											<!-- <div>CPU资源</div> -->
											<div style="">
												<h4>CPU资源</h4>
												<div style="margin: 10px 15px">
													<div>
														<el-text size="small" v-if="container.resources"> CPU需求: {{ container.resources?.requests?.cpu || '-' }} </el-text>
													</div>
													<div>
														<el-text size="small" v-if="container.resources"> CPU限制: {{ container.resources?.limits?.cpu || '-' }} </el-text>
													</div>
												</div>
											</div>
											<div style="margin-left: 85px">
												<h4>Memory资源</h4>
												<div style="margin: 10px 15px">
													<div>
														<el-text size="small" v-if="container.resources">
															Memory需求: {{ container.resources?.requests?.memory || '-' }}
														</el-text>
													</div>
													<div>
														<el-text size="small" v-if="container.resources"> Memory限制: {{ container.resources?.limits?.memory || '-' }} </el-text>
													</div>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
							<!--							<div v-else>-->
							<!--								<div v-for="c in podStore.state.podDetail.status?.conditions" :key="c.type">-->
							<!--									<h4>状态</h4>-->
							<!--									<div style="display: flex; justify-content: space-between; margin: 10px 0">-->
							<!--										<el-text size="default" type="danger">-->
							<!--											<div>{{ c.reason }}</div>-->
							<!--										</el-text>-->
							<!--										<el-text type="danger">-->
							<!--											{{ 'NotStarted' }}-->
							<!--										</el-text>-->
							<!--										<el-text type="danger">-->
							<!--											{{ 'NotReady' }}-->
							<!--										</el-text>-->
							<!--									</div>-->
							<!--									<div style="display: flex">-->
							<!--										<div>-->
							<!--											<h4>重启次数</h4>-->
							<!--											<div style="margin: 10px 0">-->
							<!--												<el-text size="large">-->
							<!--													<div v-if="data.containerStatuses">{{ c.restartCount }}</div>-->
							<!--													<div v-else>0</div>-->
							<!--												</el-text>-->
							<!--											</div>-->
							<!--										</div>-->
							<!--										<div style="margin-left: 120px">-->
							<!--											<h4>上次重启时间</h4>-->
							<!--											<div style="margin: 10px 0">-->
							<!--												<el-text size="large"> - </el-text>-->
							<!--											</div>-->
							<!--										</div>-->
							<!--									</div>-->
							<!--									<div>-->
							<!--										<h4>错误原因</h4>-->
							<!--										<div style="margin: 10px 0">-->
							<!--											<div style="color: red">-->
							<!--												{{ c.message }}-->
							<!--											</div>-->
							<!--										</div>-->
							<!--										<div style="display: flex; margin-top: 5px">-->
							<!--											&lt;!&ndash; <div>CPU资源</div> &ndash;&gt;-->
							<!--											<div style="">-->
							<!--												<h4>CPU资源</h4>-->
							<!--												<div style="margin: 10px 15px">-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources"> CPU需求: {{ container.resources?.requests?.cpu || '-' }} </el-text>-->
							<!--													</div>-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources"> CPU限制: {{ container.resources?.limits?.cpu || '-' }} </el-text>-->
							<!--													</div>-->
							<!--												</div>-->
							<!--											</div>-->
							<!--											<div style="margin-left: 85px">-->
							<!--												<h4>Memory资源</h4>-->
							<!--												<div style="margin: 10px 15px">-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources">-->
							<!--															Memory需求: {{ container.resources?.requests?.memory || '-' }}-->
							<!--														</el-text>-->
							<!--													</div>-->
							<!--													<div>-->
							<!--														<el-text size="small" v-if="container.resources"> Memory限制: {{ container.resources?.limits?.memory || '-' }} </el-text>-->
							<!--													</div>-->
							<!--												</div>-->
							<!--											</div>-->
							<!--										</div>-->
							<!--									</div>-->
							<!--								</div>-->
							<!--							</div>-->
						</el-card>
					</el-space>
				</el-tab-pane>
				<el-tab-pane label="元数据" name="second">
					<MetaDetail :metaData="podStore.state.podDetail!.metadata" />
				</el-tab-pane>
				<el-tab-pane label="环境变量" name="third">
					<el-descriptions :column="1" direction="vertical">
						<el-descriptions-item :label="'容器: ' + item.name" v-for="item in podStore.state.podDetail?.spec?.containers" :key="item.name">
							<el-card class="card" :body-style="{ height: '200px' }">
								<div v-if="item.env">
									<div v-for="(value, index) in item.env" :key="index" style="margin-bottom: 8px">
										<el-tag type="info" size="default"> {{ value }} </el-tag>
									</div>
								</div>
								<div v-else>无数据</div>
							</el-card>
						</el-descriptions-item>
					</el-descriptions>
				</el-tab-pane>
				<el-tab-pane label="Init容器" name="fourth">
					<el-table :data="podStore.state.podDetail.status?.initContainerStatuses" style="width: 100%">
						<el-table-column label="名称">
							<template #default="scope">
								<el-text :type="scope.row.ready === true ? 'success' : 'danger'">{{ scope.row.name }}</el-text>
							</template>
						</el-table-column>
						<el-table-column label="状态">
							<template #default="scope">
								<el-text :type="scope.row.ready === true ? 'success' : 'danger'" v-for="(item, key) in scope.row.state" :key="key">
									<div>{{ item.reason }}</div>
									<div style="font-size: 10px">
										{{ item.message }}
									</div>
								</el-text>
							</template>
						</el-table-column>
						<el-table-column label="镜像" prop="image" />

						<el-table-column label="重启次数" prop="restartCount" />
					</el-table>
				</el-tab-pane>
				<el-tab-pane label="监控" name="five">监控</el-tab-pane>
				<el-tab-pane label="事件" name="six">
					<el-alert title="资源事件只保存最近1小时内发生的事件" :closable="false" type="info" class="mb15" show-icon />
					<el-table :data="data.events" max-height="500px">
						<el-table-column prop="type" label="类型" width="100px">
							<template #default="scope">
								<el-button link type="primary">{{ scope.row.type }}</el-button>
							</template>
						</el-table-column>
						<el-table-column label="原因">
							<template #default="scope">
								{{ scope.row.reason }}
							</template>
						</el-table-column>
						<el-table-column label="来源">
							<template #default="scope">
								{{ scope.row.source.component }}
							</template>
						</el-table-column>
						<el-table-column prop="spec.nodeName" label="消息">
							<template #default="scope">
								{{ scope.row.message }}
							</template>
						</el-table-column>
						<el-table-column label="时间" width="180px">
							<template #default="scope">
								{{ dateStrFormat(scope.row.firstTimestamp) }}
							</template>
						</el-table-column>
					</el-table>
				</el-tab-pane>
			</el-tabs>
		</el-card>
		<YamlDialog
			v-model:dialogVisible="data.dialogVisible"
			:disabled-update="true"
			:code-data="data.codeData"
			@update="updatePod"
			v-if="data.dialogVisible"
		/>
	</div>
</template>
<script lang="ts" setup name="podDetail">
import { reactive, defineAsyncComponent, h } from 'vue';
import { ArrowLeft, CaretBottom, Edit, View, Delete } from '@element-plus/icons-vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useDeploymentApi } from '@/api/kubernetes/deployment';
import { ContainerStatus, Pod, PodCondition, PodStatus } from 'kubernetes-models/v1';
import { ReplicaSet, ReplicaSetCondition } from 'kubernetes-models/apps/v1';
import router from '@/router';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { ElMessage, ElMessageBox, TabsPaneContext } from 'element-plus';
import { usePodApi } from '@/api/kubernetes/pod';
import { podInfo } from '@/stores/pod';
import { dateStrFormat } from '@/utils/formatTime';
import { isObjectValueEqual } from '@/utils/arrayOperation';

const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const MetaDetail = defineAsyncComponent(() => import('@/components/kubernetes/metaDetail.vue'));
// ContainerStateTerminated
const route = useRoute();
const podStore = podInfo();
const k8sStore = kubernetesInfo();
const podApi = usePodApi();
const deploymentApi = useDeploymentApi();
const data = reactive({
	dialogVisible: false,
	codeData: {} as Pod,
	param: {
		cloud: k8sStore.state.activeCluster,
	},
	replicaSets: [] as ReplicaSet[],
	pods: [] as Pod[],
	iShow: false,
	activeName: 'first',
	deployment: [],
	events: [] as ReplicaSetCondition[],
	containers: podStore.state.podDetail.spec?.containers,
	intiContainers: podStore.state.podDetail.spec?.initContainers,
	containerStatuses: podStore.state.podDetail.status?.containerStatuses,
	initContainerStatus: podStore.state.podDetail.status?.initContainerStatuses,
});

const handleClick = (tab: TabsPaneContext) => {
	if (tab.paneName === 'six') {
		getEvents();
	}
};

const podRestart = (status: PodStatus) => {
	let count = 0;
	status.containerStatuses!.forEach((item) => {
		count += item.restartCount;
	});
	return count;
};
// FIXME
const podStatus = (status: PodStatus) => {
	let s = '<span style="color: green">Running</span>';
	if (status.phase === 'Running') {
		status.conditions!.forEach((item: PodCondition) => {
			if (item.status != 'True') {
				let res = '';
				status.containerStatuses?.forEach((c: ContainerStatus) => {
					if (!c.ready) {
						if (c.state?.waiting) {
							res = ` <div>${c.state.waiting.reason}</div> <div style="font-size: 10px">${c.state.waiting.message}</div>`;
							// res = `${c.state.waiting.reason}`;
						}
						if (c.state?.terminated) {
							res = `${c.state.terminated.reason}`;
							// res = 'Terminating';
						}
					}
				});
				return (s = `<span style="color: red">${res}</span>`);
			}
		});
	} else if (status.phase === 'Succeeded') {
		let res = '';
		status.containerStatuses?.forEach((c: ContainerStatus) => {
			if (!c.ready) {
				if (c.state?.terminated) {
					res = `${c.state.terminated.reason}`;
					// res = 'Terminating';
				}
			}
		});
		return (s = `<span style="color: #E6A23C">${res}</span>`);
	} else {
		let res = status.phase;
		status.containerStatuses?.forEach((c: ContainerStatus) => {
			if (!c.ready) {
				if (c.state?.waiting) {
					res = ` </div> <div>${c.state.waiting.reason}</div>`;
					// res = `${c.state.waiting.reason}`;
				}
				if (c.state?.terminated) {
					res = `${c.state.terminated.reason}`;
					// res = 'Terminating';
				}
			}
		});
		return (s = `<span style="color: red">${res}</span>`);
	}

	return s;
};

const updatePod = () => {
	// TODO 完善功能
	ElMessage.success('更新成功');
	// const updateData = YAML.load(yamlRef.value.code) as Deployment;
	// delete updateData.status;
	// delete updateData.metadata?.managedFields;
	// deploymentApi
	// 	.updateDeployment(updateData, { cloud: k8sStore.state.activeCluster })
	// 	.then((res) => {
	// 		if (res.code == 200) {
	// 			ElMessage.success('更新成功');
	// 		} else {
	// 			ElMessage.error(res.message);
	// 		}
	// 	})
	// 	.catch((e) => {
	// 		ElMessage.error(e.message);
	// 	});
	data.dialogVisible = false;
};

const getPods = async () => {
	const res = await deploymentApi.detailDeployment(
		k8sStore.state.activeDeployment.metadata!.namespace!.toString(),
		k8sStore.state.activeDeployment?.metadata!.name!.toString(),
		data.param
	);
	data.pods = res.data.pods;
	data.replicaSets = res.data.replicaSets;
};

const getEvents = async () => {
	const pod = podStore.state.podDetail;
	const res = await podApi.podEvents(pod.metadata!.namespace, pod.metadata!.name, data.param);
	data.events = res.data;
};
const jumpPodExec = () => {
	podStore.state.podShell = podStore.state.podDetail;
	router.push({
		name: 'podShell',
	});
};
const jumpPodLog = () => {
	podStore.state.podShell = podStore.state.podDetail;
	router.push({
		name: 'podLog',
	});
};
const backRoute = () => {
	mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
	router.push({
		name: 'k8sPod',
	});
};
const deletePod = async (pod: Pod) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${pod.metadata?.name}`),
			h('span', null, ' 容器. 是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(() => {
			podApi.deletePod(pod.metadata?.namespace, pod.metadata?.name, data.param);
			getPods();
			ElMessage({
				type: 'success',
				message: `${pod.metadata?.name}` + ' 已删除',
			});
			backRoute();
		})
		.catch(); // 取消
};

const showYaml = async () => {
	delete podStore.state.podDetail.metadata?.managedFields;
	data.codeData = podStore.state.podDetail;
	data.dialogVisible = true;
};
</script>
<style lang="scss">
.card {
	overflow-y: auto;
	/* 开启滚动显示溢出内容 */
}

.test {
	backdrop-filter: blur(16px) saturate(180%);
	-webkit-backdrop-filter: blur(16px) saturate(180%);
	background-color: rgba(214, 221, 210, 0.75);
	border-radius: 12px;
	border: 1px solid rgba(209, 213, 219, 0.3);
	// background-image: url('https://images.unsplash.com/photo-1519681393784-d120267933ba?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1124&q=100');
	// background-position: center;
}

.tag-center {
	display: flex;
	flex-direction: column;
	align-items: center;

	.el-tag {
		margin-bottom: 3px;
		white-space: normal;
		height: auto;
	}
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

.card-header {
	align-items: center;
}

.text {
	font-size: 14px;
}

.item {
	margin-bottom: 18px;
}

.box-card {
	width: 430px;
	margin-bottom: 10px;
	height: 400px;

	.card-header {
		height: 55px;
	}
}

.svg {
	margin-right: 5px;
}

.image {
	margin-top: 15px;
	margin-left: 5px;
}
</style>
