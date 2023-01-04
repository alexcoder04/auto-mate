package events

import (
	"os/exec"
	"strings"
	"time"

	"github.com/alexcoder04/friendly/v2"
)

// Waits until wifi state changes
// Arguments:
// - status: string ("enabled"/"disabled")
// Returns:
// - status: string ("enabled"/"disabled")
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

	errCounter := 0
	for {
		out, err := friendly.GetOutput([]string{"nmcli", "radio", "wifi"}, "")
		if err != nil {
			errCounter += 1
		}

		if errCounter == 10 {
			return map[string]any{
				"success": false,
			}
		}

		if strings.TrimSpace(out) == i["status"] {
			return map[string]any{
				"success": true,
				"status":  i["status"],
			}
		}

		time.Sleep(15 * time.Second)
	}
}
