package utils

import (
	"ezfaas/build"
	"fmt"
)

func ParseArguments(args []string) int {
	if len(args) <= 1 {
		DisplayHelp()
		return 1
	}
	param := args[1]
	switch param {
	case "build":
		fmt.Println("Building...")
		build.BuildEzfaasFunction(args[2])
	case "push":
		fmt.Println("Pushing...")
	case "deploy":
		fmt.Println("Deploying...")
	case "login":
		fmt.Println("Checking if Docker is authenticated...")
		checkDockerLogin()
	case "version":
		fmt.Println("1.0.0")
	default:
		fmt.Println("Invalid parameter " + param)
	}
	return 1
}
