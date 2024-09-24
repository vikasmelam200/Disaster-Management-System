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
type DisasterInfo struct {
	gorm.Model
	ID                uint   `json:"id" gorm:"primaryKey"`
	HighWayStatus     string `json:"highway_status"`
	FloodingInfo      string `json:"flooding_info"`
	TrafficStatus     string `json:"traffic_status"`
	TravelTimes       string `json:"travel_times"`
	RideQuality       string `json:"ride_quality"`
	WeatherConditions string `json:"weather_conditions"`
	AlternateRoads    string `json:"alternate_roads"`
	LawAndOrderIssues string `json:"law_and_order_issues"`
	Roadstatus        string `json:"road_status"`
	BestRoutes        string `json:"best_routes"`
	AffectedArea      string `json:"affected_area"`
}
