package main

import (
	"github.com/ImTheCurse/jelly/pkg/email"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := email.New("Hello test", []string{"example@gmail.com"}, "Hello World!")
	e.SendMail()
}
