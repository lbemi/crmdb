<template>
	<div class="system-dept-dialog-container">
		<el-dialog class="my-dialog" :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="formRef" :model="state.host" size="default" label-width="90px" :rules="formRules">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="分组" prop="group_id">
							<el-cascader
								:options="state.groups"
								:props="{ checkStrictly: true, value: 'id', label: 'name' }"
								placeholder="请选分组"
								clearable
								class="w100"
								v-model="state.groupId"
								@change="getGroupId"
							>
								<template #default="{ node, data }">
									<span>{{ data.name }}</span>
									<span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
								</template>
							</el-cascader>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="主机IP" prop="ip">
							<el-input v-model="state.host.ip" placeholder="请输入主机IP" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="主机描述">
							<el-input v-model="state.host.remark" placeholder="请输入描述信息" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="端口" prop="port">
							<el-input v-model.number="state.host.port" placeholder="请输端口号" clearable></el-input>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="标签" prop="labels">
							<el-tag
								style="margin-right: 3px"
								v-for="tag in state.host.labels"
								:key="tag"
								size="small"
								class="mx-1"
								closable
								:disable-transitions="false"
								@close="handleClose(tag)"
							>
								{{ tag }}
							</el-tag>
							<el-input
								v-if="state.inputVisible"
								ref="InputRef"
								v-model="state.inputValue"
								class="ml-1 w-20"
								size="small"
								@keyup.enter="handleInputConfirm"
								@blur="handleInputConfirm"
							/>
							<el-button v-else class="button-new-tag ml-1" size="small" @click="showInput"> + 新标签 </el-button>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="主机状态">
							<el-radio-group v-model="state.host.status">
								<el-radio :label="1">启用</el-radio>
								<el-radio :label="2">不启用</el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="SSH">
							<el-radio-group v-model="state.host.enable_ssh">
								<el-radio :label="1">启用</el-radio>
								<el-radio :label="2">不启用</el-radio>
							</el-radio-group>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit(formRef)" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="hostDialog">
import { nextTick, reactive, ref } from 'vue';
import { Host } from '@/types/asset/hosts';
import { Group } from '@/types/asset/group';
import { ElInput, ElMessage, FormInstance, FormRules } from 'element-plus';
import { useHostApi } from '@/api/asset/hosts';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const hostApi = useHostApi();
const InputRef = ref<InstanceType<typeof ElInput>>();
const formRef = ref<FormInstance>();
const state = reactive({
	inputValue: '',
	inputVisible: false,
	groupId: [],
	host: {
		enable_ssh: 1,
		status: 1,
		labels: ['asd', 'sad'],
		ip: '192.168.1.1',
		port: 22,
	} as Host,
	groups: [] as Group[], // 主机数据
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: Group[]) => {
	state.groups = row;
	if (type === 'edit') {
		state.dialog.title = '修改主机';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增主机';
		state.dialog.submitTxt = '新 增';
		// 清空表单，此项需加表单验证才能使用
		// nextTick(() => {
		// 	deptDialogFormRef.value.resetFields();
		// });
	}
	console.log(state.groups);
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	resetForm(formRef.value);
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate(async (valid) => {
		if (valid) {
			state.host.group_id = state.groupId.pop();

			console.log(state.host);
			if (state.dialog.title === '新增主机') {
				await hostApi
					.addHost(state.host)
					.then((res: any) => {
						ElMessage.success(res.message);
						emit('refresh');
						closeDialog();
					})
					.catch((e: any) => {
						ElMessage.error(e.message);
					});
			}
		} else {
			console.log('error submit!');
			return false;
		}
	});
};

const resetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
const props = defineProps({
	groups: Array<Group>,
});

const getGroupId = () => {
	if (state.groupId.length > 0) {
		state.host.group_id = state.groupId[state.groupId.length - 1];
	}
};
const validateIP = (rule: any, value: any, callback: any) => {
	if (value === '' || typeof value === 'undefined' || value == null) {
		callback(new Error('请输入正确的IP地址'));
	} else {
		const reg =
			/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
		if (!reg.test(value) && value !== '') {
			callback(new Error('请输入正确的IP地址'));
		} else {
			callback();
		}
	}
};

const handleClose = (tag: string) => {
	state.host.labels.splice(state.host.labels.indexOf(tag), 1);
};
const showInput = () => {
	state.inputVisible = true;
	nextTick(() => {
		InputRef.value!.input!.focus();
	});
};
//添加标签
const handleInputConfirm = () => {
	if (!state.host.labels) {
		state.host.labels = [];
	}
	if (state.inputValue) {
		state.host.labels.push(state.inputValue);
	}
	state.inputVisible = false;
	state.inputValue = '';
};
const formRules = reactive<FormRules>({
	ip: [{ required: true, validator: validateIP, trigger: 'blur' }],
	port: [{ required: true, message: '请输入端口', trigger: 'blur' }],
	group_id: [{ required: true, message: '请选择分组', trigger: 'change' }],
});

// 暴露变量
defineExpose({
	openDialog,
});
</script>
<style></style>
