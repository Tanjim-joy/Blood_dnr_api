package models

import (
	"time"

	"gorm.io/gorm"
)

type Donor struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(255);not null;index:idx_name" json:"name" binding:"required"`
	Phone        string     `gorm:"type:varchar(255);not null;index:idx_phone,unique" json:"phone" binding:"required"`
	Email        string     `gorm:"type:varchar(255);index:idx_email" json:"email"`
	BloodGroup   string     `gorm:"type:varchar(10);not null;index:idx_blood_group,length:10" json:"blood_group" binding:"required"`
	Age          int        `gorm:"type:int" json:"age"`
	Gender       string     `gorm:"type:varchar(20)" json:"gender"`
	Area         string     `gorm:"type:varchar(255);not null;index:idx_area,length:255" json:"area" binding:"required"`
	City         string     `gorm:"type:varchar(255);index:idx_city,length:255" json:"city"`
	LastDonation *time.Time `gorm:"type:datetime(3);index:idx_last_donation" json:"last_donation"`
	Available    bool       `gorm:"default:true;index:idx_available" json:"available"`
	Notes        string     `gorm:"type:text" json:"notes"`
}
