<template>
	<div>
		<el-form-item label="端口设置：" style="margin-bottom: 0">
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0"
				@click="pushPort">新增</el-button>
		</el-form-item>
		<el-form-item>
			<el-form :key="portIndex" v-for="(item, portIndex) in data.ports" :inline="true">
				<el-form-item label="名称">
					<el-input v-model="item.name" size="small" style="width: 120px" />
				</el-form-item>
				<el-form-item label="容器端口" style="margin-left: 10px">
					<el-input-number v-model="item.containerPort" size="small" :min="1" :max="65535" />
				</el-form-item>
				<el-form-item label="协议" style="margin-left: 10px">
					<el-select v-model="item.protocol" size="small" style="width: 80px">
						<el-option v-for="item in protocolType" :key="item.type" :label="item.type" :value="item.value" />
					</el-select>
					<el-button :icon="RemoveFilled" type="primary" size="small" text
						@click="data.ports.splice(portIndex, 1)"></el-button>
				</el-form-item>
				<el-form-item>
				</el-form-item>
			</el-form>
		</el-form-item>
	</div>
</template>

<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { ContainerPort } from 'kubernetes-types/core/v1';
import { onMounted, reactive } from 'vue';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;
import { deepClone } from '@/utils/other';

const data = reactive({
	loadFromParent: false,
	ports: <Array<ContainerPort>>[],
});

type propsType = {
	ports: Array<ContainerPort> | undefined;
};
const props = defineProps<propsType>();

const pushPort = () => {
	const name = uuid().toString().split('-')[1];
	data.ports.push({ name: 'p-' + name, containerPort: 80, protocol: 'TCP' });
};

onMounted(() => {
	if (props.ports != undefined) {
		data.ports = deepClone(props.ports) as Array<ContainerPort>;
	}
});

const returnPorts = () => {
	return data.ports;
};

defineExpose({
	returnPorts,
});

const protocolType = [
	{
		type: 'tcp',
		value: 'TCP',
	},
	{
		type: 'udp',
		value: 'UDP',
	},
];
</script>

<style scoped></style>
