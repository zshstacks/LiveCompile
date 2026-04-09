# LiveCompile

A simple live reload tool for Go. Watches your project for file changes, rebuilds automatically, and restarts your binary.

## Installation

```bash
go install github.com/zshstacks/LiveCompile@latest
```

Or clone and build manually:

```bash
git clone https://github.com/zshstacks/LiveCompile
cd LiveCompile
go build .
```

## Usage

In its simplest form, run it from your project directory:

```bash
LiveCompile
```

To also run your binary after each successful build:

```bash
LiveCompile -command="./myapp"
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-dir` | `.` | Directory to watch |
| `-build-dir` | `.` | Directory to run the build command in |
| `-build` | `go build` | Build command to run |
| `-command` | `` | Command to run after a successful build |
| `-exclude` | `.git,vendor` | Comma separated directories to exclude |
| `-include` | `.go` | Comma separated file extensions to watch |

## Examples

Watch the current directory and run the binary after each build:
```bash
LiveCompile -command="./myapp"
```

Watch a different directory:
```bash
LiveCompile -dir=./myapp -build-dir=./myapp -command=./myapp/myapp
```

Watch additional file types like `.env` and `.yaml`:
```bash
LiveCompile -command="./myapp" -include=".go,.env,.yaml"
```

Exclude additional directories:
```bash
LiveCompile -command="./myapp" -exclude=".git,vendor,tmp"
```

## How it works

- Watches all subdirectories recursively using [fsnotify](https://github.com/fsnotify/fsnotify)
- Debounces file change events so rapid saves only trigger one build
- Kills the running binary before starting the new one
- Handles `Ctrl+C` gracefully by cleaning up the child process before exiting

## License
[BSD-2-Clause license](./LICENSE)