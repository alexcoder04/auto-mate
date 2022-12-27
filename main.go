package main

import (
	"github.com/alexcoder04/auto-mate/actions"
)

func main() {
	sequence := []string{
		"empty",
		"qr-encode",
		"notification",
		"file-open",
	}
	args := [][]string{
		{},
		{},
		{"QR code read", "This is a test notification", "folder-new"},
		{"/tmp/8c3982ca-88e7-440e-8d9d-7813abf52a97.png"},
	}

	out := map[string]any{}
	for i, a := range sequence {
		switch a {
		case "calendar-add":
			out = actions.CalendarAdd(out, args[i]...)
		case "empty":
			out = actions.Empty(out, args[i]...)
		case "file-open":
			out = actions.FileOpen(out, args[i]...)
		case "notification":
			out = actions.Notification(out, args[i]...)
		case "qr-encode":
			out = actions.QrEncode(out, args[i]...)
		case "stop":
			out = actions.Stop(out, args[i]...)
		case "time":
			out = actions.Time(out, args[i]...)
		case "wifi":
			out = actions.Wifi(out, args[i]...)
		}

		if !out["success"].(bool) {
			return
		}
	}
}
