package amap

import (
	"github.com/alexcoder04/auto-mate/actions"
	"github.com/alexcoder04/auto-mate/events"
)

var Actions = map[string]func(map[string]any) map[string]any{
	"on.battery_level":  events.BatteryLevel,
	"on.battery_status": events.BatteryStatus,
	"on.bluetooth":      events.Bluetooth,
	"on.file_changed":   events.FileChanged,
	"on.file_created":   events.FileCreated,
	"on.network":        events.Network,
	"on.time":           events.Time,
	"on.wifi":           events.Wifi,

	"get.clipboard":     actions.GetClipboard,
	"get.user_and_host": actions.GetUserAndHost,
	"get.wifi_state":    actions.GetWifiState,

	"do.calculate":          actions.Calculate,
	"do.copy":               actions.CliboardCopy,
	"do.copy-self-destruct": actions.CliboardCopySelfDestruct,
	"do.empty":              actions.Empty,
	"do.file_open":          actions.FileOpen,
	"do.folder_create":      actions.FolderCreate,
	"do.notification":       actions.Notification,
	"do.qr_encode":          actions.QrEncode,
	"do.stop":               actions.Stop,
	"do.wifi":               actions.Wifi,
}
