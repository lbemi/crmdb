<template>
	<div class="system-user-dialog-container">
		<el-dialog title="YAML" v-model="dialogVisible" width="800px">
			<div ref="editorRef" class="editor-main" />
			<template #footer>
				<span class="dialog-footer">
					<el-button size="default" @click="handleClose">关闭</el-button>
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
const code = ref();
const dialogVisible = ref(false);

const handleClose = () => {
	dialogVisible.value = false;
	code.value = {};
};

const initEditor = () => {
	if (typeof editorView.value !== 'undefined') {
		editorView.value.destroy();
	}
	const state = EditorState.create({
		doc: YAML.dump(code.value),
		extensions: [basicSetup, oneDark, StreamLanguage.define(yaml)],
	});

	if (editorRef.value) {
		editorView.value = new EditorView({
			state,
			// state: startState,
			parent: editorRef.value,
		});
	}
};

const props = defineProps({
	code: Object,
	dialogVisible: Boolean,
});

watch(
	() => [props.code, props.dialogVisible],
	() => {
		dialogVisible.value = props.dialogVisible;
		code.value = props.code;
		initEditor();
	},
	{
		immediate: true,
		deep: true,
	}
);
onMounted(() => {
	nextTick(() => {
		initEditor();
	});
});
</script>

<style scoped lang="scss">
.editor-main {
	width: 100%;
	height: 100%;
}
</style>
