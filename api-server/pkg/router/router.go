// Save this file in ./platform/router/router.go

package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"ds/collector/pkg/auth/authenticator"
	"ds/collector/pkg/auth/callback"
	"ds/collector/pkg/auth/login"
	"ds/collector/pkg/auth/logout"
	"ds/collector/pkg/auth/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/oidc-callback", callback.Handler(auth))
	router.GET("/user", user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
