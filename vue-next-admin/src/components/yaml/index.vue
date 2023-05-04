<template>
	<div class="system-user-dialog-container">
		<el-dialog title="YAML" v-model="dialogVisible" width="769px">
			<codemirror v-model="code" :style="{ height: '100%' }" :autofocus="true" :tabSize="2" :extensions="extensions" />
			<template #footer>
				<span class="dialog-footer">
					<el-button size="default" @click="handleClose">取 消</el-button>
					<el-button type="primary" size="default" @click="props.updateResource">更新</el-button>
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
const merge = ref(false);
const handleClose = () => {
	dialogVisible.value = false;

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
		default:
			code.value = '';
	}
};

const openDialog = (data: any) => {
	dialogVisible.value = true;
	code.value += YAML.dump(data);
};

const props = defineProps({
	updateResource: Function,
	resourceType: String,
	merge: Boolean,
});

watch(
	() => [props.resourceType, props.merge],
	() => {
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
				default:
					code.value = '';
			}
		}
		if (props.merge) {
			merge.value = props.merge;
		}
	},
	{
		immediate: true,
	}
);

defineExpose({
	openDialog,
	handleClose,
	code,
});
</script>

<style scoped></style>
