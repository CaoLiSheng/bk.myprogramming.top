package blogCfg

import (
	"net/http"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	srv.Do(c, db.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		return &srv.Result{Code: http.StatusOK, Results: listBlogCfg(c)}
	})
}

