package domain

import (
	"time"
)

// Db Structures
type Users struct {
	ID          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP; NOT NULL"`
	Username    string    `gorm:"NOT NULL; unique"`
	Email       string    `gorm:"index; NOT NULL; unique"`
	Password    string    `gorm:"NOT NULL"`
	Is_admin    bool      `gorm:"NOT NULL; default:false"`
	Is_Premium  bool      `gorm:"NOT NULL; default:false"`
	Is_active   bool      `gorm:"NOT NULL; default:true"`
	Is_banned   bool      `gorm:"NOT NULL; default:false"`
	Is_verified bool      `gorm:"NOT NULL; default:false"`
}

type Profiles struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"foreignKey:UserID"`
	Name        string `gorm:""`
	Profile_img string `gorm:"default:''"`
	Bio         string `gorm:"default:''"`
	City        string `gorm:"default:''"`
	Github      string `gorm:"default:''"`
	Linkedin    string `gorm:"default:''"`
}

type Badges struct {
	ID        uint   `gorm:"primaryKey"`
	BadgeImg  uint   `gorm:"foreignKey:UserID; NOT NULL"`
	Name      uint   `gorm:"NOT NULL"`
	ShortDisc string `gorm:"NOT NULL"`
}

type UserBadges struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"foreignKey:UserID"`
	BadgeID uint `gorm:"foreignKey:BadgeID"`
}

type Skills struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"foreignKey:UserID"`
	Skill  string
	Level  int `gorm:"validate:min=0,max=10"`
}

type Followers struct {
	ID         uint      `gorm:"primaryKey"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP; NOT NULL"`
	UserID     uint      `gorm:"foreignKey:UserID"`
	FollowerID uint      `gorm:"foreignKey:UserID"`
}
