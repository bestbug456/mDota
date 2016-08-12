package command

import (
	"encoding/json"
	"math/rand"
	"testing"
)

// Internal dependences
import (
	"core"
	"core/libsvm-go"
	"core/queue"
)

func BenchmarkBigData100Request(b *testing.B) {
	b.StopTimer()
	wq := make(chan queue.WorkRequest, 100)
	modeldata, err := genericImportFromFile("../../data/model.json")
	if err != nil {
		b.Fatalf(err.Error())
	}
	var model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modeldata, &model)
	if err != nil {
		b.Fatalf(err.Error())
	}
	err = core.PrepareSVMforLargeParalelAnalysis(&model, wq, 100)
	if err != nil {
		b.Fatalf(err.Error())
	}

	var users []Userdata
	// Create 1K of user
	var usr Userdata
	for i := 0; i < 999; i++ {
		features := make([]float64, 5)
		for i := 0; i < 5; i++ {
			features[i] = (rand.Float64() * 10.0)
		}
		usr.Feature = features
		usr.Name = "Test"
		users = append(users, usr)
	}

	var work queue.WorkRequest
	response := make(chan queue.Response, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		for i := 0; i < len(users); i++ {
			work.RequestId = i
			work.Response = response
			work.Feature = core.SetupData(users[i].Feature)
			wq <- work
		}

		for i := 0; i < len(users); i++ {
			_ = <-response
		}
	}
}

func BenchmarkBigData1000Request(b *testing.B) {
	b.StopTimer()
	wq := make(chan queue.WorkRequest, 1000)
	modeldata, err := genericImportFromFile("../../data/model.json")
	if err != nil {
		b.Fatalf(err.Error())
	}
	var model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modeldata, &model)
	if err != nil {
		b.Fatalf(err.Error())
	}
	err = core.PrepareSVMforLargeParalelAnalysis(&model, wq, 1000)
	if err != nil {
		b.Fatalf(err.Error())
	}

	var users []Userdata
	// Create 1K of user
	var usr Userdata
	for i := 0; i < 999; i++ {
		features := make([]float64, 5)
		for i := 0; i < 5; i++ {
			features[i] = (rand.Float64() * 10.0)
		}
		usr.Feature = features
		usr.Name = "Test"
		users = append(users, usr)
	}

	var work queue.WorkRequest
	response := make(chan queue.Response, 1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		for i := 0; i < len(users); i++ {
			work.RequestId = i
			work.Response = response
			work.Feature = core.SetupData(users[i].Feature)
			wq <- work
		}

		for i := 0; i < len(users); i++ {
			_ = <-response
		}
	}
}

func BenchmarkBigData10000Request(b *testing.B) {
	b.StopTimer()
	wq := make(chan queue.WorkRequest, 10000)
	modeldata, err := genericImportFromFile("../../data/model.json")
	if err != nil {
		b.Fatalf(err.Error())
	}
	var model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modeldata, &model)
	if err != nil {
		b.Fatalf(err.Error())
	}
	err = core.PrepareSVMforLargeParalelAnalysis(&model, wq, 10000)
	if err != nil {
		b.Fatalf(err.Error())
	}

	var users []Userdata
	// Create 1K of user
	var usr Userdata
	for i := 0; i < 999; i++ {
		features := make([]float64, 5)
		for i := 0; i < 5; i++ {
			features[i] = (rand.Float64() * 10.0)
		}
		usr.Feature = features
		usr.Name = "Test"
		users = append(users, usr)
	}

	var work queue.WorkRequest
	response := make(chan queue.Response, 10000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		for i := 0; i < len(users); i++ {
			work.RequestId = i
			work.Response = response
			work.Feature = core.SetupData(users[i].Feature)
			wq <- work
		}

		for i := 0; i < len(users); i++ {
			_ = <-response
		}
	}
}

func BenchmarkBigData100000Request(b *testing.B) {
	b.StopTimer()
	wq := make(chan queue.WorkRequest, 100000)
	modeldata, err := genericImportFromFile("../../data/model.json")
	if err != nil {
		b.Fatalf(err.Error())
	}
	var model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modeldata, &model)
	if err != nil {
		b.Fatalf(err.Error())
	}
	err = core.PrepareSVMforLargeParalelAnalysis(&model, wq, 100000)
	if err != nil {
		b.Fatalf(err.Error())
	}

	var users []Userdata
	// Create 1K of user
	var usr Userdata
	for i := 0; i < 999; i++ {
		features := make([]float64, 5)
		for i := 0; i < 5; i++ {
			features[i] = (rand.Float64() * 10.0)
		}
		usr.Feature = features
		usr.Name = "Test"
		users = append(users, usr)
	}

	var work queue.WorkRequest
	response := make(chan queue.Response, 100000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		for i := 0; i < len(users); i++ {
			work.RequestId = i
			work.Response = response
			work.Feature = core.SetupData(users[i].Feature)
			wq <- work
		}

		for i := 0; i < len(users); i++ {
			_ = <-response
		}
	}
}

func BenchmarkBigData1000000Request(b *testing.B) {
	b.StopTimer()
	wq := make(chan queue.WorkRequest, 1000000)
	modeldata, err := genericImportFromFile("../../data/model.json")
	if err != nil {
		b.Fatalf(err.Error())
	}
	var model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modeldata, &model)
	if err != nil {
		b.Fatalf(err.Error())
	}
	err = core.PrepareSVMforLargeParalelAnalysis(&model, wq, 1000000)
	if err != nil {
		b.Fatalf(err.Error())
	}

	var users []Userdata
	// Create 1K of user
	var usr Userdata
	for i := 0; i < 999; i++ {
		features := make([]float64, 5)
		for i := 0; i < 5; i++ {
			features[i] = (rand.Float64() * 10.0)
		}
		usr.Feature = features
		usr.Name = "Test"
		users = append(users, usr)
	}

	var work queue.WorkRequest
	response := make(chan queue.Response, 1000000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		for i := 0; i < len(users); i++ {
			work.RequestId = i
			work.Response = response
			work.Feature = core.SetupData(users[i].Feature)
			wq <- work
		}
		for i := 0; i < len(users); i++ {
			_ = <-response
		}
	}
}
