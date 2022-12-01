/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="visible"
    @close="handleClose()"
    :title="title"
    style="width: 400px;"
  >
    <el-tree
      ref="menusRef"
      node-key="id"
      :data="roleList"
      :default-checked-keys="defaultCheckedRoles"
      check-strictly
      show-checkbox
    >
      <template #default="{ data: { name } }"> {{ name }}</template>
    </el-tree>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose()">取消</el-button>
        <el-button type="primary" @click="btnOk()"> 确定 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { RoleInfo } from '@/type/sys'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { userApi } from '../api'

const menusRef = ref()

const props = defineProps<{
  visible: boolean
  title: string
  userID: number
  roleList: Array<RoleInfo> | undefined
  defaultCheckedRoles: Array<number>
}>()

const data = reactive({
  roleForm: {
    role_ids: []
  }
})

const emits = defineEmits(['update:visible', 'valueChange'])

const handleClose = () => {
  emits('update:visible', false)
}

const btnOk = async () => {
  const roleIds = menusRef.value.getCheckedKeys()
  data.roleForm.role_ids = roleIds
  userApi.setUserRole
    .request({ id: props.userID }, data.roleForm)
    .then((res) => {
      handleClose()
      emits('valueChange')
      ElMessage.success(res.message)
    })
}
</script>

<style scoped lang="less"></style>
