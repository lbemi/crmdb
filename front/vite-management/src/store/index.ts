import { reject } from 'lodash';
import { App } from 'vue'
import { createStore } from 'vuex'
import { getUserLeftMenusApi } from '../request/api'

interface MenuObj {
  id: number,
  name: string,
  url: string,
  children: MenuObj[]
}

interface State {
  menus: MenuObj[]
}
const store = createStore<State>({
  state() {
    return {
      menus: []
    }
  },
  getters: {
    // 获取menus
    // getMenus(state) {
    //   return state.menus;
    // }
  },
  mutations: {
    // 更新menus
    updateMenus(state, menus) {
      state.menus = menus
    }
  },
  actions: {
    getLeftMenusApi({ commit }) {
      return new Promise((reslove, reject) => {
        getUserLeftMenusApi().then(res => {
          if (res.code === 200) {
            commit('updateMenus', res.data)
            reslove(res.data)
          } else {
            reject(res)
          }
        })
      })
    }
  },
  modules: {}
})



export const initStore = (app: App<Element>) => [
  app.use(store)
]
export default store