package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

import (
	"core/libsvm-go"
)

var (
	kLabelResult = map[float64]string{1: "support", 2: "carry", 3: "roamer", 4: "midlaner", 5: "offlaner"} // , "initiator": 0.6 to do :D
)

func TestSvm(t *testing.T) {
	// import data file
	data, err := importFromFile("./data/trainingDOTA.json")
	if err != nil {
		t.Fatalf("Error while importing file: ", err.Error())
	}

	// run setup function
	Setup(data)
	// train
	model := libSvm.NewModel(Parameters)
	model.Train(ProblemToAnalyze)
	t.Log(model)
	user := []float64{6.81, 8.26, 8.31, 5.50, 7.26}
	x := setupData(user)
	ris := model.Predict(x)
	t.Log(kLabelResult[ris])
}

func BenchmarkCore(b *testing.B) {
	b.StopTimer()

	// import data file
	data, err := importFromFile("./data/trainingDOTA.json")
	if err != nil {
		b.Fatalf("Error while importing file: ", err.Error())
	}

	user := []float64{6.81, 8.26, 8.31, 5.50, 7.26}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// run setup function
		Setup(data)
		// train
		model := libSvm.NewModel(Parameters)
		model.Train(ProblemToAnalyze)
		x := setupData(user)
		model.Predict(x)
	}
}

// Support function

// function that import data from a file
func importFromFile(pathfile string) ([]TrainsetData, error) {
	var output ExsternalTrainset
	// read the file
	data, err := ioutil.ReadFile(pathfile)
	if err != nil {
		return output.Data, err
	}
	// decode json data
	err = json.Unmarshal(data, &output)
	if err != nil {
		return output.Data, err
	}
	return output.Data, nil
}
