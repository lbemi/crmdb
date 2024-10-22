declare interface Account {
	id: number;
	created_at: string;
	updated_at: string;
	name: string;
	user_name: string;
	password: string;
	auth_method: string;
	secret: string;
	status: number;
}

declare interface HostBingAccounts {
	id: number;
	created_at: string;
	updated_at: string;
	rule_name: string;
	account_id: number[];
	resource_id: number[];
}
