package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/juby-gif/pillshare-server/internal/controllers"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func main() {

	// Load .env file using the package "github.com/joho/godotenv"
	// If any error occurs while loading, it will output the error as "Error loading .env file"
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the values of environment variables from .env file
	// Mapping the value to the corresponding keys
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	db, err := utils.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(db)

	// Initializing controllers
	c := controllers.New()

	// http.NewServeMux function to create an empty ServeMux.
	// mux.HandleFunc function for registering handler for all api requests with the URL path /.
	mux := http.NewServeMux()
	mux.HandleFunc("/", c.HandleRequests)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("SERVER_API_DOMAIN"), os.Getenv("PORT")),
		Handler: mux,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go runMainRuntimeLoop(server)

	log.Print("Server Started")

	<-done

	stopMainRuntimeLoop(server)
}

func runMainRuntimeLoop(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func stopMainRuntimeLoop(srv *http.Server) {
	log.Printf("Starting graceful shutdown now...")

	// Execute the graceful shutdown sub-routine which will terminate any
	// active connections and reject any new connections.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Printf("Graceful shutdown finished.")
	log.Print("Server Exited")
}
