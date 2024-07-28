package core

import (
	"context"

	"github.com/go-logr/logr"
)

func Hello(ctx context.Context) {
	logger := logr.FromContextOrDiscard(ctx)
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")

	source := &ClipboardEmailSource{}
	ProcessEmail(ctx, source)
}
