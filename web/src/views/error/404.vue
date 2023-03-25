<template>
  <div
    style="
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100%;
"
  >
    <div>
      <img src="@/assets/image/404.png" />
    </div>
    <div>
      <div style=" align-items: center;text-align: center;">
        <div style="font-size: 80px;">404</div>
        <div>地址输入有误，请重新输入地址~</div>
        <div>您可以先检查网址，然后重新输入</div>
        <div style="margin-top: 50px;">
          <el-button type="primary" round @click="onGoHome">返回首页</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()

const onGoHome = () => {
  router.push('/dashboard')
}
</script>
