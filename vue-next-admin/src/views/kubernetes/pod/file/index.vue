<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto">
			<el-container style="height: 500px; border: 1px solid #eee">
				<el-aside width="200px" style="background-color: rgb(238, 241, 246)">
					<el-menu :default-openeds="['1', '1-1']">
						<el-submenu index="1">
							<template v-slot:title><i class="el-icon-message"></i>pod操作</template>
							<el-submenu :index="1 - 1">
								<template v-slot:title>选择容器</template>
								<el-menu-item index="1-1-1">
									<router-link to="/files">nginx</router-link>
								</el-menu-item>
								<el-menu-item index="1-1-1">sidecar</el-menu-item>
							</el-submenu>
						</el-submenu>
					</el-menu>
				</el-aside>
				<el-main>
					<!--					<div class="video-container">-->
					<!--						<el-card shadow="hover">-->
					<div class="video-header clearfix">
						<div class="header-top">
							<el-tag style="float: left; cursor: pointer" @click="setDir(0)">根目录</el-tag>
							<el-breadcrumb separator="/" class="breadcrumb">
								<el-breadcrumb-item v-bind:key="p" v-for="(p, index) in state.path.split('/')">
									<a class="breadcrumb-link" @click="setDir(index)">{{ p }}</a>
								</el-breadcrumb-item>
							</el-breadcrumb>
							<el-button type="primary" size="default" style="margin-left: 20px;">
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

					<div v-loading="state.loading" class="video-main">
						<ul class="list">
							<li class="list-item" v-bind:key="item.name" v-for="item in state.files">
								<div class="inner">
									<el-image :src="getFsImage(item.fsType)" class="icon-thumb" fit="contain" alt
										@dblclick="gotoDir(item)" @click="selectItem(item)" :key="item.name"></el-image>
									<i class="icon-folder"></i>

								</div>
								<i class="icon-file-selected"></i>
								<div class="file-name">
									<span :title="item.name">{{ item.name }}</span>
								</div>
							</li>
						</ul>
					</div>
					<!--						</el-card>-->
					<!--					</div>-->
				</el-main>
			</el-container>
		</el-card>
	</div>
</template>
<script setup lang="ts" name="podFile">
import { onMounted, reactive } from 'vue';
import { podInfo } from '@/stores/pod';
import folder from '@/assets/folder.svg';
import file from '@/assets/file.svg';
import { kubernetesInfo } from '@/stores/kubernetes';
import { usePodApi } from '@/api/kubernetes/pod';
import { FileType } from '@/types/kubernetes/cluster';

const podApi = usePodApi();
const k8sStore = kubernetesInfo();


onMounted(() => {
	getFileList();
})
const getFileList = async () => {
	await podApi.getPodFileList("default", "nginx-2-598f88c6dc-f7fb8", "nginx-2", {
		cloud: k8sStore.state.activeCluster,
		path: state.path
	}).then((res) => {
		state.files = res.data;
		console.log(res);
	});
}
const state = reactive({
	path: '/',
	searchFileName: '',
	loading: false,
	files: [
	] as Array<FileType>,
	selected: {}
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
	state.selected = item
}
const setDir = (index: number) => {
	state.selected = {}
	if (index === 0) {
		state.path = "/"
		getFileList()
		return
	}
	//        /dev/abc
	let list = state.path.split("/")
	state.path = list.slice(0, index + 1).join("/")

	getFileList()

}
const gotoDir = (item: FileType) => {
	state.selected = {}
	if (item.fsType === "d") { //文件夹才会进入
		state.path = state.path + "/" + item.name
		state.path = state.path.replace("//", "/")
		getFileList()
	}
}
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
	height: 20px;
	margin-top: 10px;
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
	float: right;
	height: 40px;
	line-height: 40px;
	position: relative;
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
	width: 60px;
	height: 60px;
}

.file-name {
	padding-left: 10px;
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
