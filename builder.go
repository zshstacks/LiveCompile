package main

import (
	"log"
	"os"
	"os/exec"
)

func builder() {

	killBinary()

	cmd := exec.Command("go", "build", ".")
	cmd.Dir = *flagBuildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("Build failed: ", err)
		return
	}

	log.Println("Build OK")

	if *flagCommand != "" {
		runBinary(*flagCommand)
	}

}
