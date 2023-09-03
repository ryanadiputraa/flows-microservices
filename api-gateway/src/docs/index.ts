import { servers } from './server';
import { userService } from './user';

const info = {
	openapi: '3.0.3',
	info: {
		title: 'Flows Microservice',
		description: 'Microservice for Flows app',
		version: '1.0.0',
		contact: {
			name: 'Ryan Adi Putra',
			email: 'ryannadiputraa@gmail.com',
			url: 'ryanadiputra.vercel.app',
		},
	},
};

const docs = {
	...info,
	...servers,
	...userService,
};

export default docs;
