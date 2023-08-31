import { NextFunction, Request, Response } from 'express';
import { ValidationError } from 'joi';

import { HttpResponse, ResponseError } from '../types/http-response';
import { parseValidationError } from '../validation/validation';

export const errorMiddleware = async (err: any, _: Request, res: Response, next: NextFunction) => {
	if (!err) {
		next();
		return;
	}

	let resp: HttpResponse<null>;
	if (err instanceof ResponseError) {
		resp = {
			message: err.message,
			err_code: err.errCode,
			errors: {},
			data: null,
		};
		res.status(err.status).json(resp).end();
	} else if (err instanceof ValidationError) {
		const errors = parseValidationError(err.message);

		resp = {
			message: 'invalid parameters',
			err_code: 'invalid_params',
			errors: errors,
			data: null,
		};
		res.status(400).json(resp).end();
	} else {
		resp = {
			message: 'unkown errors occurred',
			err_code: 'internal_server_error',
			errors: {},
			data: null,
		};
		res.status(500).json(resp).end();
	}
};
