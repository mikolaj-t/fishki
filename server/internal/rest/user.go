package rest

import (
	"fishki/pkg/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var DOMAIN string = os.Getenv("DOMAIN")

type userHandler struct {
	service core.UserService
}

func NewUserHandler(service core.UserService) core.UserHandler {
	return &userHandler{service}
}

func (u userHandler) Setup(router *gin.Engine, authorized *gin.RouterGroup) {
	router.POST("user/login", u.Login)
	router.POST("user/register", u.Register)
	authorized.POST("user/logout", u.Logout)
	authorized.GET("users/get", u.Get)
	authorized.POST("users/update", u.Update)
	authorized.POST("users/delete", u.Delete)
}

func (u userHandler) Login(ctx *gin.Context) {
	uname := ctx.PostForm("username")
	pwd := ctx.PostForm("password")

	if uname == "" || pwd == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := u.service.Login(uname, pwd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		ctx.Status(http.StatusUnauthorized)
		return
	}
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie(sessionCookie, string(user.Session), 315360000,
		"/", DOMAIN, false, true)
	ctx.SetCookie(usernameCookie, user.Username, 315360000,
		"/", DOMAIN, false, true)
	ctx.Status(http.StatusOK)
}

func (u userHandler) Logout(ctx *gin.Context) {
	uany, _ := ctx.Get(userKey)

	user := uany.(core.User)

	err := u.service.Logout(user.Username, user.Session)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (u userHandler) Register(ctx *gin.Context) {
	uname := ctx.PostForm("username")
	email := ctx.PostForm("email")
	pwd := ctx.PostForm("password")

	if uname == "" || email == "" || pwd == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if _, err := u.service.Register(uname, email, pwd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (u userHandler) Get(ctx *gin.Context) {
	id := ctx.Query("id")
	uname := ctx.Query("uname")

	if !u.authorizedUser(ctx, uname) {
		return
	}

	var user *core.User
	var err error

	if id == "" {
		user, err = u.service.GetByUsername(uname)
	} else {
		user, err = u.service.Get(id)
	}

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	user.PasswordHash = nil

	ctx.IndentedJSON(http.StatusOK, user)
}

func (u userHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (u userHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (u userHandler) authorizedUser(ctx *gin.Context, username string) bool {
	user, _ := UserFromCtx(ctx)
	if !(user.Username == username) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return false
	}
	return true
}
