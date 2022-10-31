import axios from "axios";
import { ElMessage } from "element-plus";
const instance = axios.create({
  baseURL: "http://127.0.0.1:8080/api/v1",
  // baseURL: "http://127.0.0.1:8090/",
  timeout: 15000,
});

export interface Response {
  code: number;
  result: any;
  message?: string;
}

// request请求中间件拦截器
instance.interceptors.request.use(
  (config) => {
    let token = localStorage.getItem("token");
    if (token) {
      config.headers = config.headers || {}; // 判断是会否有headers
      config.headers["Authorization"] = token;
    }
    return config;
  },
  (err) => {
    return Promise.reject(err);
  }
);

// response返回中间件拦截器
instance.interceptors.response.use(
  (result) => {
    const data: Response = result.data
    if (data.code == 200) {
      return data
    } else {
      ElMessage.error(result.data.message)
      return Promise.reject(data);
    }
    
  },
  (e: any) => {
    if (e.message) {
      // 对响应错误做点什么
      if (e.message.indexOf("timeout") != -1) {
        ElMessage.error("网络超时");
      } else if (e.message == "Network Error") {
        ElMessage.error("网络连接错误");
      } else if (e.message.indexOf("404")) {
        ElMessage.error("请求接口找不到");
      } else {
        if (e.response.data) ElMessage.error(e.response.data.message);
        else ElMessage.error("接口路径找不到");
      }
    }
    return Promise.reject(e);
  }
);

export default instance;
