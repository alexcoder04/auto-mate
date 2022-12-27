package actions

// Does nothing
// Arguments: none
func Empty(i map[string]any, a ...string) map[string]any {
	return map[string]any{
		"success": true,
	}
}
