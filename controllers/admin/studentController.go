package admin

import (
	"gintest2/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct{}

func (con StudentController) Index(c *gin.Context) {
	//1.获取所有学生信息
	// studentList := []models.Student{}
	// models.DB.Find(&studentList)
	// c.JSON(200, gin.H{
	// 	"result": studentList,
	// })
	//2.获取所有课程信息
	// lessonList := []models.Lesson{}
	// models.DB.Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })
	//3.查询学生信息的时候 显示学生选修的课程
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Find(&studentList)
	// c.JSON(200, gin.H{
	// 	"result": studentList,
	// })
	//4.查询张三选修了什么课程
	// studentList := []models.Student{}
	// models.DB.Where("name = ?", "张三").Preload("Lesson").Find(&studentList)
	// c.JSON(200, gin.H{
	// 	"result": studentList,
	// })
	//5.查询课程被那些学生选修了
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student").Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })
	//6.计算机网络被那些学生选修了
	// lessonList := []models.Lesson{}
	// models.DB.Where("name = ? ", "计算机网络").Preload("Student").Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })
	//7.查询数据指定条件
	// lessonList := []models.Lesson{}
	// models.DB.Limit(2).Offset(2).Preload("Student").Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })
	//8.查询课程被哪些学生选修的时候去掉张三
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "name != ?", "张三").Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })
	//9.查询课程被哪些学生选修的时候去掉张三和李四
	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "name not in(?,?)", "张三", "李四").Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })
	//10.查看课程被那些学生选修 要求：学生id倒序输出   自定义加载sql
	lessonList := []models.Lesson{}
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return models.DB.Order("student.id DESC")
	}).Find(&lessonList)
	c.JSON(200, gin.H{
		"result": lessonList,
	})

}
