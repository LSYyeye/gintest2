package models

type ArticleCate struct {
	Id      int       `json:"id"` //主键
	Title   string    `json:"title`
	State   int       `json:"state"`
	Article []Article `gorm:"foreignKey:CateId"`
}

// 表示配置操作数据库的表名称
func (ArticleCate) TableName() string {
	return "article_cate"
}
