package middlewares

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils/consts"
	userconsts "Blog_API/pkg/utils/consts/user"
	"Blog_API/pkg/utils/response"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		conf := config.LocalConfig

		authHeader := c.Request().Header.Get(consts.Authorization)
		if authHeader == "" {
			return response.ErrorResponseWithStatus(c, http.StatusUnauthorized, consts.AuthorizationHeaderRequired)
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != consts.Bearer {
			return response.ErrorResponseWithStatus(c, http.StatusUnauthorized, consts.InvalidToken)
		}

		var jwtClaims types.JWTClaims

		// Parse the token with claims
		token, err := jwt.ParseWithClaims(tokenParts[1], &jwtClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(conf.JWTSecret), nil
		})
		if err != nil {
			return response.ErrorResponseWithStatus(c, http.StatusUnauthorized, consts.InvalidToken)
		}

		if claims, ok := token.Claims.(*types.JWTClaims); ok && token.Valid {
			c.Set(userconsts.UserID, claims.UserID)
			c.Set(userconsts.UserEmail, claims.UserEmail)
			return next(c)
		}

		return response.ErrorResponseWithStatus(c, http.StatusUnauthorized, consts.InvalidToken)
	}
}
