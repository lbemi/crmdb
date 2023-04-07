<template>
	<div class="layout-pd">
		<el-card shadow="hover">
			<el-backtop :right="100" :bottom="100" />
			<el-row>
				<el-col :span="3">
					<div style="height: 100%">
						<el-steps :active="active" finish-status="success" direction="vertical" style="margin-top: 5px; height: 100%" align-center>
							<el-step title="基本信息" description="Some description" />
							<el-step title="容器配置" description="Some description" />
							<el-step title="高级配置" description="Some description" />
						</el-steps>
					</div>
				</el-col>
				<el-col :span="11">
					<div style="margin-top: 10px" id="0">

						<el-form :model="form" label-width="120px">
							<el-form-item label="应用名称">
								<el-input v-model="form.name" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="副本数量">
								<el-input v-model="form.replicas" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="类型">
								<el-input v-model="form.type" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="标签">
								<Label v-model:tableData="form.labels" @on-click="getLabels" />
							</el-form-item>
							<el-form-item label="注解">
								<el-input v-model="form.annotations" size="default" style="width: 220px" />
							</el-form-item>
						</el-form>
					</div>
					<el-divider />
					<div style="margin-top: 10px" id="1" ref="one">
						<el-form :model="form" label-width="120px">
							<el-form-item label="应用名称">
								<el-input v-model="form.name" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="副本数量">
								<el-input v-model="form.replicas" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="类型">
								<el-input v-model="form.type" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="标签">
								<Label v-model:tableData="form.labels" @on-click="getLabels" />
							</el-form-item>
							<el-form-item label="注解">
								<el-input v-model="form.annotations" size="default" style="width: 220px" />
							</el-form-item>
						</el-form>
					</div>
					<el-divider />
					<div style="margin-top: 10px" id="2">
						<el-form :model="form" label-width="120px">
							<el-form-item label="应用名称">
								<el-input v-model="form.name" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="副本数量">
								<el-input v-model="form.replicas" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="类型">
								<el-input v-model="form.type" size="default" style="width: 220px" />
							</el-form-item>
							<el-form-item label="标签">
								<Label v-model:tableData="form.labels" @on-click="getLabels" />
							</el-form-item>
							<el-form-item label="注解">
								<el-input v-model="form.annotations" size="default" style="width: 220px" />
							</el-form-item>
						</el-form>
					</div>
				</el-col>
				<el-col :span="10" style="margin-left: 20px">
					<div class="btn">
						<div>
							<el-link type="primary" :underline="false" @click="jumpTo(0)" style="font-size: 13px;">基础信息</el-link>
						</div>
						<div>
							<el-link type="primary" :underline="false" @click="jumpTo(1)" style="font-size: 13px; ">容器配置</el-link>
						</div>
						<div>
							<el-link type="primary" :underline="false" @click="jumpTo(2)" style="font-size: 13px; ">高级配置</el-link>
						</div>
						<el-button @click="next" style="margin-top: 5px" size="small">下一步</el-button>
						<el-button @click="next" style="margin-top: 5px" size="small">确认</el-button>
					</div>
				</el-col>
			</el-row>
		</el-card>
	</div>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive, ref } from 'vue';

const Label = defineAsyncComponent(() => import('/@/components/label/index.vue'));
const active = ref(0);

const one = ref();
const form = reactive({
	name: '',
	replicas: 0,
	type: 'deployment',
	labels: [],
	annotations: '',
});
const getLabels = (labels: any) => {
	form.labels = labels;
};
const jumpTo = (id) => {
	active.value = id
	document.getElementById(id).scrollIntoView(true);
};
const next = () => {
	console.log(form);
	document.getElementById(active.value + '').scrollIntoView(true);
	if (active.value++ > 2) active.value = 0;
};
</script>

<style scoped>
.d2 {
	min-width: 100%;
	height: 100%;
	position: relative;
	display: flex;
	justify-content: flex-end;
}
.btn {
	position: fixed;
	right: 120px;
	top: 150px;
}
</style>
