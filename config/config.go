package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type GreeveConfig struct {
	PORT        int
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	SMTP SMTPConfig

	JWT_Secret string

	PROJECT_ID  string
	BUCKET_NAME string

	Midtrans MidtransConfig
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}
type MidtransConfig struct {
	ClientKey string
	ServerKey string
}

func InitConfig() *GreeveConfig {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var res = new(GreeveConfig)

	res.PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	res.DB_HOST = os.Getenv("DB_HOST")
	res.DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	res.DB_USER = os.Getenv("DB_USER")
	res.DB_PASSWORD = os.Getenv("DB_PASS")
	res.DB_NAME = os.Getenv("DB_NAME")

	res.SMTP.Host = os.Getenv("SMTP_HOST")
	res.SMTP.Port = os.Getenv("SMTP_PORT")
	res.SMTP.Username = os.Getenv("SMTP_USER")
	res.SMTP.Password = os.Getenv("SMTP_PASS")

	res.JWT_Secret = os.Getenv("JWT_SECRET")

	res.PROJECT_ID = os.Getenv("PROJECT_ID")
	res.BUCKET_NAME = os.Getenv("BUCKET_NAME")

	res.Midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	res.Midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")

	return res
}
