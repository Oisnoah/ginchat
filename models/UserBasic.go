package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	Email         string
	Password      string
	Phone         string
	Role          string
	ClinicID      uint
	ClientPort    string
	LoginTime     string
	HeartbeatTime string
	LogoutTime    string
	IsLogout      bool
	DeviceInfo    string
	DeviceType    string
	IsOnline      bool
}

func (UserBasic) TableName() string {
	return "user_basic"
}
