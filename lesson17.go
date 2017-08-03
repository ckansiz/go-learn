package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var pathName string = "C:\\Base\\Database"
	ticker := time.NewTicker(time.Second * 2)

	go func() {
		for t := range ticker.C {
			fmt.Printf("Time : %s\n", t)
			getFileList(pathName)
		}
	}()

	var enter string
	fmt.Println("Press Enter for Exit")
	fmt.Scanln(&enter)
	ticker.Stop()
}

func getFileList(pathName string) {
	fmt.Println("________", pathName, "___________")
	filepath.Walk(pathName, func(path string, fileInfo os.FileInfo, err error) error {
		if !fileInfo.IsDir() {
			fmt.Printf("\t%s\t%d bytes\n", fileInfo.Name(), fileInfo.Size())
		}

		return nil
	})
}
