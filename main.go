package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	OutputDst string
	SourceDst string
}

func main() {
	conf := new(config)
	flag.StringVar(&conf.OutputDst, "output", "output1", "Output Folder")
	flag.StringVar(&conf.SourceDst, "source", "datafeed", "Check Path")
	flag.Parse()

	for _, v := range checkExt(conf.SourceDst, ".zip") {
		archive, err := zip.OpenReader(v)
		if err != nil {
			panic(err)
		}
		defer archive.Close()

		for _, f := range archive.File {
			filePath := filepath.Join(conf.OutputDst, f.Name)
			fmt.Println("unzipping file ", filePath)

			if !strings.HasPrefix(filePath, filepath.Clean(conf.OutputDst)+string(os.PathSeparator)) {
				fmt.Println("invalid file path")
				return
			}
			if f.FileInfo().IsDir() {
				fmt.Println("creating directory...")
				os.MkdirAll(filePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				panic(err)
			}

			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}

			fileInArchive, err := f.Open()
			if err != nil {
				panic(err)
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				panic(err)
			}

			dstFile.Close()
			fileInArchive.Close()
		}

		if err = os.Remove(v); err != nil {
			fmt.Printf("ERROR: %s", err.Error())
		} else {
			fmt.Println("Removed ", v)
		}
	}
}

func checkExt(basePath string, ext string) []string {
	var files []string
	filepath.Walk(basePath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if strings.Contains(strings.ToLower(f.Name()), ext) {
				files = append(files, path)
			}

		}
		return nil
	})

	return files
}
