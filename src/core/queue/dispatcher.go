package queue

import (
	"core/libsvm-go"
)

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nworkers int, WorkQueue chan WorkRequest, model *libSvm.Model) error {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		worker, err := NewWorker(i+1, WorkerQueue, model)
		if err != nil {
			return err
		}
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()

	return nil
}
