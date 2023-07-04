package main

import (
	"io/ioutil"
	"path"
	"strings"
	"sync"

	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func main() {
	files, err := ioutil.ReadDir("./config")
	if err != nil {
		Sugar.Fatalw("failed to read config folder", "error", err.Error())
	}

	var wg sync.WaitGroup
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".lua") {
			continue
		}

		code, err := ioutil.ReadFile(path.Join("./config", file.Name()))
		if err != nil {
			Sugar.Warnf("failed to read config file '%s'", file.Name())
			continue
		}

		wg.Add(1)
		go func(code []byte) {
			engine := NewEngine()
			engine.Run(string(code))
			engine.Shutdown()
			wg.Done()
		}(code)
	}

	wg.Wait()
}
