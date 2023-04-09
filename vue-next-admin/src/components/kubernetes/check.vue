<template>
	<div>
		<el-tabs v-model="activeName" @tabClick="handleEdit">
			<el-tab-pane label="Http模式" name="first">
				<el-form :model="data.checkData.httpGet" label-width="120px">
					<el-form-item label="路径">
						<el-input v-model="data.checkData.httpGet!.path" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="端口">
						<el-input v-model="data.checkData.httpGet!.port" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="Http头">
						<el-input v-model="data.checkData.httpGet!.httpHeaders![0].name" placeholder="key" size="default" style="width: 100px" />
						<el-input
							v-model="data.checkData.httpGet!.httpHeaders![0].value"
							placeholder="value"
							size="default"
							style="width: 100px; margin-left: 5px"
						/>
					</el-form-item>
					<el-form-item label="延迟探测时间(s)">
						<el-input-number v-model="data.checkData.initialDelaySeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="执行探测频率(s)">
						<el-input-number v-model="data.checkData.periodSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="超时时间(s)">
						<el-input-number v-model="data.checkData.timeoutSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="健康阀值(s)">
						<el-input-number v-model="data.checkData.successThreshold" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="不健康阀值(s)">
						<el-input-number v-model="data.checkData.failureThreshold" size="default" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
			<el-tab-pane label="TCP模式" name="second">
				<el-form :model="data.checkData" label-width="120px">
					<el-form-item label="端口">
						<el-input v-model="data.checkData.tcpSocket!.port" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="延迟探测时间(s)">
						<el-input-number v-model="data.checkData.initialDelaySeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="执行探测频率(s)">
						<el-input-number v-model="data.checkData.periodSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="超时时间(s)">
						<el-input-number v-model="data.checkData.timeoutSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="健康阀值(s)">
						<el-input-number v-model="data.checkData.successThreshold" size="default" style="width: 200px" disabled />
					</el-form-item>
					<el-form-item label="不健康阀值(s)">
						<el-input-number v-model="data.checkData.failureThreshold" size="default" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
			<el-tab-pane label="Exec模式" name="third">
				<el-form :model="data.checkData.exec" label-width="120px">
					<el-form-item label="命令">
						<el-input v-model="data.checkData.exec!.command" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="延迟探测时间(s)">
						<el-input-number v-model="data.checkData.initialDelaySeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="执行探测频率(s)">
						<el-input-number v-model="data.checkData.periodSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="超时时间(s)">
						<el-input-number v-model="data.checkData.timeoutSeconds" size="default" style="width: 200px" />
					</el-form-item>
					<el-form-item label="健康阀值(s)">
						<el-input-number v-model="data.checkData.successThreshold" size="default" style="width: 200px" disabled />
					</el-form-item>
					<el-form-item label="不健康阀值(s)">
						<el-input-number v-model="data.checkData.failureThreshold" size="default" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
		</el-tabs>
	</div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { V1HTTPGetAction, V1Probe, V1TCPSocketAction } from '@kubernetes/client-node';

const data = reactive({
	checkData: {
		httpGet: {
			httpHeaders: [
				{
					name: '',
					value: '',
				},
			],
		},
		tcpSocket: {},
		exec: {},
		successThreshold: 1,
		initialDelaySeconds: 3,
		periodSeconds: 10,
		timeoutSeconds: 1,
		failureThreshold: 3,
	} as V1Probe,
});
const activeName = ref('first');
const handleEdit = () => {
	if (activeName.value === 'first') {
		data.checkData.exec = {};
		data.checkData.tcpSocket = {} as V1TCPSocketAction;
		data.checkData.initialDelaySeconds = 3;
		data.checkData.periodSeconds = 10;
		data.checkData.timeoutSeconds = 1;
		data.checkData.failureThreshold = 3;
	} else if (activeName.value === 'second') {
		data.checkData.exec = {};
		data.checkData.httpGet = {
			httpHeaders: [
				{
					name: '',
					value: '',
				},
			],
		} as V1HTTPGetAction;
		data.checkData.initialDelaySeconds = 15;
		data.checkData.periodSeconds = 10;
		data.checkData.timeoutSeconds = 1;
		data.checkData.failureThreshold = 3;
	} else if (activeName.value === 'third') {
		data.checkData.tcpSocket = {} as V1TCPSocketAction;
		data.checkData.httpGet = {
			httpHeaders: [
				{
					name: '',
					value: '',
				},
			],
		} as V1HTTPGetAction;
		data.checkData.initialDelaySeconds = 5;
		data.checkData.periodSeconds = 10;
		data.checkData.timeoutSeconds = 1;
		data.checkData.failureThreshold = 3;
	}
};
const getData = () => {
	if (activeName.value != 'first') {
		data.checkData.httpGet = {} as V1TCPSocketAction;
	}
	return data.checkData;
};
defineExpose({
	data,
	getData,
});
</script>

<style scoped></style>
