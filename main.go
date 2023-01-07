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

func fprintResultAsJSON(w io.Writer, result *whoisparser.WhoisInfo) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(&result); err != nil {
		return fmt.Errorf("failed to encode result as JSON: %w", err)
	}

	return nil
}

//nolint:varnamelen,dupl,funlen
func fprintResult(w io.Writer, result *whoisparser.WhoisInfo) {
	if result.Domain != nil {
		fmt.Fprintln(w, " === DOMAIN ===")
		fmt.Fprintf(w, "ID                  : %v\n", result.Domain.ID)
		fmt.Fprintf(w, "Domain              : %v\n", result.Domain.Domain)
		fmt.Fprintf(w, "Punycode            : %v\n", result.Domain.Punycode)
		fmt.Fprintf(w, "Name                : %v\n", result.Domain.Name)
		fmt.Fprintf(w, "Extension           : %v\n", result.Domain.Extension)
		fmt.Fprintf(w, "WhoisServer         : %v\n", result.Domain.WhoisServer)
		fmt.Fprintf(w, "Status              : %v\n", result.Domain.Status)
		fmt.Fprintf(w, "NameServers         : %v\n", result.Domain.NameServers)
		fmt.Fprintf(w, "DNSSec              : %v\n", result.Domain.DNSSec)
		fmt.Fprintf(w, "CreatedDate         : %v\n", result.Domain.CreatedDate)
		fmt.Fprintf(w, "CreatedDateInTime   : %v\n", result.Domain.CreatedDateInTime)
		fmt.Fprintf(w, "UpdatedDate         : %v\n", result.Domain.UpdatedDate)
		fmt.Fprintf(w, "UpdatedDateInTime   : %v\n", result.Domain.UpdatedDateInTime)
		fmt.Fprintf(w, "ExpirationDate      : %v\n", result.Domain.ExpirationDate)
		fmt.Fprintf(w, "ExpirationDateInTime: %v\n", result.Domain.ExpirationDateInTime)
	}

	if result.Registrar != nil {
		fmt.Fprintln(w, "\n\n === REGISTRAR ===")
		fmt.Fprintf(w, "ID          : %v\n", result.Registrar.ID)
		fmt.Fprintf(w, "Name        : %v\n", result.Registrar.Name)
		fmt.Fprintf(w, "Organization: %v\n", result.Registrar.Organization)
		fmt.Fprintf(w, "Street      : %v\n", result.Registrar.Street)
		fmt.Fprintf(w, "City        : %v\n", result.Registrar.City)
		fmt.Fprintf(w, "Province    : %v\n", result.Registrar.Province)
		fmt.Fprintf(w, "PostalCode  : %v\n", result.Registrar.PostalCode)
		fmt.Fprintf(w, "Country     : %v\n", result.Registrar.Country)
		fmt.Fprintf(w, "Phone       : %v\n", result.Registrar.Phone)
		fmt.Fprintf(w, "PhoneExt    : %v\n", result.Registrar.PhoneExt)
		fmt.Fprintf(w, "Fax         : %v\n", result.Registrar.Fax)
		fmt.Fprintf(w, "FaxExt      : %v\n", result.Registrar.FaxExt)
		fmt.Fprintf(w, "Email       : %v\n", result.Registrar.Email)
		fmt.Fprintf(w, "ReferralURL : %v\n", result.Registrar.ReferralURL)
	}

	if result.Registrant != nil {
		fmt.Fprintln(w, "\n\n === REGISTANT ===")
		fmt.Fprintf(w, "ID          : %v\n", result.Registrant.ID)
		fmt.Fprintf(w, "Name        : %v\n", result.Registrant.Name)
		fmt.Fprintf(w, "Organization: %v\n", result.Registrant.Organization)
		fmt.Fprintf(w, "Street      : %v\n", result.Registrant.Street)
		fmt.Fprintf(w, "City        : %v\n", result.Registrant.City)
		fmt.Fprintf(w, "Province    : %v\n", result.Registrant.Province)
		fmt.Fprintf(w, "PostalCode  : %v\n", result.Registrant.PostalCode)
		fmt.Fprintf(w, "Country     : %v\n", result.Registrant.Country)
		fmt.Fprintf(w, "Phone       : %v\n", result.Registrant.Phone)
		fmt.Fprintf(w, "PhoneExt    : %v\n", result.Registrant.PhoneExt)
		fmt.Fprintf(w, "Fax         : %v\n", result.Registrant.Fax)
		fmt.Fprintf(w, "FaxExt      : %v\n", result.Registrant.FaxExt)
		fmt.Fprintf(w, "Email       : %v\n", result.Registrant.Email)
		fmt.Fprintf(w, "ReferralURL : %v\n", result.Registrant.ReferralURL)
	}

	if result.Administrative != nil {
		fmt.Fprintln(w, "\n\n === ADMINISTRATIVE ===")
		fmt.Fprintf(w, "ID          : %v\n", result.Administrative.ID)
		fmt.Fprintf(w, "Name        : %v\n", result.Administrative.Name)
		fmt.Fprintf(w, "Organization: %v\n", result.Administrative.Organization)
		fmt.Fprintf(w, "Street      : %v\n", result.Administrative.Street)
		fmt.Fprintf(w, "City        : %v\n", result.Administrative.City)
		fmt.Fprintf(w, "Province    : %v\n", result.Administrative.Province)
		fmt.Fprintf(w, "PostalCode  : %v\n", result.Administrative.PostalCode)
		fmt.Fprintf(w, "Country     : %v\n", result.Administrative.Country)
		fmt.Fprintf(w, "Phone       : %v\n", result.Administrative.Phone)
		fmt.Fprintf(w, "PhoneExt    : %v\n", result.Administrative.PhoneExt)
		fmt.Fprintf(w, "Fax         : %v\n", result.Administrative.Fax)
		fmt.Fprintf(w, "FaxExt      : %v\n", result.Administrative.FaxExt)
		fmt.Fprintf(w, "Email       : %v\n", result.Administrative.Email)
		fmt.Fprintf(w, "ReferralURL : %v\n", result.Administrative.ReferralURL)
	}

	if result.Technical != nil {
		fmt.Fprintln(w, "\n\n === TECHNICAL ===")
		fmt.Fprintf(w, "ID          : %v\n", result.Technical.ID)
		fmt.Fprintf(w, "Name        : %v\n", result.Technical.Name)
		fmt.Fprintf(w, "Organization: %v\n", result.Technical.Organization)
		fmt.Fprintf(w, "Street      : %v\n", result.Technical.Street)
		fmt.Fprintf(w, "City        : %v\n", result.Technical.City)
		fmt.Fprintf(w, "Province    : %v\n", result.Technical.Province)
		fmt.Fprintf(w, "PostalCode  : %v\n", result.Technical.PostalCode)
		fmt.Fprintf(w, "Country     : %v\n", result.Technical.Country)
		fmt.Fprintf(w, "Phone       : %v\n", result.Technical.Phone)
		fmt.Fprintf(w, "PhoneExt    : %v\n", result.Technical.PhoneExt)
		fmt.Fprintf(w, "Fax         : %v\n", result.Technical.Fax)
		fmt.Fprintf(w, "FaxExt      : %v\n", result.Technical.FaxExt)
		fmt.Fprintf(w, "Email       : %v\n", result.Technical.Email)
		fmt.Fprintf(w, "ReferralURL : %v\n", result.Technical.ReferralURL)
	}

	if result.Billing != nil {
		fmt.Fprintln(w, "\n\n === BILLING ===")
		fmt.Fprintf(w, "ID          : %v\n", result.Billing.ID)
		fmt.Fprintf(w, "Name        : %v\n", result.Billing.Name)
		fmt.Fprintf(w, "Organization: %v\n", result.Billing.Organization)
		fmt.Fprintf(w, "Street      : %v\n", result.Billing.Street)
		fmt.Fprintf(w, "City        : %v\n", result.Billing.City)
		fmt.Fprintf(w, "Province    : %v\n", result.Billing.Province)
		fmt.Fprintf(w, "PostalCode  : %v\n", result.Billing.PostalCode)
		fmt.Fprintf(w, "Country     : %v\n", result.Billing.Country)
		fmt.Fprintf(w, "Phone       : %v\n", result.Billing.Phone)
		fmt.Fprintf(w, "PhoneExt    : %v\n", result.Billing.PhoneExt)
		fmt.Fprintf(w, "Fax         : %v\n", result.Billing.Fax)
		fmt.Fprintf(w, "FaxExt      : %v\n", result.Billing.FaxExt)
		fmt.Fprintf(w, "Email       : %v\n", result.Billing.Email)
		fmt.Fprintf(w, "ReferralURL : %v\n", result.Billing.ReferralURL)
	}
}
