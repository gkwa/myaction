package core

import (
	"context"

	"github.com/go-logr/logr"
)

func ProcessJsonnet(ctx context.Context, email string) {
	logger := logr.FromContextOrDiscard(ctx)

	configPath, content, err := readConfig(ctx)
	if err != nil {
		return
	}

	updatedConfig, err := updateConfig(ctx, content, email)
	if err != nil {
		return
	}

	if err := validateAndSaveConfig(ctx, configPath, updatedConfig, email); err != nil {
		logger.Error(err, "Failed to validate and save config")
	}
}

func ProcessEmail(ctx context.Context, source EmailSource) {
	logger := logr.FromContextOrDiscard(ctx)

	email, err := source.GetEmail(ctx)
	if err != nil {
		logger.Error(err, "Failed to get valid email")
		return
	}

	ProcessJsonnet(ctx, email)
}
