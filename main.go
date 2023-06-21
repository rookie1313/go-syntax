package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建Gin引擎实例
	router := gin.Default()

	// 定义异步GET请求的路由处理函数
	router.GET("/async", func(c *gin.Context) {
		// 启动一个Goroutine执行异步任务
		go func() {
			//log.Println("异步执行：" + copyContext.Request.URL.String())
			c.JSON(http.StatusOK, "异步执行")
		}()

		// 返回异步任务已启动的消息
		c.JSON(http.StatusOK, "Async task started")
	})

	// 运行Gin引擎
	router.Run(":8080")
}
