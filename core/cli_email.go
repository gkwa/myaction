package core

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	"github.com/go-logr/logr"
)

type CliEmailSource struct {
	Email string
}

func (c *CliEmailSource) GetEmail(ctx context.Context) (string, error) {
	logger := logr.FromContextOrDiscard(ctx)

	email := strings.TrimSpace(c.Email)
	if !isValidEmailOrDomain(email) {
		err := fmt.Errorf("invalid email or domain: %s", email)
		logger.Error(err, "Invalid input")
		return "", err
	}

	return email, nil
}

func isValidEmailOrDomain(input string) bool {
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}
	if input[0] == '@' {
		return isValidDomain(input[1:])
	}
	if _, err := mail.ParseAddress(input); err == nil {
		return true
	}
	return isValidDomain(input)
}

func isValidDomain(domain string) bool {
	parts := strings.Split(domain, ".")
	return len(parts) >= 2 && parts[0] != "" && parts[1] != "" && !strings.Contains(domain, " ")
}
