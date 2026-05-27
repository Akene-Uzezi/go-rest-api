package main

import (
	"database/sql"
	"log"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// @title Rest API in Gin
// @version 1.0
// @description This is a simple REST API in Gin.
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token for in the format **Bearer &lt;token&gt
type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "secret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
