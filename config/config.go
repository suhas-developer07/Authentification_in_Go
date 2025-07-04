package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	POSTGERS_URL  string
	APP_PORT      string
}

func (e *envConfig) LoadConfig(){
	err := godotenv.Load()

	if err != nil{
		log.Panic("ENV file is not loaded")
	}
	e.APP_PORT=loadstring("APP_PORT", ":8080")
	e.POSTGERS_URL=loadstring("POSTGRES_URL","postgressurl")

}

var Conifg envConfig 

func init(){
	Conifg.LoadConfig()
}

func loadstring(key string, fallback string) string {
	val, ok := os.LookupEnv(key)  //returns string and boolean
	if !ok {
		log.Panic("APP_PORT is not loaded")
		return fallback
	}
	return val

}