package actions

import "time"

// Waits until a certain time.
// Arguments:
// - time: string (HH:MM)
// Returns:
// - time: string - current time (HH:MM)
func Time(i map[string]any) map[string]any {
	if _, ok := i["time"]; !ok {
		return map[string]any{
			"success": false,
		}
	}

	for {
		if time.Now().Format("15:04") == i["time"] {
			return map[string]any{
				"success": true,
				"time":    time.Now().Format("15:04"),
			}
		}
	}
}
