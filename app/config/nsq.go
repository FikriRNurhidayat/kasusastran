package config

import "fmt"

func GetNSQAddress() string {
	return fmt.Sprintf("%s:%s",
		Getenv("NSQD_HOST", "127.0.0.1"),
		Getenv("NSQD_PORT", "4150"),
	)
}

func GetNSQLookupAddress() string {
	return fmt.Sprintf("%s:%s",
		Getenv("NSQD_LOOKUP_HOST", "127.0.0.1"),
		Getenv("NSQD_LOOKUP_PORT", "4161"),
	)
}
