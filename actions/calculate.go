package actions

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

// Get current user and host name
// Arguments:
// - expression: string
// Returns:
// - result: string
func Calculate(i map[string]any) map[string]any {
	if _, ok := i["expression"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	expr, err := govaluate.NewEvaluableExpression(i["expression"].(string))
	if err != nil {
		return map[string]any{
			"success": "false",
		}
	}
	res, err := expr.Evaluate(nil)
	if err != nil {
		return map[string]any{
			"success": "false",
		}
	}

	return map[string]any{
		"success": true,
		"result":  fmt.Sprintf("%v", res),
	}
}
