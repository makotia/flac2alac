package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func convert(filePath, outputDir string) {
	outputFilePath := path.Join(outputDir, strings.Replace(strings.Replace(filePath, ".flac", ".m4a", 1), "FLAC", "", 1))
	os.MkdirAll(filepath.Dir(outputFilePath), 0777)
	err := exec.Command("ffmpeg", "-i", filePath, "-acodec", "alac", "-vcodec", "copy", outputFilePath).Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	inputDir := os.Args[1]
	outputDir := os.Args[2]
	files := dirwalk(inputDir)
	for i := 0; i < len(files); i++ {
		fileName := files[i]
		if strings.HasSuffix(fileName, ".flac") {
			fmt.Println(fileName)
			convert(fileName, outputDir)
		}
	}
}
