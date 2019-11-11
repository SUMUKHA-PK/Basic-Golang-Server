package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var serverData *Data

// Server is the function that starts the listening server
func Server(data *Data) error {

	serverData = data
	if err := checkValidPort(data.Port); err != nil {
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

	data.Router.HandleFunc("/healthCheck", HealthCheckHandler)

	server := &http.Server{
		Handler:   data.Router,
		Addr:      ":" + data.Port,
		ConnState: updateConnectionCount,
	}

	go gracefulShutdown(server)

	log.Println("Starting Server on port " + data.Port)
	if data.HTTPS {
		return server.ListenAndServeTLS("server.crt", "server.key")
	}
	return server.ListenAndServe()
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

// HealthCheckHandler is used for pings to check health of server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health checked. OK")
	enableCors(&w)
	delete(serverData.ConnectionMap, r.RemoteAddr)
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

func updateConnectionCount(c net.Conn, s http.ConnState) {
	if s == http.StateNew {
		serverData.ConnectionMap[c.RemoteAddr().String()]++
		serverData.Count++
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
