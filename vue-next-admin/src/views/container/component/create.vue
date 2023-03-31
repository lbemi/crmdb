/** * Created by lei on 2022/11/16 */
<template>
	<div class="system-menu-dialog-container">
		<el-dialog  v-model="dialogVisable"  :close="handleClose" width="500px">
    <template #header="{ titleId, titleClass }">
      <div class="my-header">
        <h4 :id="titleId" :class="titleClass">{{ title }}</h4>
      </div>
    </template>
    <div class="dialog-body">
      <el-form
        ref="ruleFormRef"
        :model="state.cluster"
        status-icon
        label-width="80px"
        class="demo-ruleForm"
      >
        <el-form-item label="集群名称" prop="name">
          <el-input v-model="state.cluster.name" autocomplete="off" />
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
            >创建</el-button
          >
          <el-button @click="resetForm(ruleFormRef)">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-dialog>
  </div>
</template>

<script setup lang="ts" name="crateCluster">
import { reactive, ref, toRefs } from 'vue'
import { UploadFilled } from '@element-plus/icons-vue'
import { ElMessage, UploadFile } from 'element-plus'
import type { FormInstance } from 'element-plus'
import {ClusterForm} from "/@/types/cluster";
import { useClusterApi } from '/@/api/kubernetes/cluster';

const clusterApi = useClusterApi();
const ruleFormRef = ref<FormInstance>()

const state = reactive<ClusterForm>({
  cluster: {
    name: '',
  }
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
      newFormData.append('name', state.cluster.name)
      await clusterApi.createCluster(newFormData, requestConfig.headers)
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
