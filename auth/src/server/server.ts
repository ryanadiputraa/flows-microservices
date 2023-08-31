import express, { Application } from 'express';

import { Logger } from './logger';
import PublicRouter from '../routes/public-api';
import { Config } from '../types/config';
import { errorMiddleware } from '../middleware/error-middleware';

class Server {
	private config: Config;
	private server: Application;
	private logger: Logger;

	constructor(config: Config, logger: Logger) {
		this.config = config;
		this.server = express();
		this.logger = logger;

		this.server.use(express.json());

		const publicRouter = new PublicRouter(this.config, logger);
		this.server.use(publicRouter.router);
		this.server.use(errorMiddleware);
	}

	run() {
		this.server.listen(this.config.port, () => {
			this.logger.info(`http server running on port: ${this.config.port}`);
		});
	}
}

export default Server;
