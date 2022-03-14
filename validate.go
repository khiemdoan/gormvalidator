package gormvalidator

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func validate(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil || db.Statement.SkipHooks {
		return
	}
	if !db.Statement.Schema.BeforeSave && !db.Statement.Schema.BeforeCreate {
		return
	}

	_validator := validator.New()

	callMethod(db, func(value interface{}, tx *gorm.DB) (called bool) {
		db.AddError(_validator.Struct(value))

		if i, ok := value.(ValidateInterface); ok {
			called = true
			db.AddError(i.Validate(tx))
		}

		return called
	})
}
