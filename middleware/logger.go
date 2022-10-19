package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("example","testMiddlware")

		htmlType, _ := context.GetQuery("type")
		if htmlType == ""{
			log.Println("type不存在")
		}else{
			log.Println("type为"+htmlType)
		}


		// 在这里可以处理请求返回给用户之前的逻辑
		latency := time.Since(t)
		log.Print(latency)

		//查询请求状态码
		status := context.Writer.Status()
		log.Println(status)
		context.Next()
	}
}