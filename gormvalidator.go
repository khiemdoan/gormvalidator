package gormvalidator

import "gorm.io/gorm"

func RegisterCallbacks(db *gorm.DB) {
	callbackName := "gormvalidations:validate"
	createCallback := db.Callback().Create()
	if createCallback.Get(callbackName) == nil {
		createCallback.Before("gorm:before_create").Register(callbackName, validate)
	}
	updateCallback := db.Callback().Update()
	if updateCallback.Get(callbackName) == nil {
		updateCallback.Before("gorm:before_update").Register(callbackName, validate)
	}
}
