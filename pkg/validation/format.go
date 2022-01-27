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

func validHost(host string) bool {
	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	return re.MatchString(host)
}

func createFormatValidationFunction(format string) ValidationFunction {
	switch format {
	case "email":
		return func(value interface{}) error {
			str := value.(string)
			str = strings.Trim(str, " ")
			_, err := mail.ParseAddress(str)
			if err != nil {
				return fmt.Errorf("invalid email")
			}
			return nil
		}
	case "hostname":
		return func(value interface{}) error {
			str := value.(string)
			str = strings.Trim(str, " ")
			if !validHost(str) {
				return fmt.Errorf("invalid hostname")
			}
			return nil
		}
	case "ip":
		return func(value interface{}) error {
			str := value.(string)
			str = strings.Trim(str, " ")
			if net.ParseIP(str) == nil {
				return fmt.Errorf("invalid IP address")
			}
			return nil
		}
	case "uri":
		return func(value interface{}) error {
			str := value.(string)
			str = strings.Trim(str, " ")
			if _, err := url.Parse(str); err != nil {
				return fmt.Errorf("invalid URI")
			}
			return nil
		}
	default:
		return nil
	}
}

func createPatternValidationFunction(pattern string) ValidationFunction {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return emptyValidator
	}
	return func(value interface{}) error {
		str := value.(string)
		if !re.MatchString(str) {
			return fmt.Errorf("pattern %s matching failed", pattern)
		}
		return nil
	}
}
