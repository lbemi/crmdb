<template >
    <span>容器组名: {{ pod?.metadata.name }}</span>
    <el-select v-model="selectContiner" class="m-2" placeholder="选择容器" size="large">
        <el-option v-for="item in pod?.spec.containers" :key="item.name" :label="item.name" :value="item.name" />
    </el-select>
    <el-button type="primary" @click="showLogs">显示日志</el-button>
    <el-button @click="logs.log = ''">清空</el-button>
    <div ref="logs" class="logs">
        {{ logs.log }}
    </div>

</template>
<script setup lang="ts">
import { ref,reactive } from 'vue'
import { podStore } from '@/store/kubernetes/pods'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { service } from '@/request/request';
import router from '@/router';

const cloud = kubeStore().activeCluster
const selectContiner = ref()
const pod = podStore().podShell
const logs = reactive({
    log: "asds"
})

const getLog = async (namespace: string, podName: string, container: string) => {
    console.log("--------------");

    service.request({
        url: '/pod/log/' + namespace + '/' + podName + '/' + container + '?cloud=' + cloud,
        method: 'GET',
        onDownloadProgress: e => {
            console.log(e)
            const dataChunk = e.currentTarget.response
            logs.value += dataChunk
            console.log("++++++++++++++++++++", logs);

            const logDiv = document.getElementById('logs')
            if (logDiv !== undefined && logDiv?.scrollTop !== undefined) {
                logDiv.scrollTop = logDiv.scrollHeight
            }

        }
    })

}

const showLogs = () => {
    service.request({
        url: '/pod/log/' + pod?.metadata.namespace! + '/' + pod?.metadata.name! + '/' + selectContiner.value + '?cloud=' + cloud,
        method: 'GET',
        onDownloadProgress: e => {
            const dataChunk = e.currentTarget.response
            logs.log =logs.log+ dataChunk 
            console.log("++++++++++++++++++++", logs.log);

            const logDiv = document.getElementById('logs')
            if (logDiv !== undefined && logDiv?.scrollTop !== undefined) {
                console.log("+++++++!!!!!!!!");

                logDiv.scrollTop = logDiv.scrollHeight
            }

        }
    }).catch(e => {
        // router.push({name: 'pod'})
    })
}
</script>
<style lang="less">
.logs {
    overflow: auto;

    margin: 10px auto;
    min-height: 200px;
    max-height: 400px;
    border: solid 1px black;
    background-color: #454545;
    padding: 10px;

    color: #27aa5e;
    line-height: 21pt;
    white-space: pre;
    width: 90%
}
</style>