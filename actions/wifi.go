package actions

import (
	"os/exec"
	"strings"

	"github.com/alexcoder04/friendly/v2"
)

// Changes WiFi state.
// Arguments:
// - status: string ("on"/"off")
// Returns:
// - status: string ("on"/"off")
func Wifi(i map[string]any) map[string]any {
	// TODO check whether it is really on/off
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

	_, err = friendly.GetOutput([]string{"nmcli", "radio", "wifi", i["status"].(string)}, "")
	return map[string]any{
		"success": err == nil,
		"status":  i["status"],
	}
}

// Get wifi state
// Arguments: none
// Returns:
// - wifi-on: string
func GetWifiState(i map[string]any) map[string]any {
	_, err := exec.LookPath("nmcli")
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	out, err := friendly.GetOutput([]string{"nmcli", "radio", "wifi"}, "")
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	return map[string]any{
		"success": true,
		"wifi-on": strings.TrimSpace(out),
	}
}
