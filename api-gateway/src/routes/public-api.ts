import express from 'express';

import { Logger } from '../server/logger';
import { Config } from '../types/config';

class PublicRouter {
	router: express.Router;

	constructor(__: Config, _: Logger) {
		this.router = express.Router();
	}
}

export default PublicRouter;
