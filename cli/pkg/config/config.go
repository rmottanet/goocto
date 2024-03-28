package config

import (
    "os"
    "fmt"
    "log"
    
    "github.com/joho/godotenv"
)


var EnvVars map[string]string


func LoadEnv() {
    EnvVars = make(map[string]string)

    // Load environment variables from .env file, if it exists
    if _, err := os.Stat(".env"); err == nil {
        if err := godotenv.Load(); err != nil {
            log.Fatalf("Error loading .env file: %v", err)
        }
    }

    // Prompt for environment variables if they are not set
    promptIfEmpty("GITHUB_TOKEN")
    promptIfEmpty("GITHUB_USER")

    // Update EnvVars after potential prompts
    EnvVars["GITHUB_TOKEN"] = os.Getenv("GITHUB_TOKEN")
    EnvVars["GITHUB_USER"] = os.Getenv("GITHUB_USER")
}

func promptIfEmpty(key string) {
    if os.Getenv(key) == "" {
        var value string
        fmt.Printf("Please enter value for %s: ", key)
        fmt.Scanln(&value)
        os.Setenv(key, value)
        EnvVars[key] = value
    }
}
