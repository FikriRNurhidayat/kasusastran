package config

import (
	"strconv"
	"time"
)

func GetAccessTokenSecretKey() string {
	return Getenv("ACCESS_TOKEN_SECRET_KEY", "rahasia")
}

func GetAccessTokenExpirationTimeUnit() time.Duration {
	val := Getenv("ACCESS_TOKEN_EXPIRATION_UNIT", "MINUTE")
	switch val {
	case "HOUR":
		return time.Hour
	case "MINUTE":
		return time.Minute
	default:
		return time.Second
	}
}

func GetAccessTokenExpirationTimeValue() time.Duration {
	val := Getenv("ACCESS_TOKEN_EXPIRATION_VALUE", "5")
	exp, _ := strconv.ParseUint(val, 10, 32)
	return time.Duration(exp)
}

func GetAccessTokenExpirationTime() time.Duration {
	return GetAccessTokenExpirationTimeUnit() * GetAccessTokenExpirationTimeValue()
}

func GetRefreshTokenSecretKey() string {
	return Getenv("REFRESH_TOKEN_SECRET_KEY", "rahasia")
}

func GetRefreshTokenExpirationTimeUnit() time.Duration {
	val := Getenv("REFRESH_TOKEN_EXPIRATION_UNIT", "HOUR")
	switch val {
	case "HOUR":
		return time.Hour
	case "MINUTE":
		return time.Minute
	default:
		return time.Second
	}
}

func GetRefreshTokenExpirationTimeValue() time.Duration {
	val := Getenv("REFRESH_TOKEN_EXPIRATION_VALUE", "24")
	exp, _ := strconv.ParseUint(val, 10, 32)
	return time.Duration(exp)
}

func GetRefreshTokenExpirationTime() time.Duration {
	return GetRefreshTokenExpirationTimeUnit() * GetRefreshTokenExpirationTimeValue()
}
