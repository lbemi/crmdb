<template>
	<el-dialog v-model="data.visible" @close="handleClose(formRef)" :title="title" style="width: 600px">
		<el-form label-width="100px" ref="formRef" :rules="formRules" :model="data.namespace.metadata" style="max-width: 700px">
			<el-form-item label="名字:" prop="name" class="name">
				<el-input v-model="data.namespace.metadata!.name" :disabled="title === '更新命名空间'" size="small" />
			</el-form-item>
			<span class="info">长度为 1 ~ 63 个字符，只能包含数字、小写字母和中划线（-），且首尾只能是字母或数字</span>
			<el-form-item label="标签:" prop="data.metadata.labels" class="label">
				<el-table :data="table.tableData" style="width: 100%; font-size: 12px">
					<el-table-column :prop="item.prop" :label="item.label" v-for="item in table.tableHeader" :key="item.prop">
						<template #default="scope">
							<div v-show="item.editable || scope.row.editable">
								<template v-if="item.type === 'input'">
									<el-input
										style="width: 62px; font-size: 10px"
										size="small"
										v-model="scope.row[item.prop]"
										:placeholder="`${item.label}`"
										@change="handleEdit(scope.$index, scope.row)"
									/>
								</template>
							</div>
							<div v-show="!item.editable && !scope.row.editable" class="editable-row">
								<span class="editable-row-span">{{ scope.row[item.prop] }}</span>
							</div>
						</template>
					</el-table-column>
					<el-table-column label="操作" width="120px" style="display: flex; padding-left: 2px">
						<template #default="scope">
							<el-button v-show="!scope.row.editable || !scope.row.addStatus" size="small" link type="primary" @click="scope.row.editable = true"
								>编辑</el-button
							>
							<el-button v-show="scope.row.addStatus" :disabled="scope.row.key === ''" size="small" link type="primary" @click="append(scope.$index)"
								>添加</el-button
							>
							<el-button v-show="scope.row.editable && !scope.row.addStatus" size="small" link type="primary" @click="add(scope.row)">确定</el-button>
							<el-button v-show="!scope.row.addStatus && !scope.row.editable" size="small" type="danger" link @click="handleDelete(scope.$index)"
								>删除</el-button
							>
						</template>
					</el-table-column>
				</el-table>
			</el-form-item>
		</el-form>
		<template #footer>
			<span class="dialog-footer">
				<el-button size="small" @click="handleClose(formRef)">取消</el-button>
				<el-button size="small" type="primary" @click="btnOk(formRef)"> 确定 </el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script setup lang="ts" name="namespaceDialog">
import { Namespace } from 'kubernetes-types/core/v1';
import { FormRules, FormInstance, ElMessage } from 'element-plus';
import { reactive, watch, ref, onMounted } from 'vue';
import { kubernetesInfo } from '@/stores/kubernetes';
import { useNamespaceApi } from '@/api/kubernetes/namespace';

const k8sStore = kubernetesInfo();
const namespaceApi = useNamespaceApi();
const item = reactive({
	key: '',
	value: '',
	addStatus: false,
	editable: false,
});

interface label {
	key: string;
	value: string;
}
interface header {
	prop: string;
	label: string;
	editable: boolean;
	type: string;
}
const table = reactive({
	tableHeader: [
		{
			prop: 'key',
			label: '变量名称',
			editable: false,
			type: 'input',
		},
		{
			prop: 'value',
			label: '变量值',
			editable: false,
			type: 'input',
		},
	],
	tableData: [] as label[],
});
onMounted(() => {
	item.editable = true;
	item.addStatus = true;
	table.tableData.splice(table.tableData.length, 0, item);
});
const handleEdit = (index: number, row: header) => {
	row.editable = true;
};
const handleDelete = (index: number) => {
	table.tableData.splice(index, 1);
};
const append = (index: number) => {
	item.editable = false;
	table.tableData.push({
		key: item.key,
		value: item.value,
	});
	table.tableData.splice(index, 1);
	item.key = '';
	item.value = '';
	item.editable = true;
	table.tableData.splice(table.tableData.length + 1, 0, item);
};
const add = (header: header) => {
	header.editable = false;
};
const props = defineProps<{
	visible: boolean;
	title: string;
	namespace?: Namespace;
}>();

const emits = defineEmits(['update:visible', 'valueChange']);
const formRef = ref<FormInstance>();

const data = reactive({
	visible: false,
	namespace: {
		metadata: {
			name: '',
		},
	} as Namespace,
	name: '',
	labels: {} as { [index: string]: string },
});
const formRules = reactive<FormRules>({
	name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
});

const handleClose = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
	emits('update:visible', false);
};

const handleLabels = (labels: Array<{ [index: string]: string }>) => {
	data.namespace.metadata!.labels = {};
	for (const k in labels) {
		if (labels[k].key != '') {
			data.namespace.metadata!.labels[labels[k].key] = labels[k].value;
		}
	}
};
const btnOk = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate((valid) => {
		if (valid) {
			if (props.title === '更新命名空间') {
				handleLabels(table.tableData);
				namespaceApi
					.updateNamespace({ cloud: k8sStore.state.activeCluster }, data.namespace)
					.then((res) => {
						if (res.code == 200) {
							ElMessage.success(res.message);
							emits('valueChange');
							emits('update:visible', false);
						} else {
							ElMessage.error(res.message);
						}
					})
					.catch((e) => {
						ElMessage.error(e);
					});
			} else {
				handleLabels(table.tableData);
				namespaceApi
					.createNamespace({ cloud: k8sStore.state.activeCluster }, data.namespace)
					.then((res: any) => {
						ElMessage.success(res.message);
						emits('valueChange');
						emits('update:visible', false);
					})
					.catch((e) => {
						ElMessage.error(e);
					});
			}
		}
	});
};
watch(
	() => props,
	() => {
		data.visible = props.visible;
		if (props.namespace?.metadata && props.title === '更新命名空间') {
			data.namespace.metadata!.name = props.namespace.metadata.name;
			data.namespace.metadata!.labels = props.namespace.metadata.labels;

			for (let key in props.namespace.metadata.labels) {
				const l: label = {
					key: key,
					value: props.namespace.metadata.labels[key],
				};
				table.tableData.push(l);
			}
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>

<style scoped lang="scss">
.info {
	font-size: 7px;
	margin-left: 100px;
	color: #999 !important;
	margin-top: 5px;
}
.name {
	margin-bottom: 15px;
}
.label {
	margin-top: 10px;
}
</style>
