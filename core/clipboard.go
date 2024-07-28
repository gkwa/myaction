package core

import (
	"context"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/go-logr/logr"
)

type ClipboardEmailSource struct{}

func (c *ClipboardEmailSource) GetEmail(ctx context.Context) (string, error) {
	logger := logr.FromContextOrDiscard(ctx)

	content, err := clipboard.ReadAll()
	if err != nil {
		logger.Error(err, "Failed to read from clipboard")
		return "", err
	}

	email := strings.TrimSpace(content)
	if !isValidEmailOrDomain(email) {
		err := fmt.Errorf("invalid email or domain: %s", email)
		logger.Error(err, "Invalid input")
		return "", err
	}

	return email, nil
}
