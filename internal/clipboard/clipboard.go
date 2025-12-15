// Package clipboard provides clipboard image reading functionality.
package clipboard

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"golang.design/x/clipboard"
)

var (
	imageCounter atomic.Int64
	initOnce     sync.Once
	initErr      error
)

// ReadImage checks if the clipboard contains an image and saves it to a temp file.
// Returns the file path if an image was found, empty string if no image, or error.
func ReadImage() (string, error) {
	// Initialize clipboard (only once).
	initOnce.Do(func() {
		initErr = clipboard.Init()
	})
	if initErr != nil {
		return "", initErr
	}

	// Read image data from clipboard.
	imgData := clipboard.Read(clipboard.FmtImage)
	if len(imgData) == 0 {
		return "", nil // No image in clipboard.
	}

	// Save to temp file.
	count := imageCounter.Add(1)
	tmpDir := os.TempDir()
	fileName := fmt.Sprintf("pluse-clipboard-%d-%d.png", time.Now().Unix(), count)
	filePath := filepath.Join(tmpDir, fileName)

	if err := os.WriteFile(filePath, imgData, 0o644); err != nil {
		return "", err
	}

	return filePath, nil
}
