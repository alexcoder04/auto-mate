package actions

import (
	"os/exec"
)

// Opens a file in the default application.
// Arguments:
// - file: string
// Returns:
// - file: string - file opened
func FileOpen(i map[string]any) map[string]any {
	if _, ok := i["file"]; !ok {
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
	cmd := exec.Command("xdg-open", i["file"].(string))
	err = cmd.Start()
	return map[string]any{
		"success": err == nil,
		"file":    i["file"],
	}
}
