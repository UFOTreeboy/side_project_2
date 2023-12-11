package main

import "github.com/gin-gonic/gin"

func main() {
    // 初始化 Gin 引擎
    r := gin.Default()

    // 定義路由處理器
    r.GET("/", func(c *gin.Context) {
        // 返回 "Hello" 字串
        c.String(200, "Hello!肥宅們!")
    })

    // 啟動伺服器
    r.Run(":8080")
}