package validation

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

const (
	FormatEmail = iota
	FormatHostname
	FormatIP
	FormatURI
)

// EmailValidator ensures that the given value is an email
func EmailValidator(value string) error {
	value = strings.Trim(value, " ")
	_, err := mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func NewPatternValidator(pattern string) Validator[string] {
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return func(value string) error {
		if !re.MatchString(value) {
			return fmt.Errorf("pattern %s matching failed", pattern)
		}
		return nil
	}
}

func validHost(host string) bool {
	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	return re.MatchString(host)
}

func createFormatValidationFunction(format string) Validator[string] {
	switch format {
	case "email":
		return func(value string) error {
			value = strings.Trim(value, " ")
			_, err := mail.ParseAddress(value)
			if err != nil {
				return fmt.Errorf("invalid email")
			}
			return nil
		}
	case "hostname":
		return func(value string) error {
			value = strings.Trim(value, " ")
			if !validHost(value) {
				return fmt.Errorf("invalid hostname")
			}
			return nil
		}
	case "ip":
		return func(value string) error {
			value = strings.Trim(value, " ")
			if net.ParseIP(value) == nil {
				return fmt.Errorf("invalid IP address")
			}
			return nil
		}
	case "uri":
		return func(value string) error {
			value = strings.Trim(value, " ")
			if _, err := url.Parse(value); err != nil {
				return fmt.Errorf("invalid URI")
			}
			return nil
		}
	default:
		return nil
	}
}
