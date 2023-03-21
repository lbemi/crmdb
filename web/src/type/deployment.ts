export interface Label {
  app: string
}

export interface ManagedField {
  manager: string
  operation: string
  apiVersion: string
  time: string
  fieldsType: string
  fieldsV1: object
}

export interface Metadata {
  name: string
  namespace: string
  uid: string
  resourceVersion: string
  generation: number
  creationTimestamp: string
  labels: object[]
  annotations: object[]
  managedFields: object
}

export interface Selector {
  matchLabels: object[]
}

export interface Label {
  app: string
}

export interface ObjectMeta {
  creationTimestamp: string
  labels: object
}

export interface Port {
  containerPort: number
  protocol: string
}

export interface Container {
  name: string
  image: string
  ports: Port[]
  resources: object
  terminationMessagePath: string
  terminationMessagePolicy: string
  imagePullPolicy: string
}

export interface Spec {
  containers: Container[]
  restartPolicy: string
  terminationGracePeriodSeconds: number
  dnsPolicy: string
  securityContext: object
  schedulerName: string
}

export interface Template {
  metadata: ObjectMeta
  spec: Spec
}

export interface RollingUpdate {
  maxUnavailable: string
  maxSurge: string
}

export interface Strategy {
  type: string
  rollingUpdate: RollingUpdate
}

export interface DeploySpec {
  replicas: number
  selector: Selector
  template: Template
  strategy: Strategy
  revisionHistoryLimit: number
  progressDeadlineSeconds: number
}

export interface Condition {
  type: string
  status: string
  lastUpdateTime: string
  lastTransitionTime: string
  reason: string
  message: string
}

export interface Statu {
  observedGeneration: number
  replicas: number
  updatedReplicas: number
  readyReplicas: number
  unavailableReplicas: number
  availableReplicas: number
  conditions: Condition[]
}

export interface Deployment {
  metadata: Metadata
  spec: DeploySpec
  status: Statu
}


export class deploymentData {
  Deployments: Deployment[] = []
}

interface query {
  namespace: string
  cloud: string
  deploymentName?: string
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
  selectData: Deployment[] = []
  Deployments: Deployment[] = []
}
