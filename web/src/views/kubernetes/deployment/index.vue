/**
 * Created by lei on 2022/11/29
 */
<template>
  <div>
    Deployment
  </div>
</template>

<script setup lang="ts">
  import { reactive,onMounted } from "vue";
import { deploymentApi} from "../api"
  
onMounted(()=>{
  listDeployment()
})

  const query = reactive({
    namespace: "default",
    cloud: "test2"
  })
  const listDeployment = async () => {
    const res = await deploymentApi.list.request(query)
    console.log("+++++",res);
  } 
</script>

<style scoped lang="less">

</style>
