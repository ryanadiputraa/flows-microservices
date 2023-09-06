import { Request, Response } from 'express';
import axios from 'axios';

import { Logger } from '../server/logger';
import { catchServiceErr } from '../utils/service';

class AuthService {
	private baseURL = 'http://auth';
	private logger: Logger;

	constructor(logger: Logger) {
		this.logger = logger;
	}

	refreshTokens = async (req: Request, res: Response) => {
		try {
			const tokens = req.headers['authorization'];
			const [bearer, refresh_token] = tokens?.split(' ') ?? '';
			if (!bearer || bearer !== 'Bearer' || !refresh_token)
				return res.status(403).json({
					message: 'missing authorization header',
					err_code: 'forbidden',
					errors: null,
					data: null,
				});

			const resp = await axios.post(`${this.baseURL}/api/refresh`, { refresh_token });
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

export default AuthService;
