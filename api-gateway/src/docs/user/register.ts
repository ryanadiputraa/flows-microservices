export const register = {
	post: {
		tags: ['User Service'],
		description: 'register new user by email',
		operationId: 'registerUser',
		parameters: [
			{
				name: 'payload',
				in: 'path',
				required: true,
				description: 'register user payload',
				schema: {
					type: 'object',
					properties: {
						email: {
							type: 'string',
							example: 'johndoe@mail.com',
						},
						password: {
							type: 'string',
							example: 'secretpassword',
						},
						first_name: {
							type: 'string',
							example: 'john',
						},
						last_name: {
							type: 'string',
							example: 'doe',
						},
						currency: {
							type: 'string',
							example: 'USD',
							description: 'must be a valid suppoerted currency',
						},
						picture: {
							type: 'string',
							example: 'https://domain.com/img.png',
						},
					},
				},
			},
		],
		responses: {
			201: {
				description: 'User successfully regiter',
				content: {
					'application/json': {
						schema: {
							type: 'object',
							properties: {
								message: {
									type: 'string',
									example: 'user successfully register',
								},
								err_code: {
									type: 'string',
									example: null,
								},
								errors: {
									type: 'object',
									example: null,
								},
								data: {
									type: 'object',
									properties: {
										email: {
											type: 'string',
											example: 'johndoe@mail.com',
										},
										first_name: {
											type: 'string',
											example: 'john',
										},
										last_name: {
											type: 'string',
											example: 'doe',
										},
										currency: {
											type: 'string',
											example: 'USD',
											description: 'must be a valid suppoerted currency',
										},
										picture: {
											type: 'string',
											example: 'https://domain.com/img.png',
										},
									},
								},
							},
						},
					},
				},
			},
			400: {
				description: 'Invalid params',
				content: {
					'application/json': {
						schema: {
							type: 'object',
							properties: {
								message: {
									type: 'string',
									example: 'fail to register user',
								},
								err_code: {
									type: 'string',
									example: 'invalid_params',
								},
								errors: {
									type: 'object',
									properties: {
										email: {
											type: 'Array',
											example: ['email is requiured', 'email must be a valid email address'],
										},
										password: {
											type: 'Array',
											example: ['password must at least 8 characters'],
										},
										picture: {
											type: 'Array',
											example: ['picture must be a valid url'],
										},
									},
								},
								data: {
									type: 'object',
									example: null,
								},
							},
						},
					},
				},
			},
		},
	},
};
