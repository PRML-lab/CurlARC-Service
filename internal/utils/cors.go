package utils

import "os"

func GetAllowOrigins() []string {
	env := os.Getenv("ENV")
	switch env {
	case "production":
		return []string{os.Getenv("ALLOW_ORIGINS_PROD")}
	default:
		return []string{os.Getenv("ALLOW_ORIGINS_DEV")}
	}
}
