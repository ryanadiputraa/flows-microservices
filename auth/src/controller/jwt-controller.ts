import { NextFunction, Request, Response } from 'express';

import { JWTClaimsDTO, JWTTokens } from '../model/jwt';
import JWTService from '../service/jwt-service';
import { HttpResponse, ResponseError } from '../types/http-response';

class JwtController {
	private service: JWTService;

	constructor(service: JWTService) {
		this.service = service;
	}

	generateJWTTokens = async (req: Request, res: Response, next: NextFunction) => {
		let resp: HttpResponse<JWTTokens>;
		try {
			const dto: JWTClaimsDTO = req.body;
			const tokens = await this.service.generateJWTTokens(dto);
			resp = {
				message: 'jwt tokens generated',
				err_code: null,
				errors: null,
				data: tokens,
			};
			res.status(200).json(resp);
		} catch (error) {
			next(error);
		}
	};

	ParseJWTClaims = async (req: Request, res: Response, next: NextFunction) => {
		let resp: HttpResponse<JWTClaimsDTO>;
		try {
			const token = req.query['token'] ?? '';
			if (!token) throw new ResponseError(400, 'invalid_params', 'missing jwt access token');
			const claims = await this.service.ParseJWTClaims(String(token));
			resp = {
				message: 'successfully parse jwt claims',
				err_code: null,
				errors: null,
				data: claims,
			};
			res.status(200).json(resp);
		} catch (error) {
			next(error);
		}
	};
}

export default JwtController;
