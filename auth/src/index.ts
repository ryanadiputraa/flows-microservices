import dotenv from 'dotenv';

import Server from './server/server';
import WinstonLogger from './server/logger';

dotenv.config();

const config = {
	port: Number(process.env.PORT) || 8080,
	jwtSecret: process.env.JWT_SECRET,
	jwtRefreshSecret: process.env.JWT_REFRESH_SECRET,
};

const logger = new WinstonLogger();

const server = new Server(config, logger);
server.run();
