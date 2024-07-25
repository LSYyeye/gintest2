package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	//username是空接口类型
	username, _ := c.Get("username")
	fmt.Println(username)
	//将空接口类型转换为string类型
	// v, _ := username.(string)
	// c.String(200, "后台用户列表---"+v)
	//类型断言
	v, ok := username.(string)
	if ok == true {
		c.String(200, "后台用户列表---"+v)
	} else {
		c.String(200, "后台用户列表---用户数据获取失败")
	}
}
