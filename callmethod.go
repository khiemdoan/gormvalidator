package gormvalidator

import (
	"reflect"

	"gorm.io/gorm"
)

// This function is copied from [gorm](https://github.com/go-gorm/gorm/blob/master/callbacks/callmethod.go)
func callMethod(db *gorm.DB, fc func(value interface{}, tx *gorm.DB) bool) {
	tx := db.Session(&gorm.Session{NewDB: true})
	if called := fc(db.Statement.ReflectValue.Interface(), tx); !called {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			db.Statement.CurDestIndex = 0
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				fc(reflect.Indirect(db.Statement.ReflectValue.Index(i)).Addr().Interface(), tx)
				db.Statement.CurDestIndex++
			}
		case reflect.Struct:
			fc(db.Statement.ReflectValue.Addr().Interface(), tx)
		}
	}
}
