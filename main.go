package main

import (
	"strings"

	"github.com/alexcoder04/auto-mate/actions"
)

func m(b map[string]any, a map[string]any) map[string]any {
	res := map[string]any{}

	for k, v := range a {
		res[k] = v
	}

	for k, v := range b {
		vs, ok := v.(string)
		if ok && strings.HasPrefix(vs, ":->") {
			res[k] = res[v.(string)[3:]]
			delete(res, v.(string)[3:])
			continue
		}
		res[k] = v
	}
	return res
}

func main() {
	sequence := []string{
		"empty",
		"qr-encode",
		"file-open",
		"notification",
	}
	args := []map[string]any{
		{},
		{
			"data": "https://github.com/alexcoder04",
		},
		{},
		{
			"title": "QR code ready",
			"body":  ":->file",
		},
	}

	out := map[string]any{}
	for i, a := range sequence {
		switch a {
		case "calendar-add":
			out = actions.CalendarAdd(m(args[i], out))
		case "empty":
			out = actions.Empty(m(args[i], out))
		case "file-open":
			out = actions.FileOpen(m(args[i], out))
		case "notification":
			out = actions.Notification(m(args[i], out))
		case "qr-encode":
			out = actions.QrEncode(m(args[i], out))
		case "stop":
			out = actions.Stop(m(args[i], out))
		case "time":
			out = actions.Time(m(args[i], out))
		case "wifi":
			out = actions.Wifi(m(args[i], out))
		}

		if !out["success"].(bool) {
			return
		}
	}
}
