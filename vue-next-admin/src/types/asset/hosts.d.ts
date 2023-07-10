import { Group } from '@/types/asset/group';

declare interface Host {
	id: number;
	created_at: string;
	updated_at: string;
	group_id: number;
	labels: string[];
	ip: string;
	remark: string;
	port: number;
	status: number;
	enable_ssh: number;
}

interface Data {
	hosts: Host[];
	total: number;
}

export type HostResponseType = {
	code: number;
	data: Data;
	message: string;
};

interface HostTableType extends TableType {
	hosts: Host[];
	groups: Group[];
}

declare interface HostState {
	inputValue: string;
	type: string;
	defaultProps: {
		children: string;
		label: string;
	};
	groupName: string;
	groupIds: string;
	tableData: HostTableType;
}

declare interface HostParams {
	page: number;
	limit: number;
	name?: string;
	userName?: string;
}
