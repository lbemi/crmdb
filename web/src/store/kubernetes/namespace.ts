import { defineStore } from "pinia";
import { reactive, ref } from "vue";
import { Namespace } from "@/type/namespace";
import { namespacerApi } from "@/views/kubernetes/api";

export const nsStore = defineStore(
  "nsStore",
  () => {
        const namespace= ref<Array<Namespace>>([])
        const query = reactive({
            cloud: "",
          })

         const listNamespace =async (clusterName: string)=>{
            query.cloud = clusterName
            const res=  await namespacerApi.list.request(query)
            namespace.value = res.data.items
            console.log(namespace)
          }
    return { namespace, listNamespace };
  },
  {
    persist: {
      storage: localStorage,
    },
  }
);
