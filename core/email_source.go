package core

import (
	"context"
)

type EmailSource interface {
	GetEmail(ctx context.Context) (string, error)
}
