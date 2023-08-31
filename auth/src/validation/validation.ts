import Joi from 'joi';

export const validateRequest = <T>(schema: Joi.ObjectSchema, req: object): T => {
	const res = schema.validate(req, {
		abortEarly: false,
		convert: false,
	});
	if (res.error) throw res.error;
	return res.value;
};

export const parseValidationError = (errorMessage: string): object => {
	const errorObject = {};
	const errors = errorMessage.split('.');

	errors.forEach((error) => {
		const fieldErrors = error.match(/"([^"]+)"/g); // Extract field names from error message
		const key = JSON.parse(fieldErrors[0]);
		const msg = error.split('" ')?.[1];

		if (!errorObject[key]) {
			errorObject[key] = [msg];
		} else {
			errorObject[key].push(msg);
		}
	});

	return errorObject;
};
