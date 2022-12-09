/** * Created by lei on 2022/12/09 */
<template>
  <div>
    <el-dialog
      v-model="visible"
      @close="handleClose(formRef)"
      :title="title"
      style="width: 500px"
    >
      <el-form
        label-width="100px"
        ref="formRef"
        :rules="formRules"
        :model="data.metadata"
        style="max-width: 400px"
      >
        <el-form-item label="名字:" prop="name" class="name">
          <el-input v-model="data.metadata.name" />
        </el-form-item>
        <span class="info"
          >长度为 1 ~ 63 个字符，只能包含数字、小写字母和中</span
        >
        <span class="info">划线（-），且首尾只能是字母或数字</span>
        <el-form-item label="标签:" prop="data.metadata.labels" class="label">
          sdfasdf
        </el-form-item>
      </el-form>
      <Test />
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handleClose(formRef)">取消</el-button>
          <el-button type="primary" @click="btnOk(formRef)"> 确定 </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { NamespaceFormData, Namespace } from '@/type/namespace'
import type { FormRules, FormInstance } from 'element-plus'
import { reactive, watch, ref } from 'vue'
import Test from './test.vue'
const props = defineProps<{
  visible: boolean
  title: string
  namespace?: Namespace
}>()
const emits = defineEmits(['update:visible', 'valueChange'])
const formRef = ref<FormInstance>()

const data = reactive(new NamespaceFormData())
const formRules = reactive<FormRules>({
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }]
})

const handleClose = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  emits('update:visible', false)
}

const btnOk = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  console.log(data.metadata)
}

watch(
  () => {
    props.namespace
  },
  () => {
    if (props.namespace) {
      data.metadata.name = props.namespace.metadata.name
      data.metadata.labels = props.namespace.metadata.labels
    }
  }
)
</script>

<style scoped lang="less">
.info {
  font-size: 7px;
  margin-left: 100px;
  color: #999 !important;
  margin-top: 5px;
}
.name {
  margin-bottom: 5px;
}
.label {
  margin-top: 10px;
}
.table {
  thead {
    background-color: #eff3f8;
    font-size: 12px;
    text-align: left;
    padding: 0px;
    tr {
      padding: 0px;
      height: 15px;
      text-right {
        width: 100px;
      }
    }
  }
}
</style>
