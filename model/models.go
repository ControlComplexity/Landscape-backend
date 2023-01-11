package model

type Image struct {
	ID    int `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Title string
	URL   string
}
