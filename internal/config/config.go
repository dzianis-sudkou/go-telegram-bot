package config

import (
	"log"
	"os"
)

const Logger bool = true
const BotDebug bool = false

func Config(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable [%s] not found", key)
	}
	return val
}
