package model

import (
	"time"
)

type ActivityDO struct{
	ID        int64     `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Name      string    `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:用户名"`
	Price     float32   `gorm:"column:price;type:float;index;not null;default:0;comment:价格"`
	Location      string    `gorm:"column:location;type:varchar(50);index;not null;default:'';comment:用户"`
	SmallImg   string    `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:小图"`
	BigImg     string     `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:大图"`
	Description     string  `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:描述"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

