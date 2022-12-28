package actions

import "os"

// Creates a folder.
// Arguments:
// folder: string
// Returns:
// - folder: string
func FolderCreate(i map[string]any) map[string]any {
	if _, ok := i["folder"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	err := os.MkdirAll(i["folder"].(string), 0700)
	return map[string]any{
		"success": err == nil,
		"folder":  i["folder"],
	}
}
