/**
 * Created by lei on 2022/09/24
 */
<template>
  <el-row :gutter="10" class="h-full">
    <el-col :lg="16" :md="12" class="bg-blue-400 ">
        <span class="flex h-screen justify-center items-center">asdhjahdkajshdka </span>
    </el-col>
    <el-col :lg="8" :md="12" >
      <div class="flex h-screen justify-center items-center">
        <span class="  flex block  justify-center mb-2 m-40">ajsdaskdas</span>
        <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="220px" class="login-content-form"
          size="large">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="ruleForm.username" type="text" autocomplete="off" />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="ruleForm.password" type="password" autocomplete="off" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="loginFn()" class="w-60">登录</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-col>
  </el-row>

</template>


<script setup lang="ts">
import { useStore } from '@/store/usestore';
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
const router = useRouter()
// 获取pinia存储对象
const store = useStore()

let { ruleForm } = toRefs(state)
let ruleFormRef = ref();
const loginFn = () => {
  ruleFormRef.value.validate().then(() => {
    adminLoginApi(
      {
        "name": ruleForm.value.username,
        "password": ruleForm.value.password
      }
    ).then(res => {
      // 存储token
      if (res.code === 200) {
        ElMessage.success(res.message)
        sessionStorage.setItem('token', 'Bearer ' + res.data.token)
        store.getLeftMenusApi().then(res => {
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
  password: [{ required: true, validator: validatePass, trigger: 'blur' }]
})


</script>

<style scoped lang="less">

</style>
