<template>
	<el-drawer v-model="state.visible" @close="handleClose" size="60%">
		<template #header="{ titleId, titleClass }">
			<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
		</template>
		<div class="box_body">
			<el-descriptions :title="state.secret.metadata!.name" :column="2" border>
				<el-descriptions-item label="命名空间" label-align="right" align="center" label-class-name="my-label" class-name="my-content" width="150px">{{
					state.secret.metadata!.namespace
				}}</el-descriptions-item>
				<el-descriptions-item label="创建时间" label-align="right" align="center" :span="2">{{
					dateStrFormat(state.secret.metadata!.creationTimestamp)
				}}</el-descriptions-item>
				<el-descriptions-item label="标签" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.secret.metadata!.labels" :key="index">
						<el-tag class="label" type="success" size="small"> {{ key }}:{{ item }} </el-tag>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="注解" label-align="right" align="center" :span="2">
					<div v-for="(item, key, index) in state.secret.metadata!.annotations" :key="index">
						<span> {{ key }}:{{ item }} </span>
					</div>
				</el-descriptions-item>
				<el-descriptions-item label="数据" label-align="right" align="center" :span="2">
					<el-table :data="state.keyValues">
						<el-table-column prop="key" label="名称" width="120" />
						<el-table-column prop="value" label="值" />
						<el-table-column width="80" label="查看">
							<template #default="scope">
								<el-icon color="#409EFC" @click="parseData(scope.row)"><View /></el-icon>
								<el-icon @click="secData(scope.row)"><Hide /></el-icon>
								<!--								<el-button :icon="View" size="small" circle @click="parseData(scope.row)"></el-button>-->
							</template>
						</el-table-column>
					</el-table>
				</el-descriptions-item>
			</el-descriptions>
		</div>
	</el-drawer>
</template>

<script lang="ts" setup>
import { ElDrawer } from 'element-plus';
import { View, Hide } from '@element-plus/icons-vue';
import { Secret } from 'kubernetes-models/v1';
import { onMounted, reactive } from 'vue';
import { isObjectValueEqual } from '@/utils/arrayOperation';

const state = reactive({
	visible: false,
	secret: {
		metadata: {
			name: '',
			namespace: '',
		},
		data: {},
	} as Secret,
	keyValues: [] as Array<{ key: string; value: string }>,
});

const props = defineProps({
	visible: Boolean,
	secret: {
		type: Object as () => Secret,
	},
	title: String,
});

const parseData = (data: any) => {
	data.value = atob(data.value);
};
const secData = (data: any) => {
	data.value = btoa(data.value);
};

const convertConfigMapTo = () => {
	let kvs = [] as Array<{ key: string; value: string }>;
	if (state.secret.data) {
		Object.keys(state.secret.data).forEach((k) => {
			kvs.push({
				key: k,
				value: state.secret.data![k],
			});
		});
	}
	state.keyValues = kvs;
};

onMounted(() => {
	state.visible = props.visible;
	if (!isObjectValueEqual(props.secret, {})) {
		state.secret = props.secret;
		convertConfigMapTo();
	}
});

const emit = defineEmits(['update:visible']);
const handleClose = () => {
	emit('update:visible', false);
};
</script>
<style scoped>
.box_body {
	margin-left: 20px;
	margin-top: 10px;
}
.footer {
	display: flex;
	margin-top: 50px;
	/*margin-left: 80px;*/
	justify-content: center;
}
.label {
	margin-top: 3px;
	margin-bottom: 1px;
}
</style>
