package utils

import (
	"fmt"
	"os/exec"
)

func checkDockerLogin() int {
	path, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println("Docker CLI not found")
	}
	cmd := exec.Command(path, "login")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		fmt.Println("Docker login check failed. Please authenticate docker cli")
		return 1
	} else {
		fmt.Println("Docker Login Check Success")
	}
	return 0
}
