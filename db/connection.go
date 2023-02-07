package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/xavimoreno548/BookStore/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	once sync.Once
)

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_HOST")
)

var DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)  //Fix!!


func InitDB() *gorm.DB {
	var err error

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil{
		fmt.Println("DB connection error: ", err.Error())
	}

	err = DB.AutoMigrate(&models.User{},
						&models.Author{},
						&models.Editorial{},
						&models.Book{},
						&models.Favorite{},
						&models.FavoriteItems{},
						&models.Purchase{},
						&models.PurchaseItems{},
	)

	if err != nil {
		fmt.Println("Migration error: ", err.Error())
	}

	return DB
}

func GetConnDB() *gorm.DB{
	if DB == nil {
		once.Do(func() {
			DB = InitDB()
		})
	}

	return DB
}