import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { clusterInfo } from '@/type/cluster'
import { podApi } from '@/views/kubernetes/api'
import { nsStore } from './namespace'
import { kubeStore } from './kubernetes'
import { Data, Pod } from '@/type/pod'
export const podStore = defineStore('pod', () => {
  const ns = nsStore()
  const kube = kubeStore()
  const clusters = ref<Array<clusterInfo>>()
  const data = reactive(new Data())
  const podShell = ref<Pod>()
  const listPods = async () => {
    data.query.cloud = kube.activeCluster
    data.query.namespace = ns.activeNamespace
    const res = await podApi.list.request(data.query)
    data.pods = res.data.data
    data.total = res.data.total
  }
  const deletPod = async (namespace: string, podName: string) => {
    await podApi.delete.request({ "cloud": kube.activeCluster, "namespace": namespace, "podName": [podName] })
  }

  return { clusters, data, podShell, listPods, deletPod }
},
  {
    persist: {
      storage: localStorage
    }
  })

export default podStore
