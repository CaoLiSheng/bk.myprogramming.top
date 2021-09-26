package blogPost

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
		data := make(BlogPostRows, 0)
		data.ListBlogPost(c)
		return &srv.Result{Code: http.StatusOK, Results: data}
	})
}

func Add(c *gin.Context) {
	row := new(AddBlogPostRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.AddBlogPost(c)
		return &srv.Result{Code: http.StatusOK, Results: row}
	})
}

func Modify(c *gin.Context) {
	row := new(ModifyBlogPostRow)
	if srv.BadRequest(c, row) {return}
	
	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.ModifyBlogPost(c)
		return &srv.Result{Code: http.StatusOK, Results: row}
	})
}

func BindTag(c *gin.Context) {
	row := new(BlogPostTagRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.BindPostTag(c)
		return &srv.Result{Code: http.StatusOK}
	})
}

func UnbindTag(c *gin.Context) {
	row := new(BlogPostTagRow)
	if srv.BadRequest(c, row) {return}

	srv.Do(c, srv.NewJobOpts(true, true), func(c *db.Core) *srv.Result {
		row.UnbindPostTag(c)
		return &srv.Result{Code: http.StatusOK}
	})
}

func Save(c *gin.Context) {
	form := new(BlogPostSaveForm)
	if srv.BadRequest(c, form) {return}

	srv.Do(c, srv.NewJobOpts(false, true), func(c *db.Core) *srv.Result {
		path := form.RefreshAndGetPath(c)
		blogCfgDict := blogCfg.BlogCfgDictIDs{"post_base"}.GetBlogCfg(c)
		filePath := blogCfgDict["post_base"] + "/" + path
		
		f, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		srv.IsPanic(err)

		_, err = f.WriteString(form.Content)
		srv.IsPanic(err)

		return &srv.Result{Code: http.StatusOK}
	})
}
