package main

import (
	"github.com/ImTheCurse/jelly/pkg/sms"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	msg := Message.NewSMS(os.Getenv("TO_NUMBER"), "Hello World!")
	srcNumber := os.Getenv("TWILLO_PHONE_NUMBER")
	msg.Send(srcNumber)
}
