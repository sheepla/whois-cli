package resolver

import (
	"errors"
	"fmt"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

var (
	ErrWhois = errors.New("an error occurred while querying the whois server")
	ErrParse = errors.New("an error occurred while parsing whois raw record")
)

func Resolve(domain string) (*whoisparser.WhoisInfo, error) {
	raw, err := whois.Whois(domain)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrWhois, err)
	}

	result, err := whoisparser.Parse(raw)
	if err != nil {
		return &result, fmt.Errorf("%w: %s", ErrParse, err)
	}

	return &result, nil
}
