package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	v1 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		wsUpgrader := &websocket.Upgrader{}
		ws, err := wsUpgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			return
		}
		defer ws.Close()
		for {
			messageType, p, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			switch messageType {
			case websocket.TextMessage:
				fmt.Printf("处理文本消息, %s\n", string(p))
				ws.WriteMessage(websocket.TextMessage, p)
				// c.Writer.Write(p)
			case websocket.BinaryMessage:
				fmt.Println("处理二进制消息")
			case websocket.CloseMessage:
				fmt.Println("关闭websocket连接")
				return
			case websocket.PingMessage:
				fmt.Println("处理ping消息")
				ws.WriteMessage(websocket.PongMessage, []byte("ping"))
			case websocket.PongMessage:
				fmt.Println("处理pong消息")
				ws.WriteMessage(websocket.PongMessage, []byte("pong"))
			default:
				fmt.Printf("未知消息类型: %d\n", messageType)
				return
			}
		}

		context.String(0, router.BasePath())
	})
	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}

	}
}
