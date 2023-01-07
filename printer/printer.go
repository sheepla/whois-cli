package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	whoisparser "github.com/likexian/whois-parser"
	"github.com/olekukonko/tablewriter"
)

type Row struct {
	Item  string
	Value string
}

//nolint:varnamelen
func FprintResult(w io.Writer, result *whoisparser.WhoisInfo) {
	if result.Domain != nil {
		fmt.Fprint(w, "::\n:: DOMAIN\n::\n")
		fprintDomain(w, result.Domain)
		fmt.Fprintln(w, "")
	}

	if result.Registrar != nil {
		fmt.Fprint(w, "::\n:: REGISTRAR\n::\n")
		fprintContact(w, result.Registrar)
		fmt.Fprintln(w, "")
	}

	if result.Registrant != nil {
		fmt.Fprint(w, "::\n:: REGISTANT\n::\n")
		fprintContact(w, result.Registrant)
		fmt.Fprintln(w, "")
	}

	if result.Administrative != nil {
		fmt.Fprint(w, "::\n:: ADMINISTRATIVE\n::\n")
		fprintContact(w, result.Administrative)
		fmt.Fprintln(w, "")
	}

	if result.Technical != nil {
		fmt.Fprint(w, "::\n:: TECHNICAL\n::\n")
		fprintContact(w, result.Technical)
		fmt.Fprintln(w, "")
	}

	if result.Billing != nil {
		fmt.Fprint(w, "::\n:: BILLING\n::\n")
		fprintContact(w, result.Billing)
		fmt.Fprintln(w, "")
	}
}

func fprintDomain(w io.Writer, domain *whoisparser.Domain) {
	if domain == nil {
		return
	}

	table := tablewriter.NewWriter(w)
	table.SetBorders(tablewriter.Border{
		Top:    false,
		Left:   false,
		Bottom: false,
		Right:  false,
	})
	table.SetHeader([]string{"ITEM", "VALUE"})

	appendRow(table, "ID", domain.ID)
	appendRow(table, "Domain", domain.Domain)
	appendRow(table, "Punycode", domain.Punycode)
	appendRow(table, "Name", domain.Name)
	appendRow(table, "Extension", domain.Extension)
	appendRow(table, "WhoisServer", domain.WhoisServer)
	appendRow(table, "Status", domain.Status)
	appendRow(table, "NameServers", domain.NameServers)
	appendRow(table, "DNSSec", domain.DNSSec)
	appendRow(table, "CreatedDate", domain.CreatedDate)
	appendRow(table, "UpdatedDate", domain.UpdatedDate)
	appendRow(table, "ExpirationDate", domain.ExpirationDate)

	table.Render()
}

func fprintContact(w io.Writer, contact *whoisparser.Contact) {
	if contact == nil {
		return
	}

	table := tablewriter.NewWriter(w)
	table.SetBorders(tablewriter.Border{
		Top:    false,
		Left:   false,
		Bottom: false,
		Right:  false,
	})

	table.SetHeader([]string{"ITEM", "VALUE"})

	appendRow(table, "ID", contact.ID)
	appendRow(table, "Name", contact.Name)
	appendRow(table, "Organization", contact.Organization)
	appendRow(table, "Street", contact.Street)
	appendRow(table, "City", contact.City)
	appendRow(table, "Province", contact.Province)
	appendRow(table, "PostalCode", contact.PostalCode)
	appendRow(table, "Country", contact.Country)
	appendRow(table, "Phone", contact.Phone)
	appendRow(table, "PhoneExt", contact.PhoneExt)
	appendRow(table, "Fax", contact.Fax)
	appendRow(table, "FaxExt", contact.FaxExt)
	appendRow(table, "Email", contact.Email)
	appendRow(table, "ReferralURL", contact.ReferralURL)

	table.Render()
}

func toDisplayString(value any) string {
	if arr, ok := value.([]string); ok {
		if 0 < len(arr) {
			return strings.Join(arr, " ")
		}
	}

	if s, ok := value.(string); ok {
		return strings.TrimSpace(s)
	}

	if b, ok := value.(bool); ok {
		return strconv.FormatBool(b)
	}

	return ""
}

func appendRow(table *tablewriter.Table, key string, value any) {
	if table == nil {
		return
	}

	if str := toDisplayString(value); str != "" {
		table.Append([]string{key, str})
	}
}

func FprintResultAsJSON(w io.Writer, result *whoisparser.WhoisInfo) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(&result); err != nil {
		return fmt.Errorf("failed to encode result as JSON: %w", err)
	}

	return nil
}
