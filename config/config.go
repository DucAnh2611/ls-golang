package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/DucAnh2611/ls-golang/logging"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logging.Error("Failed to load dotenv!")
	}
}

func GetEnv[K any](key string, fallback K) K {
	if value, exists := os.LookupEnv(key); exists {
		switch any(fallback).(type) {
		case int:
			if v, err := strconv.Atoi(value); err == nil {
				return any(v).(K)
			}
		case int64:
			if v, err := strconv.ParseInt(value, 10, 64); err == nil {
				return any(v).(K)
			}
		case float64:
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				return any(v).(K)
			}
		case bool:
			if v, err := strconv.ParseBool(value); err == nil {
				return any(v).(K)
			}
		case string:
			return any(value).(K)
		case []string:
			return any(strings.Split(value, ",")).(K)
		}
	}

	return fallback
}
