import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userApi } from '@/views/sys/api'
export interface MenuObj {
  id: number
  name: string
  url: string
  icon: string
  children: MenuObj[]
}

export const useStore = defineStore(
  'main',
  () => {
    const menus = ref<MenuObj[]>([])
    const permissions = ref<Array<string>>([])
    const isCollapse = ref<boolean>(false)
    const getLeftMenus = async () => {
      await userApi.listMenus.request().then((res) => {
        menus.value = res.data
      })
    }
    const getUserPermissions = async () => {
      await userApi.permission.request().then((res) => {
        permissions.value = res.data
      })
    }
    return { menus, permissions, isCollapse, getLeftMenus, getUserPermissions }
  },
  {
    persist: {
      storage: localStorage
    }
  }
)
