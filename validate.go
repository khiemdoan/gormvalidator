package gormvalidator

import (
	"gorm.io/gorm"
)

func validate(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil || db.Statement.SkipHooks {
		return
	}
	if !db.Statement.Schema.BeforeSave && !db.Statement.Schema.BeforeCreate {
		return
	}

	callMethod(db, func(value interface{}, tx *gorm.DB) (called bool) {
		db.AddError(_validator.Struct(value))
		if i, ok := value.(ValidateInterface); ok {
			db.AddError(i.Validate(tx))
		}
		return true
	})
}
