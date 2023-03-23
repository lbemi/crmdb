<template>
  <!-- header有两部分 -->
  <el-card :body-style="{ padding: '0px' }">
    <div class="header">
      <div style="display: flex; align-items: center; margin-left: 10px">
        <span style="margin-left: 10px; font-size: large">运维管理平台</span>
      </div>

      <div style="display: flex; margin-right: 10px; align-items: center">
        <span style="align-items: center; margin-right: 10px">
          欢迎 {{ userInfo?.user_name }}
        </span>
        <el-dropdown>
          <span>
            <el-avatar :size="30" :src="data.circleUrl" />
          </span>
          <template #dropdown>
            <div style="margin-top: 15px; margin-left: 20px; font-size: 18px">
              {{ userInfo?.user_name }}
            </div>
            <div style="margin-top: 10px; margin-left: 20px; color: #adb0bb">
              账号ID: {{ userInfo?.id }}
            </div>

            <el-dropdown-menu>
              <el-dropdown-item :icon="UserFilled" divided>
                账号信息
              </el-dropdown-item>
              <el-dropdown-item :icon="HelpFilled"> 访问管理 </el-dropdown-item>
              <el-dropdown-item :icon="Shop" divided disabled>
                帮助设置偏好
                <el-radio-group
                  v-model="data.radio"
                  size="small"
                  style="margin-left: 18px"
                >
                  <el-radio-button label="开启" />
                  <el-radio-button label="关闭" />
                </el-radio-group>
              </el-dropdown-item>
              <el-dropdown-item divided @click="logout">退出</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { loginStore } from '@/store/login'
import { storeToRefs } from 'pinia'
import { UserFilled, HelpFilled, Shop } from '@element-plus/icons-vue'
const login = loginStore()
const { userInfo } = storeToRefs(login)
const router = useRouter()
const data = reactive({
  circleUrl:
    'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
  radio: '开启',
  activeIndex: 1,
  showMessage: false,
  loginUser: 'admin',
  userId: '123456789'
})

const logout = () => {
  // 清除本地缓存的 token 和 account
  localStorage.clear()
  // 跳转到登陆页面
  router.push('/login')
}
</script>

<style scoped type="less">
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 10px;
  width: 100%;
  height: 45px;
}
</style>
