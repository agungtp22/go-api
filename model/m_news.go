package model

type News struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(200)" json:"title"`
	Slug  string `gorm:"type:varchar(200)" json:"slug""`
	Body  string `gorm:"type:text" json:"body"`
}
