package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/config"
	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"github.com/skshahriarahmedraka/Product-Store-Application/routes"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a Authentication-Service-Using-Golang.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	config.LoadEnvVars()
	config.LoadAdmin()
	mongodatabase.DatabaseInitialization()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())

	// Set up your routes
	routes.Routes(r)

	server := &http.Server{
		Addr:    os.Getenv("HOST_ADDR"),
		Handler: r,
	}

	go func() {
		fmt.Println("🚀✨ Api is started ...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")

	return nil
}
