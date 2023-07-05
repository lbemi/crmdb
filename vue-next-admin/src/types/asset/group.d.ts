export interface Group {
	id: number;
	created_at: string;
	updated_at: string;
	name: string;
	parent_id: number;
	sequence: number;
	memo: string;
	status: string;
	children?: Group[];
}

export interface GroupData extends TableType {
	data: Group[];
}
