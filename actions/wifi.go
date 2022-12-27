package actions

import (
	"os/exec"

	"github.com/alexcoder04/friendly/v2"
)

// Changes WiFi state.
// Arguments: "on"/"off"
func Wifi(i map[string]any, a ...string) map[string]any {
	if len(a) < 1 {
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

	_, err = friendly.GetOutput([]string{"nmcli", "radio", a[0]}, "")
	return map[string]any{
		"success": err == nil,
	}
}
