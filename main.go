//nolint:exhaustivestruct,exhaustruct,gochecknoglobals,gci
package main

import (
	"fmt"
	"os"

	"sheepla/whois-cli/printer"
	"sheepla/whois-cli/resolver"

	cli "github.com/urfave/cli/v2"
)

var (
	appName        = "whois"
	appVersion     = "unknown"
	appRevision    = "unknown"
	appUsage       = "whois CLI"
	appDescription = "whois CLI"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
	exitCodeErrWhois
	exitCodeErrJSON
)

func (e exitCode) Int() int { return int(e) }

func main() {
	if err := initApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func initApp() *cli.App {
	app := &cli.App{
		Name:        appName,
		Usage:       appUsage,
		Description: appDescription,
		Action:      run,
		ArgsUsage:   "DOMAIN",
		Version:     fmt.Sprintf("%s-%s", appVersion, appRevision),
	}

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "json",
			Aliases: []string{"j"},
			Usage:   "Output in JSON format",
		},
		// &cli.BoolFlag{
		// 	Name:    "shell",
		// 	Aliases: []string{"s"},
		// 	Usage:   "Start interactive mode",
		// },
	}

	return app
}

func run(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return cli.Exit("must requires an augument", exitCodeErrArgs.Int())
	}

	if 1 < ctx.NArg() {
		return cli.Exit(
			fmt.Sprintf("too many arguments (%v)", ctx.Args().Slice()),
			exitCodeErrArgs.Int(),
		)
	}

	domain := ctx.Args().First()

	result, err := resolver.Resolve(domain)
	if err != nil {
		return cli.Exit(err, exitCodeErrWhois.Int())
	}

	if ctx.Bool("json") {
		if err := printer.FprintResultAsJSON(ctx.App.Writer, result); err != nil {
			return cli.Exit(
				err,
				exitCodeErrJSON.Int(),
			)
		}

		return cli.Exit("", exitCodeOK.Int())
	}

	printer.FprintResult(ctx.App.Writer, result)

	return cli.Exit("", exitCodeOK.Int())
}
