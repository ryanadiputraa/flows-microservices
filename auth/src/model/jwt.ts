export interface JWTTokens {
	access_token: string;
	expires_in: number;
	refresh_token: string;
}

export interface JWTClaimsDTO {
	user_id: string;
}
