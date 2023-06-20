<template>
	<div>
		<el-form-item label="环境变量：" style="margin-bottom: 0">
			<el-button
				:icon="CirclePlusFilled"
				type="primary"
				size="small"
				text
				style="padding-left: 0"
				@click="data.env.push({ type: 'custom', name: '', value: '', otherValue: '' })"
				>新增</el-button
			>
		</el-form-item>
		<el-form-item>
			<el-table
				ref="envRef"
				:data="data.env"
				style="width: 100%; font-size: 10px"
				v-show="data.env.length != 0"
				:cell-style="{ padding: '5px' }"
				:header-cell-style="{ padding: '5px' }"
			>
				<el-table-column label="类型" width="130">
					<template #default="scope">
						<el-select v-model="scope.row.type" size="small">
							<el-option v-for="item in envType" :key="item.type" :label="item.type" :value="item.value" />
						</el-select>
					</template>
				</el-table-column>

				<el-table-column label="变量名称" width="180">
					<template #default="scope">
						<el-input v-model="scope.row.name" size="small" />
					</template>
				</el-table-column>
				<el-table-column label="变量/变量引用" width="290">
					<template #default="scope">
						<el-input v-model="scope.row.value" size="small" style="width: 120px" />
						<el-input v-model="scope.row.otherValue" size="small" v-if="scope.row.type != 'custom'" style="width: 120px; margin-left: 5px" />
					</template>
				</el-table-column>
				<el-table-column>
					<template #default="scope">
						<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.env.splice(scope.$index, 1)"></el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-form-item>
	</div>
</template>

<script setup lang="ts">
import { CirclePlusFilled, RemoveFilled } from '@element-plus/icons-vue';
import { EnvVar } from 'kubernetes-types/core/v1';
import { onMounted, reactive } from 'vue';
import { deepClone } from '@/utils/other';

// FIXME 资源引用有问题，待修复
interface envImp {
	name: string;
	value: string;
	otherValue: string;
	type: 'custom' | 'secretKeyRef' | 'resourceFieldRef' | 'fieldRef' | 'configMapKeyRef';
}
const data = reactive({
	loadFromParent: false,
	env: <Array<envImp>>[],
});

type propsType = {
	env: Array<EnvVar> | undefined;
};
const props = defineProps<propsType>();

const buildEnv = () => {
	const envData = [] as EnvVar[];
	const envTup = deepClone(data.env);
	envTup.forEach((item: any, index: number) => {
		if (item.type === 'custom') {
			//自定义变量
			const envVar: EnvVar = {
				name: item.name,
				value: item.value,
			};
			envData[index] = envVar;
		} else if (item.type === 'fieldRef') {
			const envVar: EnvVar = {
				name: item.name,
				valueFrom: {
					fieldRef: {
						fieldPath: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		} else if (item.type === 'resourceFieldRef') {
			const envVar: EnvVar = {
				name: item.name,
				valueFrom: {
					resourceFieldRef: {
						containerName: item.value,
						resource: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		} else if (item.type === 'configMapKeyRef') {
			const envVar: EnvVar = {
				name: item.name,
				valueFrom: {
					configMapKeyRef: {
						name: item.value,
						key: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		} else if (item.type === 'secretKeyRef') {
			const envVar: EnvVar = {
				name: item.name,
				valueFrom: {
					secretKeyRef: {
						name: item.value,
						key: item.otherValue,
					},
				},
			};
			envData[index] = envVar;
		}
	});

	return envData;
};

const parseEnv = (envs: Array<EnvVar>) => {
	const envData = [] as envImp[];
	envs.forEach((env, index) => {
		envData.push({
			name: '',
			value: '',
			type: 'custom',
			otherValue: '',
		});
		envData[index].name = env.name;
		if (env.valueFrom) {
			if (env.valueFrom.fieldRef) envData[index].type = 'fieldRef';
			if (env.valueFrom.fieldRef?.fieldPath) envData[index].value = env.valueFrom.fieldRef.fieldPath;

			if (env.valueFrom.secretKeyRef) envData[index].type = 'secretKeyRef';
			if (env.valueFrom.secretKeyRef?.key) envData[index].otherValue = env.valueFrom.secretKeyRef.key;
			if (env.valueFrom.secretKeyRef?.name) envData[index].value = env.valueFrom.secretKeyRef.name;

			if (env.valueFrom.configMapKeyRef) envData[index].type = 'configMapKeyRef';
			if (env.valueFrom.configMapKeyRef?.key) envData[index].otherValue = env.valueFrom.configMapKeyRef.key;
			if (env.valueFrom.configMapKeyRef?.name) envData[index].value = env.valueFrom.configMapKeyRef.name;

			if (env.valueFrom.resourceFieldRef) envData[index].type = 'resourceFieldRef';
			if (env.valueFrom.resourceFieldRef?.resource) envData[index].otherValue = env.valueFrom.resourceFieldRef.resource;
			if (env.valueFrom.resourceFieldRef?.containerName) envData[index].value = env.valueFrom.resourceFieldRef.containerName;
		}
		if (env.value) {
			envData[index].value = env.value;
			envData[index].type = 'custom';
		}
	});
	data.env = envData;
};
onMounted(() => {
	if (props.env != undefined) {
		data.loadFromParent = true;
		parseEnv(props.env);
		setTimeout(() => {
			data.loadFromParent = false;
		}, 100);
	}
});
// watch(
// 	() => props.env,
// 	() => {
// 		if (props.env) {
// 			data.loadFromParent = true;
// 			parseEnv(props.env);
// 			setTimeout(() => {
// 				data.loadFromParent = false;
// 			}, 100);
// 		}
// 	},
// 	{
// 		immediate: true,
// 		deep: true,
// 	}
// );
const returnEnvs = () => {
	const env = buildEnv();
	return env;
};

defineExpose({
	returnEnvs,
});
// watch(
// 	() => data.env,
// 	() => {
// 		if (!data.loadFromParent) {
// 			const env = buildEnv();
// 			emit('updateEnv', env);
// 		}
// 	},
// 	{
// 		immediate: true,
// 		deep: true,
// 	}
// );

const envType = [
	{
		type: '配置项',
		value: 'configMapKeyRef',
	},
	{
		type: '资源引用',
		value: 'fieldRef',
	},
	{
		type: '资源引用2',
		value: 'resourceFieldRef',
	},
	{
		type: '加密字典',
		value: 'secretKeyRef',
	},
	{
		type: '自定义',
		value: 'custom',
	},
];
</script>

<style scoped></style>
