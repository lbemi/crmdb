<template>
	<div>
		<el-form-item label="启动命令：" label-width="90px" style="margin-bottom: 0">
			<template #label>
				<el-tooltip
					class="box-item"
					effect="light"
					content="自定义容器启动时运行的命令; 默认情况下，容器启动时将运行镜像默认命令"
					placement="top-start"
					raw-content
				>
					启动命令：
				</el-tooltip>
			</template>
			<el-checkbox v-model="data.set" label="开启" size="small" />
			<el-button
				v-if="data.show"
				type="info"
				v-show="data.set"
				text
				:icon="CaretTop"
				@click="data.show = !data.show"
				size="small"
				style="margin-left: 30px"
				>隐藏</el-button
			>
			<el-button
				v-else
				type="info"
				v-show="data.set"
				text
				:icon="CaretBottom"
				@click="data.show = !data.show"
				size="small"
				style="margin-left: 30px"
				>展开</el-button
			>
		</el-form-item>
		<el-form-item label="命令：" v-show="data.set && data.show" style="margin-bottom: 0" label-width="60px">
			<el-input v-model="data.commands" size="small" style="width: 200px" />
			<span style="font-size: 10px; color: rgba(22, 9, 7, 0.57); margin-left: 5px">如有多个命令请使用半角逗号（,）分隔</span>
		</el-form-item>
		<el-form-item label="参数：" v-show="data.set && data.show" style="margin-bottom: 0" label-width="60px">
			<el-input v-model="data.args" size="small" style="width: 200px" />
			<span style="font-size: 10px; color: rgba(22, 9, 7, 0.57); margin-left: 5px"> 如有多个参数请使用半角逗号（,）分隔</span>
		</el-form-item>
	</div>
</template>

<script setup lang="ts">
import { CaretBottom, CaretTop } from '@element-plus/icons-vue';
import { reactive, watch } from 'vue';
import { deepClone } from '/@/utils/other';

const data = reactive({
	set: false,
	show: true,
	commands: '',
	args: '',
	k8s: {
		commands: [],
		args: [],
	},
});
const props = defineProps({
	args: Array<String>,
	commands: Array<String>,
});

const handleArr = (source: Array<String>) => {
	const dataCopy = deepClone(source);
	let str = '';
	dataCopy.forEach((item, index) => {
		if (index == dataCopy.length - 1) {
			str = str + item;
		} else {
			str = str + item + ',';
		}
	});
	return str;
};
watch(
	() => [props.args, props.commands],
	() => {
		if (props.args) {
			data.args = handleArr(props.args);
		}
		if (props.commands) {
			data.commands = handleArr(props.commands);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
const emit = defineEmits(['updateCommand']);
watch(
	() => [data.args, data.commands, data.set],
	() => {
		if (!data.set) {
			data.k8s.args = [];
			data.k8s.commands = [];
		} else {
			if (data.args) {
				data.k8s.args = data.args.split(',');
			}
			if (props.commands) {
				data.k8s.commands = data.commands.split(',');
			}
		}
		emit('updateCommand', data.k8s);
	},
	{
		immediate: true,
		deep: true,
	}
);
</script>

<style scoped></style>
