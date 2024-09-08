package main

import (
	"flag"
	"log"
	"os"
	"quiz3/config"
	"quiz3/routes"

	"github.com/joho/godotenv"
)

func main() {
    
    migrate := flag.Bool("migrate", false, "Run database migrations")
    flag.Parse()

    
    err := godotenv.Load("config/.env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    
    config.ConnectDatabase()

    if *migrate {
        log.Println("Running migrations...")
        config.RunMigration()
        return
    }

    
    router := routes.SetupRouter()

   
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    
    log.Printf("Server running on port %s", port)
    router.Run(":" + port)
}
