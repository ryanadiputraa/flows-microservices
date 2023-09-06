export const refresh = {
	post: {
		tags: ['User Service'],
		description: 'refresh jwt tokens',
		parameters: [
			{
				name: 'Authorization',
				in: 'header',
				description: 'jwt auth refresh token',
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
				description: 'Refresh jwt tokens',
				content: {
					'application/json': {
						schema: {
							type: 'object',
							properties: {
								message: {
									type: 'string',
									example: 'user successfully sign in',
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
										access_token: {
											type: 'string',
											example:
												'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjFmZmJjYzQtNDRlNS00NjU1LTkwNmEtOTk2MWQ0Nzk2YWU0IiwiaWF0IjoxNjk0MDAzOTU3LCJleHAiOjE2OTQwMDc1NTd9.ykNPkSk54Uo0YcSma3psSMtbVH80P51qbkYqhsJCtlk',
										},
										expires_in: {
											type: 'number',
											example: 1694007557,
										},
										refresh_token: {
											type: 'string',
											example:
												'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjFmZmJjYzQtNDRlNS00NjU1LTkwNmEtOTk2MWQ0Nzk2YWU0IiwiaWF0IjoxNjk0MDAzOTU3LCJleHAiOjE2OTY1OTU5NTd9.wDrohN_LaO_jx2cu0ccLUWV_tPKLTwmTo5RDDoLfF9A',
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
