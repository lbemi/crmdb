/**
 * Created by lei on 2022/09/24
 */
<template>
  <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="120px" class="login-content-form"
    size="large">
    <el-form-item label="用户名" prop="username">
      <el-input v-model="ruleForm.username" type="text" autocomplete="off" />
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input v-model="ruleForm.password" type="password" autocomplete="off" />
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="loginFn()">登录</el-button>
    </el-form-item>
  </el-form>
</template>


<script setup lang="ts">
import { useStore } from 'vuex';
import { reactive, toRefs, ref } from 'vue';
import { useRouter } from 'vue-router';
import { adminLoginApi } from '@/request/api'
import { ElMessage } from 'element-plus';
// import { ElMessage } from 'element-plus';
const state = reactive({
  ruleForm: {
    username: 'admin',
    password: 'admin'
  }
})
//获取当前项目的路由对象
let router = useRouter()
// 获取vuex存储对象
let store = useStore()

let { ruleForm } = toRefs(state)
let ruleFormRef = ref();
const loginFn = () => {
  ruleFormRef.value.validate().then(() => {
    adminLoginApi(
      {
        "user_name": ruleForm.value.username,
        "password": ruleForm.value.password
      }
    ).then(res => {
      // 存储token
      if (res.code === 2000) {
        ElMessage.success(res.message)
        sessionStorage.setItem('token',  res.data.token)
        store.dispatch('getLeftMenusApi').then(res => {
          router.push('/home')
        })
      } else {
        ElMessage.error(res.message)
      }
    })
  }).catch(() => {
    alert('校验不通过');
  })
}

const validatePass = (rule: unknown, value: string | undefined, callback: Function): void => {
  if (!value) {
    callback('密码不能为空!')
  } else {
    callback()
  }
}
const rules = reactive({
  username: [{ required: true, message: '请输入用户名', tigger: 'blur' }],
  password: [{ validator: validatePass, trigger: 'blur' }]
})


</script>

<style scoped lang="less">

</style>
