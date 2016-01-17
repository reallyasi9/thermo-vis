package main

import (
	"time"
)

type structure struct {
	StructureID         string           `json:"structure_id"`
	Thermostats         []string         `json:"thermostats"`
	SmokeCOAlarms       []string         `json:"smoke_co_alarms"`
	Cameras             []string         `json:"cameras"`
	Away                string           `json:"away"`
	Name                string           `json:"name"`
	CountryCode         string           `json:"country_code"`
	PostalCode          string           `json:"postal_code"`
	PeakPeriodStartTime time.Time        `json:"peak_period_start_time"`
	PeakPeriodEndTime   time.Time        `json:"peak_period_end_time"`
	TimeZone            string           `json:"time_zone"`
	RHREnrollment       bool             `json:"rhr_enrollment"`
	Wheres              map[string]where `json:"wheres"`
}

type where struct {
	WhereID string `json:"where_id"`
	Name    string `json:"name"`
}

type devices struct {
	Thermostats   map[string]thermostat   `json:"thermostats"`
	SmokeCOAlarms map[string]smokeCOAlarm `json:"smoke_co_alarms"`
	Cameras       map[string]camera       `json:"cameras"`
}

type thermostat struct {
	DeviceID               string    `json:"device_id"`
	Locale                 string    `json:"locale"`
	SoftwareVersion        string    `json:"software_version"`
	StructureID            string    `json:"structure_id"`
	Name                   string    `json:"name"`
	NameLong               string    `json:"name_long"`
	LastConnection         time.Time `json:"last_connection"`
	IsOnline               bool      `json:"is_online"`
	CanCool                bool      `json:"can_cool"`
	CanHeat                bool      `json:"can_heat"`
	IsUsingEmergencyHeat   bool      `json:"is_using_emergency_heat"`
	HasFan                 bool      `json:"has_fan"`
	FanTimerActive         bool      `json:"fan_timer_active"`
	FanTimerTimeout        time.Time `json:"fan_timer_timeout"`
	HasLeaf                bool      `json:"has_leaf"`
	TemperatureScale       string    `json:"temperature_scale"`
	TargetTemperatureF     float64   `json:"target_temperature_f"`
	TargetTemperatureC     float64   `json:"target_temperature_c"`
	TargetTemperatureHighF float64   `json:"target_temperature_high_f"`
	TargetTemperatureHighC float64   `json:"target_temperature_high_c"`
	TargetTemperatureLowF  float64   `json:"target_temperature_low_f"`
	TargetTemperatureLowC  float64   `json:"target_temperature_low_c"`
	AwayTemperatureHighF   float64   `json:"away_temperature_high_f"`
	AwayTemperatureHighC   float64   `json:"away_temperature_high_c"`
	AwayTemperatureLowF    float64   `json:"away_temperature_low_f"`
	AwayTemperatureLowC    float64   `json:"away_temperature_low_c"`
	HVACMode               string    `json:"hvac_mode"`
	AmbientTemperatureF    float64   `json:"ambient_temperature_f"`
	AmbientTemperatureC    float64   `json:"ambient_temperature_c"`
	Humidity               float64   `json:"humidity"`
	HVACState              string    `json:"hvac_state"`
	WhereID                string    `json:"where_id"`
}

type smokeCOAlarm struct {
	DeviceID           string    `json:"device_id"`
	Locale             string    `json:"locale"`
	SoftwareVersion    string    `json:"software_version"`
	StructureID        string    `json:"structure_id"`
	Name               string    `json:"name"`
	NameLong           string    `json:"name_long"`
	LastConnection     time.Time `json:"last_connection"`
	IsOnline           bool      `json:"is_online"`
	BatteryHealth      string    `json:"battery_health"`
	COAlarmState       string    `json:"co_alarm_state"`
	SmokeAlarmState    string    `json:"smoke_alarm_state"`
	IsManualTestActive bool      `json:"is_manual_test_active"`
	LastManualTestTime time.Time `json:"last_manual_test_time"`
	UIColorState       string    `json:"ui_color_state"`
	WhereID            string    `json:"where_id"`
}

type camera struct {
	DeviceID              string    `json:"device_id"`
	SoftwareVersion       string    `json:"software_version"`
	StructureID           string    `json:"structure_id"`
	WhereID               string    `json:"where_id"`
	Name                  string    `json:"name"`
	NameLong              string    `json:"name_long"`
	IsOnline              bool      `json:"is_online"`
	IsStreaming           bool      `json:"is_streaming"`
	IsAudioInputEnabled   bool      `json:"is_audio_input_enabled"`
	LastIsOnlineChange    time.Time `json:"last_is_online_change"`
	IsVideoHistoryEnabled bool      `json:"is_video_history_enabled"`
	WebURL                string    `json:"web_url"`
	AppURL                string    `json:"app_url"`
	LastEvent             event     `json:"last_event"`
}

type event struct {
	HasSound         bool      `json:"has_sound"`
	HasMotion        bool      `json:"has_motion"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	URLsExpireTime   time.Time `json:"urls_expire_time"`
	WebURL           string    `json:"web_url"`
	AppURL           string    `json:"app_url"`
	ImageURL         string    `json:"image_url"`
	AnimatedImageURL string    `json:"animated_image_url"`
}
