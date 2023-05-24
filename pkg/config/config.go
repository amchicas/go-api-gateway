package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	AuthUrl    string
	ProfileUrl string
	CourseUrl  string
}

func LoadConfig() (c Config, err error) {
	err = godotenv.Load("./pkg/config/envs/dev.env")
	if err != nil {
		log.Fatalf("Some error occued .env Err: %s", err)
	}
	c = Config{Port: os.Getenv("PORT"), AuthUrl: os.Getenv("AUTH_SRV_URL"), ProfileUrl: os.Getenv("PROFILE_SRV_URL"), CourseUrl: os.Getenv("COURSE_SRV_URL")}
	return
}
