package test

import (
	srv "bk.myprogramming.top/server"

	"github.com/gin-gonic/gin"
)

// API :
func API(c *gin.Context) {
	req := new(req)
	if srv.BadRequest(c, req) {
		return
	}

	srv.Do(c, srv.NewJobOpts(true, true), handler(req))
}

