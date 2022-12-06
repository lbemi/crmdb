import { defineStore } from 'pinia'
import { ref } from 'vue'
import { clusterInfo } from '@/type/container'

export const deployStore = defineStore('deploy', () => {
  const activeCluster = ref<clusterInfo>()
  const clusters = ref<Array<clusterInfo>>()
  return { activeCluster, clusters }
})
