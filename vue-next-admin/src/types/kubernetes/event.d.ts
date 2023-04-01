import * as k8s from '@kubernetes/client-node';
declare interface EventData extends QueryType{
    events: k8s.CoreV1Event[]
}

