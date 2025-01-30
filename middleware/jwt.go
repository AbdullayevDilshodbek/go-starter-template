package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware - Faqat avtorizatsiyadan o‘tgan foydalanuvchilarni ruxsat berish
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if(r.URL.Path == "/api/v1/login") {
			next.ServeHTTP(w, r)
            return
		}
		w.Header().Set("Content-Type", "application/json")

		// Authorization header ni olish
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
			return
		}

		// "Bearer TOKEN" formatida kelishi kerak
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, `{"error": "Invalid token format"}`, http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]

		// Tokenni tekshirish
		claims, err := validateToken(token)
		if err != nil {
			http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
			return
		}

		// Token to'g'ri bo'lsa, so'rovni davom ettirish
		log.Println("User ID:", claims["user_id"]) // User ID'ni logga chiqarish (kerak bolsa)
		next.ServeHTTP(w, r)
	})
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Tokenning signing method'ini tekshirish
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// Token valid ekanligini tekshirish
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Token muddati tugaganini tekshirish
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, jwt.ErrTokenExpired
			}
		}
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}