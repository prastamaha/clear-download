package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// WalkMatch
func searchFile(dir string, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func main() {
	userHome := os.Getenv("HOME")
	downloadDir := userHome + "/Downloads"
	os.Chdir(downloadDir)

	fmt.Print("File extensions to be deleted: ")
	var exten string
	fmt.Scanf("%s", &exten)
	fileExten := "*." + exten

	pngs, err := searchFile(downloadDir, fileExten)
	if err != nil {
		log.Fatal(err)
	}
	for _, png := range pngs {
		fmt.Println(png)
	}

	var sure string = "N"
	fmt.Print("Are you sure you want to delete the files [y/N]? ")
	fmt.Scanf("%s", &sure)
	// fmt.Println(sure)

	if sure == "y" || sure == "yes" || sure == "Y" {
		for _, png := range pngs {
			e := os.Remove(png)
			if e != nil {
				log.Fatal(e)
			}
		}
		fmt.Println("Files have been deleted")
	} else if sure == "n" || sure == "no" || sure == "N" {
		fmt.Println("Canceled")
	} else {
		fmt.Println("Unknown Args")
	}

}
