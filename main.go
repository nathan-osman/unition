package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/unition/db"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := cli.NewApp()
	app.Name = "unition"
	app.Usage = "browser-based song lyric repository and live presenter"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr",
			Value:  ":http",
			EnvVar: "ADDR",
			Usage:  "HTTP address to listen on",
		},
		cli.StringFlag{
			Name:   "db-host",
			Value:  "postgres",
			EnvVar: "DB_HOST",
			Usage:  "PostgreSQL database host",
		},
		cli.IntFlag{
			Name:   "db-port",
			Value:  5432,
			EnvVar: "DB_PORT",
			Usage:  "PostgreSQL database port",
		},
		cli.StringFlag{
			Name:   "db-name",
			Value:  "postgres",
			EnvVar: "DB_NAME",
			Usage:  "PostgreSQL database name",
		},
		cli.StringFlag{
			Name:   "db-user",
			Value:  "postgres",
			EnvVar: "DB_USER",
			Usage:  "PostgreSQL database user",
		},
		cli.StringFlag{
			Name:   "db-password",
			Value:  "postgres",
			EnvVar: "DB_PASSWORD",
			Usage:  "PostgreSQL database password",
		},
		cli.BoolFlag{
			Name:   "debug",
			EnvVar: "DEBUG",
			Usage:  "enable debug mode",
		},
	}
	app.Action = func(c *cli.Context) error {

		// Initialize zap
		var cfg zap.Config
		if c.Bool("debug") {
			cfg = zap.NewDevelopmentConfig()
		} else {
			cfg = zap.NewProductionConfig()
		}
		logger, err := cfg.Build()
		if err != nil {
			return err
		}

		// Create the database connection
		conn, err := db.New(&db.Config{
			Host:     c.GlobalString("db-host"),
			Port:     c.GlobalInt("db-port"),
			Name:     c.GlobalString("db-name"),
			User:     c.GlobalString("db-user"),
			Password: c.GlobalString("db-password"),
			Logger:   logger,
		})
		if err != nil {
			return err
		}
		defer conn.Close()

		// Wait for SIGINT or SIGTERM
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
	}
}
