package actions

import (
	"image/color"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

// Generates a QR code from procided string.
// Arguments: data
func QrEncode(i map[string]any, a ...string) map[string]any {
	file := path.Join(os.TempDir(), uuid.NewString()+".png")
	err := qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.White, color.Black, file)
	return map[string]any{
		"success": err == nil,
		"file":    file,
	}
}
