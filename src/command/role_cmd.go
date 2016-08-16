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
	ModelPath   string
	Fight       float64
	Versatility float64
	Push        float64
	Support     float64
	Farm        float64
)

func init() {
	RoleCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
	RoleCmd.PersistentFlags().StringVarP(&ModelPath, "mpath", "m", "./data/model.json", "the path contain the problem model")
	RoleCmd.PersistentFlags().Float64VarP(&Fight, "fight", "f", 0.0, "Your value for fight")
	RoleCmd.PersistentFlags().Float64VarP(&Versatility, "versatility", "v", 0.0, "Your value for versatility")
	RoleCmd.PersistentFlags().Float64VarP(&Push, "push", "p", 0.0, "Your value for push")
	RoleCmd.PersistentFlags().Float64VarP(&Support, "support", "s", 0.0, "Your value for support")
	RoleCmd.PersistentFlags().Float64VarP(&Farm, "farm", "g", 0.0, "Your value for Farm")
	RoleCmd.RunE = role
}

func role(cmd *cobra.Command, args []string) error {
	modeldata, err := genericImportFromFile(ModelPath)
	if err != nil {
		return err
	}
	var Model libSvm.Model
	// decode json trainset
	err = json.Unmarshal(modeldata, &Model)
	if err != nil {
		return err
	}
	var users []Userdata
	if UserPath != "" {
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
		users = usersFile.Users
	} else {
		var user Userdata
		user.Feature = []float64{Fight, Versatility, Push, Support, Farm}
		users = append(users, user)
	}

	for i := 0; i < len(users); i++ {
		fmt.Printf("Result of analysis: since you have %g %g %g %g %g feature ", users[i].Feature[0], users[i].Feature[1], users[i].Feature[2], users[i].Feature[3], users[i].Feature[4])
		result := core.Predict(users[i].Feature, &Model)
		fmt.Printf("the system say you have %s role, congratulations!\n", kLabelResult[result])
	}
	return nil

}
