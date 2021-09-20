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
	srv.Do(c, db.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		return &srv.Result{Code: http.StatusOK, Results: ListBlogDir(c)}
	})
}

func Add(c *gin.Context) {
	row := new(addBlogDirRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, db.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		addBlogDir(c, row)
		return &srv.Result{Code: http.StatusOK, Results: row}
	})
}

func Remove(c *gin.Context) {
	row := new(blogDirRowId)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, db.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		removeBlogDir(c, row)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Modify(c *gin.Context) {
	row := new(blogDirRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, db.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		modifyBlogDir(c, row)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Solid(c *gin.Context) {
	row := new(blogDirRowId)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, db.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		dir := getBlogDir(c, row)
		path := "/" + dir.Name
		for dir.PID != 0 {
			dir = getBlogDir(c, &blogDirRowId{ID:dir.PID})
			path = "/" + dir.Name + path
		}
		postBase := blogCfg.GetBlogCfg(c, "post_base")

		srv.IsPanic(os.MkdirAll(postBase + path, os.ModePerm))

		return &srv.Result{Code: http.StatusOK}
	})
}
