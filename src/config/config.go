package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBURI = ""
	Port  = 3000
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, _ = strconv.Atoi(os.Getenv("PORT"))
	DBURI = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_NAME"))
}
