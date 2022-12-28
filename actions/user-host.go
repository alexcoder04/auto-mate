package actions

import (
	"os"
	"os/user"
)

// Get current user and host name
// Arguments: none
// Returns:
// - username: string
// - hostname: string
func GetUserAndHost(i map[string]any) map[string]any {
	user, err := user.Current()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	return map[string]any{
		"success":  true,
		"username": user.Username,
		"hostname": hostname,
	}
}
