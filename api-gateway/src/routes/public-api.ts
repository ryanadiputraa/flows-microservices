import express from 'express';

import { Logger } from '../server/logger';
import { Config } from '../types/config';
import UserSerive from '../service/user-service';
import AuthService from '../service/auth-service';

class PublicRouter {
	router: express.Router;

	constructor(__: Config, logger: Logger) {
		this.router = express.Router();
		const authService = new AuthService(logger);
		const userService = new UserSerive(logger);

		this.router.post('/auth/register', userService.register);
		this.router.post('/auth/login', userService.login);
		this.router.post('/auth/refresh', authService.refreshTokens);
		this.router.get('/api/users', userService.userInfo);
	}
}

export default PublicRouter;
