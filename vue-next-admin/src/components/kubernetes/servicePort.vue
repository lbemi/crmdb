<template>
	<div>
		<el-form-item label="端口设置" style="margin-bottom: 5px">
			<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0" @click="pushPort">新增</el-button>
		</el-form-item>
		<el-form-item>
			<el-form :key="portIndex" v-for="(item, portIndex) in data.ports" style="display: flex">
				<el-form-item label="名称">
					<el-input v-model="item.name" size="small" style="width: 90px" />
				</el-form-item>

				<!-- <el-form-item label="节点端口" style="margin-left: 10px" v-if="props.serviceType === 'NodePort'">
						<el-input-number v-model="item.nodePort" size="small" placeholder="eg: 3000" :min="1" />
					</el-form-item> -->
				<el-form-item label="服务端口" style="margin-left: 10px">
					<el-input v-model.number="item.port" size="small" placeholder="eg: 80 " />
				</el-form-item>
				<el-form-item label="容器端口" style="margin-left: 10px">
					<el-input v-model="item.targetPort" size="small" placeholder="eg: 80 或 http" />
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
		</el-form-item>
	</div>
</template>

<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { Service, ServicePort } from 'kubernetes-types/core/v1';
import { reactive, watch } from 'vue';
import jsPlumb from 'jsplumb';
import uuid = jsPlumb.jsPlumbUtil.uuid;
import { deepClone } from '/@/utils/other';

const data = reactive({
	loadFromParent: false,
	ports: <Array<ServicePort>>[],
	headless: false,
});
const props = defineProps({
	ports: Array<ServicePort>,
	serviceType: String,
});
const emit = defineEmits(['updatePort']);
const pushPort = () => {
	const name = uuid().toString().split('-')[1];
	data.ports.push({ name: 'p-' + name, port: 0, protocol: 'TCP' });
};

watch(
	() => props.ports,
	() => {
		if (props.ports) {
			data.loadFromParent = true;
			data.ports = deepClone(props.ports) as Array<ServicePort>;
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

const parseTargetPort = () => {
	data.ports.forEach((item: ServicePort) => {
		if (item.targetPort && typeof item.targetPort === 'string' && !isNaN(parseInt(item.targetPort))) {
			item.targetPort = parseInt(item.targetPort);
		}
	});
};

watch(
	() => data.ports,
	() => {
		if (!data.loadFromParent) {
			parseTargetPort();
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
