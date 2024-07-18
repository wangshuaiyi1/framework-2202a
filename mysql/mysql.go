package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Client(port int64, dbname string, username string, password string, host string, handler func(db *gorm.DB) error) error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db, err := cli.DB()
	if err != nil {
		return err
	}

	defer db.Close()

	return handler(cli)
}
