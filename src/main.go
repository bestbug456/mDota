package main

// import (
// 	"io/ioutil"
// )

import (
	"fmt"
)

// // internal support library
// import (
// 	input "command"
// )

func init() {
	PrepareMainCommand()
}

func main() {
	_, err := mDotaCmd.ExecuteC()
	if err != nil {
		fmt.Println(err)
	}

}
