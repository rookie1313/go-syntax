package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// 多种响应方式
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
	})
	// 2. struct响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name   string
			Text   string
			Number int
		}
		msg.Name = "root"
		msg.Text = "text"
		msg.Number = 1
		c.JSON(http.StatusOK, msg)
	})
	//test
	r.GET("/Test", func(c *gin.Context) {
		var msg struct {
			T string
		}
		msg.T = "something"
		c.JSON(http.StatusOK, gin.H{"Literal": msg})
	})
	// 3.XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "abc"})
	})
	// 4.YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name": "jack"})
	})
	// 5.protobuf格式,谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
