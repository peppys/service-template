package middleware

import (
	"github.com/peppys/service-template/internal/services"
	"github.com/peppys/service-template/internal/utils"
	"log"
	"net/http"
	"strings"
)

func Authorization(authService *services.AuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Println("failed parsing malformed auth token")
			next.ServeHTTP(w, r)
			return
		}
		claims, err := authService.VerifyToken(r.Context(), authHeader[1])
		if err != nil {
			log.Printf("failed verifying token: %v", authHeader[1])
			next.ServeHTTP(w, r)
			return
		}

		log.Printf("Attatching auth to request: %v", claims)
		ctx := utils.ContextWithUserClaims(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
