package middleware

import (
	"context"
	"net/http"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/jwt"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/pkg/response"
)

type AuthMiddleware struct {
	jwt jwt.JWTService
}

func NewAuthMiddleware(jwt jwt.JWTService) *AuthMiddleware {
	return &AuthMiddleware{jwt: jwt}
}

func (m *AuthMiddleware) ClaimJWTToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := m.jwt.ExtractJWTTokenHeader(r.Header)
		if err != nil {
			response.WriteErrorResponse(w, r, http.StatusUnauthorized, "missing token header", domain.UNAUTHENTICATED, nil)
			return
		}

		claims, err := m.jwt.ParseJWTClaims(r.Context(), token)
		if err != nil {
			response.WriteErrorResponse(w, r, http.StatusForbidden, "invalid jwt token", domain.FORBIDDEN, nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		next(w, r.WithContext(ctx))
	})
}
