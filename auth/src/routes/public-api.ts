import express from 'express';

import JwtController from '../controller/jwt-controller';
import JWTService from '../service/jwt-service';
import { Logger } from '../server/logger';
import { Config } from '../types/config';

class PublicRouter {
	router: express.Router;

	constructor(config: Config, logger: Logger) {
		this.router = express.Router();

		const jwtService = new JWTService(logger, config);
		const jwtController = new JwtController(jwtService);

		this.router.post('/api/tokens', jwtController.generateJWTTokens);
		this.router.get('/api/claims', jwtController.ParseJWTClaims);
	}
}

export default PublicRouter;
