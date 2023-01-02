package actions

import (
	"time"

	"golang.design/x/clipboard"
)

func copy(text string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}

	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}

func get() (string, error) {
	err := clipboard.Init()
	if err != nil {
		return "", err
	}

	return string(clipboard.Read(clipboard.FmtText)), nil
}

// Copy text to clipboard
// Arguments:
// - text: string - text to copy
// Returns: none
func CliboardCopy(i map[string]any) map[string]any {
	if _, ok := i["text"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	err := copy(i["text"].(string))
	return map[string]any{
		"success": err == nil,
	}
}

// Get data from clipboard.
// Arguments: none
// Returns:
// - text: string - text from clipboard
func GetClipboard(i map[string]any) map[string]any {
	text, err := get()

	return map[string]any{
		"success": err == nil,
		"text":    text,
	}
}

// Copy text to clipboard and delete it again after certain amount of time.
// Arguments:
// - text: string - text to copy
//  - timeout: int - clear clipboard after so many seconds
// Returns: none
func CliboardCopySelfDestruct(i map[string]any) map[string]any {
	if _, ok := i["text"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	if _, ok := i["timeout"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	prev, _ := get()
	new := i["text"].(string)
	err := copy(new)
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	if err == nil {
		go func() {
			time.Sleep(time.Duration(i["timeout"].(int)) * time.Second)
			// only overwrite if it is still our text
			now, _ := get()
			if now == new {
				copy(prev)
			}
		}()
	}

	return map[string]any{
		"success": true,
	}
}
