package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

func builder() {

	killBinary()

	cmd := exec.Command("go", "build", ".")
	cmd.Dir = *flagBuildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()

	if err := cmd.Run(); err != nil {
		color.Red("Build failed: %s", err)
		return
	}

	duration := time.Since(start)
	color.Green("Build OK (%.2fs)", duration.Seconds())

	if *flagCommand != "" {
		runBinary(*flagCommand)
	}

}
