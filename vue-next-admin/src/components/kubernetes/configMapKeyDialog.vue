<!-- <template>
	<div>
		<el-dialog v-model="dialogFormVisible" title="指定键：" width="300px">
			<el-form ref="dialogRef" :model="props." size="small">
				<el-form-item label="键名：">
					<el-select v-model="data.form.key" placeholder="选择key">
						<el-option v-for="(item, key, index) in data.keyValData" :value="key" :key="index" :label="key"> {{ key }} </el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="挂载路径：">
					<el-input v-model="data.form.path" autocomplete="off" />
					<a>使用相对路径</a>
				</el-form-item>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="dialogFormVisible = false" size="small">Cancel</el-button>
					<el-button type="primary" @click="handleConfirm" size="small"> Confirm </el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
const dialogFormVisible = ref(false);
const data = reactive({
	form: {
		key: '',
		path: '',
	},
});

const props = defineProps({
	configMap: Object,
	volumeIndex: Number,
	dialogFormVisible: Boolean,
});
</script>

<style scoped></style> -->
