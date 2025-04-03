package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"go.henil.dev/jrnl/fs"
	"go.henil.dev/jrnl/ui"
	"go.henil.dev/jrnl/utils"
)

var (
	previewFlag = flag.Bool("p", false, "renders the record before exiting")
	pickFlag    = flag.Bool("l", false, "allows to pick a record")
	todayFlag   = flag.Bool("t", false, "allows to pick records from today")
)

var logger = utils.GetDefaultLogger()

func openEditor(path string) {

	logger.Infof("Opening editor for file: %s", path)

	cmd := exec.Command("nano", "--nonewlines", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logger.Errorf("Error running editor: %s", err)
	}
}

func main() {
	flag.Parse()

	if *pickFlag || *todayFlag {

		var rootPath string = fs.GetRecordsStorageDirectoryPath()

		if *todayFlag {
			rootPath = fs.GetTodaysRecordsDirectoryPath()
		}

		filepath := ui.ShowFilePicker(rootPath)

		if filepath == "" {
			logger.Warn("No file selected")
			return
		}

		newRecord := fs.NewRecordFromPath(filepath)

		if !newRecord.Exists() {
			logger.Warn("File does not exist")
			return
		}

		renderedText, err := newRecord.GetRender()

		if err != nil {
			panic(err)
		}

		fmt.Println(renderedText)

		return
	}

	todaysRecordsDirectory := fs.NewTodaysRecordsDirectory()

	err := todaysRecordsDirectory.CreateIfNot()

	if err != nil {
		panic(err)
	}

	newRecord := fs.NewCurrentTimeRecord(todaysRecordsDirectory)

	_, err = newRecord.CreateIfNot()

	if err != nil {
		panic(err)
	}

	openEditor(newRecord.Path)

	if content, _ := newRecord.GetContent(); content == "" {
		logger.Warn("Record is empty, discarding")
		if err := newRecord.Delete(); err != nil {
			logger.Fatal("Failed to delete empty record")
			panic(err)
		}
		return
	}

	if *previewFlag {

		renderedText, err := newRecord.GetRender()

		if err != nil {
			panic(err)
		}

		fmt.Println(renderedText)
	}

}
