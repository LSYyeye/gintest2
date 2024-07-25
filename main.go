package main

import (
	"fmt"
	"gintest2/models"
	"gintest2/routers"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 中间件
func initMiddlewareOne(c *gin.Context) {
	fmt.Println("1-我是一个中间件--initMiddlewareOne")
	//调用该请求的剩余处理程序
	//可以用c.Next来统计程序执行的时间
	c.Next()
	fmt.Println("2-我是一个中间件--initMiddlewareOne")
}
func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("1-我是一个中间件--initMiddlewareTwo")
	//调用该请求的剩余处理程序
	//可以用c.Next来统计程序执行的时间
	c.Next()
	fmt.Println("2-我是一个中间件--initMiddlewareTwo")
}
func main() {
	r := gin.Default()
	//自定义模板函数  要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		//注册模板函数
		"UnixToTime": models.UnixToTime,
	})
	//加载模板：html必须要有的，紧挨着路由
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录  第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")
	//配置session中间件
	//创建基于cookie的存储引擎，secret111参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRouterInit(r)
	//全局中间件
	// r.Use(initMiddlewareOne, initMiddlewareTwo)

	r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
		fmt.Println("这是一个首页")
		c.String(200, "gin首页")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "gin新闻首页")
	})
	r.GET("/login", func(c *gin.Context) {
		c.String(200, "gin-login首页")
	})
	r.Run(":80")
}
