<!-- <template>
	<div>
		<el-dialog v-model="data.dialogVisible" :title="props.title" width="30%" :before-close="handleClose">
			<el-input-number v-model="num" :min="1" :max="10" @change="handleChange" />
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="dialogVisible = false">Cancel</el-button>
					<el-button type="primary" @click="dialogVisible = false"> Confirm </el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';

const data = reactive({
	dialogVisible: false,
	replicaset: 0,
});

const props = defineProps({
  title: String,
	replicaset: Number,
	type: String,
});

watch((),()=>{
  data.dialogVisible = props.dialogVisible

})
</script>

<style scoped></style> -->
