import { defineStore } from 'pinia'
import { ref } from 'vue'
import { clusterInfo } from '@/type/container'

export const kubeStore = defineStore(
  'kubeStore',
  () => {
    const activeCluster = ref<string>('')
    const clusters = ref<Array<clusterInfo>>()
    return { activeCluster, clusters }
  },
  {
    persist: {
      storage: localStorage
    }
  }
)
