package blogCfg

import (
	"net/http"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		data := make(BlogCfgRows, 0)
		data.ListBlogCfg(c)
		return &srv.Result{Code: http.StatusOK, Results: data}
	})
}

func Replace(c *gin.Context) {
	row := new(BlogCfgRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.ReplaceBlogCfg(c)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Remove(c *gin.Context) {
	row := new(BlogCfgRowId)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.RemoveBlogCfg(c)
		return &srv.Result{Code: http.StatusOK}
	})
}
