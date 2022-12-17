package models

type Staff struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(50)" json:"name"`
	Role    string `gorm:"type:varchar(50)" json:"role"`
	Address string `gorm:"type:varchar(50)" json:"address"`
}
