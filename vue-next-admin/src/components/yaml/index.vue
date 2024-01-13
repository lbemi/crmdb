<template>
	<div class="system-user-dialog-container">
		<el-drawer v-model="dialogVisible" size="40%" @close="handleClose()">
			<template #header>
				<h3>yaml</h3>
			</template>
			<Codemirror v-model="code" style="height: 92%;margin-left: 20px;margin-right: 15px;" :autofocus="true"
				:tabSize="2" :extensions="extensions" :disabled="disabledUpdate" />
			<div class="mt30" style="align-items: center; margin-left: 20px;">
				<el-button size="small" @click="handleClose">关闭</el-button>
				<el-button type="primary" size="small" @click="update" v-if="!disabledUpdate">更新</el-button>
			</div>

		</el-drawer>
	</div>
</template>

<script setup lang="ts">
import YAML from 'js-yaml';
import { ref, watch } from 'vue';
import { Codemirror } from 'vue-codemirror';
import { oneDark } from '@codemirror/theme-one-dark';
import { StreamLanguage } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';

const code = ref('');
const extensions = [oneDark, StreamLanguage.define(yaml)];
const dialogVisible = ref(false);
const handleClose = () => {
	emit('update:dialogVisible', false);
};

const emit = defineEmits(['update', 'update:dialogVisible']);

const update = () => {
	emit('update', code.value);
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
				case 'virtualService':
					code.value = `apiVersion: v1\nkind: virtualService\n`;
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
