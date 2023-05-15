<template>
	<div class="system-user-dialog-container">
		<el-dialog v-model="dialogVisible" width="800px" @close="handleClose()">
			<template #header>
				<h3>YAML</h3>
			</template>
			<div ref="editorRef" />
			<template #footer>
				<span class="dialog-footer">
					<el-button size="small" @click="handleClose">关闭</el-button>
					<el-button type="primary" size="small" @click="update" v-if="!disabledUpdate">更新</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import YAML from 'js-yaml';
import { nextTick, onMounted, ref, watch } from 'vue';
import { oneDark } from '@codemirror/theme-one-dark';
import { EditorState } from '@codemirror/state';
import { EditorView, basicSetup } from 'codemirror';
import { StreamLanguage } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';

const editorRef = ref();
const editorView = ref();
const code = ref('');
const dialogVisible = ref(false);

const initEditor = () => {
	if (typeof editorView.value !== 'undefined') {
		editorView.value.destroy();
	}
	const state = EditorState.create({
		doc: code.value,
		extensions: [basicSetup, oneDark, StreamLanguage.define(yaml)],
	});

	if (editorRef.value) {
		editorView.value = new EditorView({
			state,
			parent: editorRef.value,
		});
	}
};

onMounted(() => {
	nextTick(() => {
		initEditor();
	});
});

const handleClose = () => {
	emit('update:dialogVisible', false);
};

const emit = defineEmits(['update', 'update:dialogVisible']);

const update = () => {
	emit('update', editorView.value.state.doc.toString());
};

const props = defineProps({
	codeData: Object,
	dialogVisible: Boolean,
	resourceType: String,
	disabledUpdate: Boolean,
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
				case 'service':
					code.value = `apiVersion: v1\nkind: Service\n`;
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
</script>

<style scoped></style>
