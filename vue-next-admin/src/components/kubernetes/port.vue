<template>
	<div>
		<el-form-item label="端口设置：" style="margin-bottom: 0">
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="pushPort">新增</el-button>
		</el-form-item>
		<el-form-item>
			<div>
				<el-form :key="portIndex" v-for="(item, portIndex) in data.ports" style="display: flex">
					<el-form-item label="名称">
						<el-input v-model="item.name" size="small" style="width: 120px" />
					</el-form-item>
					<el-form-item label="容器端口" style="margin-left: 10px">
						<el-input-number v-model="item.containerPort" size="small" />
					</el-form-item>
					<el-form-item label="协议" style="margin-left: 10px">
						<el-select v-model="item.protocol" size="small" style="width: 80px">
							<el-option v-for="item in protocolType" :key="item.type" :label="item.type" :value="item.value" />
						</el-select>
					</el-form-item>
					<el-form-item>
						<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.ports.splice(portIndex, 1)"></el-button>
					</el-form-item>
				</el-form>
			</div>
		</el-form-item>
	</div>
</template>

<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { ContainerPort } from 'kubernetes-types/core/v1';
import { reactive, watch } from 'vue';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;
import { deepClone } from '@/utils/other';

const data = reactive({
	loadFromParent: false,
	ports: <Array<ContainerPort>>[],
});

const props = defineProps({
	ports: Array<ContainerPort>,
});

const emit = defineEmits(['updatePort']);
const pushPort = () => {
	const name = uuid().toString().split('-')[1];
	data.ports.push({ name: 'p-' + name, containerPort: 80, protocol: 'TCP' });
};

watch(
	() => props.ports,
	() => {
		if (props.ports) {
			data.loadFromParent = true;
			data.ports = deepClone(props.ports) as Array<ContainerPort>;
			setTimeout(() => {
				data.loadFromParent = false;
			}, 2);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);

watch(
	() => data.ports,
	() => {
		if (!data.loadFromParent) {
			emit('updatePort', data.ports);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);

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
