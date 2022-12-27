package actions

import "github.com/TheCreeper/go-notify"

// Sends a desktop notification.
// Arguments:
// - [title: string]
// - [body: string]
// - [icon: string]
func Notification(i map[string]any) map[string]any {
	title := "auto-mate"
	body := ""
	icon := ""

	if val, ok := i["title"]; ok {
		title = val.(string)
	}

	if val, ok := i["body"]; ok {
		body = val.(string)
	}

	if val, ok := i["icon"]; ok {
		icon = val.(string)
	}

	ntf := notify.NewNotification(title, body)
	ntf.AppIcon = icon

	_, err := ntf.Show()
	return map[string]any{
		"success": err == nil,
	}
}
