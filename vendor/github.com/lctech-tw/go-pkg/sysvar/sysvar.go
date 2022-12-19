package sysvar

import (
	"os"
	"strconv"
	"strings"
)

func Get(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetAll(key, sep string, fallback []string) []string {
	rawValue := os.Getenv(key)
	if len(rawValue) < 1 {
		return fallback
	}
	valueAll := strings.Split(rawValue, sep)
	if len(valueAll) < 1 {
		return fallback
	}
	return valueAll
}

func GetInt(key string, fallback int) int {
	got := Get(key, "fallback")
	if got == "fallback" {
		return fallback
	}
	atoi, err := strconv.Atoi(got)
	if err != nil {
		return fallback
	}
	return atoi
}

func GetForRunMode(valueForProd interface{}, valueForDev interface{}) interface{} {
	if Get("RUNMODE", "dev") == "prod" {
		return valueForProd
	}
	return valueForDev
}
