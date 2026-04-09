package main

import (
	"log"
	"os"
	"os/exec"
)

var currentProcess *exec.Cmd

func runBinary(command string)  {
	cmd := exec.Command(command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Println("Cant start the binary: ", err)
	}

	currentProcess = cmd

}

func killBinary() {
	if currentProcess != nil {
		currentProcess.Process.Kill()
		currentProcess = nil //cleanup
	}
}