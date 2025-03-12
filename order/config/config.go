package config

import (
    "log"
    "os"
    "strconv"
)

// GetEnv returns the environment value for "ENV".
// Possible values: "development" or "production".
func GetEnv() string {
    return getEnvironmentValue("ENV")
}

// GetDataSourceURL returns the environment value for "DATA_SOURCE_URL".
// This is the database connection URL.
func GetDataSourceURL() string {
    return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetPaymentServiceURL() string {
    return getEnvironmentValue("PAYMENT_SERVICE_URL")
}

// GetApplicationPort returns the port number for the application.
// It reads the "APPLICATION_PORT" environment variable and converts it to an integer.
func GetApplicationPort() int {
    portStr := getEnvironmentValue("APPLICATION_PORT")
    port, err := strconv.Atoi(portStr)
    if err != nil {
        log.Fatalf("Invalid port: %s", portStr)
    }
    return port
}

// getEnvironmentValue retrieves the value of the specified environment variable.
// If the variable is missing, it logs a fatal error and stops the application.
func getEnvironmentValue(key string) string {
    if os.Getenv(key) == "" {
        log.Fatalf("Environment variable %s is missing.", key)
    }
    return os.Getenv(key)
}