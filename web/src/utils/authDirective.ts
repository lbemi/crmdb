import type { App, DirectiveBinding } from 'vue'
import { auth } from './authFunction'

// 用户权限指令
export function authDirective(app: App) {
  // 单个权限验证（v-auth="xxx"）
  app.directive('auth', {
    mounted(el: HTMLElement, binding: DirectiveBinding) {
      if (!auth(binding.value)) {
        // console.log("禁用:",binding.value);
        parseNoAuth(el, binding)
      }
    }
  })
  // // 多个权限验证，满足一个则显示（v-auths="[xxx:xxx]"）
  // app.directive('auths', {
  //     mounted(el:HTMLElement, binding:DirectiveBinding) {
  //         if (!auths(binding.value)) {
  //             parseNoAuth(el, binding);
  //         }
  //     },
  // });
  // // 多个权限验证，全部满足则显示（v-auth-all="[xxx:xxx:xxx]"）
  // app.directive('auth-all', {
  //     mounted(el:HTMLElement, binding:DirectiveBinding) {
  //         if (!authAll(binding.value)) {
  //             parseNoAuth(el, binding);
  //         };
  //     },
  // });
}

/**
 * 处理没有权限场景
 *
 * @param el  元素
 * @param binding 绑定至
 */
const parseNoAuth = (el: HTMLElement, binding: DirectiveBinding) => {
  const { arg } = binding

  // 如果是禁用模式，则将元素禁用
  // if (arg == 'disabled') {
  el.setAttribute('disabled', 'true')
  el.classList.add('is-disabled')
  el.addEventListener('click', disableClickFn, true)
  // } else {
  //     // 移除该元素
  //     el.parentNode.removeChild(el);
  // }
}

const disableClickFn = (event: any) => {
  event && event.stopImmediatePropagation()
}
