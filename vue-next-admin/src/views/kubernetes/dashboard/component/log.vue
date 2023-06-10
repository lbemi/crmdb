/** * Created by lei on 2022/12/09 */
<template>
	<div>
		<el-table
			:data="state.events"
			:default-sort="{
				prop: 'metadata.creationTimestamp',
				order: 'descending',
			}"
			v-loading="state.loading"
		>
			<el-table-column prop="metadata.creationTimestamp" label="时间" width="180" sortable>
				<template #default="scope">
					{{ dateStrFormat(scope.row.metadata.creationTimestamp) }}
				</template>
			</el-table-column>
			<el-table-column
				label="类型"
				width="100"
				:filters="[
					{ text: 'Normal', value: 'Normal' },
					{ text: 'Warning', value: 'Warning' },
				]"
				:filter-method="filterTag"
			>
				<template #default="scope">
					<el-tag v-if="scope.row.type === 'Normal'" size="small" type="success">
						{{ scope.row.type }}
					</el-tag>
					<el-tag v-else-if="scope.row.type === 'Warning'" size="small" type="warning">
						{{ scope.row.type }}
					</el-tag>
					<el-tag v-else type="danger">
						{{ scope.row.type }}
					</el-tag>
				</template>
			</el-table-column>
			<el-table-column label="信息">
				<template #default="scope">
					{{ scope.row.metadata.uid.split('-')[0] }}<el-divider direction="vertical" /> {{ scope.row.reason }}| {{ scope.row.message
					}}<el-divider direction="vertical" />
					<span style="color: #409eff">{{ scope.row.metadata.name.split('.')[0] }}</span>
				</template></el-table-column
			>
		</el-table>
		<!-- 分页区域 -->
		<Pagination :total="state.total" @handlePageChange="handlePageChange" />
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onMounted, reactive } from 'vue';
import type { ElTableColumn } from 'element-plus';
import { PageInfo } from '/@/types/kubernetes/common';
import { kubernetesInfo } from '/@/stores/kubernetes';
import { useEventApi } from '/@/api/kubernetes/event';
import { Event } from 'kubernetes-types/core/v1';
import { dateStrFormat } from '/@/utils/formatTime';

const Pagination = defineAsyncComponent(() => import('/@/components/pagination/pagination.vue'));

const state = reactive({
	events: [] as Event[],
	query: {
		cloud: '',
		page: 1,
		limit: 10,
	},
	total: 0,
	loading: false,
});
const k8sStore = kubernetesInfo();
const eventApi = useEventApi();
const filterTag = (value: string, row: Event) => {
	return row.type === value;
};
const listEvent = () => {
	state.loading = true;
	state.query.cloud = k8sStore.state.activeCluster;
	eventApi.getEventLog('all', state.query).then((res) => {
		if (res.data.data) {
			state.events = res.data.data;
			state.total = res.data.total;
		}
	});
	state.loading = false;
};
const handlePageChange = (pageInfo: PageInfo) => {
	state.query.page = pageInfo.page;
	state.query.limit = pageInfo.limit;
	listEvent();
};
onMounted(() => {
	listEvent();
});
</script>

<style scoped lang="scss">
.el-table {
	height: calc(100vh - 260px);
}
</style>
