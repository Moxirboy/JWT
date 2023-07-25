package middleware

import (
	"JWT/internal/controller/http/v1/response"
	json "JWT/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	var (
		userId, err = json.ParseToken(headerParts[1])
	)
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
