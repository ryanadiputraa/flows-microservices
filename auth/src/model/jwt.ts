export interface JWTTokens {
	accessToken: string;
	expiresIn: number;
	refreshToken: string;
}

export interface GenerateJWTTokensDTO {
	user_id: string;
}
