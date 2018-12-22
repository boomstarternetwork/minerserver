package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/boomstarternetwork/bestore"
	"github.com/boomstarternetwork/minerserver/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	cli "gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

func main() {
	app := cli.NewApp()
	app.Name = "minerserver"
	app.Usage = ""
	app.Description = "Boomstarter minerserver."
	app.Author = "Vadim Chernov"
	app.Email = "v.chernov@boomstarter.ru"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config file",
			Value: "",
		},
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "postgres-cs",
			Usage: "postgres connection string",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "bind-addr",
			Usage: "web server bind address",
			Value: ":80",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "run-mode",
			Usage: "run mode: production or development",
			Value: "production",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "log-level",
			Usage: "log level: debug, info, warn, error, off",
			Value: "info",
		}),
	}

	app.Before = altsrc.InitInputSourceWithContext(app.Flags,
		func(c *cli.Context) (altsrc.InputSourceContext, error) {
			config := c.String("config")
			if config != "" {
				return altsrc.NewYamlSourceFromFlagFunc("config")(c)
			}
			return &altsrc.MapInputSource{}, nil
		})

	app.Action = appAction

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func appAction(c *cli.Context) error {
	connStr := c.String("postgres-cs")
	bindAddr := c.String("bind-addr")
	runMode := c.String("run-mode")
	logLevel := c.String("log-level")

	s, err := bestore.NewDBStore(connStr, runMode)
	if err != nil {
		return cli.NewExitError("failed to create store: "+err.Error(), 1)
	}

	e, err := initWebServer(s, runMode, logLevel)
	if err != nil {
		return cli.NewExitError("failed to init web server: "+err.Error(), 2)
	}

	err = e.Start(bindAddr)
	if err != nil {
		return cli.NewExitError("failed to start web server: "+
			err.Error(), 3)
	}

	return nil
}

func initWebServer(s bestore.Store, runMode string,
	logLevel string) (*echo.Echo, error) {
	e := echo.New()

	switch runMode {
	case "production":
		e.Use(middleware.Recover())
	case "development":
		e.Use(middleware.Recover())
		e.Debug = true
	case "testing":
	default:
		return nil, errors.New("invalid mode")
	}

	switch logLevel {
	case "debug", "info", "warn", "error":
		e.Use(middleware.Logger())
	}

	switch logLevel {
	case "debug":
		e.Logger.SetLevel(log.DEBUG)
	case "info":
		e.Logger.SetLevel(log.INFO)
	case "warn":
		e.Logger.SetLevel(log.WARN)
	case "error":
		e.Logger.SetLevel(log.ERROR)
	case "off":
	default:
		return nil, errors.New("invalid log level")
	}

	h := handler.NewHandler(s)

	e.HTTPErrorHandler = h.ErrorHandler

	e.GET("/projects", h.Projects)

	return e, nil
}
