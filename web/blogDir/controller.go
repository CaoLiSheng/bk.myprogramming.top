package blogDir

import (
	"net/http"
	"os"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"bk.myprogramming.top/web/blogCfg"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		data := make(BlogDirRows, 0)
		data.ListBlogDir(c)
		return &srv.Result{Code: http.StatusOK, Results: data}
	})
}

func Add(c *gin.Context) {
	row := new(AddBlogDirRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.AddBlogDir(c)
		return &srv.Result{Code: http.StatusOK, Results: row}
	})
}

func Remove(c *gin.Context) {
	row := new(BlogDirRowId)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.RemoveBlogDir(c)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Modify(c *gin.Context) {
	row := new(BlogDirRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.ModifyBlogDir(c)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Solid(c *gin.Context) {
	row := new(BlogDirRowId)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		path := row.GetBlogDirPath(c)
		blogCfgDict := blogCfg.BlogCfgDictIDs{"post_base"}.GetBlogCfg(c)
		srv.IsPanic(os.MkdirAll(blogCfgDict["post_base"] + path, os.ModePerm))

		return &srv.Result{Code: http.StatusOK, Results: path}
	})
}
