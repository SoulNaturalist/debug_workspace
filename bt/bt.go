package bt

import (
	"fmt"
	"os/exec"
	"runtime"
)

func CreateNewScreen() string {
	cmd := exec.Command("screen")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error creating screen: %s\n", err.Error())
		fmt.Println(string(stdout))
		return "error creating"
	}
	return string(stdout)

}

func KillAllScreen() string {
	if runtime.GOOS == "windows" {
		fmt.Print("Windows is not supported")
		return "Windows is not supported"
	} else {
		cmd := exec.Command("screen", "-wipe")
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error kill screen: %s\n", err.Error())
			fmt.Println(string(stdout))
			return "error kill screen"
		}
		return string(stdout)
	}
}
