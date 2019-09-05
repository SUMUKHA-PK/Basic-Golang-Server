package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Server is the function that starts the listening server
func Server(r *mux.Router, port string) error {

	if err := checkValidPort(port); err != nil {
		return err
	}
	LogFileLocation := os.Getenv("LogFileLocation")
	if LogFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LogFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	r.HandleFunc("/healthCheck", healthCheckHandler)

	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	go gracefulShutdown(server)

	log.Println("Starting Server on port " + port)
	return server.ListenAndServeTLS("server.crt", "server.key")
}

// gracefulShutdown shuts down the server on getting a ^C signal
func gracefulShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for currently serving items.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

// healthCheckHandler is used for pings to check health of server
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health checked. OK")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func checkValidPort(port string) error {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	if portInt > 65535 {
		return errors.New("Port number exceeds limit of 65535")
	}
	return nil
}
