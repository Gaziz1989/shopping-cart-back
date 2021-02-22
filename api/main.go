package main

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"landing-back/repositories"
	"landing-back/useCase/product"
)

func main() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		panic(err)
	}

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbName"), os.Getenv("dbPass")) //Build connection string

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	// db.LogMode(true)
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	// fmt.Println(defaultTableName)
	// 	return "config_trello." + defaultTableName
	// }
	// db.Debug().AutoMigrate(&NotificationConfigurations{}) //Database migration
	defer db.Close()
	fmt.Println("PostgreSQL connection established!")

	productRepo := repositories.NewPostgreSQLRepository(db)
	productService := product.NewService(productRepo)


	fmt.Printf("%+v\n", productRepo)
	fmt.Printf("%+v\n", productService)
}