// Copyright  The Shinichi Nakagawa. All rights reserved.
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"path"
	"flag"
)


func IsExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func MakeWorkDirectory(dirname string) {
	if IsExist(dirname) {
		fmt.Println("Directory delete")
		if err := os.RemoveAll(dirname); err != nil {
			fmt.Println(err)
		}
	}
	os.MkdirAll(dirname, 0777)
}

func DownloadArchives(url string, dirname string) {
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
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", dirname, filename), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		file.Close()
	}()

	file.Write(body)
}


func GetEventsFile(year int, dirname string) {
	var url string = fmt.Sprintf("http://www.retrosheet.org/events/%deve.zip", year)
	// file Download
	DownloadArchives(url, dirname)
	return
}

func GetGameLogs(year int, dirname string) {
	var url string = fmt.Sprintf("http://www.retrosheet.org/gamelogs/gl%d.zip", year)
	// file Download
	DownloadArchives(url, dirname)
	return
}

func main() {
	// Commandline Options
	var fromYear = flag.Int("f", 2010, "Season Year(From)")
	var toYear = flag.Int("t", 2014, "Season Year(To)")
	flag.Parse()

	// make dir
	var dirname string = "files"
	MakeWorkDirectory(dirname)

	// Events/Game Log download
	for year := *fromYear; year < *toYear + 1; year++ {
		fmt.Println(fmt.Sprintf("Get Retrosheet Archives(%d Season)", year))
		GetEventsFile(year, dirname)
		GetGameLogs(year, dirname)
	}

}
