package web

import (
	srv "bk.myprogramming.top/server"
	"bk.myprogramming.top/web/blogCfg"
	"bk.myprogramming.top/web/blogDir"
	"bk.myprogramming.top/web/test"

	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine) {
	superAuth := srv.BearerAuth(srv.SuperUserType)
	e.POST("/login/super", superAuth.LoginHandler)
	e.POST("/refresh/super", superAuth.RefreshHandler)
	// e.POST("/logout/super", superAuth.LogoutHandler)
	super := e.Group("/super")
	super.Use(superAuth.MiddlewareFunc())
	{
		super.GET("/test", test.API)

		super.GET("/dir/list", blogDir.List)
		super.POST("/dir/add", blogDir.Add)
		super.POST("/dir/remove", blogDir.Remove)
		super.POST("/dir/modify", blogDir.Modify)
		super.POST("/dir/solid", blogDir.Solid)

		super.GET("/cfg/list", blogCfg.List)
		super.POST("/cfg/replace", blogCfg.Replace)
		super.POST("/cfg/remove", blogCfg.Remove)
	}
}
