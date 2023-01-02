package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/alexcoder04/auto-mate/actions"
	"go.uber.org/zap"
)

type ConfigAction struct {
	Type string         `json:"type"`
	Args map[string]any `json:"args"`
}

type ConfigSequence struct {
	Name    string         `json:"name"`
	Actions []ConfigAction `json:"actions"`
}

var Actions = map[string]func(map[string]any) map[string]any{
	"on.battery-level":  actions.OnBatteryLevel,
	"on.battery-status": actions.OnBatteryStatus,
	"on.file-changed":   actions.OnFileChanged,
	"on.file-created":   actions.OnFileCreated,
	"on.time":           actions.OnTime,
	"on.wifi":           actions.OnWifi,

	"get.clipboard":     actions.GetClipboard,
	"get.user-and-host": actions.GetUserAndHost,
	"get.wifi-state":    actions.GetWifiState,

	"do.calculate":          actions.Calculate,
	"do.copy":               actions.CliboardCopy,
	"do.copy-self-destruct": actions.CliboardCopySelfDestruct,
	"do.empty":              actions.Empty,
	"do.file-open":          actions.FileOpen,
	"do.folder-create":      actions.FolderCreate,
	"do.notification":       actions.Notification,
	"do.qr-encode":          actions.QrEncode,
	"do.stop":               actions.Stop,
	"do.wifi":               actions.Wifi,
}

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func init() {
	var err error

	Logger, err = zap.NewProduction()
	if err != nil {
		println("failed to create logger:", err.Error())
		os.Exit(1)
	}

	defer Logger.Sync()

	Sugar = Logger.Sugar()
}

func m(b map[string]any, a map[string]any) map[string]any {
	res := map[string]any{}

	for k, v := range a {
		res[k] = v
	}

	for k, v := range b {
		vs, ok := v.(string)
		if ok && strings.HasPrefix(vs, ":->") {
			res[k] = res[v.(string)[3:]]
			delete(res, v.(string)[3:])
			continue
		}
		res[k] = v
	}
	return res
}

func handleSequence(sequence ConfigSequence, wg *sync.WaitGroup) {
	out := map[string]any{}
	for i, a := range sequence.Actions {
		if _, ok := Actions[a.Type]; !ok {
			Sugar.Errorw("action not found", "action", a, "sequence", sequence.Name)
			return
		}

		out = Actions[a.Type](m(a.Args, out))

		if !out["success"].(bool) {
			Sugar.Errorw("action failed", "action", a, "step", i+1, "sequence", sequence.Name)
			return
		}
	}

	Sugar.Infow("sequence finished", "sequence", sequence.Name)
	wg.Done()
}

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		Sugar.Fatalw("failed to read config file", "error", err.Error())
	}

	config := []ConfigSequence{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		Sugar.Fatalw("failed to parse config file", "error", err.Error())
	}

	var wg sync.WaitGroup
	for i, sequence := range config {
		go handleSequence(sequence, &wg)
		Sugar.Infow("started sequence", "sequence", sequence.Name, "number", i+1)
		wg.Add(1)
	}

	wg.Wait()
}
