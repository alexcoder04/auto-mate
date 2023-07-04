package logic

import (
	"strings"

	"github.com/alexcoder04/auto-mate/amap"
	lua "github.com/yuin/gopher-lua"
)

var L *lua.LState

func init() {
	L = lua.NewState()

	L.DoString(`get = {}`)
	for action, function := range amap.Actions {
		if !strings.HasPrefix(action, "get.") {
			continue
		}
		L.SetGlobal(action, L.NewFunction(func(l *lua.LState) int {
			_ = function
			// wrapper around the real go function
			return 0
		}))
	}
}
