<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<el-container style="height: 500px; border: 1px solid #eee">
				<el-aside width="200px" style="background-color: rgb(238, 241, 246)">
					<el-menu :default-openeds="['1', '1-1']">
						<el-sub-menu index="1">
							<template v-slot:title><i class="el-icon-message"></i>pod操作</template>
							<el-sub-menu index="1-1">
								<template v-slot:title>选择容器</template>
								<el-menu-item index="1-1-1">
									<router-link to="/files">nginx</router-link>
								</el-menu-item>
								<el-menu-item index="1-1-1">sidecar</el-menu-item>
							</el-sub-menu>
						</el-sub-menu>
					</el-menu>
				</el-aside>
				<el-main>
					<!-- <div class="video-container">
						<el-card shadow="hover"> -->
					<div class="video-header clearfix">
						<div class="header-top">
							<div>
								<el-tag style="float: left; cursor: pointer" @click="setDir(0)">根目录</el-tag>
								<el-breadcrumb separator="/" class="breadcrumb">
									<el-breadcrumb-item v-bind:key="p" v-for="(p, index) in state.path.split('/')">
										<a class="breadcrumb-link" @click="setDir(index)">{{ p }}</a>
									</el-breadcrumb-item>
								</el-breadcrumb>
							</div>
							<div>
								<el-button type="primary" size="default">
									<i class="el-icon-upload2"></i>
									上传
								</el-button>
								<el-button type="primary" size="default">
									<i class="el-icon-plus"></i>
									新建文件夹
								</el-button>
								<el-button type="primary" size="default">
									<i class="el-icon-delete"></i>
									删除
								</el-button>
								<el-input v-model="state.searchFileName" placeholder="输入名称"
									style="width: 150px; margin-left: 10px" size="default" clearable></el-input>
								<!-- 查询 -->
								<el-button size="default" type="primary" style="margin-left: 2px"
									icon="el-icon-search">搜索</el-button>
							</div>
						</div>
					</div>

					<div v-loading="state.loading" class="video-main">
						<!-- <ul> -->
						<div class="list-item" v-bind:key="item.name" v-for="item in state.files">
							<div>
								<div class="inner">
									<el-image :src="getFsImage(item.fsType)" class="icon-thumb" fit="contain" alt
										@dblclick="gotoDir(item)" @click="selectItem(item)" :key="item.name"></el-image>
									<i class="icon-folder"></i>
								</div>
								<div class="file-name">
									<i class="icon-file-selected"></i>
									<span :title="item.name">{{ item.name }}</span>
								</div>
							</div>
						</div>
						<!-- </ul> -->
						<!-- </div>
						</el-card> -->
					</div>
					<el-card v-show="state.selected.name !== undefined">
						<div slot="header" class="clearfix">
							<span>详情</span>
							<el-button style="float: right; padding: 3px 0" link>操作按钮</el-button>
						</div>
						<el-form :inline="true">
							<el-form-item label="名称" key="name">
								<el-input key="name" v-model="state.selected.name" placeholder="名称">{{ state.selected &&
									state.selected.name ?
									state.selected.name : '' }}</el-input>
							</el-form-item>
							<el-form-item label="权限:" key="auth">
								<span>{{ state.selected && state.selected.permissions ? state.selected.permissions : ''
								}}</span>
							</el-form-item>
							<el-form-item label="大小:" key="size">
								<span>{{ state.selected && state.selected.size ? state.selected.size : '0' }}</span>
							</el-form-item>
							<el-form-item label="用户:" key="user">
								<el-input key="user" v-model="state.selected.user" placeholder="用户">{{ state.selected &&
									state.selected.user ?
									state.selected.user : '' }}</el-input>


							</el-form-item>
							<el-form-item label="组:" key="group">
								<el-input key="group" v-model="state.selected.group" placeholder="用户">{{ state.selected &&
									state.selected.group ? state.selected.group : '' }}</el-input>

							</el-form-item>
						</el-form>
					</el-card>
				</el-main>
			</el-container>
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podFile">
import { onMounted, reactive } from 'vue';
import folder from '@/assets/folder.svg';
import file from '@/assets/file.svg';
import { kubernetesInfo } from '@/stores/kubernetes';
import { usePodApi } from '@/api/kubernetes/pod';
import { FileType } from '@/types/kubernetes/cluster';

const podApi = usePodApi();
const k8sStore = kubernetesInfo();

onMounted(() => {
	getFileList();
});
const getFileList = async () => {
	state.loading = true
	await podApi
		.getPodFileList('default', 'nginx-2-598f88c6dc-f7fb8', 'nginx-2', {
			cloud: k8sStore.state.activeCluster,
			path: state.path,
		})
		.then((res) => {
			state.files = res.data;
			console.log(res);
		});
	state.loading = false
};
const state = reactive({
	path: '/',
	searchFileName: '',
	loading: false,
	files: [] as Array<FileType>,
	selected: {} as FileType,
});

const getFsImage = (type: string) => {
	if (type === 'd') {
		//文件夹
		return folder;
	} else {
		return file;
	}
};

const selectItem = (item: FileType) => {
	state.selected = item;
};
const setDir = (index: number) => {
	// state.loading = true
	state.selected = {} as FileType;
	if (index === 0) {
		state.path = '/';
		getFileList();
		return;
	}
	//        /dev/abc
	let list = state.path.split('/');
	state.path = list.slice(0, index + 1).join('/');

	getFileList();
};
const gotoDir = (item: FileType) => {
	state.selected = {} as FileType;
	if (item.fsType === 'd') {
		//文件夹才会进入
		state.path = state.path + '/' + item.name;
		state.path = state.path.replace('//', '/');
		getFileList();
	}
};
</script>

<style lang="scss" scoped>
.clearfix:after {
	content: '';
	display: block;
	clear: both;
}

.video-container {
	min-width: 630px;
	margin: 10px;
}

.video-header {
	padding: 0 0 5px 0;
	border-bottom: 1px solid #dbdbdb;
}

.breadcrumb {
	float: left;
	// height: 20px;
	margin-top: 3px;
}

.breadcrumb-link {
	cursor: pointer;
	font-size: 16px;
}

.breadcrumb-link:hover {
	color: #409eff;
	text-decoration: underline;
}

.breadcrumb-link-active {
	// 面包屑当前激活文件夹的样式
	font-weight: 700;
}

.header-top {
	// float: right;
	// height: 40px;
	// line-height: 40px;
	// position: relative;
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.list-item {
	border: 1px solid #fff;
	box-sizing: border-box;
	position: relative;
	height: 120px;
	width: 110px;
	margin: 5px;
	display: inline-block;
	cursor: pointer;
	border-radius: 5px;
	// display: flex;
	// height: 80px;
	// width: 80px;

	// justify-content: flex-start;
}

.inner {
	height: 60px;
	width: 60px;

	margin: 5px 10px;
	padding: 10px 15px 10px 15px;
}

.icon-folder {
	display: inline-block;
	width: 60px;
	height: 60px;
	background-size: 100% 100%;
}

.icon-thumb {
	width: 50px;
	height: 50px;
}

.file-name {
	padding-left: 35px;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
	color: #424e67;
	font-size: 14px;
}

.file-name:hover {
	color: #409eff;
}

.hover-cover {
	width: 60px;
	height: 60px;
	position: absolute;
	left: 10px;
	top: 5px;
	background-color: rgb(0, 0, 0);
	opacity: 0;
	text-align: center;
	line-height: 60px;
	font-size: 12px;
}

.list-item:hover {
	background-color: #f1f5fa;
}

.icon-file-selected {
	opacity: 0.5;
}

.hover-cover {
	opacity: 0.6;
}

.icon-file-selected {
	// 小圆点默认样式
	position: absolute;
	left: 5px;
	top: 5px;
	display: inline-block;
	width: 20px;
	height: 20px;
	background-size: 100% 100%;

	opacity: 0;
}

.icon-file-selected:hover {
	opacity: 1 !important;
}

.active {
	border: 1px solid #409eff;
	border-radius: 8px;
}

.icon-file-selected {
	position: absolute;
	left: 5px;
	top: 5px;
	display: inline-block;
	width: 20px;
	height: 20px;
	background-size: 100% 100%;

	opacity: 1;
}

.icon-file-selected {
	opacity: 1 !important;
}

.loadding-message {
	color: #424e67;
	font-size: 12px;
}
</style>
