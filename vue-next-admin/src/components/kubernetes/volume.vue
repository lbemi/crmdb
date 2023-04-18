<template>
	<div>
		<el-form-item label-width="90px" style="margin-bottom: 0">
			<template #label>
				<el-tooltip
					class="box-item"
					effect="light"
					content="自定义容器启动时运行的命令; 默认情况下，容器启动时将运行镜像默认命令"
					placement="top-start"
					raw-content
				>
					数据卷：
				</el-tooltip>
			</template>
      <el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="handleSet">新增</el-button>
      <el-button
				v-if="data.show"
				type="info"
				v-show="data.volumeData.length !=0"
				text
				:icon="CaretTop"
				@click="data.show = !data.show"
				size="small"
				style="margin-left: 30px"
				>隐藏</el-button
			>
			<el-button v-else type="info"  text :icon="CaretBottom" @click="data.show = !data.show" size="small" style="margin-left: 30px"
				>展开</el-button
			>
		</el-form-item>
		<el-form-item  label-width="0">
			<el-table
				:data="data.volumeData"
        v-show="data.volumeData.length !=0 && data.show"
				style="width: 100%; font-size: 10px"
				:cell-style="{ padding: '0,5px' }"
				:row-style="{ padding: '2px' }"
				:header-cell-style="{ padding: '5px' }"
        :header-row-style="{ padding: '5px' }"
			>
				<el-table-column prop="" label="类型" width="130">
					<template #default>
						<el-select v-model="data.volumeType" size="small">
							<el-option v-for="item in data.typeList" :key="item.value" :label="item.label" :value="item.value" />
						</el-select>
					</template>
				</el-table-column>

				<el-table-column prop="name" label="名称" width="150">
					<template #default="scope">
						<el-input v-model="scope.row.name" size="small" />
					</template>
				</el-table-column>
				<el-table-column prop="" label="挂载源" width="150">
					<template #default="scope">
						<el-input v-model="scope.row.hostPath.path" size="small" placeholder="主机路径：/tmp"  v-show="data.volumeType === 'hostPath'"/>
					</template>
				</el-table-column>
				<el-table-column prop="mountPath" label="容器挂载路径" width="150">
					<template #default="scope">
						<el-input v-model="scope.row.volumeMountData.mountPath" size="small" placeholder="容器路径：/app" />
					</template>
				</el-table-column>
				<el-table-column prop="subPath" label="容器子路径" width="120">
					<template #default="scope">
						<el-input v-model="scope.row.volumeMountData.subPath" size="small" placeholder="默认为空" />
					</template>
				</el-table-column>
				<el-table-column>
					<template #default="scope">
						<el-button :icon="RemoveFilled" type="danger" size="small" text @click="data.volumeData.splice(scope.$index, 1)"></el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-form-item>
	</div>
</template>

<script setup lang="ts">
import {CaretBottom, CaretTop, CirclePlusFilled, Delete, Edit, RemoveFilled} from '@element-plus/icons-vue';
import { reactive, watch } from 'vue';
import { deepClone } from '/@/utils/other';
import { V1HostPathVolumeSource, V1Volume, V1VolumeMount } from '@kubernetes/client-node';
import { V1SecretVolumeSource } from '@kubernetes/client-node/dist/gen/model/v1SecretVolumeSource';
import { V1ConfigMapVolumeSource } from '@kubernetes/client-node/dist/gen/model/v1ConfigMapVolumeSource';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;

const data = reactive({
	set: false,
	show: true,
	volumeType: 'hostPath',
	volumeMountData: {
		name: '',
		path: '',
	},
	typeList: [
		{
			label: '主机目录',
			value: 'hostPath',
		},
		{
			label: '加密字典',
			value: 'secret',
		},
		{
			label: '配置项',
			value: 'configMap',
		},
	],
	volumeData: [],
	volumes: [] as V1Volume[],
});

const handleSet = () => {
	// data.set = !data.set;
	const name = 'volume-' + uuid().toString().split('-')[1];
	data.volumeData.push({
		type: 'hostPath',
		name: name,
		mountSource: {},
		containerPath: '',
		subPath: '',
		hostPath: {} as V1HostPathVolumeSource,
		secret: {} as V1SecretVolumeSource,
		configMap: {} as V1ConfigMapVolumeSource,
		volumeMountData: {
			name: name,
			mountPath: '',
			subPath: '',
		} as V1VolumeMount,
	});
};
const props = defineProps({});

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
// watch(
//     () => [props.args, props.commands],
//     () => {
//       if (props.args) {
//         data.args = handleArr(props.args);
//       }
//       if (props.commands) {
//         data.commands = handleArr(props.commands);
//       }
//     },
//     {
//       immediate: true,
//       deep: true,
//     }
// );
// const emit = defineEmits(['updateCommand']);
// watch(
//     () => [data.args, data.commands, data.commandSet],
//     () => {
//       if (!data.commandSet) {
//         data.k8s.args = [];
//         data.k8s.commands = [];
//       } else {
//         if (data.args) {
//           data.k8s.args = data.args.split(',');
//         }
//         if (props.commands) {
//           data.k8s.commands = data.commands.split(',');
//         }
//       }
//       emit('updateCommand', data.k8s);
//     },
//     {
//       immediate: true,
//       deep: true,
//     }
// );
</script>

<style scoped>
.el-card {
	border: none;
	box-shadow: none;
}
</style>
