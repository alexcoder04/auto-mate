package events

import (
	"errors"
	"time"

	"github.com/alexcoder04/friendly/v2"
)

func getWeekdaysList(i string) ([]time.Weekday, error) {
	weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	if i == "" {
		return weekdays, nil
	}

	if len(i) != 7 {
		return weekdays, errors.New("invalid input string")
	}

	if !friendly.SafeCharsOnly(i, []byte("01")) {
		return weekdays, errors.New("invalid input string")
	}

	res := []time.Weekday{}
	for c, v := range i {
		if v == '1' {
			res = append(res, weekdays[c])
		}
	}

	return res, nil
}

// Waits until a certain time.
// Arguments:
// - time: string ("HH:MM")
// - weekdays: string ("1111100") - starting with Mo, 1 means run, 0 not run
// Returns:
// - time: string - current time (HH:MM)
func Time(i map[string]any) map[string]any {
	if _, ok := i["time"]; !ok {
		return map[string]any{
			"success": false,
		}
	}

	if _, ok := i["weekdays"]; !ok {
		i["weekdays"] = ""
	}

	weekdays, err := getWeekdaysList(i["weekdays"].(string))
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	for {
		now := time.Now()
		if now.Format("15:04") == i["time"] && friendly.ArrayContains(weekdays, now.Weekday()) {
			return map[string]any{
				"success": true,
				"time":    time.Now().Format("15:04"),
			}
		}
		time.Sleep(time.Minute)
	}
}
