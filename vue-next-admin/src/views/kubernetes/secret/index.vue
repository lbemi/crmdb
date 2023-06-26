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
				<el-button type="primary" size="small" class="ml10" @click="createSecret" :icon="Edit">创建</el-button>
				<el-button type="danger" size="small" class="ml10" :disabled="data.selectData.length == 0" @click="deleteSecret" :icon="Delete"
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
				:data="data.Secrets"
				@selection-change="handleSelectionChange"
				style="width: 100%"
				max-height="100vh - 235px"
				v-loading="data.loading"
			>
				<el-table-column type="selection" width="35" />
				<el-table-column label="名称">
					<template #default="scope">
						<el-button :size="theme.themeConfig.globalComponentSize" type="primary" text @click="secretDetail(scope.row)">
							{{ scope.row.metadata.name }}</el-button
						>
					</template>
				</el-table-column>
				<el-table-column prop="metadata.namespace" label="命名空间" />
				<el-table-column prop="type" label="类型">
					<template #default="scope">
						{{ secretType(scope.row) }}
					</template>
				</el-table-column>
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
						<el-button link type="primary" size="small" @click="secretDetail(scope.row)">详情</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="updateSecret(scope.row)">编辑</el-button><el-divider direction="vertical" />
						<el-button link type="primary" size="small" @click="showYaml(scope.row)">查看YAML</el-button><el-divider direction="vertical" />
						<el-button :disabled="scope.row.metadata.name === 'kubernetes'" link type="danger" size="small" @click="deleteSecret(scope.row)"
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
			:resourceType="'Secret'"
			@update="updateSecretYaml"
			v-if="data.dialogVisible"
		/>
		<DrawDialog
			v-model:visible="data.draw.visible"
			:secret="data.draw.secret"
			:title="data.draw.title"
			@refresh="listSecret"
			v-if="data.draw.visible"
		/>
		<SecretDetail v-model:visible="data.detail.visible" :configMap="data.detail.secret" :title="data.detail.title" v-if="data.detail.visible" />
	</div>
</template>

<script setup lang="ts" name="k8sSecret">
import { Secret } from 'kubernetes-types/core/v1';
import { defineAsyncComponent, h, onMounted, reactive, ref } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { ElMessage, ElMessageBox } from 'element-plus';
import mittBus from '@/utils/mitt';
import { useRoute } from 'vue-router';
import { dateStrFormat } from '@/utils/formatTime';
import { PageInfo } from '@/types/kubernetes/common';
import { Edit, Delete, List } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';
import { useSecretApi } from '@/api/kubernetes/secret';

const Pagination = defineAsyncComponent(() => import('@/components/pagination/pagination.vue'));
const YamlDialog = defineAsyncComponent(() => import('@/components/yaml/index.vue'));
const DrawDialog = defineAsyncComponent(() => import('./component/draw.vue'));
const SecretDetail = defineAsyncComponent(() => import('./component/detail.vue'));

type queryType = {
	key: string;
	page: number;
	limit: number;
	cloud: string;
	name?: string;
	label?: string;
};
const k8sStore = kubernetesInfo();
const SecretApi = useSecretApi();
const route = useRoute();
const theme = useThemeConfig();
const data = reactive({
	detail: {
		title: '',
		visible: false,
		secret: {} as Secret,
	},
	draw: {
		title: '',
		visible: false,
		secret: {} as Secret,
	},
	dialogVisible: false,
	codeData: {} as Secret,
	loading: false,
	selectData: [] as Secret[],
	Secrets: [] as Secret[],
	tmpSecret: [] as Secret[],
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
	listSecret();
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
	listSecret();
};
const handleChange = () => {
	listSecret();
};
const filterSecret = (Secrets: Array<Secret>) => {
	const SecretList = [] as Secret[];
	if (data.query.type === '1') {
		Secrets.forEach((Secret: Secret) => {
			if (Secret.metadata?.name?.includes(data.query.key)) {
				SecretList.push(Secret);
			}
		});
	} else {
		Secrets.forEach((Secret: Secret) => {
			if (Secret.metadata?.labels) {
				for (let k in Secret.metadata.labels) {
					if (k.includes(data.query.key) || Secret.metadata.labels[k].includes(data.query.key)) {
						SecretList.push(Secret);
						break;
					}
				}
			}
		});
	}
	data.Secrets = SecretList;
};
const createSecret = () => {
	data.draw.visible = true;
	data.draw.title = '创建Secret';
	data.draw.secret = {};
};
const deleteSecret = (Secret: Secret) => {
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将删除 '),
			h('i', { style: 'color: teal' }, `${Secret.metadata!.name}`),
			h('span', null, ' Secret. 是否继续? '),
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
			SecretApi.deleteSecret(Secret.metadata!.namespace!, Secret.metadata!.name!, { cloud: k8sStore.state.activeCluster })
				.then((res: any) => {
					listSecret();
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
const showYaml = (Secret: Secret) => {
	data.dialogVisible = true;
	delete Secret.metadata?.managedFields;
	data.codeData = Secret;
	// yamlRef.value.openDialog(Secret);
};
const updateSecretYaml = (code: any) => {
	console.log('更新Secret', code);
};

const handleSelectionChange = () => {};
const updateSecret = (secret: Secret) => {
	data.draw.visible = true;
	data.draw.title = '修改';
	data.draw.secret = secret;
};
const secretDetail = (secret: Secret) => {
	data.detail.visible = true;
	data.detail.title = '详情';
	data.detail.secret = secret;
};
const handlePageChange = (page: PageInfo) => {
	data.query.page = page.page;
	data.query.limit = page.limit;
	listSecret();
};
const listSecret = () => {
	data.loading = true;
	SecretApi.listSecret(k8sStore.state.activeNamespace, data.query)
		.then((res: any) => {
			data.Secrets = res.data.data;
			data.tmpSecret = res.data.data;
			data.total = res.data.total;
		})
		.catch((e: any) => {
			if (e.code != 5003) ElMessage.error(e.message);
		});
	data.loading = false;
};
const secretType = (secret: Secret) => {
	switch (secret.type) {
		case 'Opaque':
			return 'Opaque';
		case 'kubernetes.io/tls':
			return 'TLS';
		case 'kubernetes.io/dockerconfigjson':
			return '镜像凭证';
		case 'kubernetes.io/basic-auth':
			return '用户密码认证';
		default:
			return secret.type;
	}
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
