package model

const (
	essayTableName       = "essay"
	swiperImageTableName = "swiper_image"
)

type Image struct {
	ID    int    `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Title string `gorm:"column:title;type:string;"`
	URL   string `gorm:"column:url;type:string;"`
}

func (Image) TableName() string {
	return swiperImageTableName
}

type EssayDO struct {
	ID    int    `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Uuid  string `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:文章编号"`
	Title string `gorm:"column:title;type:varchar(128);not null;default:'';comment:文章标题"`
}

func (EssayDO) TableName() string {
	return essayTableName
}
