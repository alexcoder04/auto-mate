package notification

import (
	"github.com/TheCreeper/go-notify"
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"show": Show,
}

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

func Show(l *lua.LState) int {
	// first arg is notification body
	ntf := notify.NewNotification("auto-mate", l.ToString(1))

	_, err := ntf.Show()
	if err != nil {
		return 0
	}

	return 0
}
