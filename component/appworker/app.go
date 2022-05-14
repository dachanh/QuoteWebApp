package appworker

import (
	"fmt"
	"sync"
)

type AppWorkers struct {
	WorkerList []*AppWorker
}

func (app *AppWorkers) SetUpWorker() *AppWorker{
	var worker = &AppWorker{}
	app.WorkerList = append(app.WorkerList,worker)
	return  worker
}

func (app *AppWorkers) Run() error{
	fmt.Println("TEST WORKER")
	go func() {
		var wg = sync.WaitGroup{}
		if len(app.WorkerList) > 0{
			for _ , w := range app.WorkerList{
				wg.Add(1)
				go w.Execute()
			}
		}
		wg.Wait()
	}()
	return nil
}