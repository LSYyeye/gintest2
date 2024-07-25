package routers

import (
	"gintest2/controllers/admin"
	"gintest2/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	//middlewares.Initmiddleware中间件
	adminRouters := r.Group("/aadmin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)
		// adminRouters.GET("/user/add", admin.UserController{}.Add)
		// adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)
		// adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		// adminRouters.POST("/user/doEdit", admin.UserController{}.DoEdit)
		// adminRouters.GET("/user/delete", admin.UserController{}.Edit)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
		adminRouters.GET(("/article/delete"), admin.ArticleController{}.Delete)

		adminRouters.GET("/student", admin.StudentController{}.Index)
		adminRouters.GET("/bank", admin.BankController{}.Index)
	}
}
