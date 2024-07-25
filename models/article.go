package models

//foreignKey:外键(如果是表名称加上Id默认不配置：ArticleCateId)
//references:主键(默认是Id 可以不配置)
type Article struct {
	Id     int    `json:"id"`
	Title  string `json:"title`
	CateId int    `json:"cateid"` //外键（ArticleCateId的时候可以直接识别，不是就需要重写外键）
	Email  string `json:"email"`
	// ArticleCate ArticleCate `gorm:"foreignKey:CateId"` //重写外键
}

// 表示配置操作数据库的表名称
func (Article) TableName() string {
	return "article"
}
