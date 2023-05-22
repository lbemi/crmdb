<template>
	<div>
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
		<el-button v-else type="info" v-show="data.set" text :icon="CaretBottom" @click="data.show = !data.show" size="small" style="margin-left: 30px"
			>展开</el-button
		>
		<el-tabs v-model="activeName" v-if="data.set" v-show="data.show">
			<el-tab-pane label="Http模式" name="httpGet" v-if="data.lifeProbe.httpGet">
				<el-form :model="data.lifeProbe.httpGet" label-width="120px" v-show="data.lifeProbe.httpGet">
					<el-form-item label="请求方式" prop="scheme">
						<el-select v-model="data.lifeProbe.httpGet.scheme" size="small">
							<el-option v-for="item in schemeType" :label="item.label" :key="item.label" :value="item.value" />
						</el-select>
					</el-form-item>
					<el-form-item label="路径">
						<el-input v-model="data.lifeProbe.httpGet.path" size="small" style="width: 200px" />
					</el-form-item>
					<el-form-item label="端口">
						<el-input v-model.number="data.lifeProbe.httpGet.port" size="small" style="width: 200px" />
					</el-form-item>
					<el-form-item label="Http头">
						<el-button
							:icon="CirclePlusFilled"
							type="primary"
							size="small"
							text
							style="padding-left: 0"
							@click="data.lifeProbe?.httpGet?.httpHeaders?.push({ name: '', value: '' })"
							>新增</el-button
						>
					</el-form-item>
					<el-form-item :key="index" v-for="(item, index) in data.lifeProbe.httpGet.httpHeaders">
						<template #label> </template>
						<el-input v-model="item.name" placeholder="key" size="small" style="width: 100px" />
						<el-input v-model="item.value" placeholder="value" size="small" style="width: 100px; margin-left: 5px" />
						<el-button
							:icon="RemoveFilled"
							type="primary"
							size="small"
							text
							@click="data.lifeProbe?.httpGet?.httpHeaders?.splice(index, 1)"
						></el-button>
					</el-form-item>
				</el-form>
			</el-tab-pane>
			<el-tab-pane label="TCP模式" name="tcpSocket" v-if="data.lifeProbe.tcpSocket">
				<el-form :model="data.lifeProbe.tcpSocket" label-width="120px" v-show="data.lifeProbe.tcpSocket">
					<el-form-item label="请求地址">
						<el-input v-model="data.lifeProbe.tcpSocket.host" placeholder="一般不填写，默认为空" size="small" style="width: 200px" />
					</el-form-item>
					<el-form-item label="端口">
						<el-input v-model.number="data.lifeProbe.tcpSocket.port" size="small" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
			<el-tab-pane label="Exec模式" name="exec">
				<el-form :model="data.lifeProbe.exec" label-width="120px" v-show="data.lifeProbe.exec">
					<el-form-item label="命令">
						<el-input v-model="data.command" size="small" style="width: 200px" />
					</el-form-item>
				</el-form>
			</el-tab-pane>
		</el-tabs>
	</div>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { isObjectValueEqual } from '/@/utils/arrayOperation';
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { deepClone } from '/@/utils/other';
import { LifecycleHandler } from 'kubernetes-types/core/v1';
import { CaretBottom, CaretTop } from '@element-plus/icons-vue';

const data = reactive({
	loadFromParent: false,
	set: false,
	show: true,
	command: '',
	lifeProbe: {
		httpGet: {
			httpHeaders: [],
			scheme: 'HTTP',
			port: 0,
		},
		tcpSocket: {
			host: '',
			port: 0,
		},
		exec: {
			command: [''],
		},
	} as V1LifecycleHandler,
});
const activeName = ref('httpGet');
const props = defineProps({
	lifeData: Object,
});

watch(
	() => props.lifeData,
	() => {
		// 数据不同则更新
		if (props.lifeData && Object.keys(props.lifeData).length != 0 && !isObjectValueEqual(props.lifeData, data.lifeProbe)) {
			data.loadFromParent = true;
			data.set = true;
			const dataCopy = deepClone(props.lifeData);
			if (dataCopy.httpGet && !isObjectValueEqual(dataCopy.httpGet, data.lifeProbe.httpGet)) {
				data.lifeProbe.httpGet = dataCopy.httpGet;
			} else if (dataCopy.tcpSocket && !isObjectValueEqual(dataCopy.tcpSocket, data.lifeProbe.tcpSocket)) {
				data.lifeProbe.tcpSocket = dataCopy.tcpSocket;
			} else if (dataCopy.exec && !isObjectValueEqual(dataCopy.exec, data.lifeProbe.exec)) {
				let str = '';
				dataCopy.exec.command.forEach((item: any, index: number) => {
					if (index == dataCopy.exec.command.length - 1) {
						str = str + item;
					} else {
						str = str + item + ',';
					}
				});
				if (str != data.command) data.command = str;
			}
			setTimeout(() => {
				data.loadFromParent = false;
			}, 100);
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
const emit = defineEmits(['updateLifeData']);

watch(
	() => [data.lifeProbe, activeName, data.set, data.command],
	() => {
		if (!data.loadFromParent && data.set) {
			const copyData = deepClone(data);
			switch (activeName.value) {
				case 'httpGet': {
					delete copyData.lifeProbe.tcpSocket;
					delete copyData.lifeProbe.exec;
					break;
				}
				case 'tcpSocket': {
					delete copyData.lifeProbe.httpGet;
					delete copyData.lifeProbe.exec;
					break;
				}
				case 'exec': {
					if (data.command.indexOf(',')) {
						copyData.lifeProbe.exec.command = data.command.split(',');
					} else {
						copyData.lifeProbe.exec.command = data.command;
					}
					delete copyData.lifeProbe.httpGet;
					delete copyData.lifeProbe.tcpSocket;
					break;
				}
			}
			emit('updateLifeData', copyData.lifeProbe);
		}

		if (!data.set) {
			emit('updateLifeData', {});
		}
	},
	{
		immediate: true,
		deep: true,
	}
);
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

<style scoped></style>
