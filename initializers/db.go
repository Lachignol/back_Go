package initializers

import (
	"fmt"
	"os"

	"github.com/backEnGO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to database")

	}
	fmt.Println("Db is connect")

}

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
