import type { App } from 'vue';
import { useUserInfo } from '/@/stores/userInfo';
import { judementSameArr } from '/@/utils/arrayOperation';

/**
 * 用户权限指令
 * @directive 单个权限验证（v-auth="xxx"）
 * @directive 多个权限验证，满足一个则显示（v-auths="[xxx,xxx]"）
 * @directive 多个权限验证，全部满足则显示（v-auth-all="[xxx,xxx]"）
 */
export function authDirective(app: App) {
	// 单个权限验证（v-auth="xxx"）
	app.directive('auth', {
		mounted(el, binding) {
			const stores = useUserInfo();
			// if (!stores.userInfos.authBtnList.some((v: string) => v === binding.value)) el.parentNode.removeChild(el); //无权限的按钮删除掉
			// 无权限的按钮置灰色
			if (!stores.userInfos.authBtnList.some((v: string) => v === binding.value)) {
				el.setAttribute('disabled', 'true');
				el.classList.add('is-disabled');
				el.addEventListener('click', disableClickFn, true);
			}
		},
	});
	// 多个权限验证，满足一个则显示（v-auths="[xxx,xxx]"）
	app.directive('auths', {
		mounted(el, binding) {
			let flag = false;
			const stores = useUserInfo();
			stores.userInfos.authBtnList.map((val: string) => {
				binding.value.map((v: string) => {
					if (val === v) flag = true;
				});
			});
			if (!flag) el.parentNode.removeChild(el);
		},
	});
	// 多个权限验证，全部满足则显示（v-auth-all="[xxx,xxx]"）
	app.directive('auth-all', {
		mounted(el, binding) {
			const stores = useUserInfo();
			const flag = judementSameArr(binding.value, stores.userInfos.authBtnList);
			if (!flag) el.parentNode.removeChild(el);
		},
	});
}

const disableClickFn = (event: any) => {
	event && event.stopImmediatePropagation()
}
