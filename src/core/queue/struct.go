package queue

type WorkRequest struct {
	Feature   map[int]float64
	Response  chan Response
	RequestId int
}

type Response struct {
	Result    float64
	RequestId int
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
	Model       libSvm.Model
}
