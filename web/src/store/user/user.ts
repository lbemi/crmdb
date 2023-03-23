import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { userApi } from '@/views/sys/api'
import router from '@/router/index'
import { ElMessage, FormInstance } from 'element-plus'
import { LoginReq, User } from '../interface/user'
import { useStore } from '../usestore'

export const userStore = defineStore('userStore', () => {
  const ruleFormRef = ref<FormInstance>()
  const ruleForm = reactive<LoginReq>({
    user_name: 'admin',
    password: 'admin',
    captcha: '12345',
    captcha_id: ''
  })

  const userInfo = ref<User>()
  const loginFn = () => {
    ruleFormRef.value
      ?.validate()
      .then(async () => {
        await userApi.login
          .request({
            ...ruleForm
          })
          .then((res) => {
            // 存储token
            //  store.getUserPermissions()
            localStorage.setItem('token', res.data.token)
            userInfo.value = res.data.user
            // store.getLeftMenusApi().then((res) => {
            router.push('/home')
            // });
            ElMessage.success(res.message)
            const userStore = useStore()
            userStore.getUserPermissions()
          })
      })
      .catch(() => {})
  }

  return { ruleFormRef, ruleForm, userInfo, loginFn }
})
