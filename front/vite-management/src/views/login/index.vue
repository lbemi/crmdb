/** * Created by lei on 2022/09/24 */
<template>
  <el-row :gutter="10" class="h-full">
    <el-col :lg="12" :md="12" class="bg-blue-400">
      <span class="flex h-screen justify-center items-center">运维管理平台
      </span>
    </el-col>
    <el-col :lg="12" :md="12">
      <div class="flex h-screen justify-center items-center">
        <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="220px" class="login-content-form"
          size="large">
          <el-form-item label="用户名" prop="user_name">
            <el-input v-model="ruleForm.user_name" type="text" autocomplete="off" />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="ruleForm.password" type="password" autocomplete="off" />
          </el-form-item>
          <!-- 登录验证码 -->
          <el-form-item label="验证码" prop="captcha">
            <el-row class="w100">
              <el-col :span="16">
                <el-input v-model="ruleForm.captcha" clearable autocomplete="off" />
              </el-col>
              <el-col class="captcha-box" :span="8">
                <img @click="onChangeCaptcha" class="captcha-img" :src="ruleForm.pic_path" alt="captcha" />
              </el-col>
            </el-row>
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
import { useStore } from "@/store/usestore";
import { reactive, toRefs, ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { adminLoginApi, getUserCaptchaApi } from "@/request/api";
import { ElMessage } from "element-plus";
// import { ElMessage } from 'element-plus';
const state = reactive({
  ruleForm: {
    user_name: "admin",
    password: "admin",
    captcha: "",
    captcha_id: "",
    pic_path: ""
  },
});
//获取当前项目的路由对象
const router = useRouter();
// 获取pinia存储对象
const store = useStore();

onMounted(() => {
  onChangeCaptcha()
})

let { ruleForm } = toRefs(state);
let ruleFormRef = ref();
const loginFn = () => {
  ruleFormRef.value
    .validate()
    .then(async () => {
      await adminLoginApi({
        user_name: ruleForm.value.user_name,
        password: ruleForm.value.password,
        captcha: "14722",
        captcha_id: "yBRooA83JKBFs9akf7kJ",
      }).then((res) => {
        // 存储token
        store.getUserPermissions()
        if (res.code === 200) {
          ElMessage.success(res.message);
          sessionStorage.setItem("token", res.data.token);
          store.getLeftMenusApi().then((res) => {
            router.push("/home");
          });
        } else {
          ElMessage.error(res.message);
        }
      });
    })
    .catch(() => {
      ElMessage.error("校验不通过");
    });
};

const validatePass = (
  rule: unknown,
  value: string | undefined,
  callback: Function
): void => {
  if (!value) {
    callback("密码不能为空!");
  } else {
    callback();
  }
};
const rules = reactive({
  user_name: [{ required: true, message: "请输入用户名", tigger: "blur" }],
  password: [{ required: true, validator: validatePass, trigger: "blur" }],
});
const onChangeCaptcha = async () => {
  await getUserCaptchaApi().then((res) => {
    console.log(res);
    ruleForm.value.pic_path = res.data.pic_path
    ruleForm.value.captcha_id = res.data.captcha_id

  })
  console.log(ruleForm.value);
}
</script>

<style scoped lang="less">
.captcha-img {
  filter: brightness(61%);
}
.captcha-box {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    .captcha-img {
        width: 90%;
        margin-left: auto;
    }
    .el-button {
        width: 90%;
        height: 100%;
    }
}
</style>
