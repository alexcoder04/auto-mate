package battery

import (
	"github.com/alexcoder04/auto-mate/lib/glbat"
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"level":  level,
	"status": status,
}

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

func getBatInfo() (glbat.Battery, error) {
	batIds, err := glbat.GetDetected()
	if err != nil {
		return glbat.Battery{}, err
	}

	return glbat.GetBat(batIds[0])
}

func level(l *lua.LState) int {
	bat, err := getBatInfo()
	if err != nil {
		l.Push(lua.LNumber(0))
		return 1
	}

	l.Push(lua.LNumber(bat.Capacity))
	return 1
}

func status(l *lua.LState) int {
	bat, err := getBatInfo()
	if err != nil {
		l.Push(lua.LString(bat.Status))
		return 1
	}

	l.Push(lua.LString(bat.Status))
	return 1
}
