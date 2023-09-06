import { users } from './info';
import { login } from './login';
import { refresh } from './refresh';
import { register } from './register';

export const userService = {
	paths: {
		'/auth/register': {
			...register,
		},
		'/auth/login': {
			...login,
		},
		'/auth/refresh': {
			...refresh,
		},
		'/auth/users': {
			...users,
		},
	},
};
