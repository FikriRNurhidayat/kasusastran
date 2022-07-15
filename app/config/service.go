package config

func GetDomain() string {
	return Getenv("DOMAIN", "api.binaracademy.com")
}

func GetService() string {
	return Getenv("SERVICE", "account")
}
