/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="visible"
    @close="handleClose"
    :title="title"
    style="width: 400px"
  >
  <el-form label-width="100px" :model="form" style="max-width: 260px">
      <el-form-item label="名字:" required>
        <el-input v-model="form.user_name"  />
      </el-form-item>
      <el-form-item label="描述:">
        <el-input v-model="form.description" required="" />
      </el-form-item>
      <el-form-item label="邮箱:">
        <el-input v-model="form.email" />
      </el-form-item>
      <el-form-item label="密码" required prop="password">
        <el-input v-model="form.password" type="password" show-password />
      </el-form-item>
      <el-form-item label="状态:">
        <el-switch
          v-model="form.status"
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
import { ref, reactive } from "vue";

defineProps<{
  visible: boolean
  // v-model:visible.isBt  //v-model传值方式
  // visibleModifiers?: {
  //   isBt:boolean
  // }
  title: string
}>();

const emits = defineEmits(["update:visible", "valueChange"]);

const handleClose = () => {
  emits("update:visible", false);
};

const form = reactive<UserForm>({
  user_name: "",
	email: "",
	mobile: "",
	status: 1,
	password: "",
  description:""
});

const btnOk = () => {
  console.log(form);
  
};
</script>

<style scoped lang="less"></style>
