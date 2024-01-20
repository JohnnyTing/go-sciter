package window

import (
	"runtime"

	"github.com/JohnnyTing/go-sciter"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

// https://github.com/golang/go/wiki/LockOSThread
// https://github.com/JohnnyTing/go-sciter/issues/201
func init() {
	runtime.LockOSThread()
}
