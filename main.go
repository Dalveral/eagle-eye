package main

import (
	"app/executor"
	"fmt"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	newPath := homeDir + "/Desktop/Screenshots"

	isFolder := executor.CheckIfDirExists(newPath)

	if isFolder == false {
		executor.ExecuteCommand("mkdir", []string{newPath})
	}

	executor.ExecuteCommand("defaults", []string{"write", "com.apple.screencapture", "location", newPath})
	executor.ExecuteCommand("killall", []string{"SystemUIServer"})
}
