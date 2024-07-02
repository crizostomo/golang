package app

import "github.com/urfave/cli"

// Generate will return the command line ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPS and Server Names on the Internet"

	return app
}