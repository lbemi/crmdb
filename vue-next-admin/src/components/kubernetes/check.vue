<template>
	<div>
		<el-form-item>
			<el-checkbox v-model="data.set" label="开启" size="small" />
			<el-button v-if="data.show" type="info" v-show="data.set" text :icon="CaretTop" @click="data.show = !data.show"
				size="small" style="margin-left: 30px">隐藏</el-button>
			<el-button v-else type="info" v-show="data.set" text :icon="CaretBottom" @click="data.show = !data.show"
				size="small" style="margin-left: 30px">展开</el-button>
		</el-form-item>
		<el-tabs v-model="activeName" v-if="data.set" v-show="data.show">
			<el-tab-pane label="Http模式" name="httpGet">
				<el-form :model="data.probe.httpGet" label-width="130px" v-show="data.probe.httpGet">
					<el-form-item label="请求方式" prop="scheme">
						<el-select v-model="data.probe!.httpGet!.scheme">
							<el-option v-for="item in schemeType" :label="item.label" :key="item.label"
								:value="item.value" />
						</el-select>
					</el-form-item>
					<el-form-item label="路径">
						<el-input v-model="data.probe.httpGet!.path" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="端口">
						<el-input v-model.number="data.probe.httpGet!.port" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="Http头">
						<el-button :icon="CirclePlusFilled" type="primary" size="small" text style="padding-left: 0"
							@click="data.probe.httpGet?.httpHeaders?.push({ name: '', value: '' })">新增</el-button>
					</el-form-item>
					<el-form-item :key="index" v-for="(item, index) in data.probe.httpGet?.httpHeaders">
						<template #label> </template>
						<el-input v-model="item.name" placeholder="key" size="small" style="width: 100px" />
						<el-input v-model="item.value" placeholder="value" size="small"
							style="width: 100px; margin-left: 5px" />
						<el-button :icon="RemoveFilled" type="primary" size="small" text
							@click="data.probe.httpGet?.httpHeaders!.splice(index, 1)"></el-button>
					</el-form-item>

					<el-form-item label="延迟探测时间(s)">
						<el-input-number v-model="data.probe.initialDelaySeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="执行探测频率(s)">
						<el-input-number v-model="data.probe.periodSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="超时时间(s)">
						<el-input-number v-model="data.probe.timeoutSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="健康阀值(s)">
						<el-input-number v-model="data.probe.successThreshold" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="不健康阀值(s)">
						<el-input-number v-model="data.probe.failureThreshold" size="default" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
			<el-tab-pane label="TCP模式" name="tcpSocket">
				<el-form :model="data.probe.tcpSocket" label-width="120px" v-show="data.probe.tcpSocket">
					<el-form-item label="请求地址">
						<el-input v-model="data.probe.tcpSocket!.host" placeholder="一般不填写，默认为空" size="default"
							style="width: 200px" />
					</el-form-item>
					<el-form-item label="端口">
						<el-input v-model="data.probe.tcpSocket!.port" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="延迟探测时间(s)">
						<el-input-number v-model="data.probe.initialDelaySeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="执行探测频率(s)">
						<el-input-number v-model="data.probe.periodSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="超时时间(s)">
						<el-input-number v-model="data.probe.timeoutSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="健康阀值(s)">
						<el-input-number v-model="data.probe.successThreshold" size="default" style="width: 200px"
							disabled />
					</el-form-item>
					<el-form-item label="不健康阀值(s)">
						<el-input-number v-model="data.probe.failureThreshold" size="default" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
			<el-tab-pane label="Exec模式" name="exec">
				<el-form :model="data.probe.exec" label-width="120px" v-show="data.probe.exec">
					<el-form-item label="命令">
						<el-input v-model="data.probe!.exec!.command" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="延迟探测时间(s)">
						<el-input-number v-model="data.probe.initialDelaySeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="执行探测频率(s)">
						<el-input-number v-model="data.probe.periodSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="超时时间(s)">
						<el-input-number v-model="data.probe.timeoutSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="健康阀值(s)">
						<el-input-number v-model="data.probe.successThreshold" size="default" style="width: 200px"
							disabled />
					</el-form-item>
					<el-form-item label="不健康阀值(s)">
						<el-input-number v-model="data.probe.failureThreshold" size="default" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
		</el-tabs>
	</div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue';
import { CaretBottom, CaretTop } from '@element-plus/icons-vue';
import { Probe } from 'kubernetes-types/core/v1';
import { isObjectValueEqual } from '@/utils/arrayOperation';
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { deepClone } from '@/utils/other';
const data = reactive({
	loadFromParent: false,
	set: false,
	show: true,
	probe: <Probe>{
		httpGet: {
			// httpHeaders: [],
			scheme: 'HTTP',
			port: 80,
		},
		tcpSocket: {
			host: '',
			port: 80,
		},
		exec: {
			command: [],
		},
		successThreshold: 1,
		initialDelaySeconds: 3,
		periodSeconds: 10,
		timeoutSeconds: 1,
		failureThreshold: 3,
	},
});
const activeName = ref('httpGet');
const props = defineProps({
	checkData: Object,
});

onMounted(() => {
	if (!isObjectValueEqual(props.checkData, {}) && props.checkData != undefined) {
		data.set = true;
		const dataCopy = deepClone(props.checkData);
		if (dataCopy.httpGet && !isObjectValueEqual(dataCopy.httpGet, data.probe.httpGet)) {
			data.probe.httpGet = dataCopy.httpGet;
		} else if (dataCopy.tcpSocket && !isObjectValueEqual(dataCopy.tcpSocket, data.probe.tcpSocket)) {
			data.probe.tcpSocket = dataCopy.tcpSocket;
		} else if (dataCopy.exec && !isObjectValueEqual(dataCopy.exec, data.probe.exec)) {
			data.probe.exec = dataCopy.exec;
		}
		if (dataCopy.failureThreshold != data.probe.failureThreshold) {
			data.probe.failureThreshold = dataCopy.failureThreshold;
		}
		if (dataCopy.initialDelaySeconds != data.probe.initialDelaySeconds) {
			data.probe.initialDelaySeconds = dataCopy.initialDelaySeconds;
		}
		if (dataCopy.periodSeconds != data.probe.periodSeconds) {
			data.probe.periodSeconds = dataCopy.periodSeconds;
		}
		if (dataCopy.successThreshold != data.probe.successThreshold) {
			data.probe.successThreshold = dataCopy.successThreshold;
		}
		if (dataCopy.timeoutSeconds != data.probe.timeoutSeconds) {
			data.probe.timeoutSeconds = dataCopy.timeoutSeconds;
		}
	}
});
// watch(
// 	() => props.checkData,
// 	() => {
// 		// 数据不同则更新
// 		if (props.checkData && Object.keys(props.checkData).length != 0 && !isObjectValueEqual(props.checkData, data.probe)) {
// 			data.set = true;
// 			data.loadFromParent = true;
// 			const dataCopy = deepClone(props.checkData);
// 			if (dataCopy.httpGet && !isObjectValueEqual(dataCopy.httpGet, data.probe.httpGet)) {
// 				data.probe.httpGet = dataCopy.httpGet;
// 			} else if (dataCopy.tcpSocket && !isObjectValueEqual(dataCopy.tcpSocket, data.probe.tcpSocket)) {
// 				data.probe.tcpSocket = dataCopy.tcpSocket;
// 			} else if (dataCopy.exec && !isObjectValueEqual(dataCopy.exec, data.probe.exec)) {
// 				data.probe.exec = dataCopy.exec;
// 			}
// 			if (dataCopy.failureThreshold != data.probe.failureThreshold) {
// 				data.probe.failureThreshold = dataCopy.failureThreshold;
// 			}
// 			if (dataCopy.initialDelaySeconds != data.probe.initialDelaySeconds) {
// 				data.probe.initialDelaySeconds = dataCopy.initialDelaySeconds;
// 			}
// 			if (dataCopy.periodSeconds != data.probe.periodSeconds) {
// 				data.probe.periodSeconds = dataCopy.periodSeconds;
// 			}
// 			if (dataCopy.successThreshold != data.probe.successThreshold) {
// 				data.probe.successThreshold = dataCopy.successThreshold;
// 			}
// 			if (dataCopy.timeoutSeconds != data.probe.timeoutSeconds) {
// 				data.probe.timeoutSeconds = dataCopy.timeoutSeconds;
// 			}
// 			setTimeout(() => {
// 				data.loadFromParent = false;
// 			}, 1);
// 		}
// 	},
// 	{
// 		immediate: true,
// 		deep: true,
// 	}
// );
// const emit = defineEmits(['updateCheckData']);

const returnHealthCheck = () => {
	const copyData = deepClone(data);
	switch (activeName.value) {
		case 'httpGet': {
			delete copyData.probe.tcpSocket;
			delete copyData.probe.exec;
			break;
		}
		case 'tcpSocket': {
			delete copyData.probe.httpGet;
			delete copyData.probe.exec;
			break;
		}
		case 'exec': {
			delete copyData.probe.httpGet;
			delete copyData.probe.tcpSocket;
			break;
		}
	}
	return { set: data.set, probe: copyData.probe };
};

defineExpose({
	returnHealthCheck,
});
//
// watch(
// 	() => [data.probe, activeName, data.set],
// 	() => {
// 		if (!data.loadFromParent && data.set) {
// 			const copyData = deepClone(data);
// 			switch (activeName.value) {
// 				case 'httpGet': {
// 					delete copyData.probe.tcpSocket;
// 					delete copyData.probe.exec;
// 					break;
// 				}
// 				case 'tcpSocket': {
// 					delete copyData.probe.httpGet;
// 					delete copyData.probe.exec;
// 					break;
// 				}
// 				case 'exec': {
// 					delete copyData.probe.httpGet;
// 					delete copyData.probe.tcpSocket;
// 					break;
// 				}
// 			}
// 			emit('updateCheckData', copyData.probe);
// 		}
// 		if (!data.set) {
// 			emit('updateCheckData', {});
// 		}
// 	},
// 	{
// 		immediate: true,
// 		deep: true,
// 	}
// );
const schemeType = [
	{
		label: 'HTTP',
		value: 'HTTP',
	},
	{
		label: 'HTTPS',
		value: 'HTTPS',
	},
];
</script>

<style scoped>
.el-form-item {
	margin-bottom: 10px;
}
</style>
