package database

import (
	"fmt"

	"github.com/sherlockhua/koala/config"
	logs "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectToMySQL creates a connection to the MySQL database.
func ConnectToMySQL(config *config.Config) (*gorm.DB, error) {
	//dsn := os.Getenv("MYSQL_DSN") // Example: "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port,
		config.MySQL.DBName, config.MySQL.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Errorf("connect to mysql failed, err:%v", err)
		return nil, err
	}

	logs.Infof("Successfully connected to MySQL!")
	return db, nil
}
