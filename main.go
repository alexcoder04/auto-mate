package main

import (
	"github.com/alexcoder04/auto-mate/actions"
)

func main() {
	sequence := []string{
		"empty",
		"notification",
	}
	args := [][]string{
		{},
		{"Hello World", "This is a test notification", "folder-new"},
	}

	out := map[string]any{}
	for i, a := range sequence {
		switch a {
		case "empty":
			out = actions.Empty(out, args[i]...)
		case "time":
			out = actions.Time(out, args[i]...)
		case "notification":
			out = actions.Notification(out, args[i]...)
		}

		if !out["success"].(bool) {
			return
		}
	}
}
