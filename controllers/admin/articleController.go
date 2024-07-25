package admin

import (
	"gintest2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleTest struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (ArticleTest) TableName() string {
	return "article"
}

type ArticleController struct {
}

func (con ArticleController) Index(c *gin.Context) {
	//查询全部数据
	// articleList := []models.Article{}
	// models.DB.Find(&articleList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": articleList,
	// })
	//查询单个数据
	// articleResult := models.Article{Id: 2}
	// models.DB.Find(&articleResult)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": articleResult,
	// })
	//查询id>2的数据
	// articleList := []models.Article{}
	// models.DB.Where("id > ?", 2).Find((&articleList))
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": articleList,
	// })
	//查询id>3并且id<7的数据
	//不包含3和7
	// articleList := []models.Article{}
	// models.DB.Where("id>3 AND id<7").Find((&articleList))
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": articleList,
	// })
	//包含3和7
	// articleList := []models.Article{}
	// models.DB.Where("id between 3 and 7").Find((&articleList))
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": articleList,
	// })
	//查询id=2或者id=3的数据
	// articleList := []models.Article{}
	// // models.DB.Where("id = ? or id = ?", 2, 3).Find(&articleList)
	// models.DB.Where("id = 2").Or("id = 3").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//查询id在3,5,6的数据
	// articleList := []models.Article{}
	// models.DB.Where("id in (?)", []int{3, 5, 6}).Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//使用like查询标题中包含 1 的内容
	// articleList := []models.Article{}
	// models.DB.Where("title like ?", "%1%").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//使用select返回部分字段,其余部分还会出现但是值为空
	// articleList := []models.Article{}
	// models.DB.Select("id,title").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//要想没有其余部分需要重新建立结构体
	// articleList := []ArticleTest{}
	// models.DB.Select("id,title").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//Order排序(可以写多个)
	// articleList := []models.Article{}
	// models.DB.Order("id desc").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//limit:在order基础上只获得前两条数据
	// articleList := []models.Article{}
	// models.DB.Order("id desc").Limit(2).Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//通过offset和limit实现分页
	// articleList := []models.Article{}
	// models.DB.Order("id desc").Offset(2).Limit(2).Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	//count统计总数
	// articleList := []models.Article{}
	// var num int64
	// models.DB.Find(&articleList).Count(&num)
	// c.JSON(200, gin.H{
	// 	"result": num,
	// })
	//使用原生sql统计表的数量
	// var num int
	// models.DB.Raw("select count(*) from article").Scan(&num)
	// c.JSON(200, gin.H{
	// 	"result": num,
	// })
	//使用原生sql查询id= 2的数据
	// articleList := []models.Article{}
	// models.DB.Raw("select * from article where id = 2").Scan(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	// articleList := []models.Article{}
	// models.DB.Raw("select * from article ").Scan(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })

	//查询文章获取文章对应的分类
	articleCateList := []models.ArticleCate{}
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(200, gin.H{
		"result": articleCateList,
	})

}
func (con ArticleController) Add(c *gin.Context) {
	article := models.Article{
		Id:     2,
		Title:  "新闻222",
		CateId: 3,
		Email:  "222@qq.com",
	}
	models.DB.Create(&article)

	c.String(http.StatusOK, "添加新闻列表--add")
}
func (con ArticleController) Edit(c *gin.Context) {

	// article := models.Article{}
	// models.DB.Where("id = ?", 1).Find(&article)
	// article.Email = "111@qq.com"
	// models.DB.Save(&article)
	//使用原生sql修改article表中的一条数据
	models.DB.Exec("update article set email = ? where id = 2", "123@qq.com")

	c.String(http.StatusOK, "修改新闻列表--edit--------")
}

func (con ArticleController) Delete(c *gin.Context) {
	// article := models.Article{}
	// models.DB.Where("id = ?", 3).Delete(&article)
	//使用原生sql删除article表中的一条数据
	models.DB.Exec("delete from article where id = ?", 5)
	c.String(200, "删除成功")
}
