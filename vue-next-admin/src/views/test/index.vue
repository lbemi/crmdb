<template>
	<div class="main">
		asdasd
		<code-mirror v-model="codeVal" basic style="height: 400px" :extensions="extensions" :phrases="phrases" />
	</div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import CodeMirror from 'vue-codemirror6';
import { oneDark } from '@codemirror/theme-one-dark';
import { json } from '@codemirror/lang-json';
import { StreamLanguage } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';
import { javascript } from '@codemirror/legacy-modes/mode/javascript';
import { Ref } from 'vue-demi';

// // 初始化
let codeVal = ref('');
// // 转成json字符串并格式化
codeVal.value = `metadata:
  name: sleep
  namespace: myistio
  uid: 9499e43d-02b9-41fb-ac83-bda4b8cb49dd
  resourceVersion: '236989499'
  generation: 1
  creationTimestamp: '2023-10-24T01:37:36Z'
  annotations:
    deployment.kubernetes.io/revision: '1'
    kubectl.kubernetes.io/last-applied-configuration: >
      {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"name":"sleep","namespace":"myistio"},"spec":{"replicas":1,"selector":{"matchLabels":{"app":"sleep"}},"template":{"metadata":{"labels":{"app":"sleep"}},"spec":{"containers":[{"command":["/bin/sleep","infinity"],"image":"curlimages/curl","imagePullPolicy":"IfNotPresent","name":"sleep","volumeMounts":[{"mountPath":"/etc/sleep/tls","name":"secret-volume"}]}],"serviceAccountName":"sleep","terminationGracePeriodSeconds":0,"volumes":[{"name":"secret-volume","secret":{"optional":true,"secretName":"sleep-secret"}}]}}}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sleep
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: sleep
    spec:
      volumes:
        - name: secret-volume
          secret:
            secretName: sleep-secret
            defaultMode: 420
            optional: true
      containers:
        - name: sleep
          image: curlimages/curl
          command:
            - /bin/sleep
            - infinity
          resources: {}
          volumeMounts:
            - name: secret-volume
              mountPath: /etc/sleep/tls
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: IfNotPresent
      restartPolicy: Always
      terminationGracePeriodSeconds: 0
      dnsPolicy: ClusterFirst
      serviceAccountName: sleep
      serviceAccount: sleep
      securityContext: {}
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
status:
  observedGeneration: 1
  replicas: 1
  updatedReplicas: 1
  readyReplicas: 1
  availableReplicas: 1
  conditions:
    - type: Available
      status: 'True'
      lastUpdateTime: '2023-10-24T01:37:45Z'
      lastTransitionTime: '2023-10-24T01:37:45Z'
      reason: MinimumReplicasAvailable
      message: Deployment has minimum availability.
    - type: Progressing
      status: 'True'
      lastUpdateTime: '2023-10-24T01:37:45Z'
      lastTransitionTime: '2023-10-24T01:37:36Z'
      reason: NewReplicaSetAvailable
      message: ReplicaSet "sleep-69cfb4968f" has successfully progressed.
`
//
// // json语言
// const lang = yaml();
// // 扩展
const extensions = [oneDark, StreamLanguage.define(javascript)];
const phrases: Ref<Record<string, string>> = ref({
	// @codemirror/view
	'Control character': '制御文字',
	// @codemirror/commands
	'Selection deleted': '選択を削除',
	// @codemirror/language
	'Folded lines': '折り畳まれた行',
	'Unfolded lines': '折り畳める行',
	to: '行き先',
	'folded code': '折り畳まれたコード',
	unfold: '折り畳みを解除',
	'Fold line': '行を折り畳む',
	'Unfold line': '行の折り畳む解除',
	// @codemirror/search
	'Go to line': '行き先の行',
	go: 'OK',
	Find: '検索',
	Replace: '置き換え',
	next: '▼',
	previous: '▲',
	all: 'すべて',
	'match case': '一致条件',
	'by word': '全文検索',
	regexp: '正規表現',
	replace: '置き換え',
	'replace all': 'すべてを置き換え',
	close: '閉じる',
	'current match': '現在の一致',
	'replaced $ matches': '$ 件の一致を置き換え',
	'replaced match on line $': '$ 行の一致を置き換え',
	'on line': 'した行',
	// @codemirror/autocomplete
	Completions: '自動補完',
	// @codemirror/lint
	Diagnostics: 'エラー',
	'No diagnostics': 'エラーなし',
});
</script>
s

<style>
/* required! */
.cm-editor {
	height: 100%;
}
</style>
