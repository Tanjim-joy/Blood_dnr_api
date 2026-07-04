package models

import (
	"gorm.io/gorm"
)

type BloodRequest struct {
	gorm.Model
	PatientName string `gorm:"type:varchar(255);not null;index:idx_patient_name,length:255" json:"patient_name" binding:"required"`
	BloodGroup  string `gorm:"type:varchar(10);not null;index:idx_blood_group_req,length:10" json:"blood_group" binding:"required"`
	Units       int    `gorm:"type:int;not null" json:"units" binding:"required"`
	Hospital    string `gorm:"type:varchar(255);not null;index:idx_hospital,length:255" json:"hospital" binding:"required"`
	Contact     string `gorm:"type:varchar(255);not null;index:idx_contact,length:255" json:"contact" binding:"required"`
	Area        string `gorm:"type:varchar(255);not null;index:idx_area_req,length:255" json:"area" binding:"required"`
	NeededBy    string `gorm:"type:varchar(255)" json:"needed_by"`
	Urgency     string `gorm:"type:varchar(20);default:'normal';index:idx_urgency" json:"urgency"` // normal | urgent | critical
	Status      string `gorm:"type:varchar(20);default:'open';index:idx_status" json:"status"`     // open | fulfilled | closed
	RequestedBy string `gorm:"type:varchar(255);index:idx_requested_by,length:255" json:"requested_by"`
}

// TableName optional
func (BloodRequest) TableName() string {
	return "blood_requests"
}
