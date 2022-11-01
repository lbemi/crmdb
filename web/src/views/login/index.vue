/** * Created by lei on 2022/09/24 */
<template>
  <el-row :gutter="0" class="h-full">
    <el-col :lg="16" :md="12" class="bg-blue-400">
      <span class="flex h-screen justify-center items-center"
        >运维管理平台
      </span>
    </el-col>
    <el-col :lg="8" :md="12">
      <div class="flex h-screen justify-center items-center">
        <el-form
          :ref="(el:FormInstance) => (login.ruleFormRef = el)"
          :model="login.ruleForm"
          :rules="rules"
          label-width="170px"
          class="login-content-form"
          size="large"
        >
          <el-form-item label="用户名" prop="user_name">
            <el-input
              v-model="login.ruleForm.user_name"
              type="text"
              autocomplete="off"
            />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input
              v-model="login.ruleForm.password"
              type="password"
              autocomplete="off"
            />
          </el-form-item>
          <!-- 登录验证码 -->
          <el-form-item label="验证码" prop="captcha">
            <el-row>
              <el-col :span="16">
                <el-input
                  v-model="login.ruleForm.captcha"
                  clearable
                  autocomplete="off"
                />
              </el-col>
              <el-col class="captcha-box" :span="8">
                <img
                  @click="onChangeCaptcha"
                  class="captcha-img"
                  :src="pic_path"
                  alt="captcha"
                />
              </el-col>
            </el-row>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="login.loginFn()" class="w-60"
              >登录</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import { loginStore } from "@/store/login";
import { reactive, ref, onMounted } from "vue";
import { FormInstance, FormRules } from "element-plus";
import { userApi } from "@/request/sys/user";
const login = loginStore();
const pic_path = ref("");
onMounted(() => {
  onChangeCaptcha();
});

const onChangeCaptcha = async () => {
  let res: any = await userApi.captcha.request();
  pic_path.value = res.data.pic_path;
  login.ruleForm.captcha_id = res.data.captcha_id;
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
const rules = reactive<FormRules>({
  user_name: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [{ required: true, validator: validatePass, trigger: "blur" }],
});
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
}
</style>
