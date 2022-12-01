/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="props.visible"
    @close="handleClose(userFormRef)"
    :title="title"
    style="width: 400px"
  >
    <el-form
      label-width="100px"
      ref="userFormRef"
      :rules="userFormRules"
      :model="form"
      style="max-width: 300px"
    >
      <el-form-item label="名字:" prop="user_name">
        <el-input v-model="form.user_name" />
      </el-form-item>
      <el-form-item label="描述:" prop="description">
        <el-input v-model="form.description" />
      </el-form-item>
      <el-form-item label="邮箱:" prop="email" autocomplete="off">
        <el-input v-model="form.email" />
      </el-form-item>
      <el-form-item
        label="密码"
        required
        prop="password"
        autocomplete="off"
        v-if="!props.data"
      >
        <el-input v-model="form.password" type="password" show-password />
      </el-form-item>
      <el-form-item
        label="确认密码"
        required
        prop="confirmPassword"
        autocomplete="off"
        v-if="!props.data"
      >
        <el-input
          v-model="form.confirmPassword"
          type="password"
          show-password
        />
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
        <el-button @click="handleClose(userFormRef)">取消</el-button>
        <el-button type="primary" @click="btnOk(userFormRef)"> 确定 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { UserForm, UserInfo } from '@/type/sys'
import { FormInstance, FormRules, ElMessage } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import { userApi } from '../../api'

const userFormRef = ref<FormInstance>()

const props = defineProps<{
  visible: boolean
  // v-model:visible.isBt  //v-model传值方式
  // visibleModifiers?: {
  //   isBt:boolean
  // }
  title: string
  data?: UserInfo | undefined
}>()

const emits = defineEmits(['update:visible', 'valueChange'])

const handleClose = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  emits('update:visible', false)
}

let form = reactive<UserForm>({
  user_name: '',
  email: '',
  status: 1,
  password: '',
  description: '',
  confirmPassword: ''
})
const validatePass = (rule: any, value: string, callback: any) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else {
    if (form.confirmPassword !== '') {
      if (!userFormRef.value) return
      userFormRef.value.validateField('confirmPassword', () => null)
    }
    callback()
  }
}
const validatePass2 = (rule: any, value: string, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== form.password) {
    callback(new Error('两次密码不匹配'))
  } else {
    callback()
  }
}

const validateEmail = (rule: any, value: string, callback: any) => {
  const regEmail =
    /^(([^<>()\\[\]\\.,;:\s@"]+(\.[^<>()\\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
  if (value) {
    if (regEmail.test(value)) {
      return callback()
    }
    callback(new Error('请输入正确的邮箱'))
  }
}

const userFormRules = reactive<FormRules>({
  user_name: [{ required: true, message: '请输入名字', trigger: 'blur' }],
  password: [
    { required: true, validator: validatePass, trigger: 'blur' },
    { min: 6, max: 12, message: '密码长度在6到12位之间', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, validator: validatePass2, trigger: 'blur' }
  ],
  email: [{ validator: validateEmail, trigger: 'blur' }]
})

const btnOk = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (!props.data) {
        userApi.addUser.request(form).then((res) => {
          handleClose(userFormRef.value)
          emits('valueChange')
          ElMessage.success(res.message)
        })
      } else {
        userApi.updateUser.request({ id: props.data.id }, form).then((res) => {
          handleClose(userFormRef.value)
          emits('valueChange')
          ElMessage.success(res.message)
        })
      }
    } else {
      ElMessage.error('请正确填写')
    }
  })
}

watch(
  () => props.data,
  () => {
    form.user_name = props.data!.user_name
    form.description = props.data!.description
    form.email = props.data!.email
    form.status = props.data!.status
  }
)
</script>

<style scoped lang="less"></style>
