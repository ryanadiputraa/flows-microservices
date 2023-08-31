import winston, { format } from 'winston';

const { combine, timestamp, printf } = format;
export interface Logger {
	debug: (msg: string) => void;
	info: (msg: string) => void;
	warn: (msg: string) => void;
	error: (msg: string) => void;
}

class WinstonLogger implements Logger {
	private logger: winston.Logger;

	private logFormat = printf(({ level, message, timestamp }) => {
		return `${timestamp} - ${level}: ${message}`;
	});

	constructor() {
		this.logger = winston.createLogger({
			level: 'info',
			format: combine(timestamp(), this.logFormat),
			transports: [new winston.transports.Console({})],
		});
	}

	debug(msg: string) {
		this.logger.debug(msg);
	}
	info(msg: string) {
		this.logger.info(msg);
	}
	warn(msg: string) {
		this.logger.warn(msg);
	}
	error(msg: string) {
		this.logger.error(msg);
	}
}

export default WinstonLogger;
