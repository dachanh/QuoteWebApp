package appworker

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Task = func()


type AppWorker struct {
	Task Task
	delay int
	idle int
}

func (w *AppWorker) SetTask(fn Task) *AppWorker{
	w.Task = fn
	return w
}

func (w *AppWorker) SetDelayTime(delay int) *AppWorker{
	w.delay = delay
	return w
}

func (w *AppWorker) SetIdleTime(idle int) *AppWorker{
	w.idle = idle
	return w
}

func (w *AppWorker) Execute(){
	time.Sleep(time.Second * time.Duration((w.delay)))
	tick := time.NewTicker(time.Second*time.Duration(w.idle))
	go w.Schedule(tick)

	sigs := make(chan os.Signal,1)
	signal.Notify(sigs, syscall.SIGINT,syscall.SIGTERM)
	<- sigs
	tick.Stop()
	os.Exit(0)
}

func(w *AppWorker) Schedule(tick *time.Ticker){
	for range tick.C{
		w.Task()
	}
}