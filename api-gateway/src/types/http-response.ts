export type ERR_CODE =
	| null
	| 'invalid_params'
	| 'unauthenticated'
	| 'forbidden'
	| 'not_found'
	| 'internal_server_error';

export class ResponseError extends Error {
	status: number;
	errCode: ERR_CODE;
	message: string;

	constructor(status: number, errCode: ERR_CODE, message: string) {
		super();
		this.status = status;
		this.errCode = errCode;
		this.message = message;
	}
}

export interface HttpResponse<T> {
	message: string;
	err_code: ERR_CODE;
	errors: object;
	data: T;
}
