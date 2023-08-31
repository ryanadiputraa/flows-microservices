import express, { Application } from 'express';

import { Logger } from './logger';

class Server {
	private server: Application;
	private logger: Logger;

	constructor(config: () => any, logger: Logger) {
		config();
		this.server = express();
		this.logger = logger;

		this.server.use(express.json());
	}

	run() {
		const port: Number = Number(process.env.PORT) || 8080;
		this.server.listen(port, () => {
			this.logger.info(`http server running on port: ${port}`);
		});
	}
}

export default Server;
