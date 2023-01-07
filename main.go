//nolint:exhaustivestruct,exhaustruct,gochecknoglobals,gci
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"sheepla/whois-cli/resolver"

	whoisparser "github.com/likexian/whois-parser"
	cli "github.com/urfave/cli/v2"
)

var (
	appName        = "whois"
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
		ArgsUsage:   "QUERY",
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

	query := ctx.Args().First()

	result, err := resolver.Resolve(query)
	if err != nil {
		return cli.Exit(err, exitCodeErrWhois.Int())
	}

	if ctx.Bool("json") {
		if err := fprintResultAsJSON(ctx.App.Writer, result); err != nil {
			return cli.Exit(
				err,
				exitCodeErrJSON.Int(),
			)
		}

		return cli.Exit("", exitCodeOK.Int())
	}

	fprintResult(ctx.App.Writer, result)

	return cli.Exit("", exitCodeOK.Int())
}

//nolint:varnamelen
func fprintResult(w io.Writer, result *whoisparser.WhoisInfo) {
	fmt.Fprintf(w, "=== DOMAIN ===\n%s %s (%s)\n",
		result.Domain.ID,
		result.Domain.Name,
		result.Domain.Status,
	)

	fmt.Fprintf(w, "=== DOMAIN ===\n%s %s (%s)\n",
		result.Domain.ID,
		result.Domain.Name,
		result.Domain.Status,
	)
}

func fprintResultAsJSON(w io.Writer, result *whoisparser.WhoisInfo) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(&result); err != nil {
		return fmt.Errorf("failed to encode result as JSON: %w", err)
	}

	return nil
}
