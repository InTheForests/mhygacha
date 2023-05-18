package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"mhygacha/project"
	"mhygacha/sr"
)

func main() {
	//初始化数据库
	project.SQLiteInit()

	//关闭gin debug
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 允许所有的跨域请求
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://webstatic.mihoyo.com"} // 指定允许的域名
	config.AllowCredentials = true

	router.Use(cors.New(config))

	//sr抽卡记录
	router.GET("/common/gacha_record/api/getGachaLog", sr.GaChaLog)

	//http服务器
	go func() {
		log.Println("开始监听80端口")
		err := router.Run(":80")
		if err != nil {
			log.Printf("80端口监听失败：%q", err)
		}
	}()
	// https服务器
	go func() {
		log.Println("开始监听443端口")
		err := router.RunTLS(":443", "./static/private.crt", "./static/private.key")
		if err != nil {
			log.Printf("443端口监听失败：%q", err)
		}
	}()

	select {}
}
