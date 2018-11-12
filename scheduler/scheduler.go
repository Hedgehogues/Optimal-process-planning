package scheduler

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"

	"github.com/mohae/deepcopy"
)

type FileSlice []File
type FilesMap map[int]File
type FilesSet map[int]bool
type File struct {
	Id   int
	Time int // in seconds. If less second, than set zero
	Size int // in megabytes. If less one, than set zero
}

func NewErrorMsg(format string, args ...interface{}) error {
	pc, fn, line, _ := runtime.Caller(1)

	fname := path.Base(runtime.FuncForPC(pc).Name())
	if cwd, err := os.Getwd(); err == nil {
		fn = strings.Replace(fn, cwd, "", 1)
	}

	x := fmt.Sprintf(format, args...)
	return fmt.Errorf("[error] in %s [%s:%d] %s", fname, fn, line, x)
}

type Scheduler struct {
	remainingFiles FileSlice
	currentFiles   FilesMap
	currentSize    int
	maxSize        int
	maxCountFiles  int
}

func checkDuplicateId(fileSlice FileSlice) error {
	fileSet := FilesSet{}
	for _, file := range fileSlice {
		if file.Id < 0 {
			return NewErrorMsg("FileId must be more or equal 0. FileId: %i", file.Id)
		}
		if _, ok := fileSet[file.Id]; ok {
			return NewErrorMsg("Duplicate file id: %i", file.Id)
		}
		fileSet[file.Id] = true
	}
	return nil
}

func checkSize(fileSlice FileSlice, maxSize int) error {
	if maxSize <= 0 {
		return NewErrorMsg("MaxSize must be more 0. MaxSize: %i", maxSize)
	}
	for _, file := range fileSlice {
		if file.Size > maxSize {
			return NewErrorMsg("Large size. FileId: %i. Size of file: %i", file.Id, file.Size)
		}
	}
	return nil
}

func NewScheduler(initFiles FileSlice, maxSize, maxCountFiles int) (*Scheduler, error) {
	err := checkSize(initFiles, maxSize)
	if err != nil {
		return nil, err
	}
	err = checkDuplicateId(initFiles)
	if err != nil {
		return nil, err
	}
	if maxCountFiles <= 0 {
		return nil, NewErrorMsg("MaxCountFiles must be more 0. MaxCountFiles: %i", maxCountFiles)
	}
	scheduler := Scheduler{
		maxSize:        maxSize,
		maxCountFiles:  maxCountFiles,
		remainingFiles: deepcopy.Copy(initFiles).(FileSlice),
	}
	sort.Slice(scheduler.remainingFiles, func(i, j int) bool {
		return scheduler.remainingFiles[i].Time < scheduler.remainingFiles[j].Time ||
			scheduler.remainingFiles[i].Time == scheduler.remainingFiles[j].Time &&
				scheduler.remainingFiles[i].Size < scheduler.remainingFiles[j].Size
	})
	return &scheduler, nil
}

func (s *Scheduler) GetCurrentFiles() FileSlice {
	var currentFiles FileSlice
	for _, value := range s.currentFiles {
		currentFiles = append(currentFiles, value)
	}
	return currentFiles
}

func (s *Scheduler) Finished(fileId int) error {
	s.currentSize -= s.currentFiles[fileId].Size
	delete(s.currentFiles, fileId)
	if _, ok := s.currentFiles[fileId]; !ok {
		return NewErrorMsg("Not found FileId: %i", fileId)
	}
	return nil
}

// Return next files
// If remaining files is empty, than it return io.EOF
func (s *Scheduler) Next() (FileSlice, error) {
	if len(s.remainingFiles) == 0 {
		return nil, io.EOF
	}
	var newFiles FileSlice
	var newRemainingFiles FileSlice
	for _, file := range s.remainingFiles {
		if file.Size+s.currentSize >= s.maxSize || len(s.remainingFiles)+len(newFiles) == s.maxCountFiles {
			newRemainingFiles = append(newRemainingFiles, file)
			continue
		}
		s.currentSize += file.Size
		if _, ok := s.currentFiles[file.Id]; ok {
			return nil, NewErrorMsg("Duplicate FileId: %i", file.Id)
		}
		s.currentFiles[file.Id] = file
		newFiles = append(newFiles, file)
	}
	s.remainingFiles = newRemainingFiles
	return newFiles, nil
}
