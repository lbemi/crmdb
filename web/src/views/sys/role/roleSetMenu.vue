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
      :data="menuList"
      :default-checked-keys="defaultCheckedMenus"
      show-checkbox
      highlight-current
      default-expand-all
    >
      <template #default="{ data: { name } }"> {{ name }}</template>
    </el-tree>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose()">取消</el-button>
        <el-button
          v-loading.fullscreen.lock="loading"
          type="primary"
          @click="btnOk()"
        >
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { MenuInfo } from '@/type/sys'
import { ElMessage } from 'element-plus'
import { ref, reactive, onUpdated } from 'vue'
import { roleApi, userApi } from '../api'
import { useStore } from '@/store/usestore'
const user = useStore()
const menusRef = ref()
const loading = ref<boolean>(false)
const props = defineProps<{
  visible: boolean
  title: string
  roleID: number
  menuList: Array<MenuInfo> | undefined
  defaultCheckedMenus: Array<number>
}>()

const data = reactive({
  menuForm: {
    menu_ids: []
  }
})

const emits = defineEmits(['update:visible', 'valueChange'])

const handleClose = () => {
  emits('update:visible', false)
}

const btnOk = async () => {
  loading.value = true
  const menuIds = menusRef.value
    .getCheckedKeys()
    .concat(menusRef.value.getHalfCheckedKeys())
  data.menuForm.menu_ids = menuIds
  roleApi.setRoleMenus
    .request({ id: props.roleID }, data.menuForm)
    .then((res) => {
      handleClose()
      emits('valueChange')
      user.getLeftMenus()
      user.getUserPermissions()
      loading.value = false
      ElMessage.success(res.message)
    })
}
</script>

<style scoped lang="less"></style>
