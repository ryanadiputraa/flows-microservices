import { NextFunction, Request, Response } from 'express';

import { GenerateJWTTokensDTO, JWTTokens } from '../model/jwt';
import JWTService from '../service/jwt-service';
import { HttpResponse } from '../types/http-response';

class JwtController {
	private service: JWTService;

	constructor(service: JWTService) {
		this.service = service;
	}

	generateJWTTokens = async (req: Request, res: Response, next: NextFunction) => {
		let resp: HttpResponse<JWTTokens>;
		try {
			const dto: GenerateJWTTokensDTO = req.body;
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
}

export default JwtController;
