import { useStore } from '@/store/usestore'
import pinia from '@/store/index'

// 单个权限验证
export function auth(value: string) {
  // console.log(store.permissions);
  const store = useStore()
  if (store.permissions) {
    return store.permissions.some((v: any) => v === value)
  }
  return false
}

// // 多个权限验证，满足一个则为 true
// export function auths(value: Array<string>) {
//     let flag = false;
// const store = useStore()

//     store.permissions.map((val: any) => {
//         value.map((v: any) => {
//             if (val === v) flag = true;
//         });
//     });
//     return flag;
// }

// // 多个权限验证，全部满足则为 true
// export function authAll(value: Array<string>) {
// const store = useStore()

//     return judementSameArr(value, store.permissions);
// }
