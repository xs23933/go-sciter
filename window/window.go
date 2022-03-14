package window

import (
	"runtime"

	"github.com/xs23933/go-sciter"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

// https://github.com/golang/go/wiki/LockOSThread
// https://github.com/xs23933/go-sciter/issues/201
func init() {
	runtime.LockOSThread()
}
