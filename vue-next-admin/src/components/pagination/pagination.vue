<template>
	<div>
		<!-- 分页区域 -->
		<el-pagination
			class="mt15"
			:small="theme.themeConfig.globalComponentSize === 'small'"
			background
			v-model:current-page="page.page"
			v-model:page-size="page.limit"
			:pager-count="5"
			:page-sizes="[10, 20, 50]"
			layout="total, sizes, prev, pager, next, jumper"
			:total="total"
			@size-change="handleSizeChange"
			@current-change="handlePageChange"
		/>
	</div>
</template>

<script setup lang="ts">
import { useThemeConfig } from '@/stores/themeConfig';
import { reactive } from 'vue';
const theme = useThemeConfig();
defineProps({ total: Number });

const page = reactive({
	page: 1,
	limit: 10,
});

const emit = defineEmits(['handlePageChange']);
const handlePageChange = () => {
	emit('handlePageChange', page);
};
const handleSizeChange = () => {
	page.page = 1;
	emit('handlePageChange', page);
};
</script>

<style scoped></style>
