package middlewares

import (
	"net/http"
	"strings"

	"github.com/DucAnh2611/ls-golang/config"
	"github.com/DucAnh2611/ls-golang/constants"
	"github.com/DucAnh2611/ls-golang/response"
	"github.com/DucAnh2611/ls-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const START_HEADER = "Bearer "

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, START_HEADER) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(constants.ErrUnauthorized))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, START_HEADER)
		tokenSecret := config.GetEnv(constants.JwtSecret, "")
		if compare := strings.Compare(tokenSecret, ""); compare == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error(constants.ErrInternal))
			return
		}

		token, err := utils.ValidateToken(tokenStr, tokenSecret)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(constants.ErrInvalidToken))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["data"] == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(constants.ErrInvalidClaims))
			return
		}

		c.Set(constants.AccessTokenPayloadKey, claims["data"].(map[string]any))

		c.Next()
	}
}
