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
		action.Mask.ButtonMu.RLock()
		for key := range action.Mask.Button {
			MaskBtn = append(MaskBtn, key)
		}
		action.Mask.ButtonMu.RUnlock()

		action.Mask.CtrlMu.RLock()
		for key := range action.Mask.Ctrl {
			MaskCtrl = append(MaskCtrl, key)
		}
		action.Mask.CtrlMu.RUnlock()

		context.String(200, string(MaskBtn)+string(MaskCtrl))
	})

}