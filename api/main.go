package main

import (
	"net/http"
	"fmt"
	"os"
	"time"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"landing-back/repositories"
	"landing-back/useCase/product"
	"landing-back/logger"
	"landing-back/api/handler"
	// "landing-back/entities"
)

type Server struct {
	router *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var start = time.Time{}
	if r.Method != "OPTIONS" {
		start = time.Now()
	}
	//Seting up CORS headers
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	s.router.ServeHTTP(w, r)
	if r.Method != "OPTIONS" {
		elapsedTime := time.Since(start)
		logger.Logg(fmt.Sprintf("Method: %s", r.Method), true)
		logger.Logg(fmt.Sprintf("RequestURI: %s", strings.Split(r.RequestURI, "?")[0]), true)
		logger.Logg(fmt.Sprintf("Total Time For Execution: %s\n", elapsedTime.String()), true)
	}
}

func main() {
	//Connecting to logfile
	err := logger.OpenLog()
	if err != nil {
		logger.Logg(err.Error(), true)
		panic(err)
	}
	defer logger.CloseLog()
	
	//Load .env file
	err = godotenv.Load()
	if err != nil {
		logger.Logg(err.Error(), true)
		panic(err)
	}

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbName"), os.Getenv("dbPass")) //Build connection string

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		logger.Logg(err.Error(), true)
		panic(err)
	}
	defer db.Close()

	// db.LogMode(true)
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	fmt.Println(defaultTableName)
	// 	return "public." + defaultTableName
	// }
	// db.Debug().AutoMigrate(&entities.Product{}) //Database migration
	fmt.Println("PostgreSQL connection established!")

	productRepo := repositories.NewPostgreSQLRepository(db)
	productService := product.NewService(productRepo)


	//Creating http routes
	router := httprouter.New()
	router.HandleOPTIONS = true

	handler.MakeProductHandlers(router, productService)

	err = http.ListenAndServe(os.Getenv("serverPort"), &Server{router})
	if err != nil {
		logger.Logg(err.Error(), true)
		panic(err)
	}
}