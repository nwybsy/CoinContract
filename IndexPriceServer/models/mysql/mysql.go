package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var ErrRecordNotFound = gorm.ErrRecordNotFound

func Init(address, user, password, dbName string) error {
	var err error
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, address, dbName)
	DB, err = gorm.Open("mysql", str)
	if err != nil {
		return err
	}
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(10)
	return nil
}

func NewTransaction() *gorm.DB {
	return DB.Begin()
}

func TransactionRollback(dbConn *gorm.DB) error {
	return dbConn.Rollback().Error
}

func TransactionCommit(dbConn *gorm.DB) error {
	return dbConn.Commit().Error
}
