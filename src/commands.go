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

func buildProgram() string {
	output, err := exec.Command("sh", "../scripts/buildscript.sh").CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	return string(output)
}
