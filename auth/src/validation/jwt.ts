import Joi from 'joi';

export const generateJWTTokensValidation = Joi.object({
	user_id: Joi.string().required(),
});
