/** * Created by lei on 2022/11/16 */
<template>
	<el-dialog v-model="dialogVisible" style="width: 500px" @close="handleClose(ruleFormRef)">
		<template #header="{ titleId, titleClass }">
			<div class="my-header">
				<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
			</div>
		</template>
		<div class="dialog-body">
			<el-form ref="ruleFormRef" :model="state" status-icon label-width="80px" class="demo-ruleForm" :rules="clusterRule">
				<el-form-item label="集群名称" prop="clusterName">
					<el-input v-model="state.clusterName" autocomplete="off" size="small" style="width: 350px" />
				</el-form-item>
				<el-form-item label="配置文件">
					<el-upload ref="uploadRef" :limit="1" drag :auto-upload="false" :on-change="handleChange" multiple style="width: 350px">
						<el-icon class="el-icon--upload"><upload-filled /></el-icon>
						<div class="el-upload__text">拖拽文件到这里 <em>或者点击上传</em></div>
						<template #tip>
							<div class="el-upload__tip">kubernetes 配置文件大小小于500kb</div>
						</template>
					</el-upload>
				</el-form-item>
			</el-form>
		</div>
		<template #footer>
			<span class="dialog-footer">
				<el-button type="primary" size="small" @click="submitForm(ruleFormRef)">创建</el-button>
				<el-button size="small" @click="resetForm(ruleFormRef)">重置</el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script setup lang="ts" name="kubernetesDialog">
import { onMounted, reactive, ref, toRefs } from 'vue';
import { UploadFilled } from '@element-plus/icons-vue';
import { ElMessage, FormRules, UploadFile } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { useClusterApi } from '@/api/kubernetes/cluster';

const clusterApi = useClusterApi();
const ruleFormRef = ref();
const newFormData = new FormData();
const uploadRef = ref();
// 定义子组件向父组件传值/事件
const emits = defineEmits(['update:dialogVisible', 'valueChange']);
// 获取父组件传递的值
const props = defineProps<{
	dialogVisible: boolean;
	title: string;
}>();

const { dialogVisible } = toRefs(props);
const state = reactive({
	clusterName: '',
});

const requestConfig = {
	headers: {
		'Content-Type': 'multipart/form-data',
	},
};

const handleChange = (file: UploadFile | undefined) => {
	if (!file) {
		ElMessage.error('请添加配置文件');
	} else {
		newFormData.append('file', file.raw || '');
	}
};

const clusterRule = reactive<FormRules>({
	clusterName: [{ required: true, message: '请输集群名称', trigger: 'blur' }],
});

const submitForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate(async (valid) => {
		if (valid) {
			newFormData.append('name', state.clusterName);
			await clusterApi
				.createCluster(newFormData, requestConfig.headers)
				.then((res: any) => {
					emits('valueChange');
					handleClose(formEl);
					ElMessage.success(res.message);
				})
				.catch((e) => {
					newFormData.delete('file');
					newFormData.delete('name');
					ElMessage.error(e.message);
					handleClose(formEl);
				});
		} else {
			ElMessage.error('请正确填写!');
			return false;
		}
	});
};

const resetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
	uploadRef.value.clearFiles();
};

const handleClose = (formEl: FormInstance | undefined) => {
	resetForm(formEl);
	emits('update:dialogVisible', false);
};

onMounted(() => {});
</script>

<style scoped lang="less"></style>
