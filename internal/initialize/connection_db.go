package initialize

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"user-service/global"
	user_models "user-service/internal/module/user/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	var err error
	DB, err = connectionDb()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = performMigration()
	if err != nil {
		log.Fatalf("Could not auto migrate: %v", err)
	}
	return DB
}

func connectionDb() (*gorm.DB, error) {
	user := global.Config.Mysql.User
	password := global.Config.Mysql.Password
	dbName := global.Config.Mysql.DbName
	host := global.Config.Mysql.Host
	port := strconv.Itoa(global.Config.Mysql.Port)
	fmt.Println(user, password, dbName, host, port)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
		Logger: nil,
	})

	if err != nil {
		return nil, err
	}

	sqlDatabase, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	sqlDatabase.SetMaxIdleConns(10)
	sqlDatabase.SetMaxOpenConns(25)
	sqlDatabase.SetConnMaxLifetime(time.Hour)

	return databaseConnection, nil
}

func performMigration() error {
	models := []interface{}{
		&user_models.User{},
	}

	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}
