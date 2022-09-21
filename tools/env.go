package tools

import "os"

func GetEnvOrDefault(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvOrDefaultBool(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}
