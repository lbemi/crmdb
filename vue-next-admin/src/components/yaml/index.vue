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
import { reactive, ref, watch } from 'vue';
import { Codemirror } from 'vue-codemirror';
import { javascript } from '@codemirror/lang-javascript';
import { oneDark } from '@codemirror/theme-one-dark';

const code = ref();
const extensions = [javascript(), oneDark];
const dialogVisible = ref(false);
const state = reactive({
	dialogVisible: false,
});

const handleClose = () => {
	dialogVisible.value = false;
};

const openDialog = (data: any) => {
	dialogVisible.value = true;
	code.value += YAML.dump(data);
};

const props = defineProps({
	updateResource: Function,
	resourceType: String,
});

watch(
	() => props.resourceType,
	() => {
		if (props.resourceType) {
			switch (props.resourceType) {
				case 'deployment':
					code.value = `apiVersion: apps/v1\nkind: Deployment\n`;
					break;
				case 'statefulSet':
					code.value = `apiVersion: apps/v1
		kind: DaemonSet\n`;
					break;
				default:
					code.value = '';
			}
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
