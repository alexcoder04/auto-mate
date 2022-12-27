package actions

import "time"

func Time(i map[string]any, a ...string) map[string]any {
	for {
		if time.Now().Format("15:04") == a[0] {
			return map[string]any{
				"success": true,
			}
		}
	}
}
