/** * Created by lei on 2022/11/16 */
<template>
  <el-dialog v-model="dialogVisable" :close="handleClose" style="width: 500px">
    <template #header="{ titleId, titleClass }">
      <div class="my-header">
        <h4 :id="titleId" :class="titleClass">{{ title }}</h4>
        <el-divider />
      </div>
    </template>
    <div class="dialog-body">
      <el-form
        ref="ruleFormRef"
        :model="data.cluster"
        status-icon
        label-width="80px"
        class="demo-ruleForm"
      >
        <el-form-item label="集群名称" prop="name">
          <el-input v-model="data.cluster.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="配置文件">
          <el-upload
            :limit="1"
            drag
            :auto-upload="false"
            :on-change="handleChange"
            multiple
          >
            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
            <div class="el-upload__text">
              Drop file here or <em>click to upload</em>
            </div>
            <template #tip>
              <div class="el-upload__tip">
                kube config files with a size less than 500kb
              </div>
            </template>
          </el-upload>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm(ruleFormRef)"
            >Submit</el-button
          >
          <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref, toRefs } from 'vue'
import { UploadFilled } from '@element-plus/icons-vue'
import { ElMessage, UploadFile } from 'element-plus'
import type { FormInstance } from 'element-plus'
import { clusterForm } from '@/type/cluster'
import { clusterApi } from '../api'

const ruleFormRef = ref<FormInstance>()

const data = reactive({
  cluster: {
    name: ''
  } as clusterForm
})

const newFormData = new FormData()

const requestConfig = {
  headers: {
    'Content-Type': 'multipart/form-data'
  }
}
const handleChange = (file: UploadFile | undefined) => {
  if (!file) {
    ElMessage.error('请添加配置文件')
  } else {
    newFormData.append('file', file.raw || '')
  }
}

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      newFormData.append('name', data.cluster.name)
      await clusterApi.create
        .requestWithHeaders(newFormData, requestConfig.headers)
        .then((res) => {
          emits('valueChange')
          handleClose()
          ElMessage.success(res.message)
        })
        .catch(() => {
          newFormData.delete('file')
          newFormData.delete('name')
        })
    } else {
      ElMessage.error('请正确填写!')
      return false
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

const props = defineProps<{
  dialogVisable: boolean
  title: string
}>()

const { dialogVisable } = toRefs(props)
const emits = defineEmits(['update:dialogVisable', 'valueChange'])

const handleClose = () => {
  emits('update:dialogVisable', false)
}
</script>

<style scoped lang="less"></style>
