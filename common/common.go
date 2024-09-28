package common

import (
	"fmt"
	"main.go/tuuz/Redis"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func PrintRedis(str string, promts ...any) {
	ppp := fmt.Sprint(promts)
	var ps Redis.PubSub
	ps.Publish(str, ppp)
	//var rs Redis.Stream
	//rs.New("knet").SetMaxLen(100).Publish(map[string]any{
	//	str: ppp,
	//})
	//fmt.Println(str, promts)
}
