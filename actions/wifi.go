package actions

import (
	"os/exec"

	"github.com/alexcoder04/friendly/v2"
)

// Changes WiFi state.
// Arguments:
// - status: string ("on"/"off")
// Retuns:
// - status: string ("on"/"off")
func Wifi(i map[string]any) map[string]any {
	if _, ok := i["status"]; !ok {
		return map[string]any{
			"success": false,
		}
	}

	_, err := exec.LookPath("nmcli")
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	_, err = friendly.GetOutput([]string{"nmcli", "radio", i["status"].(string)}, "")
	return map[string]any{
		"success": err == nil,
		"status":  i["status"],
	}
}
