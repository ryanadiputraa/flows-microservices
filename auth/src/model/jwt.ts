export interface JWTTokens {
	access_token: string;
	expires_in: number;
	refresh_token: string;
}

export interface GenerateJWTTokensDTO {
	user_id: string;
}
