package events

import (
	"github.com/alexcoder04/friendly/v2/ffiles"
	"github.com/fsnotify/fsnotify"
)

// Waits until a file is changed.
// Arguments:
// - file: string
// Returns:
// - file: string
func FileChanged(i map[string]any) map[string]any {
	if _, ok := i["file"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	exist, err := ffiles.Exists(i["file"].(string))
	if !exist || err != nil {
		return map[string]any{
			"success": false,
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}
	defer watcher.Close()

	ch := make(chan bool)

	go func(c chan bool) {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				ch <- false
				return
			}
			if event.Has(fsnotify.Write) {
				ch <- true
				return
			}
		case <-watcher.Errors:
			ch <- false
			return
		}
	}(ch)

	err = watcher.Add(i["file"].(string))
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	res := <-ch

	return map[string]any{
		"success": res,
		"file":    i["file"],
	}
}

// Waits until a file is created in a certain folder.
// Arguments:
// - folder: string
// Returns:
// - folder: string
func FileCreated(i map[string]any) map[string]any {
	if _, ok := i["folder"]; !ok {
		return map[string]any{
			"success": "false",
		}
	}

	if !ffiles.IsDir(i["folder"].(string)) {
		return map[string]any{
			"success": false,
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}
	defer watcher.Close()

	ch := make(chan string)

	go func(c chan string) {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				ch <- ""
				return
			}
			if event.Has(fsnotify.Create) {
				ch <- event.Name
				return
			}
			if event.Has(fsnotify.Write) {
				ch <- ""
				return
			}
		case <-watcher.Errors:
			ch <- ""
			return
		}
	}(ch)

	err = watcher.Add(i["folder"].(string))
	if err != nil {
		return map[string]any{
			"success": false,
		}
	}

	res := <-ch

	return map[string]any{
		"success": true,
		"path":    res,
	}
}
