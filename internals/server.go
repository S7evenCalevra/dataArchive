package internals

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	app := new(application)
	app.auth.username = os.Getenv("AUTH_USERNAME")
	app.auth.password = os.Getenv("AUTH_PASSWORD")

	if app.auth.username == "" {
		log.Fatal("basic auth username must be provided")
	}

	if app.auth.password == "" {
		log.Fatal("basic auth password must be provided")
	}

	l := log.New(os.Stdout, "API ", log.LstdFlags)
	//fmt.Println(l)
	//Create the default mux server and instanciate the routes
	servermux := http.NewServeMux()

	// routes
	//servermux.HandleFunc("/getalltable", app.basicAuth(app.GetTableinfo))

	servermux.HandleFunc("/", app.basicAuth(app.GetTableinfo))

	servermux.HandleFunc("/postrecord", app.basicAuth(app.Handler2))

	//servermux.HandleFunc("/", handlers.ServeHTTP)
	listenAddr := ":4000"

	// Uncomment for Azure Function host only
	//if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
	//	listenAddr = ":" + val
	//}

	//create a server to handle timeouts etc
	s := http.Server{
		Addr:         listenAddr,        // configure the bind address
		Handler:      servermux,         // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		l.Printf("Starting server on port %s", listenAddr)
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		//log.Println("Issue with background tasks and shutting down")
		log.Fatal(
			"There was a Issue with shutting down the server gracefully.",
		)
	}
	s.Shutdown(ctx)
}
