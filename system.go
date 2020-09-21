package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	arguments := os.Args

	value, _ := exec.Command("sh", "-c", "systemctl show -p SubState --value "+arguments[1]).Output()

	status := strings.TrimSuffix(string(value), "\n")

	if status == "running" {
		fmt.Println("proccess is run ")
	} else if status == "exited" {
		fmt.Println("process running as daemon")

	} else if status == "dead" {
		fmt.Println("process not running ")

		cmd := exec.Command("sudo", "service", "openvpn", "start")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
