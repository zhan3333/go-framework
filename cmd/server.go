package cmd

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go-framework/pkg/boot"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			booted *boot.Booted
			err    error
		)
		if booted, err = boot.Boot(boot.WithConfigFile(configFile)); err != nil {
			log.Panicf("boot: %s", err.Error())
		}
		defer booted.Destroy()

		// Create context that listens for the interrupt signal from the OS.
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		addr := fmt.Sprintf("%s:%d", booted.Config.HTTP.Host, booted.Config.HTTP.Port)
		log.Println("Server: listening on", addr)
		srv := &http.Server{
			Addr:    addr,
			Handler: http.TimeoutHandler(booted.Server, booted.Config.HTTP.Timeout, ""),
		}

		// Initializing the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Listen for the interrupt signal.
		<-ctx.Done()

		// Restore default behavior on the interrupt signal and notify user of shutdown.
		stop()
		log.Println("shutting down gracefully, press Ctrl+C again to force")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}

		log.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
