<template>
	<div class="system-user-dialog-container">
		<el-dialog v-model="dialogVisible" width="769px">
			<template #header>
				<h3>YAML</h3>
			</template>
			<codemirror v-model="code" :style="{ height: '100%' }" :autofocus="true" :tabSize="2" :extensions="extensions" />
			<template #footer>
				<span class="dialog-footer">
					<el-button size="small" @click="handleClose">取 消</el-button>
					<el-button type="primary" size="small" @click="update">更新</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import YAML from 'js-yaml';
import { ref, watch } from 'vue';
import { Codemirror } from 'vue-codemirror';
import { oneDark } from '@codemirror/theme-one-dark';
import { StreamLanguage, foldGutter } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';

const code = ref('');
const extensions = [oneDark, StreamLanguage.define(yaml), foldGutter()];
const dialogVisible = ref(false);
const handleClose = () => {
	// dialogVisible.value = false;
	emit('update:dialogVisible', false);

	// switch (props.resourceType) {
	// 	case 'deployment':
	// 		code.value = `apiVersion: apps/v1\nkind: Deployment\n`;
	// 		break;
	// 	case 'statefulSet':
	// 		code.value = `apiVersion: apps/v1\nkind: DaemonSet\n`;
	// 		break;
	// 	case 'pod':
	// 		code.value = `apiVersion: v1\nkind: Pod\n`;
	// 		break;
	// 	case 'node':
	// 		code.value = `apiVersion: v1\nkind: Node\n`;
	// 		break;
	// 	case 'ingress':
	// 		code.value = `apiVersion: networking.k8s.io/v1\nkind: Ingress\n`;
	// 		break;
	// 	default:
	// 		code.value = '';
	// }
};

const emit = defineEmits(['update', 'update:dialogVisible']);

const update = () => {
	emit('update', code.value);
};

// const openDialog = (data: any) => {
// 	dialogVisible.value = true;
// 	code.value += YAML.dump(data);
// };

const props = defineProps({
	codeData: Object,
	dialogVisible: Boolean,
	resourceType: String,
});

watch(
	() => [props.resourceType, props.codeData],
	() => {
		dialogVisible.value = props.dialogVisible;
		if (props.resourceType) {
			switch (props.resourceType) {
				case 'deployment':
					code.value = `apiVersion: apps/v1\nkind: Deployment\n`;
					break;
				case 'statefulSet':
					code.value = `apiVersion: apps/v1\nkind: DaemonSet\n`;
					break;
				case 'pod':
					code.value = `apiVersion: v1\nkind: Pod\n`;
					break;
				case 'node':
					code.value = `apiVersion: v1\nkind: Node\n`;
					break;
				case 'ingress':
					code.value = `apiVersion: networking.k8s.io/v1\nkind: Ingress\n`;
					break;
				default:
					code.value = '';
			}
		}
		if (props.codeData) {
			code.value += YAML.dump(props.codeData);
		}
	},
	{
		immediate: true,
	}
);

// defineExpose({
// 	// openDialog,
// 	handleClose,
// 	code,
// });
</script>

<style scoped></style>
