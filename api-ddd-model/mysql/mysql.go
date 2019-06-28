package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetConnectionDB() (db *gorm.DB, err error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"user",
		"user",
		"message.com",
		"3306",
		"message",
	))
}
