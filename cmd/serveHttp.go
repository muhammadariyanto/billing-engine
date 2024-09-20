package cmd

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"log"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/muhammadariyanto/billing-engine/config"
	"github.com/muhammadariyanto/billing-engine/port"

	// repo
	billingRepository "github.com/muhammadariyanto/billing-engine/internal/repository/billing"
	customerRepository "github.com/muhammadariyanto/billing-engine/internal/repository/customer"
	loanRepository "github.com/muhammadariyanto/billing-engine/internal/repository/loan"

	// service
	billingService "github.com/muhammadariyanto/billing-engine/internal/service/billing"
	customerService "github.com/muhammadariyanto/billing-engine/internal/service/customer"
	loanService "github.com/muhammadariyanto/billing-engine/internal/service/loan"

	// handler
	billingHandler "github.com/muhammadariyanto/billing-engine/port/handler/billing"
	customerHandler "github.com/muhammadariyanto/billing-engine/port/handler/customer"
	loanHandler "github.com/muhammadariyanto/billing-engine/port/handler/loan"
)

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}

var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "Start HTTP server",
	Long:  `Start Boilerplate HTTP server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Init config
		config, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("Unable to load configuration and secret: %v", err)
		}

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Init repository
		customerRepo := customerRepository.New()
		loanRepo := loanRepository.New()
		billingRepo := billingRepository.New()

		// Init service
		customerSvc := customerService.New(customerRepo, loanRepo, billingRepo)
		loanSvc := loanService.New(loanRepo, customerRepo, billingRepo)
		billingSvc := billingService.New(billingRepo, loanRepo)

		// Init handler
		customerHdlr := customerHandler.New(customerSvc)
		loanHdlr := loanHandler.New(loanSvc, billingSvc)
		billingHdlr := billingHandler.New(billingSvc)

		// Init router
		routerModule := port.RouterModule{
			Cfg:             config,
			CustomerHandler: customerHdlr,
			LoanHandler:     loanHdlr,
			BillingHandler:  billingHdlr,
		}

		r := port.NewRouter(routerModule)

		server := &netHttp.Server{Addr: ":3000", Handler: r}
		// Start the server in a goroutine
		go func() {
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, netHttp.ErrServerClosed) {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Set up signal capturing
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive our signal.
		<-quit

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline or until all connections have returned.
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

		log.Println("Server gracefully stopped")
	},
}
