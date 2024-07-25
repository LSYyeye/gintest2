package admin

import (
	"fmt"
	"gintest2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

// func (con UserController) Index(c *gin.Context) {
// 	// c.String(200, "后台用户列表")
// 	con.success(c)
// }

// func (con UserController) Add(c *gin.Context) {
// 	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
// }
// func (con UserController) DoUpload(c *gin.Context) {
// 	//文件夹用年月日。文件名用年月日时分秒
// 	username := c.PostForm("username")

// 	//1.获取上传的文件
// 	file, err := c.FormFile("face")

// 	if err == nil {
// 		//2.获取后缀名 判断类型是否正确   .jpg .png .gif .jpeg
// 		extName := path.Ext(file.Filename)
// 		allowExtMap := map[string]bool{
// 			".jpg":  true,
// 			".png":  true,
// 			".gif":  true,
// 			".jpeg": true,
// 		}
// 		if _, ok := allowExtMap[extName]; !ok {
// 			c.String(200, "上传的图片不合法 ")
// 		}
// 		//3.创建图片保存目录  static/upload/20240712
// 		day := models.GetDay()
// 		dir := "./static/upload/" + day
// 		err := os.MkdirAll(dir, 0666)
// 		if err != nil {
// 			fmt.Println(err)
// 			c.String(200, "MkdirAll失败")
// 			return
// 		}
// 		//4.生成文件名称            时间戳+后缀名： 1111111111111.jpg
// 		unix := models.GetUnix() //转成string类型
// 		fileName := strconv.FormatInt(unix, 10) + extName
// 		//5.文件保存目录dst    执行上传
// 		dst := path.Join(dir, fileName)
// 		c.SaveUploadedFile(file, dst)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success":  true,
// 		"username": username,
// 	})

// 	//同名字的多个文件上传
// 	// username := c.PostForm("username")
// 	// form, _ := c.MultipartForm()
// 	// files := form.File["face[]"]
// 	// for _, file := range files {
// 	// 	dst := path.Join("./static/upload", file.Filename)
// 	// 	//上传文件至指定目录
// 	// 	c.SaveUploadedFile(file, dst)
// 	// }
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"success":  true,
// 	// 	"username": username,
// 	// })

// 	// 单个文件上传
// 	// username := c.PostForm("username")
// 	// file, err := c.FormFile("face")
// 	// //file.Filename  获取文件名称
// 	// dst := path.Join("./static/upload", file.Filename)
// 	// if err == nil {
// 	// 	c.SaveUploadedFile(file, dst)
// 	// }
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"success":  true,
// 	// 	"username": username,
// 	// 	"dst":      dst,
// 	// })
// 	// // c.String(200, "执行上传")
// }
// func (con UserController) Edit(c *gin.Context) {
// 	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
// }
// func (con UserController) DoEdit(c *gin.Context) {
// 	// c.String(200, "执行修改")
// 	username := c.PostForm("username")
// 	file1, err1 := c.FormFile("face1")
// 	file2, err2 := c.FormFile("face2")
// 	//file.Filename  获取文件名称
// 	dst1 := path.Join("./static/upload", file1.Filename)
// 	if err1 == nil {
// 		c.SaveUploadedFile(file1, dst1)
// 	}
// 	dst2 := path.Join("./static/upload", file2.Filename)
// 	if err2 == nil {
// 		c.SaveUploadedFile(file2, dst2)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success":  true,
// 		"username": username,
// 		"dst1":     dst1,
// 		"dst2":     dst2,
// 	})

// }
// 查询
func (con UserController) Index(c *gin.Context) {
	//查询数据库
	userList := []models.User{}
	models.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
	//查询age大于20的用户
	// userList := []models.User{}
	// models.DB.Where("age>30").Find(&userList)
	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })
}

// 增加
func (con UserController) Add(c *gin.Context) {
	user := models.User{
		Username: "itying",
		Age:      33,
		Email:    "2222@QQ.COM",
		// AddTime:  1000000111,
		AddTime: int(models.GetUnix()),
	}
	models.DB.Create(&user)
	fmt.Println(user)

	c.String(200, "数据增加用户成功")
}

// 修改
func (con UserController) Edit(c *gin.Context) {
	// //查询id=4的数据
	// user := models.User{Id: 4}
	// models.DB.Find(&user)
	// //修改id=4的数据
	// user.Email = "123@qq.com"
	// user.AddTime = int(models.GetUnix())
	// models.DB.Save(&user)
	//修改id=5的数据

	//更新单个列
	// user := models.User{}
	// models.DB.Model(&user).Where("id = ?", 6).Update("username", "哈哈哈哈")
	//更新多个列
	user := models.User{}
	models.DB.Where("id = ?", 3).Find(&user)
	user.Username = "丽丽就开始"
	user.AddTime = int(models.GetUnix())
	user.Email = "444@qq.com"
	models.DB.Save(&user)
	c.String(200, "修改用户成功")
}

// 删除
func (con UserController) Delete(c *gin.Context) {
	// user := models.User{Id: 5}
	// models.DB.Delete(&user)

	user := models.User{}
	models.DB.Where("username = ?", "哈哈哈哈").Delete(&user)
	c.String(200, "删除用户")
}
