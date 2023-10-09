import { Request, Response } from 'express';
import axios from 'axios';

import { catchServiceErr } from '../utils/service';
import { Logger } from '../server/logger';

class TransactionSerive {
	private baseURL = 'http://transaction';
	private logger: Logger;

	constructor(logger: Logger) {
		this.logger = logger;
	}

	addTransaction = async (req: Request, res: Response) => {
		try {
			const resp = await axios.post(
				`${this.baseURL}/api/transactions`,
				{ ...req.body },
				{
					headers: {
						Authorization: req.headers.authorization,
					},
				}
			);
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

	getTransactionSummary = async (req: Request, res: Response) => {
		try {
			const resp = await axios.get(`${this.baseURL}/api/transactions`, {
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

export default TransactionSerive;
