// This the core package
// for Support Vector Machine
// Dota matchmaking alghorimt
// created by Danilo 'bestbug'
// and Franca 'forcolotta'
package core

// Standard lib
import (
	"math"
)

// Third party library
import (
	"core/libsvm-go"
)

import (
	"core/queue"
)

const (
	kC     = 100
	kGamma = 0.1
)

var (
	Parameters       *libSvm.Parameter
	ProblemToAnalyze *libSvm.Problem
	model            *libSvm.Model
	kLabel           = map[string]float64{"support": 1, "carry": 2, "roamer": 3, "midlaner": 4, "offlaner": 5} // , "initiator": 0.6 to do :D
)

func Train() *libSvm.Model {
	model := libSvm.NewModel(Parameters)
	model.Train(ProblemToAnalyze)
	return model
}

func Predict(user []float64, model *libSvm.Model) float64 {
	x := SetupData(user)
	return model.Predict(x)
}

// Define parameters and problem
func Setup(Traindata []TrainsetData) error {
	for i := 0; i < len(Traindata); i++ {
		Traindata[i].Input = normalizeData(Traindata[i].Input)
	}
	Parameters = libSvm.NewParameter()
	Parameters.Gamma = kGamma
	Parameters.C = kC
	Parameters.QuietMode = true
	// L total number of sample
	// I number of iterations
	ProblemToAnalyze = &libSvm.Problem{L: 25, I: 0}

	ProblemToAnalyze.Y = nil      // label
	ProblemToAnalyze.X = nil      // slice of int with the starting index of XSpace
	ProblemToAnalyze.XSpace = nil // slice of snode, snode have 2 fileds: Index is the index and Value is the actual value the last element must have -1 in index
	j := 0
	for _, singleDataTrain := range Traindata {
		for i := 0; i < len(singleDataTrain.Input); i++ {
			var singleNode libSvm.Snode
			singleNode.Index = i + 1
			singleNode.Value = singleDataTrain.Input[i]
			ProblemToAnalyze.XSpace = append(ProblemToAnalyze.XSpace, singleNode)
		}
		var singleNode libSvm.Snode
		singleNode.Index = -1
		ProblemToAnalyze.XSpace = append(ProblemToAnalyze.XSpace, singleNode)

		ProblemToAnalyze.X = append(ProblemToAnalyze.X, j)
		ProblemToAnalyze.Y = append(ProblemToAnalyze.Y, kLabel[singleDataTrain.Output[0]])

		j += 6

	}

	//fmt.Println("X: ", ProblemToAnalyze.X, "\nY: ", ProblemToAnalyze.Y, "\nXSpace: ", ProblemToAnalyze.XSpace, "\n len Xspace", len(ProblemToAnalyze.XSpace), " len X: ", len(ProblemToAnalyze.X), " len y: ", len(ProblemToAnalyze.Y))

	return nil
}

func SetupData(dataToAnalyze []float64) map[int]float64 {
	dataToAnalyze = normalizeData(dataToAnalyze)
	dataMap := make(map[int]float64, 0)
	for i := 0; i < len(dataToAnalyze); i++ {
		dataMap[i+1] = dataToAnalyze[i]
	}
	return dataMap
}

func normalizeData(slice []float64) []float64 {
	mean := mean(slice)
	sd := standardDeviation(slice, mean)
	for i := 0; i < len(slice); i++ {
		slice[i] = (slice[i] - mean) / sd
	}
	return slice
}

func mean(slice []float64) float64 {
	total := 0.0
	for i := 0; i < len(slice); i++ {
		total += slice[i]
	}
	total = total / float64(len(slice))
	return total
}

func standardDeviation(slice []float64, mean float64) float64 {
	total := 0.0
	for i := 0; i < len(slice); i++ {
		total += (slice[i] - mean) * (slice[i] - mean)
	}
	total = total / float64(len(slice))
	return math.Sqrt(total)
}

func PrepareSVMforLargeParalelAnalysis(modelIn *libSvm.Model, WorkQueue chan queue.WorkRequest, nworkers int) error {

	model = modelIn
	queue.StartDispatcher(nworkers, WorkQueue, model)
	return nil
}
