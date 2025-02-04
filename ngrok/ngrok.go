package ngrok

import (
	"fmt"
	"os"
	"os/exec"
)

func Tunnel(token string, port string) {
	app := "ngrok"

	ngrokToken := token
	os.Setenv("ngrok_token", ngrokToken)

	cmd := exec.Command(app, "config", "add-authtoken", ngrokToken)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error adding auth token: %s\n", err.Error())
		fmt.Println(string(stdout))
		return
	}

	fmt.Println("Auth token added successfully.")

	cmd2 := exec.Command(app, "http", port)
	stdout2, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Printf("Error starting ngrok: %s\n", err.Error())
		fmt.Println(string(stdout2))
		return
	}

	fmt.Println("Ngrok tunnel started successfully.")
	fmt.Println(string(stdout2))
}
