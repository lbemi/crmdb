<template>
	<div class="system-user-dialog-container">
		<el-dialog title="YAML" v-model="dialogVisible" width="769px">
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
import { javascript } from '@codemirror/lang-javascript';
import { oneDark } from '@codemirror/theme-one-dark';
import { basicSetup, EditorView } from 'codemirror';
import { EditorState, Extension } from '@codemirror/state';
import {
	crosshairCursor,
	drawSelection,
	dropCursor,
	highlightActiveLine,
	highlightActiveLineGutter,
	highlightSpecialChars,
	lineNumbers,
	rectangularSelection,
	GutterMarker,
	keymap,
} from '@codemirror/view';
import {
	StreamLanguage,
	bracketMatching,
	defaultHighlightStyle,
	foldGutter,
	foldKeymap,
	indentOnInput,
	syntaxHighlighting,
} from '@codemirror/language';
import { autocompletion, closeBrackets } from '@codemirror/autocomplete';
import { lintGutter, linter } from '@codemirror/lint';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';
const editorRef = ref();
const editorView = ref();
const code = ref();
const dialogVisible = ref(false);

const handleClose = () => {
	dialogVisible.value = false;
};
const myTheme = EditorView.theme(
	{
		'&': {
			color: 'white',
			backgroundColor: '#034',
		},
		'.cm-content': {
			caretColor: '#0e9',
		},
		'&.cm-focused .cm-cursor': {
			borderLeftColor: '#0e9',
		},
		'&.cm-focused .cm-selectionBackground, ::selection': {
			backgroundColor: '#074',
		},
		'.cm-gutters': {
			backgroundColor: '#045',
			color: '#ddd',
			border: 'none',
		},
	},
	{ dark: true }
);

const basicSetup1: Extension = (() => [
	lineNumbers(), // 显示行
	highlightActiveLineGutter(),
	highlightSpecialChars(),
	foldGutter(), // 折叠
	drawSelection(),
	dropCursor(),
	EditorState.allowMultipleSelections.of(true), // 多行选择
	syntaxHighlighting(defaultHighlightStyle, { fallback: true }),
	bracketMatching(),
	closeBrackets(),
	autocompletion(), // 自动完成
	rectangularSelection(),
	crosshairCursor(),
	highlightActiveLine(),
	lintGutter(),
	StreamLanguage.define(yaml),
])();
const initEditor = () => {
	if (typeof editorView.value !== 'undefined') {
		editorView.value.destroy();
	}
	const startState = EditorState.create({
		doc: YAML.dump(code.value),
		extensions: [basicSetup, oneDark, StreamLanguage.define(yaml), foldGutter()],
	});

	if (editorRef.value) {
		editorView.value = new EditorView({
			state: startState,
			parent: editorRef.value,
		});
	}
};

const props = defineProps({
	code: Object,
	dialogVisible: Boolean,
});

watch(
	() => props,
	() => {
		dialogVisible.value = props.dialogVisible;
		code.value = props.code;
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

function highlightSelectionMatches(): Extension {
	throw new Error('Function not implemented.');
}
</script>

<style scoped lang="scss">
.editor-main {
	width: 100%;
	height: 100%;
}
</style>
