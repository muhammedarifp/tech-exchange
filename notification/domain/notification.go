package domain

import "time"

type Notifications struct {
	ID           uint      `gorm:"id primaryKey"`
	UserID       uint      `gorm:"column:user_id; NOT NULL"`
	CreateAt     time.Time `gorm:"default:CURRENT_TIMESTAMP; NOT NULL"`
	Title        string    `gorm:"column:title"`
	Body         string    `gorm:"column:body"`
	Is_importent bool      `gorm:"default:false"`
}
