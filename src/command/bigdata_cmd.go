package command

// Standard lib
import (
	"encoding/json"
	"fmt"
	"log"
)

// Exsternal dependences libraries
import (
	"github.com/spf13/cobra"
)

// Internal dependences
import (
	"core"
	"core/queue"
)

var BigDataCmd = &cobra.Command{
	Use:     "bigData",
	Short:   "Setup mdota in big data mode",
	Long:    `Setup mdota in big data mode in order to analyze a lot of user using scaling logic`,
	Example: "mdota bigData -u /path/to/users.json -t /path/to/train.json -w worker_number",
}

var (
	WorkerNumber int
)

func init() {
	BigDataCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
	BigDataCmd.PersistentFlags().StringVarP(&TrainPath, "tpath", "t", "", "the path contain the trainset")
	BigDataCmd.PersistentFlags().IntVarP(&WorkerNumber, "worker", "w", 100, "the number of worker to setup")
	BigDataCmd.RunE = bigdata
}

func bigdata(cmd *cobra.Command, args []string) error {
	wq := make(chan queue.WorkRequest, 1000)
	tsdata, err := genericImportFromFile(TrainPath)
	if err != nil {
		return err
	}
	var Trainset core.ExsternalTrainset
	// decode json trainset
	err = json.Unmarshal(tsdata, &Trainset)
	if err != nil {
		return err
	}
	err = core.PrepareSVMforLargeParalelAnalysis(Trainset.Data, wq, WorkerNumber)
	if err != nil {
		return err
	}
	userFileData, err := genericImportFromFile(UserPath)
	if err != nil {
		return err
	}

	var usersFile UsersFile
	// decode json user
	err = json.Unmarshal(userFileData, &usersFile)
	if err != nil {
		return err
	}
	var work queue.WorkRequest
	users := usersFile.Users
	response := make(chan queue.Response, 1000)
	for i := 0; i < len(users); i++ {
		work.RequestId = i
		work.Response = response
		work.Feature = core.SetupData(users[i].Feature)
		wq <- work
	}
	for i := 0; i < len(users); i++ {
		result := <-response
		fmt.Printf("Result of analysis for the user %s: since you have %g %g %g %g %g feature ", users[i].Name, users[i].Feature[0], users[i].Feature[1], users[i].Feature[2], users[i].Feature[3], users[i].Feature[4])
		log.Printf("Result of analysis %d report you have %s role, congratulations! ", result.RequestId, kLabelResult[result.Result])
	}

	return nil

}
