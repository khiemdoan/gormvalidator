package gormvalidator

import (
	"reflect"

	"gorm.io/gorm"
)

func callMethod(db *gorm.DB, fc func(value interface{}, tx *gorm.DB) bool) {
	tx := db.Session(&gorm.Session{NewDB: true})

	value := db.Statement.ReflectValue
	if called := fc(value.Interface(), tx); !called {
		switch value.Kind() {
		case reflect.Slice, reflect.Array:
			db.Statement.CurDestIndex = 0
			for i := 0; i < value.Len(); i++ {
				fc(reflect.Indirect(value.Index(i)).Addr().Interface(), tx)
				db.Statement.CurDestIndex++
			}
		case reflect.Struct:
			fc(value.Addr().Interface(), tx)
		}
	}
}
