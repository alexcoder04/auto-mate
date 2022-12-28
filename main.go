package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/alexcoder04/auto-mate/actions"
)

type ConfigAction struct {
	Type string         `json:"type"`
	Args map[string]any `json:"args"`
}

var Actions = map[string]func(map[string]any) map[string]any{
	"calendar-add":      actions.CalendarAdd,
	"empty":             actions.Empty,
	"file-open":         actions.FileOpen,
	"folder-create":     actions.FolderCreate,
	"notification":      actions.Notification,
	"on.battery":        actions.OnBattery,
	"on.battery-status": actions.OnBatteryStatus,
	"on.file-changed":   actions.OnFileChanged,
	"on.time":           actions.OnTime,
	"on.wifi":           actions.OnWifi,
	"qr-encode":         actions.QrEncode,
	"stop":              actions.Stop,
	"wifi":              actions.Wifi,
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

func handleSequence(sequence []ConfigAction, wg *sync.WaitGroup) {
	out := map[string]any{}
	for _, a := range sequence {
		if _, ok := Actions[a.Type]; !ok {
			fmt.Printf("Action %s not found\n", a)
			return
		}

		out = Actions[a.Type](m(a.Args, out))

		if !out["success"].(bool) {
			fmt.Printf("Action %s failed\n", a)
			return
		}
	}
	wg.Done()
}

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("Failed to read config file: %s\n", err.Error())
		return
	}

	config := [][]ConfigAction{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err.Error())
		return
	}

	var wg sync.WaitGroup
	for _, sequence := range config {
		go handleSequence(sequence, &wg)
		wg.Add(1)
	}

	wg.Wait()
}
