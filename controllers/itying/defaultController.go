package itying

import (
	"fmt"
	"gintest2/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(c *gin.Context) {
	//设置sessions
	session := sessions.Default(c)
	session.Set("username", "张三")
	session.Save() //设置session的时候必须调用

	//c.SetCookie(name,value string,maxAge int,path,domain string,secure,httpOnly bool)
	//第一个参数name是key
	//第二个参数value
	//第三个参数maxAge是过期时间(秒)
	//第四个参数path是cookie的路径
	//第五个参数domain是cookie的路径Domain作用域 本地调试配置成localhost，正式上线配置成域名
	//第六个参数secure：值为true时，cookie在http无效，在https有效
	//第七个参数httpOnly：为true只有后端，false前段通过js也可以接触cookie
	//设置cookie
	c.SetCookie("username", "张三", 3600, "/a", "localhost", false, true)
	//过期时间演示
	// c.SetCookie("hobby", "吃饭睡觉", 5, "/a", "localhost", false, true)
	fmt.Println(models.UnixToTime(1720748930))
	c.HTML(http.StatusOK, "default/aindex.html", gin.H{
		"msg": "我是一个msg",
		"t":   1720748930,
	})
}
func (con DefaultController) News(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "cookie="+username)
	// hobby, _ := c.Cookie("hobby")
	// c.String(200, "cookie=%v-----hobby=%v", username, hobby)
}

func (con DefaultController) Shop(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "cookie="+username)
	// hobby, _ := c.Cookie("hobby")
	// c.String(200, "cookie=%v-----hobby=%v", username, hobby)
}
func (con DefaultController) DeleteCookie(c *gin.Context) {
	//删除cookie-----删除整个
	c.SetCookie("username", "张三", -1, "/a", "localhost", false, true)
	c.String(200, "删除成功")
}
func (con DefaultController) session(c *gin.Context) {
	//获取sessions
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "username=%v", username)
}
