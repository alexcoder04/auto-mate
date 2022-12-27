package actions

import "github.com/TheCreeper/go-notify"

func Notification(i map[string]any, a ...string) map[string]any {
	var ntf notify.Notification

	switch len(a) {
	case 0:
		return map[string]any{
			"success": false,
		}
	case 1:
		ntf = notify.NewNotification(a[0], "")
	case 2:
		ntf = notify.NewNotification(a[0], a[1])
	case 3:
		ntf = notify.NewNotification(a[0], a[1])
		ntf.AppIcon = a[2]
	}

	_, err := ntf.Show()
	return map[string]any{
		"success": err == nil,
	}
}
