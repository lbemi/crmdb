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

interface HosTableType extends TableType {
	data: Host[];
}

declare interface HostState {
	tableData: HosTableType;
}
