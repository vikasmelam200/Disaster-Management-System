package models

import (
	"gorm.io/gorm"
)

type Disaster struct {
	gorm.Model
	Type              string `json:"type"`
	Description       string `json:"description"`
	Location          string `json:"location"`
	Severity          string `json:"severity"`
	HighWayImpact     bool   `json:"highway_impact"`
	EnvironmentImpact bool   `json:"environment_impact"`
	LifeImpact        bool   `json:"life_impact"`
}
