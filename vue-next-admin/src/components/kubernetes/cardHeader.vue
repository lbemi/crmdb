<template>
	<div class="mb15" style="display: flex; align-items: center">
		<el-text class="mx-1" :size="theme.themeConfig.globalComponentSize">命名空间：</el-text>
		<el-select
			v-model="k8sStore.state.activeNamespace"
			style="max-width: 180px"
			:size="theme.themeConfig.globalComponentSize"
			class="m-2"
			placeholder="Select"
			@change="prop.refresh"
			><el-option key="all" label="所有命名空间" value="all"></el-option>
			<el-option v-for="item in k8sStore.state.namespace" :key="item.metadata?.name" :label="item.metadata?.name" :value="item.metadata!.name!" />
		</el-select>
		<el-icon
			color="#409EFC"
			:size="theme.themeConfig.globalComponentSize"
			style="margin-left: 8px; margin-right: 10px"
			@click="k8sStore.listNamespace()"
			><Refresh
		/></el-icon>
		<el-input
			v-model="k8sStore.state.search.value"
			placeholder="输入标签或者名称"
			:size="theme.themeConfig.globalComponentSize"
			clearable
			@change="prop.search"
			style="width: 350px; margin-left: 10px"
		>
			<template #prepend>
				<el-select
					v-model="k8sStore.state.search.type"
					placeholder="输入标签或者名称"
					style="width: 80px"
					:size="theme.themeConfig.globalComponentSize"
				>
					<el-option label="标签" value="0" :size="theme.themeConfig.globalComponentSize" />
					<el-option label="名称" value="1" :size="theme.themeConfig.globalComponentSize" />
				</el-select>
			</template>
			<template #append>
				<el-button :size="theme.themeConfig.globalComponentSize" @click="prop.search">
					<el-icon>
						<ele-Search />
					</el-icon>
					查询
				</el-button>
			</template>
		</el-input>
		<el-button type="danger" :size="theme.themeConfig.globalComponentSize" class="ml10" :disabled="selectStatus" @click="prop.deleteFunc"
			>批量删除</el-button
		>
		<el-button type="success" :size="theme.themeConfig.globalComponentSize" @click="prop.refresh" style="margin-left: 10px">
			<el-icon>
				<ele-RefreshRight />
			</el-icon>
			刷新
		</el-button>
	</div>
</template>

<script setup lang="ts">
import { kubernetesInfo } from '@/stores/kubernetes';
import { Refresh } from '@element-plus/icons-vue';
import { useThemeConfig } from '@/stores/themeConfig';

const k8sStore = kubernetesInfo();
const theme = useThemeConfig();

const prop = defineProps({
	selectStatus: Boolean,
	refresh: Function,
	search: Function,
	deleteFunc: Function,
});
</script>

<style scoped></style>
