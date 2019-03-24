package runner

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

type Runner struct {
	interupt chan os.Signal
	complete chan error
	timeout <-chan time.Time
	tasks []func(int)
}


var ErrorTimeout = errors.New("received timeout")
var ErrorInterupt = errors.New("received interupt")

func New(d time.Duration) *Runner  {
	return &Runner{
		interupt:make(chan os.Signal,1),
		complete:make(chan error),
		timeout:time.After(d),
	}
}

func (runner *Runner) Add(tasks ...func(int))  {
	runner.tasks = append(runner.tasks,tasks...)

}

func (runner *Runner) Start() error {

	signal.Notify(runner.interupt,os.Interrupt)

	go func() {
		runner.complete <- runner.Run()
	}()

	select {
	case err := <- runner.complete:
		return err
	case <- runner.timeout:
		return ErrorTimeout
	}

}

func (runner *Runner) Run() error  {
	for id,task := range runner.tasks {
		if runner.GotInterupt() {
			return ErrorInterupt
		}
		task(id)
	}

	return nil

}

func (runner *Runner) GotInterupt() bool {
	select {

	case <- runner.interupt:
		signal.Stop(runner.interupt)
		return true

	default:
		return false
	}
}