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
  uid: string
  resourceVersion: string
  creationTimestamp: string
  labels: object
  annotations: object
  managedFields: ManagedField[]
}

export interface Taint {
  key: string
  value: string
  effect: string
}

export interface Spec {
  podCIDR: string
  podCIDRs: string[]
  taints: Taint[]
  unschedulable?: boolean
}

export interface Capacity {
  cpu: string
  'ephemeral-storage': string
  'hugepages-2Mi': string
  memory: string
  pods: string
}

export interface Allocatable {
  cpu: string
  'ephemeral-storage': string
  'hugepages-2Mi': string
  memory: string
  pods: string
}

export interface Condition {
  type: string
  status: string
  lastHeartbeatTime: string
  lastTransitionTime: string
  reason: string
  message: string
}

export interface Addresse {
  type: string
  address: string
}

export interface KubeletEndpoint {
  port: number
}

export interface DaemonEndpoint {
  kubeletEndpoint: KubeletEndpoint
}

export interface NodeInfo {
  machineID: string
  systemUUID: string
  bootID: string
  kernelVersion: string
  osImage: string
  containerRuntimeVersion: string
  kubeletVersion: string
  kubeProxyVersion: string
  operatingSystem: string
  architecture: string
}

export interface Image {
  names: string[]
  sizeBytes: number
}

export interface Statu {
  capacity: Capacity
  allocatable: Allocatable
  conditions: Condition[]
  addresses: Addresse[]
  daemonEndpoints: DaemonEndpoint
  nodeInfo: NodeInfo
  images: Image[]
}

export interface Node {
  metadata: Metadata
  spec: Spec
  status: Statu
}

export class NodeData {
  nodes: Node[] = []
}
