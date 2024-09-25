package models

import "gorm.io/gorm"

type AtmsService struct {
	gorm.Model
	ID             uint   `json:"id" gorm:"primaryKey"`
	Beneficiary    string `json:"beneficiary"`
	ServiceDetails string `json:"service_details"`
	DeliveryMedia  string `json:"delivery_media"`
}
type AtmsEnhancement struct {
	gorm.Model
	ID              uint   `json:"id" gorm:"primarykey"`
	Facility        string `json:"facility"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	ControlCenterID uint   `json:"control_center_id"`
}
