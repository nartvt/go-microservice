package model

import "time"

type Promotion struct {
	Id                     int
	Code                   string
	PromotionType          string
	Value                  int
	CreatedBy              int
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
	DeletedAt              *time.Time
	ActiveFrom             *time.Time
	ActiveTo               *time.Time
	DailyActiveFrom        int
	DailyActiveTo          int
	MaxActiveTime          int
	MaxDailyActiveTime     int
	PerUserActiveTime      int
	PerUserDailyActiveTime int
	Active                 bool
}
