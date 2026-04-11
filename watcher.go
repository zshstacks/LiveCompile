package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

func watcher(ch chan string) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	excludedDir := strings.Split(*flagExclude, ",")
	includedExts := strings.Split(*flagInclude, ",")

	//walk and add every dir to the watcher
	filepath.WalkDir(*flagDir, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			excluded := false
			for _, dir := range excludedDir {
				if strings.Contains(s, dir) {
					excluded = true
					break
				}
			}

			if !excluded {
				w.Add(s)
				log.Println("Watching dir:", s)
			}

		}

		return nil
	})

	color.Cyan("Watching for changes...")

	//block and lister fore events
	for event := range w.Events {
		//filter .go files
		matched := false
		for _, ext := range includedExts {
			if strings.HasSuffix(event.Name, ext) {
				matched = true
				break
			}
		}
		if !matched {
			continue
		}

		if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename {
			color.Yellow("File changed: %s", event.Name)
			ch <- event.Name
		}
	}
}
