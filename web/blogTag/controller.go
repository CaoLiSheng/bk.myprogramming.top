package blogTag

import (
	"net/http"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		data := make(BlogTags, 0)
		data.ListBlogTags(c)
		return &srv.Result{Code: http.StatusOK, Results: data}
	})
}

func Add(c *gin.Context) {
	row := new(AddBlogTag)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.AddBlogTag(c)
		return &srv.Result{Code: http.StatusOK, Results: row}
	})
}

func Remove(c *gin.Context) {
	row := new(BlogTagID)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.RemoveBlogTag(c)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Modify(c *gin.Context) {
	row := new(BlogTag)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.ModifyBlogTag(c)
		return &srv.Result{Code: http.StatusOK}
	})
}
