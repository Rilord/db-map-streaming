package apis

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
	Context *gin.Context
	Orm     *gorm.DB
	Errors  error
}

func (e *Api) AddErrors(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else {
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}

}

func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

func (e *Api) Bind(d inter)

func Init() *gin.Engine {

}
