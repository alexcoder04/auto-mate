package actions

import (
	"os/exec"
)

// Opens a file in the default application.
// Arguments: filename
func FileOpen(i map[string]any, a ...string) map[string]any {
	if len(a) < 1 {
		return map[string]any{
			"success": "false",
		}
	}

	_, err := exec.LookPath("xdg-open")
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	// TODO add to friendly
	cmd := exec.Command("xdg-open", a[0])
	err = cmd.Start()
	return map[string]any{
		"success": err == nil,
	}
}
