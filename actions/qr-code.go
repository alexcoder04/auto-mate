package actions

import (
	"image/color"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

// Generates a QR code from procided string.
// Arguments:
// - data: string
// Returns:
// - data: string - original data
// - file: string - file where the qr code was saved
func QrEncode(i map[string]any) map[string]any {
	if _, ok := i["data"]; !ok {
		return map[string]any{
			"success": false,
		}
	}

	file := path.Join(os.TempDir(), uuid.NewString()+".png")
	err := qrcode.WriteColorFile(i["data"].(string), qrcode.Medium, 256, color.White, color.Black, file)
	return map[string]any{
		"success": err == nil,
		"file":    file,
		"data":    i["data"],
	}
}
