package main

import (
	"context"
	"distributed_locking/driver"
	"distributed_locking/handler"
	"distributed_locking/repository"
	"distributed_locking/usecase"
	"log"
	"net/http"
	"os"
)

// Dependency Injection
var ctx = context.Background()
var rds = driver.NewRedisClient()
var db = driver.NewGormDatabase()

var promoRepository = repository.NewPromoRepository(db, &rds)
var promoUsecase = usecase.NewUsecase(promoRepository)
var promoHandler = handler.NewHandler(promoUsecase)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Router
	http.HandleFunc("/promo", promoHandler.GetPromoByCode)
	http.HandleFunc("/promo-claim", promoHandler.DecrimentPromoByCode)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
