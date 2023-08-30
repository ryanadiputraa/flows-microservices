import express, { Application } from 'express';
import dotenv from 'dotenv';

dotenv.config();
const port: Number = Number(process.env.PORT) || 8080;

const server: Application = express();
server.use(express.json());

export const serveHttp = (callback: () => void) => {
	server.listen(port, callback);
};
