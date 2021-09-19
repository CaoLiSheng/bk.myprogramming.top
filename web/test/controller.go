package test

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"

	"github.com/gin-gonic/gin"
)

// API :
func API(c *gin.Context) {
	req := new(req)
	if srv.BadRequest(c, req) {
		return
	}

	srv.Do(c, db.NewJobOpts(true, true), handler(req))
}

