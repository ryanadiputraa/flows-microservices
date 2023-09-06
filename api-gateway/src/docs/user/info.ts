export const users = {
	get: {
		tags: ['User Service'],
		description: 'get user info',
		parameters: [
			{
				name: 'Authorization',
				in: 'header',
				description: 'jwt auth token',
				required: true,
				schema: {
					type: 'string',
					example:
						'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjFmZmJjYzQtNDRlNS00NjU1LTkwNmEtOTk2MWQ0Nzk2YWU0IiwiaWF0IjoxNjk0MDAzOTU3LCJleHAiOjE2OTQwMDc1NTd9.ykNPkSk54Uo0YcSma3psSMtbVH80P51qbkYqhsJCtlk',
				},
			},
		],
		responses: {
			200: {
				description: 'Fetch user info',
				content: {
					'application/json': {
						schema: {
							type: 'object',
							properties: {
								message: {
									type: 'string',
									example: 'fetch user info',
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
										id: {
											type: 'string',
											example: '61ffbcc4-44e5-4655-906a-9961d4796ae4',
										},
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
			401: {
				description: 'Invalid params',
				content: {
					'application/json': {
						schema: {
							type: 'object',
							properties: {
								message: {
									type: 'string',
									example: 'missing auth header',
								},
								err_code: {
									type: 'string',
									example: 'unauthenticated',
								},
								errors: {
									type: 'object',
									example: null,
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
