package database

import(
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB(){
	dsn := "host=localhost user=admin password=admin123 dbname=studentdb port=5432 sslmode=disable"
	db,err :=gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!= nil{
		log.Fatal("Failed to connect to database! \n",err)
	}

	fmt.Println("Database connected successfull")

	DB = db
}