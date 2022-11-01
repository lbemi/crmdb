<template>
  <!-- header有两部分 -->
  <!-- <el-card style="display: flex; width: 100%;height: 180px;" > -->
    <!-- <img src="../../../static/huawei.logo.png" /> -->
    <div>
      <span style="font-size: large; margin-left: 30px">运维管理平台</span>
    </div>
    <!-- <img src="../../../static/huawei.logo.png" /> -->
  
  <div style="display: flex; align-items: center">
    <span
      style="
        font-size: large;
        margin-right: 20px;
        margin-top: 8px;
        cursor: pointer;
      "
      @click="handleMessage"
    >
      <el-icon>
        <component is="Message"></component>
      </el-icon>
    </span>

    <el-drawer
      v-model="data.showMessage"
      title="站内信"
      direction="rtl"
      size="20%"
    >
      <div></div>
    </el-drawer>

    <el-dropdown>
      <span
        style="
          font-size: small;
          margin-right: 25px;
          color: #adb0bb;
          cursor: pointer;
        "
      >
        工具
        <el-icon>
          <component is="CaretBottom"></component>
        </el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>百宝箱</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>

    <el-dropdown>
      <span
        style="
          font-size: small;
          margin-right: 30px;
          color: #adb0bb;
          cursor: pointer;
        "
      >
        支持
        <el-icon>
          <component is="CaretBottom"></component>
        </el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>售后支持</el-dropdown-item>
          <el-dropdown-item>待办</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>

    <div
      style="
        vertical-align: middle;
        margin-top: 2px;
        margin-right: 28px;
        cursor: pointer;
      "
    >
      <el-dropdown>
        <span>
          <el-avatar :size="30" :src="data.circleUrl" />
        </span>

        <template #dropdown>
          <div style="margin-left: 20px; font-size: 18px; margin-top: 15px">
            {{ userInfo?.user_name }}
          </div>
          <div style="margin-left: 20px; margin-top: 10px; color: #adb0bb;">
            账号ID: {{ userInfo?.id }}
          </div>

          <el-dropdown-menu>
            <el-dropdown-item divided>
              <el-icon>
                <component is="UserFilled"></component>
              </el-icon>
              账号信息
            </el-dropdown-item>
            <el-dropdown-item>
              <el-icon> <component is="HelpFilled"></component> </el-icon
              >访问管理
            </el-dropdown-item>
            <el-dropdown-item divided disabled>
              <el-icon> <component is="Shop"></component></el-icon>帮助设置偏好
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
  
<!-- </el-card> -->
</template>

<script setup lang="ts">
import { reactive } from "vue";
import { useRouter } from "vue-router";
import { loginStore } from "@/store/login";
import { storeToRefs } from "pinia";
const login = loginStore();
const { userInfo } = storeToRefs(login);
const router = useRouter();
const data = reactive({
  circleUrl:
    "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
  radio: "开启",
  activeIndex: 1,
  showMessage: false,
  loginUser: "admin",
  userId: "123456789",
});

const handleMessage = () => {
  data.showMessage = true;
};
const logout = () => {
  // 清除本地缓存的 token 和 account
  localStorage.clear();
  // 跳转到登陆页面
  router.push("/login");
};
</script>

<style scoped>
.header-input {
  margin-right: 30px;
  width: 200px;
  transition: width 0.5s;
}
.header-input:hover {
  width: 350px;
}

.el-menu-class {
  height: 51px;
}

.header-bottom {
  font-size: small;
  margin-right: 25px;
  color: #adb0bb;
}
</style>
