package command

// // Standard lib
// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// )

// // Exsternal dependences libraries
// import (
// 	"github.com/spf13/cobra"
// )

// // Internal dependences
// import (
// 	"core"
// 	"core/queue"
// )

// var BigDataCmd = &cobra.Command{
// 	Use:     "bigData",
// 	Short:   "Setup mdota in big data mode",
// 	Long:    `Setup mdota in big data mode in order to analyze a lot of user using scaling logic`,
// 	Example: "mdota bigData -u /path/to/users.json -t /path/to/train.json -w worker_number",
// }

// var (
// 	UserPath     string
// 	TrainPath    string
// 	WorkerNumber int
// 	kLabelResult = map[float64]string{1: "support", 2: "carry", 3: "roamer", 4: "midlaner", 5: "offlaner"} // , "initiator": 0.6 to do :D
// )

// func init() {
// 	BigDataCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
// 	BigDataCmd.PersistentFlags().StringVarP(&TrainPath, "tpath", "t", "", "the path contain the trainset")
// 	BigDataCmd.PersistentFlags().IntVarP(&WorkerNumber, "worker", "w", 100, "the number of worker to setup")
// 	BigDataCmd.RunE = bigdata
// }

// func bigdata(cmd *cobra.Command, args []string) error {
// 	wq := make(chan queue.WorkRequest, 1000)
// 	tsdata, err := genericImportFromFile(TrainPath)
// 	if err != nil {
// 		return err
// 	}
// 	var Trainset core.ExsternalTrainset
// 	// decode json trainset
// 	err = json.Unmarshal(tsdata, &Trainset)
// 	if err != nil {
// 		return err
// 	}
// 	err = core.PrepareSVMforLargeParalelAnalysis(Trainset.Data, wq, WorkerNumber)
// 	if err != nil {
// 		return err
// 	}
// 	userFileData, err := genericImportFromFile(UserPath)
// 	if err != nil {
// 		return err
// 	}

// 	var user []Userdata
// 	// decode json user
// 	err = json.Unmarshal(userFileData, &user)
// 	if err != nil {
// 		return err
// 	}
// 	var work queue.WorkRequest
// 	response := make(chan queue.Response, 1000)
// 	for i := 0; i < len(user); i++ {
// 		work.RequestId = i
// 		work.Response = response
// 		work.Feature = setupData(user[i].Feature)
// 		wq <- work
// 	}
// 	for i := 0; i < len(user); i++ {
// 		result := <-response
// 		log.Printf("Result of analysis %d report you have %s role, congratulations! ", result.RequestId, kLabelResult[result.Result])
// 	}

// 	return nil

// }

// func setupData(dataToAnalyze []float64) map[int]float64 {
// 	dataMap := make(map[int]float64, 0)
// 	for i := 0; i < len(dataToAnalyze); i++ {
// 		dataMap[i+1] = dataToAnalyze[i]
// 	}
// 	return dataMap
// }

// // function that import data from a file
// func genericImportFromFile(pathfile string) ([]byte, error) {
// 	// read the file
// 	data, err := ioutil.ReadFile(pathfile)
// 	if err != nil {
// 		return data, err
// 	}
// 	return data, nil
// }
