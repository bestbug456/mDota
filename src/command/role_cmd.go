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

var RoleCmd = &cobra.Command{
	Use:     "role",
	Short:   "retrive a role for a user",
	Long:    `Perform a full analyze on a user using an existing model and etiquete him with the apropriate class.`,
	Example: "mdota role -u /path/to/user.json -m /path/to/model.json",
}

var (
	ModelPath string
)

func init() {
	RoleCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
	RoleCmd.PersistentFlags().StringVarP(&ModelPath, "mpath", "m", "", "the path contain the problem model")

	RoleCmd.RunE = role
}

func role(cmd *cobra.Command, args []string) error {
	tsdata, err := genericImportFromFile(ModelPath)
	if err != nil {
		return err
	}
	var Model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(tsdata, &Model)
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
	for i := 0; i < len(users); i++ {
		fmt.Printf("Result of analysis for the user %s: since you have %g %g %g %g %g feature ", users[i].Name, users[i].Feature[0], users[i].Feature[1], users[i].Feature[2], users[i].Feature[3], users[i].Feature[4])
		result := core.Predict(users[i].Feature, &Model)
		fmt.Printf("the system say you have %s role, congratulations!\n", kLabelResult[result])
	}
	return nil

}
