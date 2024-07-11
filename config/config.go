package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
	//Load environment variables from .env file
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}

	// Get environment variables
	serverName = os.Getenv("DB_HOST")
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
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
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		serverName, port, username, password, databaseName)
}

func OpenConnection(connectionString string) (*sql.DB, error) {
	// Log the connection string (without the password)
	log.Printf("Connecting to DB with connection string: host=%s port=%d user=%s dbname=%s",
		serverName, port, username, databaseName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening the database connection: %v", err)
		return nil, err
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error verifying the database connection: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

//func OpenConnection(connectionString string) (*sql.DB, error) {
//	db, err := sql.Open("postgres", connectionString)
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}

func GetSecretKey() string {
	return secretKey
}
