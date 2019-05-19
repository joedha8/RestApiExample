package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/renosyah/RestApiExample/middleware"
	"github.com/renosyah/RestApiExample/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbPool  *sql.DB
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()

		http.Handle("/", r)

		api_router := r.PathPrefix("/api/v1").Subrouter()
		api_router.Use(middleware.AuthenticationMiddleware)
		api_router.Handle("/ping", router.HandlerFunc(router.HandlerPing)).Methods(http.MethodPost)

		port := viper.GetInt("app.port")

		srv := &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			ReadTimeout:  time.Duration(viper.GetInt("app.read_timeout")) * time.Second,
			WriteTimeout: time.Duration(viper.GetInt("app.write_timeout")) * time.Second,
			Handler:      r,
		}
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println(err)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	viper.SetConfigType("toml")

	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {

			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/config")
		viper.SetConfigName(".config")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
