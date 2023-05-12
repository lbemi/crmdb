<template>
	<div class="system-user-dialog-container">
		<el-dialog v-model="dialogVisible" width="800px" :title="title" @close="handleClose()">
			<el-form ref="ruleFormRef" :model="data.service" label-width="120px" class="demo-ruleForm" status-icon>
				<el-row :gutter="35">
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="服务名称" prop="name">
							<el-input v-model="data.service.metadata!.name" />
						</el-form-item>
					</el-col>

					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="命名空间">
							<el-select v-model="data.service.metadata!.namespace" placeholder="指定命名空间">
								<el-option
									v-for="item in k8sStore.state.namespace"
									:key="item.metadata?.name"
									:label="item.metadata?.name"
									:value="item.metadata!.name!"
								/>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="命名空间">
							<el-select v-model="data.service.metadata!.namespace" placeholder="Activity zone">
								<el-option label="Zone one" value="shanghai" />
								<el-option label="Zone two" value="beijing" />
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>

			<template #footer>
				<span class="dialog-footer">
					<el-button size="small" @click="handleClose">关闭</el-button>
					<el-button type="primary" size="small" @click="update">确定</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { V1Service } from '@kubernetes/client-node';
import { reactive, ref, watch } from 'vue';
import { kubernetesInfo } from '/@/stores/kubernetes';

const code = ref('');
const dialogVisible = ref(false);
const k8sStore = kubernetesInfo();
const handleClose = () => {
	emit('update:dialogVisible', false);
};

const emit = defineEmits(['update', 'update:dialogVisible']);

const update = () => {
	emit('update', code.value);
};

const data = reactive({
	service: <V1Service>{
		metadata: {
			name: '',
			namespace: 'default',
			labels: {},
			annotations: {},
		},
		spec: {
			selector: {},
			type: 'ClusterIP',
		},
	},
});

const props = defineProps({
	title: String,
	codeData: Object,
	dialogVisible: Boolean,
});

watch(
	() => props,
	() => {
		dialogVisible.value = props.dialogVisible;
	},
	{
		immediate: true,
	}
);
</script>

<style scoped></style>
