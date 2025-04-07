package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BaseIP       string
	ServerPort   string
	DBUser       string
	DBPass       string
	DBHost       string
	DBName       string
	FrontendDev  string
	FrontendProd string
)

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(" No se pudo cargar el archivo .env. Asegúrate de que exista y tenga las variables necesarias.")
	}

	FrontendDev = os.Getenv("FRONTEND_URL_DEV")
	if FrontendDev == "" {
		log.Fatal(" FRONTEND_URL_DEV no está definida en el archivo .env")
	}

	FrontendProd = os.Getenv("FRONTEND_URL_PROD")
	if FrontendProd == "" {
		log.Fatal(" FRONTEND_URL_PROD no está definida en el archivo .env")
	}

	BaseIP = os.Getenv("BASE_IP")
	if BaseIP == "" {
		log.Fatal(" BASE_IP no está definida en el archivo .env")
	}

	ServerPort = os.Getenv("PORT")
	if ServerPort == "" {
		log.Fatal(" PORT no está definido en el archivo .env")
	}

	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBHost = os.Getenv("DB_HOST")
	DBName = os.Getenv("DB_NAME")

	if DBUser == "" || DBPass == "" || DBHost == "" || DBName == "" {
		log.Fatal("Faltan credenciales de la base de datos en el archivo .env")
	}
}
