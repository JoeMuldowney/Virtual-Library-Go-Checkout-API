package config

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
	"strconv"
)

var (
	serverName   string
	username     string
	password     string
	databaseName string
	port         int
	secretKey    string
)

func init() {
	// Load environment variables from .env file
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}

	// Get environment variables
	serverName = os.Getenv("DB_HOST")
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	databaseName = os.Getenv("DB_NAME")
	port = getEnvAsInt("DB_PORT", 5432)
	secretKey = os.Getenv("SECRET_KEY")
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
func GetConnectionString() string {
	return fmt.Sprintf("server=%s;user=%s;password=%s;database=%s;port=%d",
		serverName, username, password, databaseName, port)
}

func OpenConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func GetSecretKey() string {
	return secretKey
}
