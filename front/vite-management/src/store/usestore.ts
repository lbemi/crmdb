import { defineStore } from "pinia";
import { getUserLeftMenusApi, getUserPermissionApi } from '../request/api'

type MenuObj = {
  id: number,
  name: string,
  url: string,
  children: MenuObj[]
}

interface State {
  menus: MenuObj[],
  permissions: Array<string>
}

export const useStore = defineStore('main', {
  state: (): State => {
    return {
      menus: [],
      permissions: []
    }
  },
  actions: {
    async getLeftMenusApi() {
      await getUserLeftMenusApi().then(res => {
        this.menus = res.data
      })
    },
    async getUserPermissions() {
      await getUserPermissionApi().then(res => {
        this.permissions = res.data
      })
    }
  },
  persist: true
})