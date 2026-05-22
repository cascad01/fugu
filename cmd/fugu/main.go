package main

import (
	"gitlab.tech.mvideo.ru/mvideoru/debug/processor"
	"log"
	"os"
)

func main() {
	dirs := os.Args[1:]
	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	for _, dir := range dirs {
		if err := processor.ProcessDir(dir); err != nil {
			log.Fatalf("getter-gen: %s: %v", dir, err)
		}
	}
}
