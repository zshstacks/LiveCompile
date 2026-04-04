package main

import (
	"log"
	"os"
	"os/exec"
)

func builder() {

	cmd := exec.Command("go", "build", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("Build failed: ", err)
		return
	}

	log.Println("Build OK")

}