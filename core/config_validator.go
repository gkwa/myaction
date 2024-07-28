package core

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
)

func validateAndSaveConfig(ctx context.Context, configPath, updatedConfig string) error {
	logger := logr.FromContextOrDiscard(ctx)

	tempDir, err := os.MkdirTemp("", "gmailctl-config")
	if err != nil {
		logger.Error(err, "Failed to create temp directory")
		return err
	}
	defer os.RemoveAll(tempDir)

	tempConfigPath := filepath.Join(tempDir, "config.jsonnet")
	if err := os.WriteFile(tempConfigPath, []byte(updatedConfig), 0o644); err != nil {
		logger.Error(err, "Failed to write temp config file")
		return err
	}

	gmailctlPath := viper.GetString("gmailctl-path")
	cmd := exec.Command(gmailctlPath, "test", "--filename", tempConfigPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err, "Config validation failed", "output", string(output))
		return fmt.Errorf("config validation failed: %s", string(output))
	}

	if err := os.WriteFile(configPath, []byte(updatedConfig), 0o644); err != nil {
		logger.Error(err, "Failed to write updated config file")
		return err
	}

	logger.Info("Config updated successfully")
	return nil
}
