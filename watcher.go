package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watcher(ch chan string) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	//walk and add every dir to the watcher
	filepath.WalkDir(".", func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir(){
			w.Add(s)
		}

		return nil
	})

	log.Println("Watching for changes...")

	//block and lister fore events
	for event := range w.Events {
	
		//filter .go files 
		if !strings.HasSuffix(event.Name, ".go") {
			continue
		}
		if event.Op&fsnotify.Write == fsnotify.Write {
			log.Println("File changed: ", event.Name)
			ch <- event.Name
		}
	}
}