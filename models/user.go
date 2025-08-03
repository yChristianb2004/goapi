package models

import "time"

type User struct {
    ID              uint      `gorm:"primaryKey"`
    Name            string    `gorm:"not null"`
    Email           string    `gorm:"unique;not null"`
    Password        string
    Role            string    `gorm:"not null"`
    IsEmailVerified bool      `gorm:"default:false"`
    CreatedAt       time.Time
}
