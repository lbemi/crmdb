import { deploymentApi } from './../../views/kubernetes/api'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { clusterInfo } from '@/type/cluster'
import { Deployment } from '@/type/deployment'
import { kubeStore } from './kubernetes'

export const deployStore = defineStore('deploy', () => {
  const kube = kubeStore()
  const activeCluster = ref<clusterInfo>()
  const clusters = ref<Array<clusterInfo>>()

  const deleteDeoloyments = (deploymentNames: [Deployment]) => {
    deploymentNames.forEach((item) => {
      deploymentApi.delete.request({
        cloud: kube.activeCluster,
        namespace: item.metadata.namespace,
        deploymentName: item.metadata.name
      })
    })
  }
  return { activeCluster, clusters, deleteDeoloyments }
})
