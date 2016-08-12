package core

import "core/libsvm-go"

type TrainsetData struct {
	Input  []float64 `json:"i"`
	Output []string  `json:"o"`
}

type ExsternalTrainset struct {
	Data []TrainsetData `json:"training_data"`
}

type ExsternalModel struct {
	Data libSvm.Model `json:"model"`
}
