package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Item struct {
	gorm.Model
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ExpiresAt int    `json:"expiresAt"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB

func init() {
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Println(e)
	// }

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	_ = pq.Efatal

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Item{})
}

func GetDB() *gorm.DB {
	return db
}
