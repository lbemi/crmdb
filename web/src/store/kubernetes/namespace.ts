import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { Namespace } from '@/type/namespace'
import { namespacerApi } from '@/views/kubernetes/api'
import { kubeStore } from './kubernetes'

const kube = kubeStore()
export const nsStore = defineStore(
  'nsStore',
  () => {
    const namespace = ref<Array<Namespace>>([])
    const activeNamespace = ref<string>('default')
    const query = reactive({
      cloud: ''
    })

    const listNamespace = async () => {
      console.log('*(*(*(*', kube.activeCluster)

      if (kube.activeCluster) {
        query.cloud = kube.activeCluster
      } else {
        console.log('error')
        return
      }
      console.log('*(*(*(*', kube.activeCluster)
      const res = await namespacerApi.list.request(query)
      namespace.value = res.data.items
    }
    return { namespace, listNamespace, activeNamespace }
  },
  {
    persist: {
      storage: localStorage
    }
  }
)
