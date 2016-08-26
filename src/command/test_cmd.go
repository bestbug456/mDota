package command

// Standard lib
import (
	"encoding/json"
	"fmt"
)

// Exsternal dependences libraries
import (
	"github.com/spf13/cobra"
)

// Internal dependences
import (
	"core"
	"core/libsvm-go"
)

var TestCmd = &cobra.Command{
	Use:     "test",
	Short:   "Test a model",
	Long:    `Perform a model test using a testSet and see how much prediction the acutal model make wrong.`,
	Example: "mdota analyze -u /path/to/user.json -t /path/to/train.json",
}

var (
	TestPath string
)

func init() {
	TestCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
	TestCmd.PersistentFlags().StringVarP(&TestPath, "tspath", "c", "", "the path contain the testset")
	TestCmd.PersistentFlags().StringVarP(&ModelPath, "mpath", "m", "", "the path contain the problem model")
	TestCmd.RunE = test
}

func test(cmd *cobra.Command, args []string) error {

	userFileData, err := genericImportFromFile(TestPath)
	if err != nil {
		return err
	}
	var testFile TestFile
	// decode json user
	err = json.Unmarshal(userFileData, &testFile)
	if err != nil {
		return err
	}
	users := testFile.Users
	modelData, err := genericImportFromFile(ModelPath)
	if err != nil {
		return err
	}
	var model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modelData, &model)
	if err != nil {
		return err
	}
	for i := 0; i < len(users); i++ {
		result := core.Predict(users[i].Feature, &model)
		if users[i].Name[0] != kLabelResult[result] {
			fmt.Println("Error: expected %s but having %s", users[i].Name[0], kLabelResult[result])
		}
	}
	fmt.Println("Done")
	return nil
}
