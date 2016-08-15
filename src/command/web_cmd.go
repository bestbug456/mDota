package command

// Standard lib
import (
	"encoding/json"
	"fmt"
	"strconv"
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

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var WebCmd = &cobra.Command{
	Use:     "web",
	Short:   "start web analyzer",
	Long:    `Start a web server that accept user data and etiquete him with the apropriate class.`,
	Example: "mdota web -m /path/to/model.json",
}

//var (
//	ModelPath string
//)

func init() {
	// RoleCmd.PersistentFlags().StringVarP(&UserPath, "upath", "u", "", "the path contain user to analyze")
	WebCmd.PersistentFlags().StringVarP(&ModelPath, "mpath", "m", "", "the path contain the problem model")

	WebCmd.RunE = web
}

func web(cmd *cobra.Command, args []string) error {
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

	//userFileData, err := genericImportFromFile(UserPath)
	//if err != nil {
	//	return err
	//}
	//var usersFile UsersFile
	//// decode json user
	//err = json.Unmarshal(userFileData, &usersFile)
	//if err != nil {
	//	return err
	//}
	//users := usersFile.Users
	//for i := 0; i < len(users); i++ {
	//	fmt.Printf("Result of analysis for the user %s: since you have %g %g %g %g %g feature ", users[i].Name, users[i].Feature[0], users[i].Feature[1], users[i].Feature[2], users[i].Feature[3], users[i].Feature[4])
	//	result := core.Predict(users[i].Feature, &Model)
	//	fmt.Printf("the system say you have %s role, congratulations!\n", kLabelResult[result])
	//}
	//return nil

	fmt.Printf("Starting web server")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/analyze", func(c *gin.Context) {

		f1, err1 := strconv.ParseFloat(c.Query("f1"), 64)
		f2, err2 := strconv.ParseFloat(c.Query("f2"), 64)
		f3, err3 := strconv.ParseFloat(c.Query("f3"), 64)
		f4, err4 := strconv.ParseFloat(c.Query("f4"), 64)
		f5, err5 := strconv.ParseFloat(c.Query("f5"), 64)

		if (err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil) {
			c.String(http.StatusOK, "Value Error")
			return
		}

		Feature := []float64{f1, f2, f3, f4, f5}
		fmt.Printf("Result of analysis for the user %s: since you have %g %g %g %g %g feature ", "Web User", Feature[0], Feature[1], Feature[2], Feature[3], Feature[4])
		result := core.Predict(Feature, &Model)
		fmt.Printf("the system say you have %s role, congratulations!\n", kLabelResult[result])
		c.HTML(http.StatusOK, "result.tmpl.html", gin.H{
			"result": kLabelResult[result],
		})
	})

	router.Run(":" + port)

	return nil

}
