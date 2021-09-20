package srv

import (
	"log"
	"net/http"

	"bk.myprogramming.top/db"

	"github.com/gin-gonic/gin"
)

func (r *Result) Send(c *gin.Context) {
	if r.Err != nil {
		log.Println("error occurred:\n", r)
		r.Message = r.Err.Error()
	}

	c.JSON(r.Code, r)
}

func Do(c *gin.Context, opts *JobOptions, job Job) (result *Result) {
	dbOpts := db.NewJobOpts(func (core *db.Core) {
		result = job(core)
	}, func (err error) {
		result = new(Result)
		result.Code = http.StatusServiceUnavailable
		result.Err = err
	})

	if opts.Simple {
		MustGet(c).DoSimple(dbOpts)
	} else {
		MustGet(c).Do(dbOpts)
	}

	if opts.Auto {
		result.Send(c)
	}

	return
}

func NewJobOpts(simple, auto bool) *JobOptions {
	return &JobOptions{Simple: simple, Auto: auto}
}
