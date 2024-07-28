package core

import (
	"context"
	"os"
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/mitchellh/go-homedir"
)

func readConfig(ctx context.Context) (string, []byte, error) {
	logger := logr.FromContextOrDiscard(ctx)

	home, err := homedir.Dir()
	if err != nil {
		logger.Error(err, "Error getting home directory")
		return "", nil, err
	}

	configPath := filepath.Join(home, ".gmailctl", "config.jsonnet")

	content, err := os.ReadFile(configPath)
	if err != nil {
		logger.Error(err, "Error reading file")
		return "", nil, err
	}

	return configPath, content, nil
}
