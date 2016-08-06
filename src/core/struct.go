package core

type TrainsetData struct {
	Input  []float64 `json:"i"`
	Output []string  `json:"o"`
}

type ExsternalTrainset struct {
	Data []TrainsetData `json:"training_data"`
}
