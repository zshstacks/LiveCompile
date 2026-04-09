package main

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func builder() {

	killBinary()

	cmd := exec.Command("go", "build", ".")
	cmd.Dir = *flagBuildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("Build failed: %s", err)
		return
	}

	color.Green("Build OK")

	if *flagCommand != "" {
		runBinary(*flagCommand)
	}

}
