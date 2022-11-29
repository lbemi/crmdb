import { defineStore } from "pinia";
import { reactive, ref } from "vue";
import { clusterInfo } from "@/type/container";

export const kubeStore = defineStore(
  "kubeStore",
  () => {
    const activeCluster = ref<clusterInfo>();
    const clusters = ref<Array<clusterInfo>>();

    return { activeCluster, clusters };
  },
  {
    persist: {
      storage: localStorage,
    },
  }
);
