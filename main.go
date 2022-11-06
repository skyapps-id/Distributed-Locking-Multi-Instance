package main

import (
	"context"
	"dle/redis"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var ctx = context.Background()
var rds = redis.NewRedisClient()

func decriment() error {
	rs := rds.Redsync()
	mutex := rs.NewMutex("promo-lock")

	mutex.Lock()
	val, err := rds.Get(ctx, "PROMO")
	if err != nil {
		fmt.Println("not found value")
		return errors.New("not found value")
	}

	toInt, _ := strconv.Atoi(val.(string))
	if toInt <= 0 {
		fmt.Println("promo is over")
		return errors.New("promo is over")
	}
	toInt--
	if err := rds.Set(ctx, "PROMO", toInt, 0); err != nil {
		fmt.Println(err)
	}
	mutex.Unlock()
	return nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Init Promo
	if err := rds.Set(ctx, "PROMO", 30000, 0); err != nil {
		fmt.Println(err)
	}

	getPromo := func(w http.ResponseWriter, req *http.Request) {
		val, err := rds.Get(ctx, "PROMO")
		if err != nil {
			fmt.Println("not found value")
		}
		io.WriteString(w, fmt.Sprintf("Promo Quota : %s", val))
	}

	submitPromo := func(w http.ResponseWriter, req *http.Request) {
		log.Println(time.Now().Format("15:04:05"))
		err := decriment()
		if err != nil {
			io.WriteString(w, err.Error())
		} else {
			io.WriteString(w, "ok!")
		}
	}

	http.HandleFunc("/get-promo", getPromo)
	http.HandleFunc("/submit-promo", submitPromo)

	log.Println("Listing for Rest API " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
