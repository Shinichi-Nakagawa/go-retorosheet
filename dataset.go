// Copyright  The Shinichi Nakagawa. All rights reserved.
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

const (
	// ProjectRootDir : Project Root
	ProjectRootDir = "."
	// InputDirName : inputfile directory
	InputDirName = "./files"
	// OutputDirName : outputfile directory
	OutputDirName = "./csv"
	// CwPath : chadwick path
	CwPath = "/usr/local/bin"
	// CwEvent : cwevent command
	CwEvent = "%s/cwevent -q -n -f 0-96 -x 0-62 -y %d ./%d*.EV* > %s/events-%d.csv"
	// CwGame : cwgame command
	CwGame = "%s/cwgame -q -n -f 0-83 -y %d ./%d*.EV* > %s/games-%d.csv"
)

// ParseCsv a parse to eventfile(output:csv file)
func ParseCsv(command string, rootDir string, inputDir string) {
	os.Chdir(inputDir)
	out, err := exec.Command("sh", "-c", command).Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
	os.Chdir(rootDir)

}

func main() {
	// Commandline Options
	var fromYear = flag.Int("f", 2010, "Season Year(From)")
	var toYear = flag.Int("t", 2014, "Season Year(To)")
	flag.Parse()

	// path
	rootDir, _ := filepath.Abs(ProjectRootDir)
	inputDir, _ := filepath.Abs(InputDirName)
	outputDir := OutputDirName

	wait := new(sync.WaitGroup)
	// Generate URL
	commandList := []string{}
	for year := *fromYear; year < *toYear+1; year++ {
		commandList = append(commandList, fmt.Sprintf(CwEvent, CwPath, year, year, outputDir, year))
		wait.Add(1)
		commandList = append(commandList, fmt.Sprintf(CwGame, CwPath, year, year, outputDir, year))
		wait.Add(1)
	}
	for _, command := range commandList {
		fmt.Println(command)
		go func(command string) {
			ParseCsv(command, rootDir, inputDir)
			wait.Done()
		}(command)
	}
	wait.Wait()

}
