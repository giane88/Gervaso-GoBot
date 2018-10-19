package main

import (
	"log"
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
	log.Print("Starting build")
	output, err := exec.Command("sh", "../scripts/buildscript.sh").CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	log.Print("build complete")
	return "Build completata con successo" + string(output)
}
