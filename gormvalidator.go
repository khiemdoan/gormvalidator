package gormvalidator

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var _validator *validator.Validate

func RegisterCallbacks(db *gorm.DB) {
	_validator = validator.New()

	callbackName := "gormvalidations:validate"
	createCallback := db.Callback().Create()
	if createCallback.Get(callbackName) == nil {
		createCallback.After("gorm:before_create").Before("gorm:create").Register(callbackName, validate)
	}
	updateCallback := db.Callback().Update()
	if updateCallback.Get(callbackName) == nil {
		updateCallback.After("gorm:before_update").Before("gorm:update").Register(callbackName, validate)
	}
}
