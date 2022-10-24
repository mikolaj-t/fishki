package rest

import (
	"crypto/subtle"
	"fishki/pkg/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const (
	authenticatedKey = "authenticated"
	userKey          = "user"
	sessionCookie    = "sessionID"
	usernameCookie   = "username"
)

type SessionMiddleware struct {
	Service core.UserService
}

func (s SessionMiddleware) Session(ctx *gin.Context) {
	sessionId, err := ctx.Cookie(sessionCookie)
	username, err := ctx.Cookie(usernameCookie)

	if sessionId == "" || username == "" {
		ctx.Status(http.StatusUnauthorized)
		fmt.Println("1 S", sessionId, " U", username)
		ctx.Abort()
		return
	}

	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		fmt.Println("2", err)
		ctx.Abort()
		return
	}

	user, err := s.Service.GetByUsername(username)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		fmt.Fprintln(os.Stderr, err)
		ctx.Abort()
		return
	}

	if subtle.ConstantTimeCompare(user.Session, []byte(sessionId)) != 1 {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Set(authenticatedKey, true)
	ctx.Set(userKey, user)

	ctx.Next()
}

func UserFromCtx(ctx *gin.Context) (*core.User, bool) {
	u, b := ctx.Get(userKey)
	if !b {
		return nil, b
	}
	user := u.(*core.User)
	return user, b
}
