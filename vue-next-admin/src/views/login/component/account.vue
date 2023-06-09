<template>
	<el-form size="large" class="login-content-form">
		<el-form-item class="login-animation1">
			<el-input text :placeholder="$t('message.account.accountPlaceholder1')" v-model="state.ruleForm.user_name" clearable autocomplete="off">
				<template #prefix>
					<el-icon class="el-input__icon"><ele-User /></el-icon>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation2">
			<el-input
				:type="state.isShowPassword ? 'text' : 'password'"
				:placeholder="$t('message.account.accountPlaceholder2')"
				v-model="state.ruleForm.password"
				autocomplete="off"
			>
				<template #prefix>
					<el-icon class="el-input__icon"><ele-Unlock /></el-icon>
				</template>
				<template #suffix>
					<i
						class="iconfont el-input__icon login-content-password"
						:class="state.isShowPassword ? 'icon-yincangmima' : 'icon-xianshimima'"
						@click="state.isShowPassword = !state.isShowPassword"
					>
					</i>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation3">
			<el-col :span="15">
				<el-input
					text
					maxlength="5"
					:placeholder="$t('message.account.accountPlaceholder3')"
					v-model="state.ruleForm.captcha"
					clearable
					autocomplete="off"
				>
					<template #prefix>
						<el-icon class="el-input__icon"><ele-Position /></el-icon>
					</template>
				</el-input>
			</el-col>
			<el-col :span="1"></el-col>
			<el-col :span="8">
				<!-- <el-button class="login-content-code" v-waves>1234</el-button> -->
				<div><img @click="onChangeCaptcha" class="captcha-img" :src="state.capthca.path" alt="captcha" /></div>
			</el-col>
		</el-form-item>
		<el-form-item class="login-animation4">
			<el-button type="primary" class="login-content-submit" round v-waves @click="onSignIn" :loading="state.loading.signIn">
				<span>{{ $t('message.account.accountBtnText') }}</span>
			</el-button>
		</el-form-item>
	</el-form>
</template>

<script setup lang="ts" name="loginAccount">
import { reactive, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { useI18n } from 'vue-i18n';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import { initFrontEndControlRoutes } from '/@/router/frontEnd';
import { initBackEndControlRoutes } from '/@/router/backEnd';
import { Session } from '/@/utils/storage';
import { formatAxis } from '/@/utils/formatTime';
import { NextLoading } from '/@/utils/loading';
import { useLoginApi } from '/@/api/login/index';
// 定义变量内容
const loginApi = useLoginApi();
const { t } = useI18n();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const route = useRoute();
const router = useRouter();
const state = reactive({
	isShowPassword: false,
	ruleForm: {
		user_name: 'lbemi_admin',
		password: 'Admin_1234',
		captcha: '12345',
		captcha_id: '',
	},
	capthca: {
		path: '',
	},
	loading: {
		signIn: false,
	},
});

// 时间获取
const currentTime = computed(() => {
	return formatAxis(new Date());
});
// 验证码
const onChangeCaptcha = async () => {
	await loginApi.getCaptcha().then((res) => {
		state.ruleForm.captcha_id = res.data.captcha_id;
		state.capthca.path = res.data.pic_path;
	});
};
onMounted(() => {
	onChangeCaptcha();
});

// 登录
const onSignIn = async () => {
	state.loading.signIn = true;
	await loginApi
		.signIn(state.ruleForm)
		.then(async (res) => {
			// 存储 token 到浏览器缓存
			Session.set('token', res.data.token);
			Session.set('userInfo', res.data.user);
			// 模拟数据，对接接口时，记得删除多余代码及对应依赖的引入。用于 `/src/stores/userInfo.ts` 中不同用户登录判断（模拟数据）
			Session.set('userName', state.ruleForm.user_name);
			if (!themeConfig.value.isRequestRoutes) {
				// 前端控制路由，2、请注意执行顺序
				const isNoPower = await initFrontEndControlRoutes();
				signInSuccess(isNoPower);
			} else {
				// 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
				// 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
				const isNoPower = await initBackEndControlRoutes();
				// 执行完 initBackEndControlRoutes，再执行 signInSuccess
				signInSuccess(isNoPower);
			}
		})
		.catch((e) => {
			if (e.code === 1003) {
				onChangeCaptcha();
			}
			ElMessage.error(e.message);
		});
	state.loading.signIn = false;
	// // 存储 token 到浏览器缓存
	// Session.set('token', res.token);
	// // 模拟数据，对接接口时，记得删除多余代码及对应依赖的引入。用于 `/src/stores/userInfo.ts` 中不同用户登录判断（模拟数据）
	// Cookies.set('userName', state.ruleForm.user_name);
	// if (!themeConfig.value.isRequestRoutes) {
	// 	// 前端控制路由，2、请注意执行顺序
	// 	const isNoPower = await initFrontEndControlRoutes();
	// 	signInSuccess(isNoPower);
	// } else {
	// 	// 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
	// 	// 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
	// 	const isNoPower = await initBackEndControlRoutes();
	// 	// 执行完 initBackEndControlRoutes，再执行 signInSuccess
	// 	signInSuccess(isNoPower);
	// }
};
// 登录成功后的跳转
const signInSuccess = (isNoPower: boolean | undefined) => {
	if (isNoPower) {
		ElMessage.warning('抱歉，您没有登录权限');
		Session.clear();
	} else {
		// 初始化登录成功时间问候语
		let currentTimeInfo = currentTime.value;
		// 登录成功，跳到转首页
		// 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
		if (route.query?.redirect) {
			router.push({
				path: <string>route.query?.redirect,
				query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
			});
		} else {
			router.push('/');
		}
		// 登录成功提示
		const signInText = t('message.signInText');
		ElMessage.success(`${currentTimeInfo}，${signInText}`);
		// 添加 loading，防止第一次进入界面时出现短暂空白
		NextLoading.start();
	}
	state.loading.signIn = false;
};
</script>

<style scoped lang="scss">
.captcha-img {
	margin-left: auto;
	width: 90%;
	align-items: center;
}
.login-content-form {
	margin-top: 20px;
	@for $i from 1 through 4 {
		.login-animation#{$i} {
			opacity: 0;
			animation-name: error-num;
			animation-duration: 0.5s;
			animation-fill-mode: forwards;
			animation-delay: calc($i/10) + s;
		}
	}
	.login-content-password {
		display: inline-block;
		width: 20px;
		cursor: pointer;
		&:hover {
			color: #909399;
		}
	}
	.login-content-code {
		width: 100%;
		padding: 0;
		font-weight: bold;
		letter-spacing: 5px;
	}
	.login-content-submit {
		width: 100%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}
}
</style>
