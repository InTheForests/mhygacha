package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"mhygacha/project"
	"mhygacha/sr"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("  FFFFFF   OOOOO   RRRRRR    EEEEEEE  SSSSSS   TTTTTTT ")
	fmt.Println("  FF      OO   OO  RR   RR   EE       SS         TT   ")
	fmt.Println("  FFFFF   OO   OO  RRRRRR    EEEEE     SSSSS     TT   ")
	fmt.Println("  FF      OO   OO  RR  RR    EE            SS    TT   ")
	fmt.Println("  FF      OO   OO  RR   RR   EE       S    SS    TT   ")
	fmt.Println("  FF       OOOO0   RR    RR  EEEEEEE  SSSSSS     TT   ")
	fmt.Println("-----------------------------------------------------------")
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
	certData, _ := staticFiles.ReadFile("static/private.crt")
	keyData, _ := staticFiles.ReadFile("static/private.key")
	cert, _ := tls.X509KeyPair(certData, keyData)
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	}
	go func() {
		log.Println("开始监听443端口")

		// 使用你的tls.Config对象来启动服务器
		server := &http.Server{
			Addr:      ":443",
			Handler:   router,
			TLSConfig: cfg,
		}

		err := server.ListenAndServeTLS("", "")
		if err != nil {
			log.Printf("443端口监听失败：%q", err)
		}
	}()

	select {}
}
