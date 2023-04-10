<template>
		<el-form :key="portIndex" v-for="(item, portIndex) in data.labels" style="display: flex">
			<el-form-item label="key">
				<el-input placeholder="key" v-model="item.key" size="small" style="width: 120px" />
			</el-form-item>
			<el-form-item label="value" style="margin-left: 10px">
				<el-input placeholder="value" v-model="item.value" size="small" />
			</el-form-item>
			<el-form-item>
				<el-button :icon="RemoveFilled" type="primary" size="small" text @click="data.labels.splice(portIndex, 1)"></el-button>
			</el-form-item>
		</el-form>
</template>

<script setup lang="ts">
import { RemoveFilled } from '@element-plus/icons-vue';
import { reactive, watch } from 'vue';

interface label {
	key: string;
	value: string;
}
const data = reactive({
	labels: [] as label[],
});
const props = defineProps({
	labelsData: Array
});
const handleLabels = () => {
    const labelsTup = []
    Object.keys(data.labels).forEach((key)=>{
        labelsTup.push({key:data.labels[key]})
    })
    return labelsTup
};
const emit = defineEmits(['updateLabels'])
watch(
	() => props.labelsData,
	(value,oldValue) => {
      // console.log("----", value,"000:",oldValue)
      // data.labels.push({key:'',value:''})
		if (props.labelsData) {
        data.labels = props.labelsData;
			}
		},{
        immediate: true
    }
);
watch(
    () => data.labels,
    (value,oldValue) => {
        if(data.labels) {
            const labels = handleLabels()
            emit('updateLabels',labels)
        }
    },{
        immediate:true
    }
);
</script>

<style scoped></style>
