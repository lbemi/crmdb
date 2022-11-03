/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="visible"
    @close="handleClose"
    title="Shipping address"
  >
    <el-form label-width="100px" :model="userForm" style="max-width: 260px">
      <el-form-item label="名字:" required>
        <el-input v-model="userForm?.user_name" disabled />
      </el-form-item>
      <el-form-item label="描述:">
        <el-input v-model="userForm?.description" required="" />
      </el-form-item>
      <el-form-item label="邮箱:">
        <el-input v-model="userForm?.email" />
      </el-form-item>
      <el-form-item label="状态:">
        <el-switch
          v-model="userForm?.status"
          class="ml-2"
          style="--el-switch-on-color: #409eff; --el-switch-off-color: #ff4949"
          :active-value="1"
          :inactive-value="0"
          size="large"
          inline-prompt
          active-text="启用"
          inactive-text="禁用"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="btnOk"> 确定 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { UserForm } from "@/type/user";
import { ref } from "vue";

const emits = defineEmits(["update:visible", "valueChange"]);

const handleClose = () => {
  emits("update:visible",false);
};

const props = defineProps({
  userFormVisible: {
    type: Boolean,
    default: false
  }
});
const userForm = ref<UserForm>();

const btnOk = () => {};
</script>

<style scoped lang="less"></style>
