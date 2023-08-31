import dotenv from 'dotenv';

import Server from './server/server';
import WinstonLogger from './server/logger';

const logger = new WinstonLogger();

const server = new Server(dotenv.config, logger);
server.run();
