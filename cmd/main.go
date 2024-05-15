package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/CP-Payne/ecommerce-server/internal/api"
	"github.com/CP-Payne/ecommerce-server/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//func init() {
//	fmt.Println("TEST init")
//}

func main() {

	port := flag.String("port", "3000", "Define the port to listen on")
	env := flag.String("env", "dev", "Define the environment: dev/development (tracelevel logging) or debug (debug level logging) or prod (info level logging)")
	flag.Parse()

	cfg := config.NewConfig(*port, *env)

	r := api.NewRouter(cfg)

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			cfg.Logger.Info("Server shutdown complete")
		} else if err != nil {
			//log.Errorf("Server error: %v", err)
			cfg.Logger.WithFields(
				log.Fields{
					"err": err,
				}).Error("Failed to start server")
			os.Exit(1)
		}
	}()

	cfg.Logger.WithFields(log.Fields{
		"port": cfg.Port,
		"env":  cfg.Env,
	}).Info("Server started")
	<-killSig

	cfg.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		cfg.Logger.WithFields(
			log.Fields{
				"err": err,
			}).Error("Server shutdown failed")
		os.Exit(1)
	}

	cfg.Logger.Info("Server shutdown complete")

}
