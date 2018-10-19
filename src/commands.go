package main

import (
	"os"
	"os/exec"
)

func execPullCommand() string {
	output, err := exec.Command("git", "pull").CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	return string(output)
}
