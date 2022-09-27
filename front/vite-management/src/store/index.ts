import { App } from 'vue'
import { createPinia } from 'pinia'


const pinia = createPinia()

export const initStore = (app: App<Element>) => [
  app.use(pinia)
]
export default pinia