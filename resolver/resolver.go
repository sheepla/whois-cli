package resolver

import (
	"errors"
	"fmt"
	"time"

	"github.com/domainr/whois"
	whoisparser "github.com/likexian/whois-parser"
)

var (
	ErrRequest = errors.New("failed to create whois request")
	ErrFetch   = errors.New("an error occurred while retrieving results from the whois server")
	ErrParse   = errors.New("an error occurred while parsing whois raw record")
)

const defaultTimeout = 10 * time.Second

type Option struct {
	Timeout time.Duration
}

func NewOption(timeout time.Duration) *Option {
	return &Option{
		Timeout: timeout,
	}
}

func Resolve(query string, opt *Option) (*whoisparser.WhoisInfo, error) {
	if opt == nil {
		opt = NewOption(defaultTimeout)
	}

	req, err := whois.NewRequest(query)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrRequest, err)
	}

	resp, err := whois.NewClient(opt.Timeout).Fetch(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFetch, err)
	}

	result, err := whoisparser.Parse(string(resp.Body))
	if err != nil {
		return &result, fmt.Errorf("%w: %s", ErrParse, err)
	}

	return &result, nil
}
