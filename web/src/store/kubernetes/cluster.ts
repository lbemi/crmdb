import { defineStore } from 'pinia'
import { ref } from 'vue'
import { clusterInfo } from '@/type/cluster'
import { clusterApi } from '@/views/container/api'

// export const clusterStore = defineStore(
//   'clusterStore',
//   () => {
//     const activeCluster = ref<string>('')
//     const clusters = ref<Array<clusterInfo>>()

//     const listCluster = async () => {
//       const res = await clusterApi.list.request()
//       clusters.value = res.data.items
//     }
//     return { activeCluster, clusters, listCluster }
//   },
//   {
//     persist: {
//       storage: localStorage
//     }
//   }
// )
