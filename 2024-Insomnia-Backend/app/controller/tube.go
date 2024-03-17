package controller

import (
	"Insomnia/app/common/tube"
	"Insomnia/app/response"
	"github.com/gin-gonic/gin"
)

type Tube struct{}
type Token struct {
	token string
}

func (t *Tube) GetQNToken(c *gin.Context) {
	response.OkData(c, Token{tube.GetQNToken()})
	return
}
