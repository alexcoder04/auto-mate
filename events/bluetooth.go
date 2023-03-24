package events

import (
	"os/exec"
	"strings"
	"time"

	"github.com/alexcoder04/friendly/v2"
)

// Waits until bluetooth state changes
// Arguments:
// - status: string ("enabled"/"disabled")
// Returns:
// - status: string ("enabled"/"disabled")
func Bluetooth(i map[string]any) map[string]any {
	if _, ok := i["status"]; !ok {
		return map[string]any{
			"success": false,
		}
	}

	_, err := exec.LookPath("bluetooth")
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	errCounter := 0
	for {
		out, err := friendly.GetOutput([]string{"bluetooth"}, "")
		if err != nil {
			errCounter += 1
		}

		if errCounter == 10 {
			return map[string]any{
				"success": false,
			}
		}

		status := strings.TrimSpace(strings.Split(strings.Split(out, "=")[1], "(")[0])

		if status == "on" && i["status"] == "enabled" {
			return map[string]any{
				"success": true,
				"status":  i["status"],
			}
		}

		if status == "off" && i["status"] == "disabled" {
			return map[string]any{
				"success": true,
				"status":  i["status"],
			}
		}

		time.Sleep(15 * time.Second)
	}
}
