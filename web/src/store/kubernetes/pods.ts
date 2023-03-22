import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { clusterInfo } from '@/type/cluster'
import { podApi } from '@/views/kubernetes/api'
import { nsStore } from './namespace'
import { kubeStore } from './kubernetes'
import { Data } from '@/type/pod'
export const podStore = defineStore('pod', () => {
  const ns = nsStore()
  const kube = kubeStore()
  const clusters = ref<Array<clusterInfo>>()
  const data = reactive(new Data())

  const listPods = async () => {
    data.query.cloud = kube.activeCluster
    data.query.namespace = ns.activeNamespace
    const res = await podApi.list.request(data.query)
    data.pods = res.data.data
    data.total = res.data.total
  }

  return { clusters, listPods, data }
})

export default podStore
