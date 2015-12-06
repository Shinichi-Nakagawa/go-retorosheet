// Copyright  The Shinichi Nakagawa. All rights reserved.
// license that can be found in the LICENSE file.

package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
)

const (
	// DirName a saved download file directory
	DirName = "files"
)

// IsExist return a file exist
func IsExist(filename string) bool {

	// Exists check
	_, err := os.Stat(filename)
	return err == nil
}

// MakeWorkDirectory a make work directory
func MakeWorkDirectory(dirname string) {
	if IsExist(dirname) {
		return
	}
	os.MkdirAll(dirname, 0777)
}

// DownloadArchives a download files for retrosheet
func DownloadArchives(url string, dirname string) string {

	// get archives
	fmt.Println(fmt.Sprintf("download: %s", url))
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("status: %s", response.Status))

	// download
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, filename := path.Split(url)
	fmt.Println(filename)
	fullfilename := fmt.Sprintf("%s/%s", dirname, filename)
	file, err := os.OpenFile(fullfilename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		file.Close()
	}()

	file.Write(body)

	return fullfilename

}

// Unzip return error a open read & write archives
func Unzip(fullfilename string, outputdirectory string) error {
	r, err := zip.OpenReader(fullfilename)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(outputdirectory, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetEventsFileURL return a events file URL
func GetEventsFileURL(year int) string {
	return fmt.Sprintf("http://www.retrosheet.org/events/%deve.zip", year)
}

// GetGameLogsURL return a game logs URL
func GetGameLogsURL(year int) string {
	return fmt.Sprintf("http://www.retrosheet.org/gamelogs/gl%d.zip", year)
}

func main() {
	// Commandline Options
	var fromYear = flag.Int("f", 2010, "Season Year(From)")
	var toYear = flag.Int("t", 2014, "Season Year(To)")
	flag.Parse()

	MakeWorkDirectory(DirName)

	wait := new(sync.WaitGroup)
	// Generate URL
	urls := []string{}
	for year := *fromYear; year < *toYear+1; year++ {
		urls = append(urls, GetEventsFileURL(year))
		wait.Add(1)
		urls = append(urls, GetGameLogsURL(year))
		wait.Add(1)
	}

	// Download files
	for _, url := range urls {
		go func(url string) {
			fullfilename := DownloadArchives(url, DirName)
			err := Unzip(fullfilename, DirName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			wait.Done()
		}(url)
	}
	wait.Wait()

}
