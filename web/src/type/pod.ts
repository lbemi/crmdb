export interface OwnerReference {
  apiVersion: string
  kind: string
  name: string
  uid: string
  controller: boolean
  blockOwnerDeletion: boolean
}

export interface Metadata {
  name: string
  generateName: string
  namespace: string
  uid: string
  resourceVersion: string
  creationTimestamp: string
  labels: object[]
  annotations: object[]
  ownerReferences: OwnerReference[]
}

export interface ServiceAccountToken {
  expirationSeconds: number
  path: string
}

export interface Source {
  serviceAccountToken: ServiceAccountToken
}

export interface Projected {
  sources: Source[]
  defaultMode: number
}

export interface Volume {
  name: string
  projected: Projected
}

export interface Port {
  containerPort: number
  protocol: string
}

export interface Resource {}

export interface VolumeMount {
  name: string
  readOnly: boolean
  mountPath: string
}

export interface Container {
  name: string
  image: string
  ports: Port[]
  resources: Resource
  volumeMounts: VolumeMount[]
  terminationMessagePath: string
  terminationMessagePolicy: string
  imagePullPolicy: string
}

export interface SecurityContext {}

export interface Toleration {
  key: string
  operator: string
  effect: string
  tolerationSeconds: number
}

export interface Spec {
  volumes: Volume[]
  containers: Container[]
  restartPolicy: string
  terminationGracePeriodSeconds: number
  dnsPolicy: string
  serviceAccountName: string
  serviceAccount: string
  nodeName: string
  securityContext: SecurityContext
  schedulerName: string
  tolerations: Toleration[]
  priority: number
  enableServiceLinks: boolean
  preemptionPolicy: string
}

export interface Condition {
  type: string
  status: string
  lastProbeTime?: any
  lastTransitionTime: string
}

export interface PodIP {
  ip: string
}

export interface Running {
  startedAt: string
}

export interface State {
  running: Running
}

export interface LastState {}

export interface ContainerStatuse {
  name: string
  state: State
  lastState: LastState
  ready: boolean
  restartCount: number
  image: string
  imageID: string
  containerID: string
  started: boolean
}

export interface Status {
  phase: string
  conditions: Condition[]
  hostIP: string
  podIP: string
  podIPs: PodIP[]
  startTime: string
  containerStatuses: ContainerStatuse[]
  qosClass: string
}

export interface Pod {
  metadata: Metadata
  spec: Spec
  status: Status
}

export class podData {
  Pods: Pod[] = []
}

interface query {
  namespace: string
  cloud: string
  podName?: string
  page: number
  limit: number
}

export class Data {
  query: query = {
    namespace: '',
    cloud: '',
    page: 1,
    limit: 10
  }

  total = 0
  loading = false
  selectData: Pod[] = []
  pods: Pod[] = []
}
