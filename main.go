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
		"on-time",
		"on-wifi",
		"qr-encode",
		"file-open",
		"notification",
		"on-file-changed",
		"notification",
	}
	args := []map[string]any{
		{
			"time":     "12:43",
			"weekdays": "1101111",
		},
		{
			"status": "enabled",
		},
		{
			"data": "https://github.com/alexcoder04",
		},
		{},
		{
			"title": "QR code ready",
			"body":  ":->file",
		},
		{
			"file": "/home/alex/.bashrc",
		},
		{
			"title": "bashrc changed",
		},
	}

	out := map[string]any{}
	for i, a := range sequence {
		switch a {
		case "on-battery":
			out = actions.OnBattery(m(args[i], out))
		case "on-battery-status":
			out = actions.OnBatteryStatus(m(args[i], out))
		case "calendar-add":
			out = actions.CalendarAdd(m(args[i], out))
		case "empty":
			out = actions.Empty(m(args[i], out))
		case "file-open":
			out = actions.FileOpen(m(args[i], out))
		case "on-file-changed":
			out = actions.OnFileChanged(m(args[i], out))
		case "notification":
			out = actions.Notification(m(args[i], out))
		case "qr-encode":
			out = actions.QrEncode(m(args[i], out))
		case "stop":
			out = actions.Stop(m(args[i], out))
		case "on-time":
			out = actions.OnTime(m(args[i], out))
		case "wifi":
			out = actions.Wifi(m(args[i], out))
		case "on-wifi":
			out = actions.OnWifi(m(args[i], out))
		}

		if !out["success"].(bool) {
			return
		}
	}
}
