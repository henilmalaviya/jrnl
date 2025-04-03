package fs

import (
	"os"
	"path"
	"strings"
	"time"

	"go.henil.dev/jrnl/utils"
)

func getAppDirectoryPath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	return path.Join(homeDir, ".jrnl")
}

func GetRecordsStorageDirectoryPath() string {
	return path.Join(getAppDirectoryPath(), "records")
}

func GetTodaysRecordsDirectoryPath() string {
	return path.Join(GetRecordsStorageDirectoryPath(), utils.FormatDate(time.Now()))
}

type RecordDirectory struct {
	Path string
}

func (d *RecordDirectory) CreateIfNot() error {
	return os.MkdirAll(d.Path, os.ModePerm)
}

func (d *RecordDirectory) Exists() bool {
	_, err := os.Stat(d.Path)
	return !os.IsNotExist(err)
}

func (d *RecordDirectory) GetRecords() []Record {

	files, err := os.ReadDir(d.Path)

	if err != nil {
		panic(err)
	}

	var records []Record

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		record := NewRecord(*d, fileName)

		records = append(records, record)
	}

	return records
}

func NewRecordDirectory(date string) RecordDirectory {
	return RecordDirectory{
		Path: path.Join(GetRecordsStorageDirectoryPath(), date),
	}
}

func NewTodaysRecordsDirectory() RecordDirectory {
	return NewRecordDirectory(utils.FormatDate(time.Now()))
}

type Record struct {
	Path string
}

func (r *Record) CreateIfNot() (*os.File, error) {
	return os.Create(r.Path)
}

func (r *Record) Exists() bool {
	_, err := os.Stat(r.Path)
	return !os.IsNotExist(err)
}

func (r *Record) GetParentDirectoryPath() string {
	return path.Dir(r.Path)
}

func (r *Record) GetContent() (string, error) {
	content, err := os.ReadFile(r.Path)

	return string(content), err
}

func (r *Record) Delete() error {
	return os.Remove(r.Path)
}

func (r *Record) GetRecordShortName() string {
	return path.Base(path.Dir(r.Path)) + "/" + path.Base(r.Path)
}

func (r *Record) GetRender() (string, error) {

	content, err := r.GetContent()

	content = "# " + r.GetRecordShortName() + "\n" + content

	if err != nil {
		return "", err
	}

	return utils.GetDefaultRenderer().Render(strings.TrimSpace(content))
}

func NewRecordFromPath(path string) Record {
	return Record{
		Path: path,
	}
}

func NewRecord(recordDir RecordDirectory, time string) Record {
	return NewRecordFromPath(path.Join(recordDir.Path, time))
}

func NewCurrentTimeRecord(recordDir RecordDirectory) Record {
	return NewRecord(recordDir, utils.FormatTime(time.Now()))
}
