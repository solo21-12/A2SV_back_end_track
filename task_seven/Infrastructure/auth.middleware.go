package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

func AuthMiddleware() gin.HandlerFunc {
	// this middleware checks if the user is authenticated
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		_, err := GetClaims(authHeader)

		if err != nil {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func RoleBasedMiddleWare(roles ...string) gin.HandlerFunc {
	// this middleware checks if the user is authorized to perform an action
	return func(ctx *gin.Context) {

		claims, err := GetClaims(ctx.GetHeader("Authorization"))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		userRole := claims.Role

		method := ctx.Request.Method

		switch method {
		case POST, DELETE, PUT, PATCH:
			authrized := false

			for _, role := range roles {
				if role == userRole {
					authrized = true
					break
				}
			}

			if !authrized {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden only admin user can perform this action"})
				ctx.Abort()
				return
			}
		}

		ctx.Next()

	}
}
