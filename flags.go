package main

import "flag"

var (
	flagDir      = flag.String("dir", ".", "directory to watch")
	flagCommand  = flag.String("command", "", "command to run after build")
	flagBuild    = flag.String("build", "go build", "build command to run")
	flagBuildDir = flag.String("build-dir", ".", "directory to run build command in")
	flagExclude  = flag.String("exclude", ".git,vendor", "comma separated directories to exclude")
	flagInclude  = flag.String("include", ".go", "comma separated file extensions to watch e.g. .go,.env,.yaml")
	flagEnv      = flag.String("env", "", "comma separated env vars e.g. PORT=8080,ENV=dev")
)
