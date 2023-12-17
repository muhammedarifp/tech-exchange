package domain

import "time"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "Pending"
	PaymentStatusSuccess   PaymentStatus = "Success"
	PaymentStatusFailed    PaymentStatus = "Failed"
	PaymentStatusCancelled PaymentStatus = "Cancelled"
)

// Defign databse models

type RazorpayAccount struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null;unique"`
	RazorpayID string `gorm:"not null"`
	Email      string `gorm:"email"`
}

type Plans struct {
	ID          uint    `gorm:"primaryKey"`
	PlanID      string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name        string  `gorm:"type:varchar(255)"`
	Description string  `gorm:"type:varchar(255)"`
	Interval    int     `gorm:"not null"`
	Period      string  `gorm:"type:varchar(50)"`
	Amount      float64 `gorm:"not null"`
	IsActive    bool    `gorm:"not null; default:true"`
}

type Subscription struct {
	ID              uint      `gorm:"primaryKey"`
	SubscriptionID  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	CustomerID      string    `gorm:"not null"`
	PlanID          string    `gorm:"type:varchar(100);not null"`
	Status          string    `gorm:"type:varchar(50);not null"`
	StartingDate    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	NextBillingDate time.Time `gorm:"not null"`
}

type Payment struct {
	ID             uint      `gorm:"primaryKey"`
	SubscriptionID uint      `gorm:"not null"`
	Amount         float64   `gorm:"not null"`
	PaymentDate    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status         string    `gorm:"index; not null"` // Payment status
}
