import { NodeSpec, NodeStatus } from 'kubernetes-types/core/v1';
import { ObjectMeta } from 'kubernetes-types/meta/v1';
import { UploadFile } from 'element-plus';

export interface ClusterInfo {
	id: number;
	created_at: string;
	updated_at: string;
	name: string;
	version: string;
	runtime: string;
	service_cidr: string;
	pod_cidr: string;
	cni: string;
	proxy_mode: string;
	status: boolean;
	nodes: number;
	internal_ip: string;
	cpu: number;
	memory: number;
}
export interface ClusterForm {
	name: string;
	kube_config: UploadFile;
}

export interface Node {
	apiVersion?: string;
	kind?: string;
	metadata?: ObjectMeta;
	spec?: NodeSpec;
	status?: NodeStatus;
	usage?: {
		cpu: number;
		memory: number;
		pod: number;
	};
	discriminator: string | undefined;
	attributeTypeMap: Array<{
		name: string;
		baseName: string;
		type: string;
	}>;
	static getAttributeTypeMap(): {
		name: string;
		baseName: string;
		type: string;
	}[];
}

export interface FileType {
	name: string;
	isDir: boolean;
	fsType: string;
	size: number;
	lastModify: string;
	user: string;
	group: string;
}