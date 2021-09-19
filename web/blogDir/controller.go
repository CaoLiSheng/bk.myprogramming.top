package blogDir

import (
	"net/http"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
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
	row := new(removeBlogDirRow)
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

}
