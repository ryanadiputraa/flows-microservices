import jwt from 'jsonwebtoken';
import { ValidationError } from 'joi';

import { GenerateJWTTokensDTO, JWTTokens } from '../model/jwt';
import { Logger } from '../server/logger';
import { generateJWTTokensValidation } from '../validation/jwt';
import { validateRequest } from '../validation/validation';
import { Config } from '../types/config';

class JWTService {
	private log: Logger;
	private secret: string;
	private refreshSecret: string;

	constructor(logger: Logger, config: Config) {
		this.log = logger;
		this.secret = config.jwtSecret;
		this.refreshSecret = config.jwtRefreshSecret;
	}

	generateJWTTokens = async (dto: GenerateJWTTokensDTO): Promise<JWTTokens> => {
		try {
			const claims = validateRequest<GenerateJWTTokensDTO>(generateJWTTokensValidation, dto);
			const access_token = jwt.sign(claims, this.secret, { expiresIn: '1h' });
			const refresh_token = jwt.sign(claims, this.refreshSecret, { expiresIn: '720h' });
			const expires_in = Math.floor(Date.now() / 1000) + 3600;

			return {
				access_token,
				expires_in,
				refresh_token,
			};
		} catch (error) {
			if (error instanceof ValidationError) {
				this.log.warn('generate access token: ' + error.message);
			} else {
				this.log.error(error);
			}
			throw error;
		}
	};
}

export default JWTService;
