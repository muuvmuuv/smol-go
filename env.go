package main

import "os"

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var port = getEnv("PORT", "1717")
var mode = getEnv("GIN_MODE", "develop")
var devEnv = mode != "release"
