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
  labels: { [index: string]: string }
  managedFields: ManagedField[]
}

export interface Spec {
  finalizers: string[]
}

export interface Statu {
  phase: string
}

export interface Namespace {
  metadata: Metadata
  spec: Spec
  status: Statu
}
export interface NamespaceForm {
  name: string
  labels: { [index: string]: string }
}

export class NamespaceData {
  titile = ''
  visible = false
  loading = false
  namespace?: Namespace
  selectData: Namespace[] = []
}

export class NamespaceFormData {
  metadata: NamespaceForm = {
    name: '',
    labels: {}
  }
}
