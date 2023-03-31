import { UploadFile } from 'element-plus'

export interface ClusterInfo {
    id: number
    created_at: string
    updated_at: string
    name: string
    version: string
    runtime: string
    service_cidr: string
    pod_cidr: string
    cni: string
    proxy_mode: string
    status: boolean
    nodes: number
    internal_ip: string
    cpu: number
    memory: number
}
export interface ClusterForm {
    name: string
    kube_config: UploadFile
}
