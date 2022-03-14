package gormvalidator

import "gorm.io/gorm"

type ValidateInterface interface {
	Validate(*gorm.DB) error
}
