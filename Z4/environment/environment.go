package environment

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gitlab.com/greyxor/slogor"
)

type RuntimeEnvironment string

const (
	Production  RuntimeEnvironment = "production"
	Development RuntimeEnvironment = "development"
)

type Environment struct {
	DSN          string
	ENV          RuntimeEnvironment
	FRONTEND_URL string
	Logger       *slog.Logger
}

func Initialize() Environment {
	err := godotenv.Load()

	if err != nil {
		e := fmt.Errorf("Error loading .env file: %w", err)
		log.Println(e)
	}

	env := parseRuntimeEnvironment(getRequiredEnv("ENV"))

	return Environment{
		DSN: getRequiredEnv("DATABASE_URI"),
		ENV: env,
		// FRONTEND_URL: "http://192.168.117.3:3000",
		FRONTEND_URL: os.Getenv("FRONTEND_URL"),
		Logger:       initializeLogger(env),
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

func initializeLogger(env RuntimeEnvironment) *slog.Logger {
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}

	var handler slog.Handler = slogor.NewHandler(
		os.Stderr,
		slogor.SetLevel(slog.LevelDebug),
		slogor.SetTimeFormat(time.DateTime),
	)

	if env == Production {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}
