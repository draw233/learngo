package util

import "os"

func GetEnv(env string) string{
	x := os.Getenv(env)
	return x
}