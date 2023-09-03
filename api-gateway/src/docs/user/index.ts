import { register } from './register';

export const userService = {
	paths: {
		'/auth/register': {
			...register,
		},
	},
};
