interface ManagedField {
  manager: string
  operation: string
  apiVersion: string
  time: string
  fieldsType: string
}

interface Metadata {
  name: string
  namespace: string
  uid: string
  resourceVersion: string
  creationTimestamp: string
  managedFields: ManagedField[]
}

interface InvolvedObject {
  kind: string
  namespace: string
  name: string
  uid: string
  apiVersion: string
  resourceVersion: string
  fieldPath: string
}

interface Source {
  component: string
  host: string
}

export interface Event {
  metadata: Metadata
  involvedObject: InvolvedObject
  reason: string
  message: string
  source: Source
  firstTimestamp: string
  lastTimestamp: string
  count: number
  type: string
  eventTime?: string
  reportingComponent: string
  reportingInstance: string
}
interface query {
  namespace: string
  cloud: string
  name?: string
}

export class EventData {
  events: Event[] = []
  query: query = {
    namespace: '',
    cloud: ''
  }
  loading = false
}
