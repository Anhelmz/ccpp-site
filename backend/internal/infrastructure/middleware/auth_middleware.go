package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const appMetadataNamespace = "https://ccphnompenh-com/app_metadata"

// Auth0Claims represents the JWT claims we expect after Auth0 login.
type Auth0Claims struct {
	Sub         string                 `json:"sub"`
	Email       string                 `json:"email"`
	AppMetadata map[string]interface{} `json:"https://ccphnompenh-com/app_metadata"`
	jwt.RegisteredClaims
}

// Toggle to bypass auth for testing; set to false to enforce tokens.
const bypassAuth = false

// Auth0Middleware validates our signed JWT and enforces role/site access.
func Auth0Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bypassAuth {
			c.Next()
			return
		}

		tokenString := extractToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			c.Abort()
			return
		}

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "your-secret-key" // development fallback
		}

		token, err := jwt.ParseWithClaims(tokenString, &Auth0Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Auth0Claims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		if !isAuthorized(claims) {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

// OptionalAuth allows requests without auth but attaches user if token is valid.
func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bypassAuth {
			c.Next()
			return
		}

		tokenString := extractToken(c)
		if tokenString == "" {
			c.Next()
			return
		}

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "your-secret-key"
		}

		token, err := jwt.ParseWithClaims(tokenString, &Auth0Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(*Auth0Claims); ok {
				c.Set("user", claims)
			}
		}

		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func isAuthorized(claims *Auth0Claims) bool {
	siteID := os.Getenv("SITE_ID")
	if siteID == "" {
		siteID = "ccphnompenh-com"
	}

	if role, exists := claims.AppMetadata["role"]; exists {
		if roleStr, ok := role.(string); ok && roleStr == "admin" {
			return true
		}
	}

	if sites, exists := claims.AppMetadata["sites"]; exists {
		if sitesList, ok := sites.([]interface{}); ok {
			for _, site := range sitesList {
				if siteStr, ok := site.(string); ok && siteStr == siteID {
					return true
				}
			}
		}
	}

	return false
}
