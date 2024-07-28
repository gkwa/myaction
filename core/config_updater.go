package core

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/google/go-jsonnet"
)

func updateConfig(ctx context.Context, content []byte, email string) (string, error) {
	logger := logr.FromContextOrDiscard(ctx)

	vm := jsonnet.MakeVM()

	newRule := fmt.Sprintf(`{
   	actions: {
   		labels: [
   			"Subscriptions"
   		]
   	},
   	filter: {
   		from: "%s"
   	}
   }`, email)

	jsonnetCode := fmt.Sprintf(`
   	local config = %s;
   	config + {
   		rules: config.rules + [%s]
   	}
   `, string(content), newRule)

	result, err := vm.EvaluateAnonymousSnippet("", jsonnetCode)
	if err != nil {
		logger.Error(err, "Error evaluating Jsonnet")
		return "", err
	}

	return result, nil
}
