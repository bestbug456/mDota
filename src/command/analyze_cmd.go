package command

// Standard lib
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Exsternal dependences libraries
import (
	"github.com/spf13/cobra"
)

// Internal dependences
import (
	"core"
)

var AnalyzeCmd = &cobra.Command{
	Use:     "analyze",
	Short:   "Analyze a user",
	Long:    `Perform a full analyze on a full analyze on a user and etiquete him with the apropriate class.`,
	Example: "mdota analyze -u /path/to/user.json -t /path/to/train.json",
}

var (
	UserPath     string
	TrainPath    string
	kLabelResult = map[float64]string{1: "support", 2: "carry", 3: "roamer", 4: "midlaner", 5: "offlaner"} // , "initiator": 0.6 to do :D

)

func init() {
	AnalyzeCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
	AnalyzeCmd.PersistentFlags().StringVarP(&TrainPath, "tpath", "t", "", "the path contain the trainset")
	AnalyzeCmd.RunE = analyze
}

func analyze(cmd *cobra.Command, args []string) error {
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
	err = core.Setup(Trainset.Data)
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
	users := usersFile.Users
	model := core.Train()
	for i := 0; i < len(users); i++ {
		fmt.Printf("Result of analysis for the user %s: since you have %g %g %g %g %g feature ", users[i].Name, users[i].Feature[0], users[i].Feature[1], users[i].Feature[2], users[i].Feature[3], users[i].Feature[4])
		result := core.Predict(users[i].Feature, model)
		fmt.Printf("the system say you have %s role, congratulations!\n", kLabelResult[result])
	}
	return nil

}

// function that import data from a file
func genericImportFromFile(pathfile string) ([]byte, error) {
	// read the file
	data, err := ioutil.ReadFile(pathfile)
	if err != nil {
		return data, err
	}
	return data, nil
}
