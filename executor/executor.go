package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(command string, args []string) {
	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))

}

func CheckIfDirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
