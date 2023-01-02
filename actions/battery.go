package actions

import (
	"time"

	"github.com/alexcoder04/glbat"
)

// Waits until battery is at certain level
// Arguments:
// - level: int
// - type: string ("lower"/"higher")
// Returns:
// - level: int - current battery level
func OnBattery(i map[string]any) map[string]any {
	batIds, err := glbat.GetDetected()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	errCount := 0
	for {
		bat, err := glbat.GetBat(batIds[0])
		if err != nil {
			errCount += 1
			if errCount == 5 {
				return map[string]any{
					"success": false,
				}
			}
			continue
		}

		switch i["type"] {
		case "higher":
			if bat.Capacity > i["level"].(int) {
				return map[string]any{
					"success": true,
					"level":   bat.Capacity,
				}
			}
		case "lower":
			if bat.Capacity < i["level"].(int) {
				return map[string]any{
					"success": true,
					"level":   bat.Capacity,
				}
			}
		default:
			return map[string]any{
				"success": false,
			}
		}

		time.Sleep(30 * time.Second)
	}
}

// Waits until battery is (dis)charging/full
// Arguments:
// - status: string ("Charging"/"Discharging"/"Full")
// Returns:
// - status: string - curent battery status ("Charging"/"Discharging"/"Full")
func OnBatteryStatus(i map[string]any) map[string]any {
	batIds, err := glbat.GetDetected()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	errCount := 0
	for {
		bat, err := glbat.GetBat(batIds[0])
		if err != nil {
			errCount += 1
			if errCount == 5 {
				return map[string]any{
					"success": false,
				}
			}
			continue
		}

		if bat.Status == i["status"].(string) {
			return map[string]any{
				"success": true,
				"status":  i["status"].(string),
			}
		}

		time.Sleep(30 * time.Second)
	}
}
