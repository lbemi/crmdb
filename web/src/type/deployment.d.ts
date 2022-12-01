export interface Label {
  app: string
}

export interface ManagedField {
  manager: string
  operation: string
  apiVersion: string
  time: string
  fieldsType: string
  fieldsV1: FieldsV1
}

export interface Metadata {
  name: string
  namespace: string
  uid: string
  resourceVersion: string
  generation: number
  creationTimestamp: string
  labels: Label
  annotations: Annotation
  managedFields: ManagedField[]
}

export interface MatchLabel {
  app: string
}

export interface Selector {
  matchLabels: MatchLabel
}

export interface Label {
  app: string
}

export interface ObjectMeta {
  creationTimestamp: any
  labels: object
}

export interface Port {
  containerPort: number
  protocol: string
}

export interface Resource {}

export interface Container {
  name: string
  image: string
  ports: Port[]
  resources: Resource
  terminationMessagePath: string
  terminationMessagePolicy: string
  imagePullPolicy: string
}

export interface SecurityContext {}

export interface Spec {
  containers: Container[]
  restartPolicy: string
  terminationGracePeriodSeconds: number
  dnsPolicy: string
  securityContext: SecurityContext
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

export interface Spec {
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
  availableReplicas: number
  conditions: Condition[]
}

export interface Deployment {
  metadata: Metadata
  spec: Spec
  status: Statu
}
