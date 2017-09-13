package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"project/config"

	"github.com/fsnotify/fsnotify"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please, specify the config file to be loaded.")
		os.Exit(1)

	} else {
		config := config.GetConfig(os.Args[1])

		fmt.Println("Watching dir:", config.Dir)
		fmt.Println(" - Create cmd:", config.Create.Cmd)
		startWatch(config)

	}

}

func startWatch(config config.Config) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// log.Println("event:", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					// log.Println("Arquivo CRIADO:", event.Name)
					log.Println("Running command: ", config.Create.Cmd)
					cmd := exec.Command("sh", "-c", config.Create.Cmd)
					err := cmd.Run()
					if err != nil {
						log.Println("ERRO CREATE: ", err)
					}

				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					// log.Println("Arquivo REMOVIDO:", event.Name)
					log.Println("Running command: ", config.Delete.Cmd)

					cmd := exec.Command("sh", "-c", config.Delete.Cmd)
					err := cmd.Run()
					if err != nil {
						log.Println("ERRO DELETE: ", err)
					}

				} else if event.Op&fsnotify.Write == fsnotify.Write {
					// log.Println("Arquivo ALTERADO:", event.Name)
					log.Println("Running command: ", config.Change.Cmd)

					cmd := exec.Command("sh", "-c", config.Change.Cmd)
					err := cmd.Run()
					if err != nil {
						log.Println("ERRO CHANGE: ", err)
					}

				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(config.Dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
