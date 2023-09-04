import express from 'express';

import { Logger } from '../server/logger';
import { Config } from '../types/config';
import UserSerive from '../service/user-service';

class PublicRouter {
	router: express.Router;

	constructor(__: Config, logger: Logger) {
		this.router = express.Router();
		const userService = new UserSerive(logger);

		this.router.post('/auth/register', userService.register);
		this.router.post('/auth/login', userService.login);
		this.router.get('/api/users', userService.userInfo);
	}
}

export default PublicRouter;
