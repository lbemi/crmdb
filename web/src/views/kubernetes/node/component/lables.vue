/** * Created by lei on 2022/12/15 */
<template>
  <el-dialog
    v-model="visible"
    :title="title"
    width="30%"
    :before-close="handleClose"
  >
    <Lables v-model:tableData="labels" />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose()">Cancel</el-button>
        <el-button type="primary" @click="handleConfirm()"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import Lables from '@/component/label/index.vue'
import { Node } from '@/type/node'
import { ref, watch } from 'vue'
interface label {
  key: string
  value: string
}
const props = defineProps<{
  visible: boolean
  title: string
  data: Node
}>()

const labels = ref<label[]>([])

const emits = defineEmits(['update:visible', 'valuechange'])
const handleClose = () => {
  emits('update:visible', false)
}
const handleConfirm = () => {
  console.log('处理....')
  handleClose()
}
watch(
  () => props.data,
  () => {
    for (let key in props.data.metadata.labels) {
      const l: label = {
        key: key,
        value: props.data.metadata.labels[key]
      }
      labels.value.push(l)
    }
  }
)
</script>
<style scoped lang="less"></style>
