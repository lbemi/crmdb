import { defineStore } from "pinia";
import { reactive, ref } from "vue";
import { userApi } from "@/views/sys/api";
import router from "@/router/index";
import { ElMessage, ElNotification, FormInstance } from "element-plus";
import { LoginReq, User } from "./interface/user";
import { useStore } from "./usestore";

export const loginStore = defineStore(
  "login",
  () => {
    const ruleFormRef = ref<FormInstance>();
    const ruleForm = reactive<LoginReq>({
      user_name: "admin",
      password: "admin",
      captcha: "12345",
      captcha_id: "",
    });

    const userInfo = ref<User>();
    const loginFn = () => {
      ruleFormRef.value
        ?.validate()
        .then(async () => {
          await userApi.login
            .request({
              ...ruleForm,
            })
            .then((res) => {
              localStorage.setItem("token", res.data.token);
              userInfo.value = res.data.user;
              router.push("/home");
              ElMessage.success(`欢迎${userInfo.value?.user_name}!`);
              const userStore = useStore();
              userStore.getUserPermissions();
            });
        })
        .catch(() => {});
    };

    return { ruleFormRef, ruleForm, userInfo, loginFn };
  },
  {
    persist: {
      storage: localStorage,
      paths: ["userInfo"],
    },
  }
);
