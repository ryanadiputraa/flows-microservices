import winston, { format } from 'winston';

const { combine, timestamp, printf } = format;

const logFormat = printf(({ level, message, timestamp }) => {
	return `${timestamp} - ${level}: ${message}`;
});

export const logger = winston.createLogger({
	level: 'info',
	format: combine(timestamp(), logFormat),
	transports: [new winston.transports.Console({})],
});
