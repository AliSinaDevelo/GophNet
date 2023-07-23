package main

import(
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port	int
	ReadTimeout	time.Duration
	WriteTimeout time.Duration
	MaxIdleConns	int
	MaxIdleTime	time.Duration
}

func LoadConfig() Config {
	// TODO :load sever config from environment variables

}

func getEnv(key, defaultValue string) string {
	
} 