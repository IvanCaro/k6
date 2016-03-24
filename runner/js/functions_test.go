package js

import (
	"github.com/robertkrimen/otto"
	"testing"
	"time"
)

func makeTestSleepFunc() (func(time.Duration), <-chan time.Duration) {
	ch := make(chan time.Duration)
	fn := func(d time.Duration) {
		go func() {
			ch <- d
		}()
	}

	return fn, ch
}

func TestJSSleep(t *testing.T) {
	sleep, times := makeTestSleepFunc()

	vm := otto.New()
	vm.Set("sleep", jsSleepFactory(sleep))
	_, err := vm.Run(`sleep(1)`)
	if err != nil {
		t.Error("JS Error", err)
	}

	d := <-times
	if d != time.Duration(1)*time.Second {
		t.Error("Wrong amount of sleep", d)
	}
}

func TestJSSleepFraction(t *testing.T) {
	sleep, times := makeTestSleepFunc()

	vm := otto.New()
	vm.Set("sleep", jsSleepFactory(sleep))
	_, err := vm.Run(`sleep(0.1)`)
	if err != nil {
		t.Error("JS Error", err)
	}

	d := <-times
	if d != time.Duration(100)*time.Millisecond {
		t.Error("Wrong amount of sleep", d)
	}
}