package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wPan/v1/Authority"
	"wPan/v1/Cache"
	"wPan/v1/Config"
	"wPan/v1/Models"
	"wPan/v1/Nsq"
	"wPan/v1/Router"
)

func main() {
	gin.SetMode(Config.ServerSetting.RunMode)

	r := gin.Default()

	if err := Models.InitMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	if err := Redis.InitRedis(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	Nsq.InitNSQ() //初始化消息队列

	Authority.InitCasbin()

	Models.DB.AutoMigrate(&Models.User{})

	r = Authority.PolicyHandler(r)

	r = Router.ResRouter(r)

	_ = r.Run()
}
