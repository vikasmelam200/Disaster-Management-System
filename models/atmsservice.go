package models

import "gorm.io/gorm"

type AtmsService struct {
	gorm.Model
	ID             uint   `json:"id" gorm:"primaryKey"`
	Beneficiary    string `json:"beneficiary"`
	ServiceDetails string `json:"service_details"`
	DeliveryMedia  string `json:"delivery_media"`
}
