/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="visible"
    @close="handleClose(roleFormRef)"
    :title="title"
    style="width: 400px"
  >
    <el-form
      label-width="100px"
      ref="roleFormRef"
      :rules="roleFormRules"
      :model="form"
      style="max-width: 300px"
    >
      <el-form-item label="名字:" prop="name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="描述:" prop="memo">
        <el-input v-model="form.memo" />
      </el-form-item>
      <el-form-item label="排序:" prop="sequence">
        <el-input-number v-model="form.sequence" :min="1" :max="10" />
      </el-form-item>
      <el-form-item label="父角色:">
        <el-select
          ref="selectRef"
          v-model="form.parent_id"
          clearable
          placeholder="请选择"
          value-key="id"
        >
          <el-option
            v-for="item in roleList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
            :disabled="item.name === form.name ? true : false"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="状态:">
        <el-switch
          v-model="form.status"
          class="ml-2"
          style="--el-switch-on-color: #409eff; --el-switch-off-color: #ff4949"
          :active-value="1"
          :inactive-value="2"
          size="large"
          inline-prompt
          active-text="启用"
          inactive-text="禁用"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose(roleFormRef)">取消</el-button>
        <el-button type="primary" @click="btnOk(roleFormRef)"> 确定 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { UserForm, RoleInfo, RoleFrom } from "@/type/sys";
import { FormInstance, FormRules, ElMessage } from "element-plus";
import { ref, reactive, watch } from "vue";
import { roleApi } from "../../api";

const roleFormRef = ref<FormInstance>();

const props = defineProps<{
  visible: boolean;
  title: string;
  data?: RoleInfo | undefined;
  roleList: Array<RoleInfo>;
}>();

const emits = defineEmits(["update:visible", "valueChange"]);

const handleClose = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  emits("update:visible", false);
};

let form = reactive<RoleFrom>({
  name: "",
  status: 1,
  memo: "",
  sequence: 1,
});

const roleFormRules = reactive<FormRules>({
  name: [{ required: true, message: "请输入名字", trigger: "blur" }],
});

const btnOk = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (!props.data) {
        roleApi.add.request(form).then((res) => {
          handleClose(roleFormRef.value);
          emits("valueChange");
          ElMessage.success(res.message);
        });
      } else {
        roleApi.update.request({ id: props.data.id }, form).then((res) => {
          handleClose(roleFormRef.value);
          emits("valueChange");
          ElMessage.success(res.message);
        });
      }
    } else {
      ElMessage.error("请正确填写");
    }
  });
};

watch(
  () => props.data,
  () => {
    console.log("*****",props);
    form.name = props.data!.name;
    form.memo = props.data!.memo;
    form.status = props.data!.status;
    form.sequence = props.data!.sequence;
    form.parent_id = props.data!.parent_id;
  }
);
</script>

<style scoped lang="less"></style>
