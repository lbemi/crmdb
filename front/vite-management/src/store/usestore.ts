import { defineStore } from "pinia";
import { getUserLeftMenusApi } from '../request/api'

type MenuObj ={
  id: number,
  name: string,
  url: string,
  children: MenuObj[]
}

interface State {
  menus: MenuObj[]
}

export const useStore = defineStore('main',{
  state: ():State=>{
    return {
      menus: []
    }
  },
  actions: {
    getLeftMenusApi() {
      return new Promise((reslove, reject) => {
        getUserLeftMenusApi().then(res => {
          if (res.code === 200) {
            this.menus = res.data
            reslove(res.data)
          } else {
            reject(res)
          }
        })
      })
    }
  },
})