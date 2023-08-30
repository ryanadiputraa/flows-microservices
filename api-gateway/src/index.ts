import { serveHttp } from './server/server';
import { logger } from './server/logger';

serveHttp(() => logger.info('http server running...'));
