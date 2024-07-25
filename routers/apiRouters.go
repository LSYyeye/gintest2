package routers

import (
	"gintest2/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/aapi")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.Userist)
		apiRouters.GET("/cart", api.ApiController{}.Cart)
	}
}
