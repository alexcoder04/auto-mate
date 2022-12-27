package actions

import "os"

// Stops executing
// Arguments: none
// Returns: none
func Stop(i map[string]any) map[string]any {
	os.Exit(0)
	return map[string]any{}
}
