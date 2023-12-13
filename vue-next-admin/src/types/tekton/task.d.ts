// interface TaskDefinition {
// 	namespace: string;
// 	uid: string;
// 	resourceVersion: string;
// 	generation: number;
// 	creationTimestamp: string;
// 	annotations: annotations;
// 	spec: {
// 		params: {
// 			name: string;
// 			type: string;
// 		}[];
// 		steps: {
// 			name: string;
// 			image: string;
// 			computeResources: any;
// 			script: string;
// 		}[];
// 	};
// }

export interface Metadata {
	name: string;
	namespace: string;
	uid: string;
	resourceVersion: string;
	generation: number;
	creationTimestamp: string;
	annotations: annotations;
}

export interface annotations {
	[key: string]: string;
}

export interface Param {
	name: string;
	type: string;
}

export interface ComputeResource {}

export interface Step {
	name: string;
	image: string;
	computeResources: ComputeResource;
	script: string;
}

export interface Spec {
	params: Param[];
	steps: Step[];
}

export interface Task {
	kind: string;
	apiVersion: string;
	metadata: Metadata;
	spec: Spec;
}
