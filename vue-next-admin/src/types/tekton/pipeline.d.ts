export interface metadata {
	generation: number;
	uid: string;
	resourceVersion: string;
	name: string;
	namespace: string;
	creationTimestamp: string;
}

export interface paramsItem {
	name: string;
	value: string;
}

export interface specItem {
	description: string;
	workspaces: Array<workspacesItem>;
	params: Array<paramsItem>;
	tasks: Array<tasksItem>;
}

export interface taskRefItem {
	kind: string;
	name: string;
}

export interface managedFieldsItem {
	apiVersion: string;
	manager: string;
	fieldsV1: fieldsV1Item;
	time: string;
	operation: string;
	fieldsType: string;
}

export interface workspacesItem {
	name: string;
	description: string;
}

export interface tasksItem {
	taskRef: taskRefItem;
	name: string;
	workspaces: Array<workspacesItem>;
	params: Array<paramsItem>;
	runAfter: Array[string];
}

export interface Pipeline {
	metadata: metadata;
	apiVersion: string;
	kind: string;
	spec: Array<specItem>;
}
