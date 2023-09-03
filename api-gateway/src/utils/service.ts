import { HttpResponse } from '../types/http-response';

export const catchServiceErr = <T>(error: any): { status: number; resp: HttpResponse<T> } => {
	if (error?.response) {
		const status = error.response.status;
		const data: HttpResponse<T> = error.response.data ?? {};
		return { status, resp: data };
	}
	return {
		status: 500,
		resp: {
			message: 'internal server error',
			err_code: 'internal_server_error',
			errors: null,
			data: null,
		},
	};
};
