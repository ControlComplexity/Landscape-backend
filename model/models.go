package model

type Image struct {
	ID    int    `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Title string `gorm:"column:title;type:string;"`
	URL   string `gorm:"column:url;type:string;"`
}

func (Image) TableName() string {
	return "swiper_image"
}
