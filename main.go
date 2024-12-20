package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"main.go/action"
	"main.go/config/app_conf"
	"main.go/netReceiver"
	"main.go/netSender"
	"main.go/netTcp"
	"main.go/route"
	"net"
)

func main() {

	netSender.Ctx.Ready()
	netReceiver.Crx.Ready()

	Act := &action.Action{}
	go Act.MainRun()

	sudp := netTcp.ServerUDP{
		SendServer: &net.UDPAddr{
			IP:   net.ParseIP("10.0.0.90"),
			Port: 6666,
		},
	}
	go sudp.Start()

	stcp := netTcp.ServerTcp{
		SendServer: &net.TCPAddr{
			IP:   net.ParseIP("10.0.0.90"),
			Port: 6666,
		},
	}
	go stcp.Start()

	go netTcp.Start()
	go netTcp.Start()
	go netTcp.Start()
	go netTcp.Start()

	go route.MainWsRouter()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
