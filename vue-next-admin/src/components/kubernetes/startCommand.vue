<template>
	<el-form>
		<el-form-item label="启动命令：" label-width="100px" style="margin-bottom: 0">
			<template #label>
				<el-tooltip class="box-item" effect="light" content="自定义容器启动时运行的命令; 默认情况下，容器启动时将运行镜像默认命令"
					placement="top-start" raw-content>
					启动命令：
				</el-tooltip>
			</template>
			<el-checkbox v-model="data.set" label="开启" size="small" />
			<el-button v-if="data.show" type="info" v-show="data.set" text :icon="CaretTop" @click="data.show = !data.show"
				size="small" style="margin-left: 30px">隐藏</el-button>
			<el-button v-else type="info" v-show="data.set" text :icon="CaretBottom" @click="data.show = !data.show"
				size="small" style="margin-left: 30px">展开</el-button>
		</el-form-item>
		<el-form-item label-width="100px">
			<el-form-item label="命令：" v-show="data.set && data.show" label-width="60px" class="nested-item">
				<el-input v-model="data.commands" size="small" style="width: 500px" type="textarea" />
				<div style="font-size: 13px; color: rgba(22, 9, 7, 0.57); margin-left: 5px">如有多个命令请使用半角逗号（,）分隔</div>
			</el-form-item>
			<el-form-item label="参数：" v-show="data.set && data.show" label-width="60px" class="nested-item">
				<el-input v-model="data.args" size="small" style="width: 500px" type="textarea" />
				<span style="font-size: 13px; color: rgba(22, 9, 7, 0.57); margin-left: 5px"> 如有多个参数请使用半角逗号（,）分隔</span>
			</el-form-item>
		</el-form-item>
	</el-form>
</template>

<script setup lang="ts">
import { CaretBottom, CaretTop } from '@element-plus/icons-vue';
import { onMounted, reactive } from 'vue';
import { deepClone } from '@/utils/other';

// FIXME args无法清空
const data = reactive({
	loadFromParent: false,
	set: false,
	show: true,
	commands: '',
	args: '',
	k8s: {
		commands: [''],
		args: [''],
	},
});

type propsType = {
	args: Array<String> | undefined;
	commands: Array<String> | undefined;
};
const props = defineProps<propsType>();

const handleArr = (source: Array<String>) => {
	const dataCopy = deepClone(source);
	let str = '';
	dataCopy.forEach((item: string, index: number) => {
		if (index == dataCopy.length - 1) {
			str = str + item;
		} else {
			str = str + item + ',';
		}
	});
	return str;
};
onMounted(() => {
	if (props.args || props.args != undefined) {
		data.set = true; //当本地数据为空，但是传递过来数据不为空时则显示
		data.args = handleArr(props.args);
	}
	if (props.commands || props.commands != undefined) {
		data.set = true; //当本地数据为空，但是传递过来数据不为空时则显示
		data.commands = handleArr(props.commands);
	}
});

const returnStartCommand = () => {
	if (!data.set) {
		data.k8s.args = [];
		data.k8s.commands = [];
	} else {
		if (data.args.length > 0) {
			data.k8s.args = data.args.split(',');
		}
		if (data.commands.length > 0) {
			data.k8s.commands = data.commands.split(',');
		}
	}
	return { set: data.set, commands: data.k8s.commands, args: data.k8s.args };
};

defineExpose({
	returnStartCommand,
});
</script>

<style scoped>
.nested-item {
	margin-bottom: 18px;
	margin-top: 9px;
}
</style>
