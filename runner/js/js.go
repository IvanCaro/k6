package js

import (
	"github.com/loadimpact/speedboat/runner"
	"github.com/robertkrimen/otto"
	"time"
)

type JSRunner struct {
	BaseVM *otto.Otto
	Script *otto.Script
}

func New() (r *JSRunner, err error) {
	r = &JSRunner{}

	// Create a base VM
	r.BaseVM = otto.New()

	// Bridge basic functions
	r.BaseVM.Set("sleep", jsSleepFactory(time.Sleep))

	return r, nil
}

func (r *JSRunner) Load(filename, src string) (err error) {
	r.Script, err = r.BaseVM.Compile(filename, src)
	return err
}

func (r *JSRunner) RunVU() <-chan runner.Result {
	out := make(chan runner.Result)

	go func() {
		vm := r.BaseVM.Copy()
		vm.Run(r.Script)
		close(out)
	}()

	return out
}