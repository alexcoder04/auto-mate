package actions

import "os"

// Stops executing
// Arguments: none
func Stop(i map[string]any, a ...string) map[string]any {
	os.Exit(0)
	return map[string]any{}
}
