import { defineStore } from "pinia";
import { reactive, ref } from "vue";
import { getUserLeftMenusApi, getUserPermissionApi } from "../request/api";

export interface MenuObj  {
  id: number;
  name: string;
  url: string;
  children: MenuObj[];
};


export const useStore = defineStore(
  "main",
  () => {
    const menus = ref<MenuObj[]>([]);
    const permissions = ref<Array<string>>([]);

    const getLeftMenus = async () => {
      await getUserLeftMenusApi().then((res) => {
        menus.value = res.data;
      });
    };
    const getUserPermissions = async () => {
      await getUserPermissionApi().then((res) => {
        console.log("保存permission....",res.data);
        
        permissions.value = res.data;
      });
      console.log("保存结果查询Permission:",permissions.value);
      
    };
    return { menus, permissions, getLeftMenus, getUserPermissions };
  },
  {
    persist: {
      storage: localStorage,
    },
  }
);
