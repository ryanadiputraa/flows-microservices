import express from 'express';

import { Logger } from '../server/logger';
import { Config } from '../types/config';
import UserSerive from '../service/user-service';
import AuthService from '../service/auth-service';
import TransactionSerive from '../service/transaction-service';

class PublicRouter {
	router: express.Router;

	constructor(__: Config, logger: Logger) {
		this.router = express.Router();
		const authService = new AuthService(logger);
		const userService = new UserSerive(logger);
		const transactionService = new TransactionSerive(logger);

		// auth
		this.router.post('/auth/register', userService.register);
		this.router.post('/auth/login', userService.login);
		this.router.post('/auth/refresh', authService.refreshTokens);

		// user
		this.router.get('/api/users', userService.userInfo);

		// transaction
		this.router.get('/api/transactions', transactionService.listTransactions);
		this.router.post('/api/transactions', transactionService.addTransaction);
		this.router.get('/api/transactions/summary', transactionService.getTransactionSummary);
	}
}

export default PublicRouter;
