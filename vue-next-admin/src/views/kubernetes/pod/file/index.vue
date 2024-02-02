<template>
	<div class="layout-padding container">
		<el-card shadow="hover" class="layout-padding-auto card">
			<template #header>
				<div class="header-top">
					<div>
						<a>容器组名: {{ pod?.metadata?.name }}：</a>
						<el-select v-model="state.selectContainer" placeholder="选择容器" size="small" @change="getFileList">
							<el-option v-for="item in pod?.spec?.containers" :key="item.name" :label="item.name"
								:value="item.name" />
						</el-select>
					</div>
					<div>
						<el-tag style="float: left; cursor: pointer; margin-left: 15px" @click="setDir(0)">根目录</el-tag>
						<el-breadcrumb separator="/" class="breadcrumb">
							<el-breadcrumb-item v-bind:key="p" v-for="(p, index) in state.path.split('/')">
								<a class="breadcrumb-link" @click="setDir(index)">{{ p }}</a>
							</el-breadcrumb-item>
						</el-breadcrumb>
					</div>
				</div>
			</template>
			<el-row :gutter="20">
				<el-col :span="18">
					<div v-loading="state.loading" class="video-main">
						<div class="list">
							<div :class="getListClass(item.name)" v-bind:key="item.name" v-for="item in state.files">
								<div class="inner">
									<el-image :src="getFsImage(item.fsType)" class="icon-thumb" fit="contain" alt
										@dblclick="gotoDir(item)" @click="selectItem(item)" :key="item.name"></el-image>
									<i class="icon-folder"></i>
								</div>
								<i class="icon-file-selected"></i>
								<div class="file-name">
									<span @click="selectItem(item)" :title="item.name">{{ item.name }}</span>
								</div>
							</div>
						</div>
					</div>
				</el-col>
				<el-col :span="6" style="display: flex">
					<el-divider direction="vertical" style="height: 100%; margin-right: 30px" />
					<div>
						<el-form>
							<div>
								<div>
									<el-button type="primary" size="small"> 上传 </el-button>
									<el-button type="success" size="small" @click="createDir"> 新建文件夹 </el-button>
									<el-button :disabled="state.selected.name === undefined" type="danger" size="small"
										@click="deleteFileOrDir"> 删除 </el-button>
								</div>
								<div style="margin-top: 10px; margin-bottom: 10px">
									<el-input v-model="state.searchFileName" placeholder="输入名称" style="width: 150px"
										size="small" clearable></el-input>
									<el-button size="small" type="primary" style="margin-left: 2px"
										:icon="Search">搜索</el-button>
								</div>
							</div>

							<el-text type="info">属性</el-text>
							<el-form-item label="名称:" key="name">
								<el-input key="name" v-model="state.selected.name" placeholder="名称" size="small"
									style="width: 150px" :disabled="!state.edit">
								</el-input>
								<el-button v-if="!state.edit" type="primary" text size="small"
									@click="state.edit = !state.edit">编辑</el-button>
								<el-button v-if="state.edit" type="primary" text size="small"
									@click="updateFileName">确定</el-button>
								<el-button v-if="state.edit" type="primary" text size="small"
									style="margin: 0; padding-left: 0" @click="state.edit = !state.edit">取消</el-button>
							</el-form-item>
							<el-form-item label="权限:" key="auth">
								<span>{{ state.selected && state.selected.permissions ? state.selected.permissions : ''
								}}</span>
							</el-form-item>
							<el-form-item label="大小:" key="size">
								<span>{{ state.selected && state.selected.size ? state.selected.size : '0' }}</span>
							</el-form-item>
							<el-form-item label="用户:" key="user">
								<el-input key="user" size="small" v-model="state.selected.user" placeholder="用户">{{
									state.selected && state.selected.user ? state.selected.user : ''
								}}</el-input>
							</el-form-item>
							<el-form-item label="属组:" key="group">
								<el-input key="group" size="small" v-model="state.selected.group" placeholder="用户">{{
									state.selected && state.selected.group ? state.selected.group : ''
								}}</el-input>
							</el-form-item>
							<el-form-item label="上次操作时间:"> {{ state.selected.lastModify }} </el-form-item>
							<el-form-item label="操作:">
								<el-button type="primary" text v-if="state.selected.fsType === '-'"
									@click="readFile">预览文件</el-button>
								<el-button type="primary" text v-if="state.selected.fsType === '-'"
									@click="downloadFile">下载</el-button>
								<el-link type="primary" v-if="state.selected.fsType === '-'" target="_blank"
									:href="getDowloadLink()">下载2</el-link>
							</el-form-item>
						</el-form>
					</div>
				</el-col>
			</el-row>
		</el-card>
		<el-drawer v-model="state.dialogVisible" size="40%">
			<template #header>
				<h3>{{ state.selected.name }} 预览</h3>
			</template>
			<Codemirror v-loading="state.loading" v-model="state.fileInfo"
				style="height: 92%; margin-left: 20px; margin-right: 15px" :indent-with-tab="true" :tabSize="2"
				:extensions="extensions" :disabled="true" />
			<div class="mt30" style="align-items: center; margin-left: 20px">
				<el-button size="small" @click="handleClose">关闭</el-button>
			</div>
		</el-drawer>
	</div>
</template>
<script setup lang="ts" name="podFile">
import { h, onMounted, reactive } from 'vue';
import folder from '@/assets/folder.svg';
import { Codemirror } from 'vue-codemirror';
import file from '@/assets/file.svg';
import { kubernetesInfo } from '@/stores/kubernetes';
import { usePodApi } from '@/api/kubernetes/pod';
import { FileType } from '@/types/kubernetes/cluster';
import { oneDark } from '@codemirror/theme-one-dark';
import { Search } from '@element-plus/icons-vue';
import { podInfo } from '@/stores/pod';
import { deepClone } from '@/utils/other';
import { ElMessage, ElMessageBox } from 'element-plus';
import { saveAs } from 'file-saver';

const pod = podInfo().state.podShell;
const extensions = [oneDark];
const podApi = usePodApi();
const k8sStore = kubernetesInfo();

onMounted(() => {
	getFileList();
});
const getFileList = async () => {
	state.loading = true;
	await podApi
		.getPodFileList(pod.metadata?.namespace, pod?.metadata?.name, state.selectContainer, {
			cloud: k8sStore.state.activeCluster,
			path: state.path,
		})
		.then((res) => {
			state.files = res.data;
		});
	state.loading = false;
};

const readFile = async () => {
	if (state.selected.name === undefined || state.selected.fsType === 'd') {
		return;
	}

	let file = getFileFilePath();
	state.dialogVisible = true;
	state.loading = true;
	await podApi
		.readPodFileInfo(pod.metadata?.namespace, pod?.metadata?.name, state.selectContainer, {
			cloud: k8sStore.state.activeCluster,
			file: file,
		})
		.then((res) => {
			state.fileInfo = res.data;
		});
	state.loading = false;
};

const handleClose = () => {
	state.dialogVisible = false;
	state.fileInfo = '';
};
const state = reactive({
	path: '/',
	searchFileName: '',
	loading: false,
	files: [] as Array<FileType>,
	selected: {} as FileType,
	dialogVisible: false,
	fileInfo: '',
	selectContainer: pod.spec?.containers[0]?.name,
	srcName: '',
	edit: false,
});
const updateFileName = () => {
	const src = (state.path + state.srcName).replace('//', '/');
	const dst = getFileFilePath();
	ElMessageBox({
		title: '提示',
		message: h('p', null, [
			h('span', null, '此操作将 '),
			h('i', { style: 'color: teal' }, `${src}`),
			h('span', null, '重命名为 '),
			h('i', { style: 'color: teal' }, `${dst}`),
			h('span', null, '  是否继续? '),
		]),
		buttonSize: 'small',
		showCancelButton: true,
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
		draggable: true,
	})
		.then(async () => {
			await podApi
				.updateFileName(pod.metadata?.namespace, pod?.metadata?.name, state.selectContainer, {
					cloud: k8sStore.state.activeCluster,
					src: src,
					dst: dst,
				})
				.then(() => {
					ElMessage.success('操作成功');
					state.edit = false;
					getFileList();
					state.selected = {} as FileType;
				})
				.catch((e) => {
					ElMessage.error(e.message);
				});
		})
		.catch(() => {
			ElMessage.info('取消');
		});
};
const getFsImage = (type: string) => {
	if (type === 'd') {
		//文件夹
		return folder;
	} else {
		return file;
	}
};

const selectItem = (item: FileType) => {
	state.selected = deepClone(item) as FileType;
	state.srcName = state.selected.name;
};
const setDir = (index: number) => {
	state.selected = {} as FileType;
	if (index === 0) {
		state.path = '/';
		getFileList();
		return;
	}
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

const createDir = async () => {
	await podApi
		.createDir(pod.metadata?.namespace, pod?.metadata?.name, state.selectContainer, {
			cloud: k8sStore.state.activeCluster,
			path: state.path,
		})
		.then(() => {
			ElMessage.success('操作成功');
			getFileList();
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
};
const deleteFileOrDir = async () => {
	const dst = getFileFilePath();
	await podApi
		.removeFileOrDir(pod.metadata?.namespace, pod?.metadata?.name, state.selectContainer, {
			cloud: k8sStore.state.activeCluster,
			dst: dst,
		})
		.then(() => {
			ElMessage.success('操作成功');
			getFileList();
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
};

const downloadFile = async () => {
	const file = getFileFilePath();
	await podApi
		.downloadFile(pod.metadata?.namespace, pod?.metadata?.name, state.selectContainer, {
			cloud: k8sStore.state.activeCluster,
			file: file.slice(1),
		})
		.then((res: any) => {
			console.log(res);

			let str = new Blob([res]);
			saveAs(str, state.selected.name);
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
};

//http://localhost:8080/api/v1/pods/namespaces/default/files/nginx-2-598f88c6dc-f7fb8/nginx-2/download?cloud=311&file=aa.txt
const getDowloadLink = () => {
	const file = getFileFilePath().slice(1);
	const href = 'http://localhost:8080/api/v1/pods/namespaces/' + pod.metadata!.namespace + '/files/' + pod.metadata!.name + '/' + state.selectContainer + '/download?cloud=' + k8sStore.state.activeCluster + '&file=' + file;
	return href;
}

const getListClass = (itemName: string) => {
	if (state.selected.name && state.selected.name === itemName) {
		return 'list-item selected';
	} else {
		return 'list-item';
	}
};

const getFileFilePath = () => {
	let filePath = state.path + '/' + state.selected.name;
	return filePath.replace('//', '/');
};
</script>

<style lang="scss" scoped>
.breadcrumb {
	float: left;
	// height: 20px;
	margin-top: 4px;
}

.breadcrumb-link {
	cursor: pointer;
	font-size: 16px;
}

.breadcrumb-link:hover {
	color: #409eff;
	text-decoration: underline;
}

.header-top {
	display: flex;
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
	padding-top: 5px;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
	color: #424e67;
	font-size: 14px;
	width: 110px;
	text-align: center;
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

.selected {
	background-color: #f1f5fa;
	border: 1px solid #409eff;
	border-radius: 8px;
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

.card {
	overflow-y: auto;
	/* 开启滚动显示溢出内容 */
}
</style>
