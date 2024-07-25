package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "张三")

	cCp := c.Copy()
	//定义一个goroutine统计日志
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
