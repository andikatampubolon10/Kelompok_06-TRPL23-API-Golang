package entity

import "time"

type Users struct {
    Id               uint64    `gorm:"primary_key;autoIncrement" json:"id"`
    Name             string    `gorm:"type:varchar(255);not null" json:"name"`
    Email            string    `gorm:"type:varchar(255);unique;not null" json:"email"`
    EmailVerifiedAt  time.Time `gorm:"type:timestamp" json:"email_verified_at"`
    Password         string    `gorm:"type:varchar(255);not null" json:"-"` 
    RememberToken    string    `gorm:"type:varchar(100)" json:"remember_token"`
    CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}
