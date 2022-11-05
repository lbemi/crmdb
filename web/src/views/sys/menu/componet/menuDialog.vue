/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="visible"
    @close="handleClose(menuFormRef)"
    :title="title"
    style="width: 400px"
  >
    <el-form
      label-width="100px"
      ref="menuFormRef"
      :rules="menuFormRules"
      :model="form"
      style="max-width: 300px"
    >
      <el-form-item label="菜单类型:">
        <el-radio-group v-model.number="form.menu_type">
          <el-radio-button label="1">菜单</el-radio-button>
          <el-radio-button label="2">按钮</el-radio-button>
          <el-radio-button label="3">API</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="名称:" prop="name" required>
        <el-input v-model="form.name" />
      </el-form-item>

      <el-form-item label="上一级:">
        <el-tree-select
          ref="selectRef"
          v-model="form.parent_id"
          :data="menuList"
          :props="{ value: 'id', label: 'name' }"
          check-strictly
          clearable
        />
      </el-form-item>

      <el-form-item label="描述:">
        <el-input v-model="form.memo" required="" />
      </el-form-item>
      <el-form-item label="排序值:">
        <el-input-number v-model="form.sequence" :min="1" :max="10" />
      </el-form-item>
      <el-form-item label="URL:">
        <el-input v-model="form.url" />
      </el-form-item>

      <el-form-item
        v-if="form.menu_type == 3 || form.menu_type == 2"
        id="ss"
        v-model="form.menu_type"
        label="请求方式:"
      >
        <el-select
          v-model="form.method"
          clearable
          placeholder="请选择"
          value-key="value"
        >
          <el-option
            v-for="item in methodOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>

      <el-form-item
        v-if="form.menu_type == 1"
        v-model="form.menu_type"
        label="图标:"
      >
        <el-input v-model="form.icon" required="" />
      </el-form-item>
      <el-form-item
        v-if="form.menu_type == 2 "
        v-model="form.menu_type"
        label="Code:"
      >
        <el-input v-model="form.code" required="" />
      </el-form-item>
      <el-form-item label="状态:">
        <el-switch
          v-model="form.status"
          class="ml-2"
          style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
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
        <el-button @click="handleClose(menuFormRef)">取消</el-button>
        <el-button type="primary" @click="btnOk(menuFormRef)"> 确定 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { UserForm, MenuInfo, MenuFrom } from "@/type/sys";
import { FormInstance, FormRules, ElMessage } from "element-plus";
import { ref, reactive, watch } from "vue";
import { menuApi } from "../../api";

const methodOptions = [
  {
    value: "GET",
    label: "GET",
  },
  {
    value: "POST",
    label: "POST",
  },
  {
    value: "DELETE",
    label: "DELETE",
  },
  {
    value: "PUT",
    label: "PUT",
  },
];
const menuFormRef = ref<FormInstance>();

const props = defineProps<{
  visible: boolean;
  // v-model:visible.isBt  //v-model传值方式
  // visibleModifiers?: {
  //   isBt:boolean
  // }
  title: string;
  data?: MenuInfo | undefined;
  menuList: Array<MenuInfo>;
}>();

const emits = defineEmits(["update:visible", "valueChange"]);

const handleClose = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  emits("update:visible", false);
};

let form = reactive<MenuFrom>({
  code: "",
  status: 1,
  memo: "",
  url: "",
  name: "",
  sequence: 1,
  menu_type: 1,
  method: "",
});

const menuFormRules = reactive<FormRules>({
  name: [{ required: true, message: "请输入名字", trigger: "blur" }],
});

const btnOk = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (!props.data) {
        menuApi.add.request(form).then((res) => {
          handleClose(menuFormRef.value);
          emits("valueChange");
          ElMessage.success(res.message);
        });
      } else {
        console.log(form);
        menuApi.update.request({ id: props.data.id }, form).then((res) => {
          handleClose(menuFormRef.value);
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
    form.menu_type = props.data!.menu_type;
    form.name = props.data!.name;
    form.memo = props.data!.memo;
    form.status = props.data!.status;
    form.sequence = props.data!.sequence;
    form.parent_id = props.data?.parent_id;
    form.icon = props.data?.icon;
    form.method = props.data?.method;
    form.url = props.data!.url;
    form.code = props.data?.code;
  }
);
</script>

<style scoped lang="less"></style>
