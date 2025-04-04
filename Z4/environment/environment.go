package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type RuntimeEnvironment string

const (
	Production  RuntimeEnvironment = "production"
	Development RuntimeEnvironment = "development"
)

type Environment struct {
	DSN string
	ENV RuntimeEnvironment
}

func Initialize() Environment {
	err := godotenv.Load()

	if err != nil {
		e := fmt.Errorf("Error loading .env file: %w", err)
		log.Println(e)
	}

	return Environment{
		DSN: getRequiredEnv("DATABASE_URI"),
		ENV: parseRuntimeEnvironment(getRequiredEnv("ENV")),
	}
}

func getRequiredEnv(s string) string {
	x := os.Getenv(s)

	if x == "" {
		panic(fmt.Errorf("Missing env variable: %s", s))
	}

	return x
}

func parseRuntimeEnvironment(s string) RuntimeEnvironment {
	switch s {
	case "production":
		return Production
	case "development":
		return Development
	default:
		panic(fmt.Errorf("Invalid runtime environment: %s", s))
	}
}
