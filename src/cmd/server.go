/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sky0621/cv-admin/src/driver"
	"github.com/sky0621/cv-admin/src/rest"
	"github.com/sky0621/cv-admin/src/swagger"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
		 * DB
		 */
		dbClient, closeFunc := driver.MustNewClient()
		defer closeFunc()

		/*
		 * API Server
		 */
		e := echo.New()
		swagger.RegisterHandlers(e, rest.NewRESTService(dbClient))
		if err := http.ListenAndServe(":3000", e); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
