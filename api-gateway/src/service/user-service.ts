import { Request, Response } from 'express';
import axios from 'axios';

import { catchServiceErr } from '../utils/service';
import { Logger } from '../server/logger';

class UserSerive {
	private baseURL = 'http://user';
	private logger: Logger;

	constructor(logger: Logger) {
		this.logger = logger;
	}

	register = async (req: Request, res: Response) => {
		try {
			const resp = await axios.post(`${this.baseURL}/auth/register`, { ...req.body });
			return res.status(resp.status).json(resp.data);
		} catch (error) {
			const { status, resp } = catchServiceErr(error);
			if (status >= 500 && status < 600) {
				this.logger.error(error);
			} else {
				this.logger.warn(error);
			}
			res.status(status).json(resp);
		}
	};

	login = async (req: Request, res: Response) => {
		try {
			const resp = await axios.post(`${this.baseURL}/auth/login`, { ...req.body });
			return res.status(resp.status).json(resp.data);
		} catch (error) {
			const { status, resp } = catchServiceErr(error);
			if (status >= 500 && status < 600) {
				this.logger.error(error);
			} else {
				this.logger.warn(error);
			}
			res.status(status).json(resp);
		}
	};

	userInfo = async (req: Request, res: Response) => {
		try {
			const resp = await axios.get(`${this.baseURL}/api/users`, {
				headers: {
					Authorization: req.headers.authorization,
				},
			});
			return res.status(resp.status).json(resp.data);
		} catch (error) {
			const { status, resp } = catchServiceErr(error);
			if (status >= 500 && status < 600) {
				this.logger.error(error);
			} else {
				this.logger.warn(error);
			}
			res.status(status).json(resp);
		}
	};
}

export default UserSerive;
