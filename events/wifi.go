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

// Waits until certain network is connected/disconnected
// Arguments:
// - status: string ("enabled"/"disabled")
// - network: string
// Returns:
// - status: string ("enabled"/"disabled")
// - network: string
func Network(i map[string]any) map[string]any {
	if _, ok := i["network"]; !ok {
		return map[string]any{
			"success": false,
		}
	}

	if _, ok := i["status"]; !ok {
		i["status"] = "enabled"
	}

	_, err := exec.LookPath("nmcli")
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	errCounter := 0
	for {
		out, err := friendly.GetOutput([]string{"nmcli", "--terse", "--fields", "NAME,TYPE", "connection", "show", "--active"}, "")
		if err != nil {
			errCounter += 1
		}

		if errCounter == 10 {
			return map[string]any{
				"success": false,
			}
		}

		connected := []string{}
		for _, line := range strings.Split(out, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			parts := strings.Split(line, ":")
			if !strings.HasSuffix(parts[1], "-wireless") {
				continue
			}
			connected = append(connected, parts[0])
		}

		if i["status"] == "disabled" {
			if !friendly.ArrayContains(connected, i["network"].(string)) {
				return map[string]any{
					"success": true,
					"status":  i["status"],
					"network": i["network"],
				}
			}
		}

		if i["status"] == "enabled" {
			if friendly.ArrayContains(connected, i["network"].(string)) {
				return map[string]any{
					"success": true,
					"status":  i["status"],
					"network": i["network"],
				}
			}
		}

		time.Sleep(15 * time.Second)
	}
}
