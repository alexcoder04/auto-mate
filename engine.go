package main

import (
	"github.com/alexcoder04/auto-mate/mods/battery"
	"github.com/alexcoder04/auto-mate/mods/notification"
	lua "github.com/yuin/gopher-lua"
)

type LuaEngine struct {
	State *lua.LState
	On    bool
}

func (engine *LuaEngine) Run(code string) {
	err := engine.State.DoString(code)
	if err != nil {
		Sugar.Warnw("error in lua code", "error", err.Error())
	}
}

func (engine *LuaEngine) Shutdown() {
	engine.State.Close()
	engine.On = false
}

func NewEngine() LuaEngine {
	engine := LuaEngine{}
	engine.State = lua.NewState()
	engine.On = true

	engine.State.PreloadModule("dev", func(l *lua.LState) int {
		mod := l.SetFuncs(l.NewTable(), map[string]lua.LGFunction{
			"log": func(l *lua.LState) int {
				Sugar.Infow(l.ToString(1), "from", "lua")
				return 0
			},
		})
		l.Push(mod)
		return 1
	})

	engine.State.PreloadModule("battery", battery.Loader)
	engine.State.PreloadModule("notification", notification.Loader)

	return engine
}
