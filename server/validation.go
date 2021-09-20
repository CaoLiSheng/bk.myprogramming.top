package srv

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, obj interface{}) bool {
	err := c.ShouldBind(obj)
	if err != nil {
		result := new(Result)
		result.Code = http.StatusBadRequest
		result.Err = errors.New("请求参数错误【" + err.Error() + "】")
		result.Send(c)
		return true
	}
	return false
}

func SendBadRequest(c *gin.Context, err error) {
	result := new(Result)
	result.Code = http.StatusBadRequest
	result.Err = errors.New("请求参数错误【" + err.Error() + "】")
	result.Send(c)
}

func IsPanic(err error) {
	if err != nil {
		panic(err)
	}
}
