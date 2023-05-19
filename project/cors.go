package project

import "github.com/gin-gonic/gin"

func PassCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		//你问我为什么不直接用*,要加几个指定域名?这你得去问mhy为什么要加Access-Control-Allow-Credentials了，*只是在debug使用所以也得加
		allowedOrigins := []string{
			"https://webstatic.mihoyo.com",
			"https://api-takumi.mihoyo.com",
			"http://127.0.0.1",
			"https://127.0.0.1",
		}

		requestOrigin := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == requestOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
