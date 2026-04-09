package main

import "flag"

var (
	flagDir      = flag.String("dir", ".", "directory to watch")
	flagCommand  = flag.String("command", "", "command to run after build")
	flagBuild    = flag.String("build", "go build", "build command to run")
	flagBuildDir = flag.String("build-dir", ".", "directory to run build command in")
)
