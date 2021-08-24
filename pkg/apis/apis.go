package apis

import (
	"github.com/Evolt0/gw-kit/pkg/apis/base"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	base.Routes(engine)
}
