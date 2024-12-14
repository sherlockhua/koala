package env

import "os"

const (
	KOALA_ENV = "KOALA_ENV"
)

const (
	PRODUCTION  = "PRODUCTION"
	STAGING     = "STAGING"
	DEVELOPMENT = "DEVELOPMENT"
	TESTING     = "TESTING"
	LOCAL       = "LOCAL"
	DEFAULT_ENV = PRODUCTION
)

func GetKoalaEnv() string {
	result := os.Getenv(KOALA_ENV)
	if result == "" {
		result = DEFAULT_ENV
	}
	return result
}

func IsOnline() bool {
	env := GetKoalaEnv()
	return env == PRODUCTION || env == STAGING
}

func IsTest() bool {
	return GetKoalaEnv() == TESTING
}

func IsDevelopment() bool {
	env := GetKoalaEnv()
	return env == DEVELOPMENT || env == LOCAL
}
