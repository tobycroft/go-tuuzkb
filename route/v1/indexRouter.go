package v1

import (
	"github.com/gin-gonic/gin"
)

func IndexRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

}
