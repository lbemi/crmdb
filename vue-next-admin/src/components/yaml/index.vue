<template>
	<div class="system-user-dialog-container">
		<el-dialog title="YAML" v-model="state.dialogVisible" width="769px">
			<codemirror v-model="code" :style="{ height: '100%' }" :autofocus="true" :tabSize="2" :extensions="extensions" />
			<template #footer>
				<span class="dialog-footer">
					<el-button size="default" @click="handleClose">取 消</el-button>
					<el-button type="primary" size="default">更新</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import YAML from 'js-yaml';
import { reactive, ref } from 'vue';
import { Codemirror } from 'vue-codemirror';
import { javascript } from '@codemirror/lang-javascript';
import { oneDark } from '@codemirror/theme-one-dark';
const code = ref();
const extensions = [javascript(), oneDark];
const state = reactive({
	dialogVisible: false,
});

const handleClose = () => {
	state.dialogVisible = false;
};

const openDialog = (data: any) => {
	state.dialogVisible = true;
	code.value = `apiVersion: apps/v1
kind: Deployment\n`;
	code.value += YAML.dump(data);
};

defineExpose({
	openDialog,
});
</script>

<style scoped></style>
