package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/action"
	"main.go/netSender"
)

func IndexRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		netSender.Ctx.CmdGetInfo()
		MaskBtn := []byte{}
		MaskCtrl := []byte{}
		action.Mask.Button.Range(func(key, value interface{}) bool {
			MaskBtn = append(MaskBtn, key.(byte))
			return true
		})

		action.Mask.Ctrl.Range(func(key, value interface{}) bool {
			MaskCtrl = append(MaskCtrl, key.(byte))
			return true
		})

		context.String(200, string(MaskBtn)+string(MaskCtrl))
	})

}
