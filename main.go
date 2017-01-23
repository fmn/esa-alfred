package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/yuichiro-h/go-esa"
)

const (
	version = "0.0.1"
)

func newEsaClient() *esa.Client {
	accessToken := os.Getenv(envNameEsaAccessToken)
	return esa.New(&esa.Config{AccessToken: accessToken})
}

func main() {
	app := cli.NewApp()
	app.Name = "esa-alfred"
	app.Version = version
	app.Commands = []cli.Command{
		{
			Name:   "search",
			Action: searchCommand,
		},
	}
	app.Run(os.Args)
}
