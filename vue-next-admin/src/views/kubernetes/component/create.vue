/** * Created by lei on 2022/11/16 */
<template>
	<el-dialog v-model="dialogVisible" :close="handleClose(ruleFormRef)" style="width: 500px">
		<template #header="{ titleId, titleClass }">
			<div class="my-header">
				<h4 :id="titleId" :class="titleClass">{{ title }}</h4>
				<el-divider />
			</div>
		</template>
		<div class="dialog-body">
			<el-form ref="ruleFormRef" status-icon label-width="80px" class="demo-ruleForm" :rules="clusterRule">
				<el-form-item label="集群名称" prop="clusterName">
					<el-input v-model="clusterName" autocomplete="off" />
				</el-form-item>
				<el-form-item label="配置文件">
					<el-upload :limit="1" drag :auto-upload="false" :on-change="handleChange" multiple>
						<el-icon class="el-icon--upload"><upload-filled /></el-icon>
						<div class="el-upload__text">Drop file here or <em>click to upload</em></div>
						<template #tip>
							<div class="el-upload__tip">kube config files with a size less than 500kb</div>
						</template>
					</el-upload>
				</el-form-item>

				<el-form-item>
					<el-button type="primary" @click="submitForm(ruleFormRef)">Submit</el-button>
					<el-button @click="resetForm(ruleFormRef)">Reset</el-button>
				</el-form-item>
			</el-form>
		</div>
	</el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref, toRefs } from 'vue';
import { UploadFilled } from '@element-plus/icons-vue';
import { ElMessage, FormRules, UploadFile } from 'element-plus';
import type { FormInstance } from 'element-plus';
import { useClusterApi } from '/@/api/kubernetes/cluster';

const clusterApi = useClusterApi();
const ruleFormRef = ref<FormInstance>();
const newFormData = new FormData();

// 定义子组件向父组件传值/事件
const emits = defineEmits(['update:dialogVisible', 'valueChange']);

// 获取父组件传递的值
const props = defineProps<{
	dialogVisible: boolean;
	title: string;
}>();

const { dialogVisible } = toRefs(props);
const clusterName = ref('');

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
	clusterName: [{ required: true, message: '请输入名字', trigger: 'blur' }],
});
const submitForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate(async (valid) => {
		if (valid) {
			newFormData.append('name', clusterName.value);
			await clusterApi
				.createCluster(newFormData, requestConfig.headers)
				.then(() => {
					emits('valueChange');
					handleClose(formEl);
					ElMessage.success('添加成功');
				})
				.catch(() => {
					newFormData.delete('file');
					newFormData.delete('name');
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
};

const handleClose = (formEl: FormInstance | undefined) => {
	resetForm(formEl);
	emits('update:dialogVisible', false);
};
</script>

<style scoped lang="less"></style>
