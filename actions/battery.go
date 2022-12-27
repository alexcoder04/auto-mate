package actions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/alexcoder04/friendly/v2/ffiles"
)

func getBatteryPath() (string, error) {
	for i := range []int{0, 1, 2, 3, 4} {
		if ffiles.IsDir(fmt.Sprintf("/sys/class/power_supply/BAT%d", i)) {
			_, err := ioutil.ReadFile(fmt.Sprintf("/sys/class/power_supply/BAT%d/capacity", i))
			if err != nil {
				continue
			}
			return fmt.Sprintf("/sys/class/power_supply/BAT%d", i), nil
		}
	}
	return "", errors.New("no battery found")
}

// Waits until battery is at certain level
// Arguments:
// - level: int
// - type: string ("lower"/"higher")
// Returns:
// - level: int - current battery level
func OnBattery(i map[string]any) map[string]any {
	batPath, err := getBatteryPath()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	errCount := 0
	for {
		data, err := ioutil.ReadFile(batPath + "/capacity")
		if err != nil {
			errCount += 1
			if errCount == 5 {
				return map[string]any{
					"success": false,
				}
			}
			continue
		}
		cap, err := strconv.Atoi(string(data))
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
			if cap > i["level"].(int) {
				return map[string]any{
					"success": true,
					"level":   cap,
				}
			}
		case "lower":
			if cap < i["level"].(int) {
				return map[string]any{
					"success": true,
					"level":   cap,
				}
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
	batPath, err := getBatteryPath()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	errCount := 0
	for {
		data, err := ioutil.ReadFile(batPath + "/status")
		if err != nil {
			errCount += 1
			if errCount == 5 {
				return map[string]any{
					"success": false,
				}
			}
			continue
		}

		if strings.TrimSpace(string(data)) == i["status"].(string) {
			return map[string]any{
				"success": true,
				"status":  i["status"].(string),
			}
		}

		time.Sleep(30 * time.Second)
	}
}
