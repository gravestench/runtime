package pkg

import (
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	rt := New()

	go func() {
		time.Sleep(time.Second * 3)
		rt.Shutdown()
	}()

	rt.Run()
}
