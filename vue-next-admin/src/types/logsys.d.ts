declare interface LoginLog {
	id: number;
	created_at: string;
	updated_at: string;
	username: string;
	status: string;
	ipaddr: string;
	loginLocation: string;
	browser: string;
	os: string;
	platform: string;
	loginTime: string;
	remark: string;
	msg: string;
}

declare interface OperatorLog {
	id: number;
	created_at: string;
	updated_at: string;
	title: string;
	businessType: string;
	method: string;
	name: string;
	url: string;
	ip: string;
	location: string;
	param: string;
	status: number;
	errMsg: string;
}
