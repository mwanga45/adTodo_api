package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)
func LoadEnv()  {
  if err := godotenv.Load(".env");err != nil{
   log.Fatal("no dotenv file  found ")
  }
}