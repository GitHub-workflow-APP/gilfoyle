package cmd

import (
	"context"
	"fmt"
	"github.com/dreamvo/gilfoyle/api"
	"github.com/dreamvo/gilfoyle/api/db"
	"github.com/dreamvo/gilfoyle/config"
	"github.com/dreamvo/gilfoyle/ent/migrate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
)

var httpPort int

func init() {
	// Register command
	rootCmd.AddCommand(serveCmd)

	// Register flags
	serveCmd.PersistentFlags().IntVarP(&httpPort, "port", "p", 3000, "HTTP port")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launch the REST API",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.InitClient(config.GetConfig())
		if err != nil {
			log.Fatalf("failed opening connection: %v", err)
		}
		defer db.Client.Close()

		// run the auto migration tool.
		if err := db.Client.Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		router := gin.Default()

		api.RegisterRoutes(router, api.RouterOptions{
			ExposeSwaggerUI: config.GetConfig().Settings.ExposeSwaggerUI,
		})

		// Launch web server
		if err := router.Run(fmt.Sprintf(":%d", httpPort)); err != nil {
			fmt.Printf("error while launching web server: %e\n", err)
		}
	},
}
