package domain

type JWTTokens struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type JWTClaims struct {
	UserID string `json:"user_id"`
}

type ServiceResponse[T any] struct {
	Message string              `json:"message"`
	ErrCode string              `json:"err_code"`
	Errors  map[string][]string `json:"errros"`
	Data    T                   `json:"data"`
}
